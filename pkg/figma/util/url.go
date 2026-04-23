package util

import (
	"fmt"
	"net/url"
	"strings"
)

type FigmaURL struct {
	FileKey string
	NodeID  string
	Branch  string
}

func ParseFigmaURL(rawURL string) (*FigmaURL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if !strings.Contains(u.Host, "figma.com") {
		return nil, fmt.Errorf("not a figma URL")
	}

	parts := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid figma URL path: %s", u.Path)
	}

	result := &FigmaURL{}

	switch parts[0] {
	case "file", "design", "proto", "board":
		if len(parts) < 2 {
			return nil, fmt.Errorf("missing file key in URL")
		}
		result.FileKey = parts[1]
	default:
		return nil, fmt.Errorf("unknown figma URL type: %s", parts[0])
	}

	nodeID := u.Query().Get("node-id")
	if nodeID != "" {
		result.NodeID = strings.ReplaceAll(nodeID, "-", ":")
	}

	return result, nil
}

func BuildFigmaFileURL(fileKey string) string {
	return fmt.Sprintf("https://www.figma.com/file/%s", fileKey)
}

func BuildFigmaNodeURL(fileKey, nodeID string) string {
	encodedNodeID := strings.ReplaceAll(nodeID, ":", "-")
	return fmt.Sprintf("https://www.figma.com/file/%s?node-id=%s", fileKey, encodedNodeID)
}

func ExtractFileKeyFromURL(rawURL string) (string, error) {
	parsed, err := ParseFigmaURL(rawURL)
	if err != nil {
		return "", err
	}
	return parsed.FileKey, nil
}
