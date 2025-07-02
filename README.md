# Estimation Bot

A command-line tool that assists with agile story point estimation by analyzing design documents hosted on GitHub using Google's Gemini AI.

## Overview

The Estimation Bot fetches design documents from GitHub URLs and uses the Gemini API to provide detailed story point estimations. It analyzes complexity, uncertainty, effort requirements, dependencies, and testing needs to suggest appropriate story points following agile estimation practices.

## Features

- 🔗 **GitHub Integration**: Fetch design documents directly from GitHub URLs
- 🤖 **AI-Powered Analysis**: Uses Google Gemini API for intelligent estimation
- 📊 **Structured Output**: Provides detailed analysis with reasoning, risk factors, and recommendations
- 🔢 **Fibonacci Scale**: Uses standard Fibonacci sequence for story points (1, 2, 3, 5, 8, 13)
- ⚙️ **Customizable Prompts**: Easily modify the AI prompt template for your team's needs
- 🎯 **Batch Processing**: Analyze multiple design documents in a single run

## Prerequisites

- **Go 1.21+**: The application is built with Go 1.21
- **Gemini API Key**: Obtain from [Google AI Studio](https://makersuite.google.com/app/apikey)
- **Internet Connection**: Required for GitHub API and Gemini API access

## Installation

### Option 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/derekwaynecarr/estimation-bot.git
cd estimation-bot

# Download dependencies
make deps

# Build the binary
make build
```

### Option 2: Using Go Install

```bash
go install github.com/derekwaynecarr/estimation-bot@latest
```

## Configuration

### API Key Setup

You can provide your Gemini API key in two ways:

1. **Environment Variable** (Recommended):
   ```bash
   export GEMINI_API_KEY="your-api-key-here"
   ```

2. **Command Line Flag**:
   ```bash
   ./estimation-bot -api-key="your-api-key-here" -urls="..."
   ```

## Usage

### Basic Usage

```bash
# Analyze a single design document
./estimation-bot -urls="https://github.com/owner/repo/blob/main/design.md"

# Analyze multiple documents
./estimation-bot -urls="https://github.com/owner/repo/blob/main/design1.md,https://github.com/owner/repo/blob/main/design2.md"
```

### Advanced Usage

```bash
# Use a different Gemini model
./estimation-bot -urls="https://github.com/owner/repo/blob/main/design.md" -model="gemini-1.5-flash"

# Specify API key inline
./estimation-bot -urls="https://github.com/owner/repo/blob/main/design.md" -api-key="your-key"
```

### Command Line Options

| Flag | Description | Default |
|------|-------------|---------|
| `-urls` | Comma-delimited list of GitHub URLs to design documents | Required |
| `-api-key` | Gemini API key (alternative to GEMINI_API_KEY env var) | - |
| `-model` | Gemini model to use | `gemini-1.5-pro` |

### Supported URL Formats

The tool supports various GitHub URL formats:

- Regular GitHub file URLs: `https://github.com/owner/repo/blob/main/file.md`
- Raw GitHub URLs: `https://raw.githubusercontent.com/owner/repo/main/file.md`
- GitHub API URLs: `https://api.github.com/repos/owner/repo/contents/file.md`

## Example Output

```
=== Processing URL 1/1 ===
URL: https://github.com/example/repo/blob/main/feature-design.md

=== ESTIMATION RESULT ===

## Story Point Estimation Analysis

### Recommended Story Points: 5 points

### Executive Summary
This feature involves implementing a new user authentication system with OAuth integration, requiring moderate complexity backend changes and comprehensive testing.

### Detailed Reasoning
The implementation requires:
- New authentication service integration
- Database schema modifications
- Frontend login flow updates
- Security testing and validation

### Complexity Analysis
- **Technical Complexity**: Medium - OAuth integration with existing auth system
- **Integration Complexity**: High - Multiple touchpoints across backend and frontend
- **Testing Complexity**: High - Security and integration testing requirements

### Risk Factors
- OAuth provider API changes or limitations
- Existing authentication system compatibility
- Security vulnerability introduction

### Effort Breakdown
- Backend Development: 40%
- Frontend Development: 25%
- Testing: 25%
- Documentation: 10%

### Dependencies
- OAuth provider approval and setup
- Security team review

### Recommendations
- **Should this be broken down?** No - Appropriate size for a single sprint
- **Prerequisites**: Security team consultation before implementation

### Assumptions Made
- OAuth provider (Google/GitHub) is already approved
- Existing authentication system can be extended
- Team has OAuth implementation experience

==================================================
```

## Customizing the Estimation Prompt

The AI prompt template is located in `prompt_template.go` and can be customized for your team's specific needs.

### Key Areas to Customize:

1. **Story Point Scale**: Modify the point values and descriptions
2. **Analysis Factors**: Add domain-specific considerations
3. **Output Format**: Adjust the structured response format
4. **Context**: Change the AI persona or add company-specific context

### Example Customization:

```go
// Add your team's specific considerations
const EstimationPromptTemplate = `You are a senior engineer at [YOUR_COMPANY] with expertise in [YOUR_DOMAIN]...

Consider our specific factors:
1. **Legacy System Impact**: How will this affect our legacy systems?
2. **Compliance Requirements**: Are there regulatory considerations?
3. **Team Expertise**: Does the team have required skills?
...`
```

## Development

### Building

```bash
# Build the application
make build

# Clean build artifacts
make clean

# Run tests
make test

# Install to $GOPATH/bin
make install
```

### Project Structure

```
estimation-bot/
├── main.go              # Main application logic
├── prompt_template.go   # Customizable AI prompt template
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── Makefile            # Build and development commands
└── README.md           # This file
```

## Troubleshooting

### Common Issues

1. **"API key not provided" error**:
   - Ensure `GEMINI_API_KEY` environment variable is set, or use `-api-key` flag

2. **"GitHub API returned status 404" error**:
   - Verify the GitHub URL is accessible and correctly formatted
   - Check if the repository is public or if you have access

3. **Build errors with Go versions**:
   - Ensure you're using Go 1.21 or later
   - Run `go mod tidy` to clean up dependencies

4. **Network connectivity issues**:
   - Verify internet connection for GitHub and Gemini API access
   - Check if corporate firewall blocks API access
