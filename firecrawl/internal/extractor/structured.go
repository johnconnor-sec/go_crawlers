package extractor

import "encoding/json"

type StructuredData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToJSON(title, content string) ([]byte, error) {
	data := StructuredData{
		Title:   title,
		Content: content,
	}
	return json.Marshal(data)
}