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
	var codeBlockContent []string
	var codeBlockLang string

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip navigation elements
		if trimmedLine == "" ||
			strings.HasPrefix(trimmedLine, "►") ||
			strings.HasPrefix(trimmedLine, "▼") ||
			strings.HasPrefix(trimmedLine, "→") {
			continue
		}

		// Handle code block start
		if strings.HasPrefix(trimmedLine, "```") {
			if !inCodeBlock {
				// Starting a new code block
				if len(currentSection) > 0 {
					contentLines = append(contentLines, strings.Join(currentSection, " "))
					currentSection = nil
				}
				inCodeBlock = true
				codeBlockLang = strings.TrimPrefix(trimmedLine, "```")
				continue
			} else {
				// Ending a code block
				inCodeBlock = false
				// Preserve the original formatting of code blocks
				if len(codeBlockContent) > 0 {
					if codeBlockLang != "" {
						contentLines = append(contentLines, "```"+codeBlockLang)
					} else {
						contentLines = append(contentLines, "```")
					}
					contentLines = append(contentLines, codeBlockContent...)
					contentLines = append(contentLines, "```")
					codeBlockContent = nil
				}
				continue
			}
		}

		// Handle content inside code blocks
		if inCodeBlock {
			// Preserve the original line including indentation
			codeBlockContent = append(codeBlockContent, line)
			continue
		}

		// Handle headers
		if strings.HasPrefix(trimmedLine, "#") {
			if len(currentSection) > 0 {
				contentLines = append(contentLines, strings.Join(currentSection, " "))
				currentSection = nil
			}
			contentLines = append(contentLines, trimmedLine)
			continue
		}

		// Handle lists
		if strings.HasPrefix(trimmedLine, "-") || strings.HasPrefix(trimmedLine, "*") || 
		   strings.HasPrefix(trimmedLine, "+") || regexp.MustCompile(`^\d+\.`).MatchString(trimmedLine) {
			if len(currentSection) > 0 {
				contentLines = append(contentLines, strings.Join(currentSection, " "))
				currentSection = nil
			}
			contentLines = append(contentLines, trimmedLine)
			continue
		}

		// Accumulate regular paragraph text
		currentSection = append(currentSection, trimmedLine)
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
