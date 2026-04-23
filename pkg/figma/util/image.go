package util

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type ImageUploadResult struct {
	ImageRef  string
	Format    string
	SizeBytes int
}

func FetchImageFromURL(ctx context.Context, imageURL string) ([]byte, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil)
	if err != nil {
		return nil, "", fmt.Errorf("create request: %w", err)
	}

	httpClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("fetch image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("read image data: %w", err)
	}

	contentType := resp.Header.Get("Content-Type")
	format := detectImageFormat(contentType, imageURL, data)

	return data, format, nil
}

func detectImageFormat(contentType, url string, data []byte) string {
	if strings.Contains(contentType, "png") {
		return "png"
	}
	if strings.Contains(contentType, "jpeg") || strings.Contains(contentType, "jpg") {
		return "jpg"
	}
	if strings.Contains(contentType, "gif") {
		return "gif"
	}
	if strings.Contains(contentType, "webp") {
		return "webp"
	}

	urlLower := strings.ToLower(url)
	if strings.HasSuffix(urlLower, ".png") {
		return "png"
	}
	if strings.HasSuffix(urlLower, ".jpg") || strings.HasSuffix(urlLower, ".jpeg") {
		return "jpg"
	}
	if strings.HasSuffix(urlLower, ".gif") {
		return "gif"
	}
	if strings.HasSuffix(urlLower, ".webp") {
		return "webp"
	}

	if len(data) >= 4 {
		if bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47}) {
			return "png"
		}
		if bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}) {
			return "jpg"
		}
		if bytes.HasPrefix(data, []byte{0x47, 0x49, 0x46}) {
			return "gif"
		}
	}

	return "png"
}

func FormatToMIMEType(format string) string {
	switch strings.ToLower(format) {
	case "jpg", "jpeg":
		return "image/jpeg"
	case "gif":
		return "image/gif"
	case "webp":
		return "image/webp"
	default:
		return "image/png"
	}
}

func ImageToBase64(data []byte, format string) string {
	mimeType := FormatToMIMEType(format)
	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, encoded)
}

func BuildImageFillPaint(imageRef string) types.Paint {
	return types.Paint{
		Type:      "IMAGE",
		ScaleMode: "FILL",
		ImageRef:  imageRef,
	}
}

func BuildImageFillPaintWithMode(imageRef, scaleMode string) types.Paint {
	return types.Paint{
		Type:      "IMAGE",
		ScaleMode: scaleMode,
		ImageRef:  imageRef,
	}
}

type DownloadedAsset struct {
	NodeID   string
	URL      string
	Data     []byte
	Format   string
	FileName string
}

func DownloadAssets(ctx context.Context, assets map[string]string) ([]DownloadedAsset, []error) {
	var results []DownloadedAsset
	var errs []error

	for nodeID, assetURL := range assets {
		data, format, err := FetchImageFromURL(ctx, assetURL)
		if err != nil {
			errs = append(errs, fmt.Errorf("node %s: %w", nodeID, err))
			continue
		}
		results = append(results, DownloadedAsset{
			NodeID:   nodeID,
			URL:      assetURL,
			Data:     data,
			Format:   format,
			FileName: fmt.Sprintf("%s.%s", sanitizeFilename(nodeID), format),
		})
	}

	return results, errs
}

func sanitizeFilename(name string) string {
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		" ", "_",
	)
	return replacer.Replace(name)
}
