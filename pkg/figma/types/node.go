package types

type Node struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	Visible             *bool                  `json:"visible,omitempty"`
	Type                string                 `json:"type"`
	PluginData          interface{}            `json:"pluginData,omitempty"`
	SharedPluginData    interface{}            `json:"sharedPluginData,omitempty"`
	ComponentPropertyReferences map[string]string `json:"componentPropertyReferences,omitempty"`
	BoundVariables      map[string]interface{} `json:"boundVariables,omitempty"`
	ExplicitVariableModes map[string]string    `json:"explicitVariableModes,omitempty"`

	Children            []Node                 `json:"children,omitempty"`

	Locked              *bool                  `json:"locked,omitempty"`
	Fills               []Paint                `json:"fills,omitempty"`
	Strokes             []Paint                `json:"strokes,omitempty"`
	StrokeWeight        *float64               `json:"strokeWeight,omitempty"`
	StrokeAlign         string                 `json:"strokeAlign,omitempty"`
	StrokeDashes        []float64              `json:"strokeDashes,omitempty"`
	CornerRadius        *float64               `json:"cornerRadius,omitempty"`
	RectangleCornerRadii []float64             `json:"rectangleCornerRadii,omitempty"`
	CornerSmoothing     *float64               `json:"cornerSmoothing,omitempty"`
	Opacity             *float64               `json:"opacity,omitempty"`
	AbsoluteBoundingBox *Rectangle             `json:"absoluteBoundingBox,omitempty"`
	AbsoluteRenderBounds *Rectangle            `json:"absoluteRenderBounds,omitempty"`
	Constraints         *LayoutConstraint      `json:"constraints,omitempty"`
	BlendMode           string                 `json:"blendMode,omitempty"`
	IsMask              *bool                  `json:"isMask,omitempty"`
	Effects             []Effect               `json:"effects,omitempty"`
	ExportSettings      []ExportSetting        `json:"exportSettings,omitempty"`
	PreserveRatio       *bool                  `json:"preserveRatio,omitempty"`
	LayoutAlign         string                 `json:"layoutAlign,omitempty"`
	LayoutGrow          *float64               `json:"layoutGrow,omitempty"`
	LayoutPositioning   string                 `json:"layoutPositioning,omitempty"`
	RelativeTransform   *Transform             `json:"relativeTransform,omitempty"`
	Size                *Vector                `json:"size,omitempty"`
	MinWidth            *float64               `json:"minWidth,omitempty"`
	MaxWidth            *float64               `json:"maxWidth,omitempty"`
	MinHeight           *float64               `json:"minHeight,omitempty"`
	MaxHeight           *float64               `json:"maxHeight,omitempty"`

	BackgroundColor     *Color                 `json:"backgroundColor,omitempty"`
	PrototypeStartNodeID string               `json:"prototypeStartNodeID,omitempty"`
	FlowStartingPoints  []FlowStartingPoint    `json:"flowStartingPoints,omitempty"`
	PrototypeDevice     *PrototypeDevice       `json:"prototypeDevice,omitempty"`

	LayoutMode          string                 `json:"layoutMode,omitempty"`
	LayoutWrap          string                 `json:"layoutWrap,omitempty"`
	PrimaryAxisSizingMode string              `json:"primaryAxisSizingMode,omitempty"`
	CounterAxisSizingMode string             `json:"counterAxisSizingMode,omitempty"`
	PrimaryAxisAlignItems string             `json:"primaryAxisAlignItems,omitempty"`
	CounterAxisAlignItems string             `json:"counterAxisAlignItems,omitempty"`
	CounterAxisAlignContent string           `json:"counterAxisAlignContent,omitempty"`
	PaddingLeft         *float64               `json:"paddingLeft,omitempty"`
	PaddingRight        *float64               `json:"paddingRight,omitempty"`
	PaddingTop          *float64               `json:"paddingTop,omitempty"`
	PaddingBottom       *float64               `json:"paddingBottom,omitempty"`
	HorizontalPadding   *float64               `json:"horizontalPadding,omitempty"`
	VerticalPadding     *float64               `json:"verticalPadding,omitempty"`
	ItemSpacing         *float64               `json:"itemSpacing,omitempty"`
	CounterAxisSpacing  *float64               `json:"counterAxisSpacing,omitempty"`
	LayoutGrids         []LayoutGrid           `json:"layoutGrids,omitempty"`
	OverflowDirection   string                 `json:"overflowDirection,omitempty"`
	ItemReverseZIndex   *bool                  `json:"itemReverseZIndex,omitempty"`
	StrokesIncludedInLayout *bool             `json:"strokesIncludedInLayout,omitempty"`

	Characters          string                 `json:"characters,omitempty"`
	Style               *TypeStyle             `json:"style,omitempty"`
	CharacterStyleOverrides []int              `json:"characterStyleOverrides,omitempty"`
	StyleOverrideTable  map[string]TypeStyle   `json:"styleOverrideTable,omitempty"`
	LineTypes           []string               `json:"lineTypes,omitempty"`
	LineIndentations    []int                  `json:"lineIndentations,omitempty"`

	ComponentID         string                 `json:"componentId,omitempty"`
	IsExposedInstance   *bool                  `json:"isExposedInstance,omitempty"`
	ExposedInstances    []string               `json:"exposedInstances,omitempty"`
	ComponentProperties map[string]ComponentProperty `json:"componentProperties,omitempty"`
	Overrides           []Override             `json:"overrides,omitempty"`

	ComponentSetID      string                 `json:"componentSetId,omitempty"`
	Description         string                 `json:"description,omitempty"`
	DocumentationLinks  []DocumentationLink    `json:"documentationLinks,omitempty"`
	Remote              *bool                  `json:"remote,omitempty"`
	Key                 string                 `json:"key,omitempty"`
	Annotations         []Annotation           `json:"annotations,omitempty"`
	Measurements        []Measurement          `json:"measurements,omitempty"`

	Reactions           []Reaction             `json:"reactions,omitempty"`

	ScrollBehavior      string                 `json:"scrollBehavior,omitempty"`
}

