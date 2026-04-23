package types

type UploadImageResponse struct {
	ImageRef string `json:"imageRef"`
	Error    string `json:"error,omitempty"`
}

type SetNodeFillRequest struct {
	NodeID    string  `json:"nodeId"`
	ImageRef  string  `json:"imageRef"`
	ScaleMode string  `json:"scaleMode"`
}

type UploadImageFromURLResult struct {
	ImageRef  string `json:"image_ref"`
	Format    string `json:"format"`
	SizeBytes int    `json:"size_bytes"`
}
