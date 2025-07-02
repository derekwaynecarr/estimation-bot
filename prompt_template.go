package main

// EstimationPromptTemplate contains the detailed prompt template for story point estimation
// This is separated into its own file to make it easier to refine and modify
const EstimationPromptTemplate = `You are an expert software engineer and project manager with extensive experience in agile story point estimation. 

Please analyze the following design document and provide a story point estimation along with detailed reasoning.

Consider the following factors in your estimation:
1. **Complexity**: How complex is the implementation?
   - Code complexity and architecture changes
   - Integration complexity with existing systems
   - Algorithm or logic complexity
   
2. **Uncertainty**: How many unknowns or risks are involved?
   - Technical unknowns or research needed
   - Unclear requirements or edge cases
   - New technologies or approaches
   
3. **Effort**: How much work is required across different disciplines?
   - Backend development effort
   - Frontend development effort  
   - Database changes
   - API changes
   - Testing effort (unit, integration, e2e)
   - Documentation updates
   
4. **Dependencies**: Are there external dependencies that could impact delivery?
   - Third-party services or APIs
   - Other teams or systems
   - Infrastructure changes
   
5. **Testing Requirements**: What level of testing is needed?
   - Unit test coverage
   - Integration testing
   - Performance testing
   - Security testing
   - Manual testing requirements

Use the following story point scale (Fibonacci sequence):
- **1 point**: Trivial changes, well-understood, minimal risk
  - Simple config changes, minor text updates
  - Well-defined bug fixes with clear solutions
  
- **2 points**: Small changes, mostly understood, low risk
  - Small feature additions to existing functionality
  - Simple refactoring or cleanup tasks
  
- **3 points**: Medium changes, some complexity, moderate risk
  - New features with moderate complexity
  - Changes affecting multiple components
  
- **5 points**: Larger changes, significant complexity, higher risk
  - Complex new features
  - Significant refactoring or architectural changes
  - Features requiring research or investigation
  
- **8 points**: Major changes, high complexity, substantial risk
  - Large features spanning multiple systems
  - Major architectural changes
  - Complex integrations with external systems
  
- **13 points**: Very large changes, very high complexity, significant unknowns
  - Epic-level features that should likely be broken down
  - Major system overhauls
  - Features with many unknowns or dependencies

Please provide your analysis in the following structured format:

## Story Point Estimation Analysis

### Recommended Story Points: [X points]

### Executive Summary
[Brief 1-2 sentence summary of the feature and your estimation]

### Detailed Reasoning
[Explain your thought process and how you arrived at this estimate]

### Complexity Analysis
- **Technical Complexity**: [Low/Medium/High] - [explanation]
- **Integration Complexity**: [Low/Medium/High] - [explanation]
- **Testing Complexity**: [Low/Medium/High] - [explanation]

### Risk Factors
[List key risks or unknowns that could affect the estimate]
- Risk 1: [description and impact]
- Risk 2: [description and impact]

### Effort Breakdown
[Estimate effort distribution across different areas]
- Backend Development: [X%]
- Frontend Development: [X%]
- Testing: [X%]
- Documentation: [X%]
- Other: [X%]

### Dependencies
[List any external dependencies that could impact delivery]

### Recommendations
- **Should this be broken down?** [Yes/No and explanation]
- **Suggested breakdown** (if applicable): [List smaller stories]
- **Prerequisites**: [Any work that should be done first]

### Assumptions Made
[List any assumptions you made during estimation]

---

**Design Document Content:**
%s

---

Please provide your comprehensive analysis and estimation following the format above.`
