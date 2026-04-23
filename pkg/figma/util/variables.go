package util

import (
	"fmt"
	"strings"

	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type DesignToken struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Value        interface{} `json:"value"`
	CollectionName string    `json:"collection"`
	ModeName     string      `json:"mode"`
	Description  string      `json:"description,omitempty"`
	Scopes       []string    `json:"scopes,omitempty"`
}

func ExtractDesignTokens(resp *types.LocalVariablesResponse) []DesignToken {
	var tokens []DesignToken

	collectionNames := make(map[string]string)
	for id, col := range resp.Meta.VariableCollections {
		collectionNames[id] = col.Name
		_ = id
	}

	modeNames := make(map[string]map[string]string)
	for _, col := range resp.Meta.VariableCollections {
		modeNames[col.ID] = make(map[string]string)
		for _, mode := range col.Modes {
			modeNames[col.ID][mode.ModeID] = mode.Name
		}
	}

	for _, variable := range resp.Meta.Variables {
		colName := collectionNames[variable.VariableCollectionID]
		col, hasCol := resp.Meta.VariableCollections[variable.VariableCollectionID]

		for modeID, value := range variable.ValuesByMode {
			modeName := ""
			if hasCol {
				modeName = modeNames[col.ID][modeID]
			}

			tokens = append(tokens, DesignToken{
				Name:           variable.Name,
				Type:           variable.ResolvedType,
				Value:          value,
				CollectionName: colName,
				ModeName:       modeName,
				Description:    variable.Description,
				Scopes:         variable.Scopes,
			})
		}
	}

	return tokens
}

func FormatColorValue(v interface{}) string {
	colorMap, ok := v.(map[string]interface{})
	if !ok {
		return fmt.Sprintf("%v", v)
	}

	r, _ := colorMap["r"].(float64)
	g, _ := colorMap["g"].(float64)
	b, _ := colorMap["b"].(float64)
	a, _ := colorMap["a"].(float64)

	rInt := int(r * 255)
	gInt := int(g * 255)
	bInt := int(b * 255)

	if a == 1.0 || a == 0 {
		return fmt.Sprintf("#%02X%02X%02X", rInt, gInt, bInt)
	}
	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", rInt, gInt, bInt, a)
}

func TokensToCSSVariables(tokens []DesignToken, modeFilter string) string {
	var sb strings.Builder
	sb.WriteString(":root {\n")

	for _, token := range tokens {
		if modeFilter != "" && token.ModeName != modeFilter {
			continue
		}

		varName := "--" + strings.ReplaceAll(strings.ToLower(token.Name), " ", "-")
		varName = strings.ReplaceAll(varName, "/", "-")

		var valueStr string
		switch token.Type {
		case "COLOR":
			valueStr = FormatColorValue(token.Value)
		case "FLOAT":
			valueStr = fmt.Sprintf("%v", token.Value)
		case "STRING":
			valueStr = fmt.Sprintf("%v", token.Value)
		case "BOOLEAN":
			valueStr = fmt.Sprintf("%v", token.Value)
		default:
			valueStr = fmt.Sprintf("%v", token.Value)
		}

		sb.WriteString(fmt.Sprintf("  %s: %s;\n", varName, valueStr))
	}

	sb.WriteString("}\n")
	return sb.String()
}

func TokensToJSON(tokens []DesignToken) map[string]interface{} {
	result := make(map[string]interface{})

	for _, token := range tokens {
		parts := strings.Split(token.Name, "/")
		current := result

		for i, part := range parts {
			if i == len(parts)-1 {
				current[part] = map[string]interface{}{
					"$value":       token.Value,
					"$type":        strings.ToLower(token.Type),
					"$description": token.Description,
				}
			} else {
				if _, exists := current[part]; !exists {
					current[part] = make(map[string]interface{})
				}
				if nested, ok := current[part].(map[string]interface{}); ok {
					current = nested
				}
			}
		}
	}

	return result
}