type LayoutConstraint struct {
	Vertical   string `json:"vertical"`
	Horizontal string `json:"horizontal"`
}

type ComponentProperty struct {
	Type         string      `json:"type"`
	Value        interface{} `json:"value"`
	PreferredValues []ComponentPropertyPreferredValue `json:"preferredValues,omitempty"`
	BoundVariables map[string]VariableAlias `json:"boundVariables,omitempty"`
}

type Override struct {
	ID             string   `json:"id"`
	OverriddenFields []string `json:"overriddenFields"`
}

type Reaction struct {
	Action  *Action  `json:"action,omitempty"`
	Trigger *Trigger `json:"trigger,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Action struct {
	Type                string   `json:"type"`
	URL                 string   `json:"url,omitempty"`
	DestinationID       string   `json:"destinationID,omitempty"`
	Navigation          string   `json:"navigation,omitempty"`
	Transition          *Transition `json:"transition,omitempty"`
	PreserveScrollPosition *bool `json:"preserveScrollPosition,omitempty"`
	OverlayRelativePosition *Vector `json:"overlayRelativePosition,omitempty"`
	ResetVideoPosition  *bool    `json:"resetVideoPosition,omitempty"`
	ResetScrollPosition *bool    `json:"resetScrollPosition,omitempty"`
	ResetInteractiveComponents *bool `json:"resetInteractiveComponents,omitempty"`
	MediaAction         string   `json:"mediaAction,omitempty"`
	VariableID          string   `json:"variableId,omitempty"`
	VariableModeID      string   `json:"variableModeId,omitempty"`
}

type Transition struct {
	Type      string   `json:"type"`
	Easing    Easing   `json:"easing"`
	Duration  float64  `json:"duration"`
	Direction string   `json:"direction,omitempty"`
	MatchLayers *bool  `json:"matchLayers,omitempty"`
}

type Easing struct {
	Type           string    `json:"type"`
	EasingFunctionCubicBezier *CubicBezier `json:"easingFunctionCubicBezier,omitempty"`
	EasingFunctionSpring *Spring `json:"easingFunctionSpring,omitempty"`
}

type CubicBezier struct {
	X1 float64 `json:"x1"`
	Y1 float64 `json:"y1"`
	X2 float64 `json:"x2"`
	Y2 float64 `json:"y2"`
}

type Spring struct {
	Mass        float64 `json:"mass"`
	Stiffness   float64 `json:"stiffness"`
	Damping     float64 `json:"damping"`
}

type Trigger struct {
	Type    string   `json:"type"`
	Delay   *float64 `json:"delay,omitempty"`
	Timeout *float64 `json:"timeout,omitempty"`
	KeyCodes []int   `json:"keyCodes,omitempty"`
	MediaHitTime *float64 `json:"mediaHitTime,omitempty"`
	Interactions []Interaction `json:"interactions,omitempty"`
}

type Interaction struct {
	Action  Action  `json:"action"`
	Trigger Trigger `json:"trigger"`
}
