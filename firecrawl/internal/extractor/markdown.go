package extractor

import (
	"fmt"
  "regexp"
  "strings"
)

func ToMarkdown(title, content string) string {
	// First remove all JavaScript
	jsPattern := `<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>`
	re := regexp.MustCompile(jsPattern)
	cleanContent := re.ReplaceAllString(content, "")

	// Remove common navigation elements and metadata
	removePatterns := []string{
		`Skip to content`,
		`Table of contents`,
		`Navigation`,
		`Search`,
		`Copyright`,
		`©.*\d{4}`,
		`var\s+\w+\s*=.*?;`,
		`function\s*\(.*?\)`,
	}

	for _, pattern := range removePatterns {
		re := regexp.MustCompile(`(?i)` + pattern)
		cleanContent = re.ReplaceAllString(cleanContent, "")
	}

	// Split content and process line by line
	lines := strings.Split(cleanContent, "\n")
	var contentLines []string
	var inCodeBlock bool
	var currentSection []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip navigation elements
		if line == "" ||
			strings.HasPrefix(line, "►") ||
			strings.HasPrefix(line, "▼") ||
			strings.HasPrefix(line, "→") {
			continue
		}

		// Handle code blocks
		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
			contentLines = append(contentLines, line)
			continue
		}

		// Preserve formatting inside code blocks
		if inCodeBlock {
			contentLines = append(contentLines, line)
			continue
		}

		// Handle headers
		if strings.HasPrefix(line, "#") {
			if len(currentSection) > 0 {
				contentLines = append(contentLines, strings.Join(currentSection, " "))
				currentSection = nil
			}
			contentLines = append(contentLines, line)
			continue
		}

		// Handle lists
		if strings.HasPrefix(line, "-") || strings.HasPrefix(line, "*") || 
		   strings.HasPrefix(line, "+") || regexp.MustCompile(`^\d+\.`).MatchString(line) {
			if len(currentSection) > 0 {
				contentLines = append(contentLines, strings.Join(currentSection, " "))
				currentSection = nil
			}
			contentLines = append(contentLines, line)
			continue
		}

		// Accumulate regular paragraph text
		currentSection = append(currentSection, line)
	}

	// Add any remaining section
	if len(currentSection) > 0 {
		contentLines = append(contentLines, strings.Join(currentSection, " "))
	}

	// Join all lines with proper spacing
	finalContent := strings.Join(contentLines, "\n\n")

	return fmt.Sprintf("# %s\n\n%s",
		strings.TrimSpace(title),
		strings.TrimSpace(finalContent))
}
