package extractor

import "encoding/json"

type StructuredData struct {
	Title   string `json:"title"`
  Header  string `json:"header"`
  Paragraph string `json:"paragraph"`
  List    []string `json:"list"`
	Content string `json:"content"`
}

func ToJSON(title, content, paragraph string, list []string, header string) ([]byte, error) {
	data := StructuredData{
		Title:   title,
		Content: content,
    Header:  header,
    Paragraph: paragraph,
    List:    []string{},

	}
	return json.Marshal(data)
}
