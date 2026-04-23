package types

type File struct {
	Name         string                 `json:"name"`
	Role         string                 `json:"role"`
	LastModified string                 `json:"lastModified"`
	EditorType   string                 `json:"editorType"`
	ThumbnailURL string                 `json:"thumbnailUrl"`
	Version      string                 `json:"version"`
	Document     Node                   `json:"document"`
	Components   map[string]Component   `json:"components"`
	ComponentSets map[string]ComponentSet `json:"componentSets"`
	SchemaVersion int                   `json:"schemaVersion"`
	Styles       map[string]Style       `json:"styles"`
	MainFileKey  string                 `json:"mainFileKey,omitempty"`
	Branches     []Branch               `json:"branches,omitempty"`
}

type FileNodes struct {
	Name         string                 `json:"name"`
	Role         string                 `json:"role"`
	LastModified string                 `json:"lastModified"`
	EditorType   string                 `json:"editorType"`
	ThumbnailURL string                 `json:"thumbnailUrl"`
	Err          string                 `json:"err,omitempty"`
	Nodes        map[string]FileNode    `json:"nodes"`
}

type FileNode struct {
	Document      Node                   `json:"document"`
	Components    map[string]Component   `json:"components"`
	ComponentSets map[string]ComponentSet `json:"componentSets"`
	SchemaVersion int                    `json:"schemaVersion"`
	Styles        map[string]Style       `json:"styles"`
}

type Branch struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
	LastModified string `json:"last_modified"`
	LinkAccess   string `json:"link_access"`
}

type Component struct {
	Key            string `json:"key"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ComponentSetID string `json:"componentSetId,omitempty"`
	DocumentationLinks []DocumentationLink `json:"documentationLinks,omitempty"`
	Remote         bool   `json:"remote"`
}

type ComponentSet struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DocumentationLinks []DocumentationLink `json:"documentationLinks,omitempty"`
	Remote      bool   `json:"remote"`
}

type Style struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Remote      bool   `json:"remote"`
	StyleType   string `json:"styleType"`
}

type FileMetadata struct {
	Name         string   `json:"name"`
	Role         string   `json:"role"`
	LastModified string   `json:"lastModified"`
	EditorType   string   `json:"editorType"`
	ThumbnailURL string   `json:"thumbnailUrl"`
	Version      string   `json:"version"`
	Branches     []Branch `json:"branches,omitempty"`
}

type ImageResponse struct {
	Err    string            `json:"err,omitempty"`
	Images map[string]string `json:"images"`
}

type ImageFillsResponse struct {
	Meta struct {
		Images map[string]string `json:"images"`
	} `json:"meta"`
}

type FileVersion struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Label       string `json:"label"`
	Description string `json:"description"`
	User        User   `json:"user"`
}

type FileVersionsResponse struct {
	Versions   []FileVersion `json:"versions"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

type Pagination struct {
	Cursor string `json:"cursor,omitempty"`
}
