package extractor

import (
	"fmt"
  "regexp"
  "strings"
)

func ToMarkdown(title, content string) string {
  patterns := []string{
    `Skip to content`,
    `Table of contents`,
    `Navigation`,
    `Search`,
    `Copyright`,
    `©.*\d{4}`,
    `var \s+\w+\s*=.*?;`,  // Remove JavaScript
    `function \s*\(.*?\)`, // Remove JavaScript functions
  }

  cleanContent := content
  for _, pattern := range patterns {
    re := regexp.MustCompile(`(?i)` + pattern)
    cleanContent = re.ReplaceAllString(cleanContent, "")
  }

  lines := strings.Split(cleanContent, "\n")
  var contentLines []string

  for _, line := range lines {
    line = strings.TrimSpace(line)
    if line == "" ||
      strings.HasPrefix(line, "►") ||
      strings.HasPrefix(line, "▼") ||
      strings.HasPrefix(line, "→") {
      continue
    }
    contentLines = append(contentLines, line)
  }

  finalContent := strings.Join(contentLines, " ")
  // finalContent = strings.Join(strings.Fields(finalContent), " ")

  return fmt.Sprintf("# %s\n\n%s",
    strings.TrimSpace(title),
    finalContent,
    )
}
