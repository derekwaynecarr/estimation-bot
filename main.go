package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Config holds the application configuration
type Config struct {
	URLs      []string
	APIKey    string
	ModelName string
}

// GitHubContent represents the structure of GitHub API responses for file content
type GitHubContent struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

func main() {
	config := parseFlags()

	if len(config.URLs) == 0 {
		log.Fatal("No URLs provided. Use -urls flag to specify design document URLs.")
	}

	if config.APIKey == "" {
		log.Fatal("Gemini API key not provided. Set GEMINI_API_KEY environment variable or use -api-key flag.")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.APIKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel(config.ModelName)

	// Process each URL
	for i, url := range config.URLs {
		fmt.Printf("\n=== Processing URL %d/%d ===\n", i+1, len(config.URLs))
		fmt.Printf("URL: %s\n\n", url)

		content, err := fetchGitHubContent(url)
		if err != nil {
			log.Printf("Error fetching content from %s: %v", url, err)
			continue
		}

		estimation, err := getEstimation(ctx, model, content)
		if err != nil {
			log.Printf("Error getting estimation for %s: %v", url, err)
			continue
		}

		fmt.Println("=== ESTIMATION RESULT ===")
		fmt.Println(estimation)
		fmt.Println("\n" + strings.Repeat("=", 50))
	}
}

func parseFlags() Config {
	var config Config

	urlsFlag := flag.String("urls", "", "Comma-delimited list of GitHub URLs to design documents")
	apiKeyFlag := flag.String("api-key", "", "Gemini API key (can also be set via GEMINI_API_KEY env var)")
	modelFlag := flag.String("model", "gemini-1.5-pro", "Gemini model to use")

	flag.Parse()

	if *urlsFlag != "" {
		config.URLs = strings.Split(*urlsFlag, ",")
		// Trim whitespace from each URL
		for i, url := range config.URLs {
			config.URLs[i] = strings.TrimSpace(url)
		}
	}

	// API key from flag takes precedence, then environment variable
	if *apiKeyFlag != "" {
		config.APIKey = *apiKeyFlag
	} else {
		config.APIKey = os.Getenv("GEMINI_API_KEY")
	}

	config.ModelName = *modelFlag

	return config
}

func fetchGitHubContent(url string) (string, error) {
	// Convert GitHub URL to API URL if it's a regular GitHub file URL
	apiURL := convertToAPIURL(url)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// If it's a GitHub API response, parse the JSON and decode content
	if strings.Contains(apiURL, "api.github.com") {
		var ghContent GitHubContent
		if err := json.Unmarshal(body, &ghContent); err != nil {
			return "", fmt.Errorf("failed to parse GitHub API response: %w", err)
		}

		if ghContent.Encoding == "base64" {
			decoded, err := base64.StdEncoding.DecodeString(ghContent.Content)
			if err != nil {
				return "", fmt.Errorf("failed to decode base64 content: %w", err)
			}
			return string(decoded), nil
		}

		return ghContent.Content, nil
	}

	// For raw GitHub URLs, return content directly
	return string(body), nil
}

func convertToAPIURL(url string) string {
	// Convert github.com URLs to api.github.com URLs
	if strings.Contains(url, "github.com") && !strings.Contains(url, "api.github.com") {
		// Example: https://github.com/owner/repo/blob/main/file.md
		// Convert to: https://api.github.com/repos/owner/repo/contents/file.md

		url = strings.Replace(url, "github.com", "api.github.com/repos", 1)
		url = strings.Replace(url, "/blob/", "/contents/", 1)

		// Remove the branch name (assumes main/master, could be enhanced)
		parts := strings.Split(url, "/contents/")
		if len(parts) == 2 {
			// Remove branch name from the path
			pathParts := strings.SplitN(parts[1], "/", 2)
			if len(pathParts) == 2 {
				url = parts[0] + "/contents/" + pathParts[1]
			}
		}
	}

	// If it's already a raw GitHub URL, use it directly
	if strings.Contains(url, "raw.githubusercontent.com") {
		return url
	}

	return url
}

func getEstimation(ctx context.Context, model *genai.GenerativeModel, content string) (string, error) {
	prompt := fmt.Sprintf(EstimationPromptTemplate, content)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	// Extract text from the response
	var result strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			result.WriteString(string(textPart))
		}
	}

	return result.String(), nil
}
