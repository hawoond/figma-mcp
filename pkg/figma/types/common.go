package types

type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Transform [2][3]float64

type ExportSetting struct {
	Suffix     string     `json:"suffix"`
	Format     string     `json:"format"`
	Constraint Constraint `json:"constraint"`
}

type Constraint struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type Paint struct {
	Type               string             `json:"type"`
	Visible            *bool              `json:"visible,omitempty"`
	Opacity            *float64           `json:"opacity,omitempty"`
	Color              *Color             `json:"color,omitempty"`
	BlendMode          string             `json:"blendMode,omitempty"`
	GradientHandlePositions []Vector      `json:"gradientHandlePositions,omitempty"`
	GradientStops      []ColorStop        `json:"gradientStops,omitempty"`
	ScaleMode          string             `json:"scaleMode,omitempty"`
	ImageTransform     *Transform         `json:"imageTransform,omitempty"`
	Scaling            *float64           `json:"scaling,omitempty"`
	Rotation           *float64           `json:"rotation,omitempty"`
	ImageRef           string             `json:"imageRef,omitempty"`
	Filters            *ImageFilters      `json:"filters,omitempty"`
	GifRef             string             `json:"gifRef,omitempty"`
	BoundVariables     map[string]interface{} `json:"boundVariables,omitempty"`
}

type ColorStop struct {
	Position float64 `json:"position"`
	Color    Color   `json:"color"`
}

type ImageFilters struct {
	Exposure    *float64 `json:"exposure,omitempty"`
	Contrast    *float64 `json:"contrast,omitempty"`
	Saturation  *float64 `json:"saturation,omitempty"`
	Temperature *float64 `json:"temperature,omitempty"`
	Tint        *float64 `json:"tint,omitempty"`
	Highlights  *float64 `json:"highlights,omitempty"`
	Shadows     *float64 `json:"shadows,omitempty"`
}

type Effect struct {
	Type      string   `json:"type"`
	Visible   bool     `json:"visible"`
	Radius    float64  `json:"radius"`
	Color     *Color   `json:"color,omitempty"`
	BlendMode string   `json:"blendMode,omitempty"`
	Offset    *Vector  `json:"offset,omitempty"`
	Spread    *float64 `json:"spread,omitempty"`
	ShowShadowBehindNode *bool `json:"showShadowBehindNode,omitempty"`
}

type TypeStyle struct {
	FontFamily          string   `json:"fontFamily"`
	FontPostScriptName  string   `json:"fontPostScriptName"`
	ParagraphSpacing    float64  `json:"paragraphSpacing,omitempty"`
	ParagraphIndent     float64  `json:"paragraphIndent,omitempty"`
	ListSpacing         float64  `json:"listSpacing,omitempty"`
	Italic              bool     `json:"italic"`
	FontWeight          float64  `json:"fontWeight"`
	FontSize            float64  `json:"fontSize"`
	TextCase            string   `json:"textCase,omitempty"`
	TextDecoration      string   `json:"textDecoration,omitempty"`
	TextAutoResize      string   `json:"textAutoResize,omitempty"`
	TextTruncation      string   `json:"textTruncation,omitempty"`
	MaxLines            *int     `json:"maxLines,omitempty"`
	TextAlignHorizontal string   `json:"textAlignHorizontal"`
	TextAlignVertical   string   `json:"textAlignVertical"`
	LetterSpacing       float64  `json:"letterSpacing"`
	Fills               []Paint  `json:"fills,omitempty"`
	Hyperlink           *Hyperlink `json:"hyperlink,omitempty"`
	OpentypeFlags       map[string]int `json:"opentypeFlags,omitempty"`
	LineHeightPx        float64  `json:"lineHeightPx"`
	LineHeightPercent   float64  `json:"lineHeightPercent,omitempty"`
	LineHeightPercentFontSize float64 `json:"lineHeightPercentFontSize,omitempty"`
	LineHeightUnit      string   `json:"lineHeightUnit"`
}

type Hyperlink struct {
	Type  string `json:"type"`
	URL   string `json:"url,omitempty"`
	NodeID string `json:"nodeID,omitempty"`
}

type LayoutGrid struct {
	Pattern     string  `json:"pattern"`
	SectionSize float64 `json:"sectionSize"`
	Visible     bool    `json:"visible"`
	Color       Color   `json:"color"`
	Alignment   string  `json:"alignment"`
	GutterSize  float64 `json:"gutterSize"`
	Offset      float64 `json:"offset"`
	Count       int     `json:"count"`
}

type PrototypeDevice struct {
	Type       string     `json:"type"`
	Size       *Vector    `json:"size,omitempty"`
	PresetIdentifier string `json:"presetIdentifier,omitempty"`
	Rotation   string     `json:"rotation"`
}

type FlowStartingPoint struct {
	NodeID string `json:"nodeID"`
	Name   string `json:"name"`
}

type DocumentationLink struct {
	URI string `json:"uri"`
}

type VariableAlias struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type ComponentPropertyDefinition struct {
	Type         string      `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	VariantOptions []string  `json:"variantOptions,omitempty"`
	PreferredValues []ComponentPropertyPreferredValue `json:"preferredValues,omitempty"`
}

type ComponentPropertyPreferredValue struct {
	Type string `json:"type"`
	Key  string `json:"key"`
}

type Annotation struct {
	Label      string              `json:"label"`
	Properties []AnnotationProperty `json:"properties"`
}

type AnnotationProperty struct {
	Type string `json:"type"`
}

type Measurement struct {
	ID       string          `json:"id"`
	Start    MeasurementAxis `json:"start"`
	End      MeasurementAxis `json:"end"`
	Offset   MeasurementOffset `json:"offset"`
}

type MeasurementAxis struct {
	NodeID string  `json:"nodeId"`
	Side   string  `json:"side"`
}

type MeasurementOffset struct {
	Type   string  `json:"type"`
	Fixed  *float64 `json:"fixed,omitempty"`
}
