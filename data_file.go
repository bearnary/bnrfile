package bnrfile

type FileData struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	FullPath    string `json:"full_path"`
	Size        int64  `json:"size"`
	ContentType string `json:"content_type"`
}
