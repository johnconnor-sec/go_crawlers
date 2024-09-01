package extractor

import (
	"fmt"
	"strings"
)

func ToMarkdown(title, content string) string {
	return fmt.Sprintf("# %s\n\n%s", title, strings.TrimSpace(content))
}