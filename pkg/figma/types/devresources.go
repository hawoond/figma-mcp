package types

type DevResource struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	FileKey   string `json:"file_key"`
	NodeID    string `json:"node_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DevResourcesResponse struct {
	DevResources []DevResource `json:"dev_resources"`
}

type CreateDevResourcesRequest struct {
	DevResources []CreateDevResource `json:"dev_resources"`
}

type CreateDevResource struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	FileKey string `json:"file_key"`
	NodeID  string `json:"node_id"`
}

type UpdateDevResourcesRequest struct {
	DevResources []UpdateDevResource `json:"dev_resources"`
}

type UpdateDevResource struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type CreateDevResourcesResponse struct {
	DevResources []DevResource `json:"dev_resources"`
	Errors       []DevResourceError `json:"errors,omitempty"`
}

type DevResourceError struct {
	FileKey string `json:"file_key"`
	NodeID  string `json:"node_id"`
	Error   string `json:"error"`
}

type ActivityLog struct {
	ID        string      `json:"id"`
	Timestamp string      `json:"timestamp"`
	Actor     ActivityActor `json:"actor"`
	Action    ActivityAction `json:"action"`
	Entity    ActivityEntity `json:"entity"`
	Context   interface{} `json:"context,omitempty"`
}

type ActivityActor struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type ActivityAction struct {
	Type string `json:"type"`
}

type ActivityEntity struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type ActivityLogsResponse struct {
	ActivityLogs []ActivityLog `json:"activity_logs"`
	Cursor       string        `json:"cursor,omitempty"`
	NextPage     bool          `json:"next_page"`
}

type LibraryAnalyticsComponentsResponse struct {
	Meta struct {
		Components []LibraryAnalyticsComponent `json:"components"`
		Cursor     string                      `json:"cursor,omitempty"`
		NextPage   bool                        `json:"next_page"`
	} `json:"meta"`
}

type LibraryAnalyticsComponent struct {
	Key              string `json:"key"`
	FileKey          string `json:"file_key"`
	NodeID           string `json:"node_id"`
	Name             string `json:"name"`
	Detachments      int    `json:"detachments"`
	Insertions       int    `json:"insertions"`
	TeamID           string `json:"team_id"`
	ComponentSetName string `json:"component_set_name,omitempty"`
}

type LibraryAnalyticsStylesResponse struct {
	Meta struct {
		Styles   []LibraryAnalyticsStyle `json:"styles"`
		Cursor   string                  `json:"cursor,omitempty"`
		NextPage bool                    `json:"next_page"`
	} `json:"meta"`
}

type LibraryAnalyticsStyle struct {
	Key       string `json:"key"`
	FileKey   string `json:"file_key"`
	NodeID    string `json:"node_id"`
	Name      string `json:"name"`
	StyleType string `json:"style_type"`
	Detachments int  `json:"detachments"`
	Insertions  int  `json:"insertions"`
	TeamID    string `json:"team_id"`
}

type LibraryAnalyticsVariablesResponse struct {
	Meta struct {
		Variables []LibraryAnalyticsVariable `json:"variables"`
		Cursor    string                     `json:"cursor,omitempty"`
		NextPage  bool                       `json:"next_page"`
	} `json:"meta"`
}

type LibraryAnalyticsVariable struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CollectionID string `json:"collection_id"`
	ResolvedType string `json:"resolved_type"`
	Usages       int    `json:"usages"`
}
