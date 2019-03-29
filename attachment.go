package garoon

type Attachment struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	ContentType string `json:"contentType"`
	Size        string `json:"size"`
}
