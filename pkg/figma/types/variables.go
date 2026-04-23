package types

type LocalVariablesResponse struct {
	Status int  `json:"status"`
	Error  bool `json:"error"`
	Meta   struct {
		Variables           map[string]Variable           `json:"variables"`
		VariableCollections map[string]VariableCollection `json:"variableCollections"`
	} `json:"meta"`
}

type PublishedVariablesResponse struct {
	Status int  `json:"status"`
	Error  bool `json:"error"`
	Meta   struct {
		Variables           map[string]PublishedVariable           `json:"variables"`
		VariableCollections map[string]PublishedVariableCollection `json:"variableCollections"`
	} `json:"meta"`
}

type Variable struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Key                  string                 `json:"key"`
	VariableCollectionID string                 `json:"variableCollectionId"`
	ResolvedType         string                 `json:"resolvedType"`
	ValuesByMode         map[string]interface{} `json:"valuesByMode"`
	Remote               bool                   `json:"remote"`
	Description          string                 `json:"description"`
	HiddenFromPublishing bool                   `json:"hiddenFromPublishing"`
	Scopes               []string               `json:"scopes"`
	CodeSyntax           VariableCodeSyntax     `json:"codeSyntax"`
}

type PublishedVariable struct {
	ID                   string `json:"id"`
	SubscribedID         string `json:"subscribed_id"`
	Name                 string `json:"name"`
	Key                  string `json:"key"`
	VariableCollectionID string `json:"variableCollectionId"`
	ResolvedType         string `json:"resolvedType"`
	UpdatedAt            string `json:"updatedAt"`
}

type VariableCollection struct {
	ID                   string              `json:"id"`
	Name                 string              `json:"name"`
	Key                  string              `json:"key"`
	Modes                []VariableMode      `json:"modes"`
	DefaultModeID        string              `json:"defaultModeId"`
	Remote               bool                `json:"remote"`
	HiddenFromPublishing bool                `json:"hiddenFromPublishing"`
	VariableIDs          []string            `json:"variableIds"`
}

type PublishedVariableCollection struct {
	ID           string `json:"id"`
	SubscribedID string `json:"subscribed_id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	UpdatedAt    string `json:"updatedAt"`
}

type VariableMode struct {
	ModeID string `json:"modeId"`
	Name   string `json:"name"`
}

type VariableCodeSyntax struct {
	WEB     string `json:"WEB,omitempty"`
	ANDROID string `json:"ANDROID,omitempty"`
	IOS     string `json:"iOS,omitempty"`
}

type PostVariablesRequest struct {
	VariableCollections []VariableCollectionChange `json:"variableCollections,omitempty"`
	VariableModes       []VariableModeChange       `json:"variableModes,omitempty"`
	Variables           []VariableChange           `json:"variables,omitempty"`
	VariableModeValues  []VariableModeValue        `json:"variableModeValues,omitempty"`
}

type VariableCollectionChange struct {
	Action              string `json:"action"`
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	InitialModeID       string `json:"initialModeId,omitempty"`
	HiddenFromPublishing *bool `json:"hiddenFromPublishing,omitempty"`
}

type VariableModeChange struct {
	Action               string `json:"action"`
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	VariableCollectionID string `json:"variableCollectionId,omitempty"`
}

type VariableChange struct {
	Action               string             `json:"action"`
	ID                   string             `json:"id,omitempty"`
	Name                 string             `json:"name,omitempty"`
	VariableCollectionID string             `json:"variableCollectionId,omitempty"`
	ResolvedType         string             `json:"resolvedType,omitempty"`
	Description          string             `json:"description,omitempty"`
	HiddenFromPublishing *bool              `json:"hiddenFromPublishing,omitempty"`
	Scopes               []string           `json:"scopes,omitempty"`
	CodeSyntax           *VariableCodeSyntax `json:"codeSyntax,omitempty"`
}

type VariableModeValue struct {
	VariableID           string      `json:"variableId"`
	ModeID               string      `json:"modeId"`
	Value                interface{} `json:"value"`
}

type PostVariablesResponse struct {
	Status int  `json:"status"`
	Error  bool `json:"error"`
	Meta   struct {
		TemporaryIDToRealID map[string]string `json:"temporaryIdToRealId"`
	} `json:"meta"`
}
