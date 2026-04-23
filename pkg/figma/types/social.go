package types

type User struct {
	ID     string `json:"id"`
	Handle string `json:"handle"`
	ImgURL string `json:"img_url"`
	Email  string `json:"email,omitempty"`
}

type Comment struct {
	ID          string       `json:"id"`
	UUID        string       `json:"uuid,omitempty"`
	FileKey     string       `json:"file_key"`
	ParentID    string       `json:"parent_id,omitempty"`
	User        User         `json:"user"`
	CreatedAt   string       `json:"created_at"`
	ResolvedAt  string       `json:"resolved_at,omitempty"`
	Message     string       `json:"message"`
	ClientMeta  *ClientMeta  `json:"client_meta,omitempty"`
	OrderID     string       `json:"order_id"`
	Reactions   []CommentReaction `json:"reactions,omitempty"`
}

type ClientMeta struct {
	X        *float64  `json:"x,omitempty"`
	Y        *float64  `json:"y,omitempty"`
	NodeID   []string  `json:"node_id,omitempty"`
	NodeOffset *Vector `json:"node_offset,omitempty"`
}

type CommentReaction struct {
	User      User   `json:"user"`
	CreatedAt string `json:"created_at"`
	Emoji     string `json:"emoji"`
}

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type PostCommentRequest struct {
	Message    string      `json:"message"`
	ClientMeta *ClientMeta `json:"client_meta,omitempty"`
	CommentID  string      `json:"comment_id,omitempty"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeamProjectsResponse struct {
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}

type ProjectFilesResponse struct {
	Name  string        `json:"name"`
	Files []ProjectFile `json:"files"`
}

type ProjectFile struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
	LastModified string `json:"last_modified"`
	Branches     []Branch `json:"branches,omitempty"`
}

type Webhook struct {
	ID          string `json:"id"`
	EventType   string `json:"event_type"`
	TeamID      string `json:"team_id"`
	ClientID    string `json:"client_id,omitempty"`
	Endpoint    string `json:"endpoint"`
	Passcode    string `json:"passcode"`
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
	ProtocolVersion string `json:"protocol_version"`
}

type WebhookRequest struct {
	EventType   string `json:"event_type"`
	TeamID      string `json:"team_id"`
	Endpoint    string `json:"endpoint"`
	Passcode    string `json:"passcode"`
	Description string `json:"description,omitempty"`
}

type WebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type WebhookPayloadsResponse struct {
	Payloads []WebhookPayload `json:"payloads"`
}

type WebhookPayload struct {
	WebhookID string      `json:"webhook_id"`
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
}

type TeamComponentsResponse struct {
	Meta struct {
		Components []TeamComponent `json:"components"`
		Cursor     string          `json:"cursor,omitempty"`
	} `json:"meta"`
}

type TeamComponent struct {
	Key            string `json:"key"`
	FileKey        string `json:"file_key"`
	NodeID         string `json:"node_id"`
	ThumbnailURL   string `json:"thumbnail_url"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	User           User   `json:"user"`
	ContainingFrame ContainingFrame `json:"containing_frame"`
	ContainingPage  ContainingPage  `json:"containing_page"`
}

type ContainingFrame struct {
	NodeID         string `json:"nodeId"`
	Name           string `json:"name"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	PageID         string `json:"pageId"`
	PageName       string `json:"pageName"`
}

type ContainingPage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeamComponentSetsResponse struct {
	Meta struct {
		ComponentSets []TeamComponentSet `json:"component_sets"`
		Cursor        string             `json:"cursor,omitempty"`
	} `json:"meta"`
}

type TeamComponentSet struct {
	Key          string `json:"key"`
	FileKey      string `json:"file_key"`
	NodeID       string `json:"node_id"`
	ThumbnailURL string `json:"thumbnail_url"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	User         User   `json:"user"`
}

type TeamStylesResponse struct {
	Meta struct {
		Styles []TeamStyle `json:"styles"`
		Cursor string      `json:"cursor,omitempty"`
	} `json:"meta"`
}

type TeamStyle struct {
	Key          string `json:"key"`
	FileKey      string `json:"file_key"`
	NodeID       string `json:"node_id"`
	ThumbnailURL string `json:"thumbnail_url"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	User         User   `json:"user"`
	StyleType    string `json:"style_type"`
	SortPosition string `json:"sort_position"`
}
