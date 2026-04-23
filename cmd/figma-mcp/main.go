package main

import (
	"fmt"
	"os"

	"github.com/hawoond/figma-mcp/pkg/figma"
	"github.com/hawoond/figma-mcp/pkg/figma/util"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	figma  *figma.Client
	editor *util.Editor
	mcp    *server.MCPServer
}

func NewServer(token string) *Server {
	figmaClient := figma.New(token)
	editor := util.NewEditor(figmaClient.Files, figmaClient.Variables)

	mcpServer := server.NewMCPServer(
		"figma-mcp",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	s := &Server{
		figma:  figmaClient,
		editor: editor,
		mcp:    mcpServer,
	}

	s.registerTools()
	return s
}

func (s *Server) registerTools() {
	s.mcp.AddTool(
		mcp.NewTool("figma_get_file",
			mcp.WithDescription("Get a Figma file by key. Returns the full document tree including all nodes, styles, and components."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key (from the URL)")),
			mcp.WithNumber("depth", mcp.Description("How deep to traverse the document tree")),
			mcp.WithString("ids", mcp.Description("Comma-separated list of node IDs to retrieve")),
			mcp.WithString("geometry", mcp.Description("Set to 'paths' to export vector data")),
		),
		s.handleGetFile,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_nodes",
			mcp.WithDescription("Get specific nodes from a Figma file by their IDs."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("node_ids", mcp.Required(), mcp.Description("Comma-separated list of node IDs to retrieve")),
			mcp.WithNumber("depth", mcp.Description("How deep to traverse each node's subtree")),
		),
		s.handleGetFileNodes,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_images",
			mcp.WithDescription("Export nodes from a Figma file as images (PNG, JPG, SVG, PDF)."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("node_ids", mcp.Required(), mcp.Description("Comma-separated list of node IDs to export")),
			mcp.WithString("format", mcp.Description("Image format: png, jpg, svg, pdf (default: png)")),
			mcp.WithNumber("scale", mcp.Description("Scale factor (0.01 to 4, default: 1)")),
			mcp.WithBoolean("use_absolute_bounds", mcp.Description("Use absolute bounding box for export")),
		),
		s.handleGetImages,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_image_fills",
			mcp.WithDescription("Get URLs for all image fills in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetImageFills,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_versions",
			mcp.WithDescription("Get version history of a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetFileVersions,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_metadata",
			mcp.WithDescription("Get metadata of a Figma file (name, last modified, thumbnail URL, etc.)."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetFileMetadata,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_search_nodes",
			mcp.WithDescription("Search for nodes in a Figma file by name (case-insensitive partial match)."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("name", mcp.Required(), mcp.Description("Node name to search for")),
		),
		s.handleSearchNodes,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_nodes_by_type",
			mcp.WithDescription("Get all nodes of a specific type from a Figma file (e.g., TEXT, FRAME, COMPONENT, INSTANCE, RECTANGLE)."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("node_type", mcp.Required(), mcp.Description("Node type: DOCUMENT, CANVAS, FRAME, GROUP, VECTOR, BOOLEAN_OPERATION, STAR, LINE, ELLIPSE, REGULAR_POLYGON, RECTANGLE, TABLE, TABLE_CELL, TEXT, SLICE, COMPONENT, COMPONENT_SET, INSTANCE, STICKY, SHAPE_WITH_TEXT, CONNECTOR, SECTION")),
		),
		s.handleGetNodesByType,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_structure",
			mcp.WithDescription("Get a structural overview of a Figma file up to a specified depth."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithNumber("max_depth", mcp.Description("Maximum depth to traverse (default: 3)")),
		),
		s.handleGetFileStructure,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_export_frames",
			mcp.WithDescription("Export all frames and components from a Figma file as images."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("format", mcp.Description("Image format: png, jpg, svg, pdf (default: png)")),
			mcp.WithNumber("scale", mcp.Description("Scale factor (default: 1)")),
		),
		s.handleExportFrames,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_export_node_as_image",
			mcp.WithDescription("Export a specific node from a Figma file as an image and return the download URL."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("node_id", mcp.Required(), mcp.Description("The node ID to export")),
			mcp.WithString("format", mcp.Description("Image format: png, jpg, svg, pdf (default: png)")),
			mcp.WithNumber("scale", mcp.Description("Scale factor (default: 1)")),
		),
		s.handleExportNodeAsImage,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_fetch_image_from_url",
			mcp.WithDescription("Fetch an image from a URL and return it as base64-encoded data. Useful for uploading external images to Figma."),
			mcp.WithString("url", mcp.Required(), mcp.Description("The URL of the image to fetch")),
		),
		s.handleFetchImageFromURL,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_search_text",
			mcp.WithDescription("Search for text content within all text nodes in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("text", mcp.Required(), mcp.Description("Text string to search for")),
		),
		s.handleSearchTextInFile,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_parse_url",
			mcp.WithDescription("Parse a Figma URL to extract the file key and node ID."),
			mcp.WithString("url", mcp.Required(), mcp.Description("The Figma URL to parse")),
		),
		s.handleParseFigmaURL,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_comments",
			mcp.WithDescription("Get all comments on a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithBoolean("as_markdown", mcp.Description("Return comments in Markdown format")),
		),
		s.handleGetComments,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_post_comment",
			mcp.WithDescription("Post a comment on a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("message", mcp.Required(), mcp.Description("The comment message")),
			mcp.WithString("comment_id", mcp.Description("Parent comment ID to reply to")),
		),
		s.handlePostComment,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_delete_comment",
			mcp.WithDescription("Delete a comment from a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("comment_id", mcp.Required(), mcp.Description("The comment ID to delete")),
		),
		s.handleDeleteComment,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_me",
			mcp.WithDescription("Get the current authenticated user's information."),
		),
		s.handleGetMe,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_team_projects",
			mcp.WithDescription("Get all projects in a Figma team."),
			mcp.WithString("team_id", mcp.Required(), mcp.Description("The Figma team ID")),
		),
		s.handleGetTeamProjects,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_project_files",
			mcp.WithDescription("Get all files in a Figma project."),
			mcp.WithString("project_id", mcp.Required(), mcp.Description("The Figma project ID")),
		),
		s.handleGetProjectFiles,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_team_components",
			mcp.WithDescription("Get all published components in a Figma team library."),
			mcp.WithString("team_id", mcp.Required(), mcp.Description("The Figma team ID")),
		),
		s.handleGetTeamComponents,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_components",
			mcp.WithDescription("Get all components defined in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetFileComponents,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_component",
			mcp.WithDescription("Get details of a specific component by its key."),
			mcp.WithString("key", mcp.Required(), mcp.Description("The component key")),
		),
		s.handleGetComponent,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_component_sets",
			mcp.WithDescription("Get all component sets (variants) defined in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetFileComponentSets,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_team_styles",
			mcp.WithDescription("Get all published styles in a Figma team library."),
			mcp.WithString("team_id", mcp.Required(), mcp.Description("The Figma team ID")),
		),
		s.handleGetTeamStyles,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_file_styles",
			mcp.WithDescription("Get all styles defined in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetFileStyles,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_local_variables",
			mcp.WithDescription("Get all local variables and variable collections in a Figma file. Requires Enterprise plan."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetLocalVariables,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_published_variables",
			mcp.WithDescription("Get all published variables from a Figma file. Requires Enterprise plan."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetPublishedVariables,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_variable_summary",
			mcp.WithDescription("Get a human-readable summary of all variable collections and their variables in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
		),
		s.handleGetVariableSummary,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_export_design_tokens",
			mcp.WithDescription("Export design tokens (variables) from a Figma file as CSS custom properties or JSON."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("format", mcp.Description("Output format: css, json, raw (default: css)")),
			mcp.WithString("mode", mcp.Description("Filter tokens by mode name (e.g., 'Light', 'Dark')")),
		),
		s.handleExportDesignTokens,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_create_variable",
			mcp.WithDescription("Create a new variable in a Figma file. Requires Enterprise plan and edit access."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("collection_id", mcp.Required(), mcp.Description("The variable collection ID")),
			mcp.WithString("name", mcp.Required(), mcp.Description("Variable name")),
			mcp.WithString("resolved_type", mcp.Required(), mcp.Description("Variable type: BOOLEAN, FLOAT, STRING, COLOR")),
		),
		s.handleCreateVariable,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_create_variable_collection",
			mcp.WithDescription("Create a new variable collection in a Figma file. Requires Enterprise plan and edit access."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("name", mcp.Required(), mcp.Description("Collection name")),
		),
		s.handleCreateVariableCollection,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_webhooks",
			mcp.WithDescription("Get all webhooks for a Figma team."),
			mcp.WithString("team_id", mcp.Required(), mcp.Description("The Figma team ID")),
		),
		s.handleGetWebhooks,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_create_webhook",
			mcp.WithDescription("Create a new webhook for a Figma team."),
			mcp.WithString("team_id", mcp.Required(), mcp.Description("The Figma team ID")),
			mcp.WithString("event_type", mcp.Required(), mcp.Description("Event type: FILE_UPDATE, FILE_DELETE, FILE_VERSION_UPDATE, FILE_COMMENT, LIBRARY_PUBLISH, TEAM_COMPONENT")),
			mcp.WithString("endpoint", mcp.Required(), mcp.Description("The webhook endpoint URL")),
			mcp.WithString("passcode", mcp.Required(), mcp.Description("Passcode for webhook verification")),
			mcp.WithString("description", mcp.Description("Webhook description")),
		),
		s.handleCreateWebhook,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_delete_webhook",
			mcp.WithDescription("Delete a Figma webhook."),
			mcp.WithString("webhook_id", mcp.Required(), mcp.Description("The webhook ID to delete")),
		),
		s.handleDeleteWebhook,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_get_dev_resources",
			mcp.WithDescription("Get dev resources (links) attached to nodes in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("node_ids", mcp.Description("Comma-separated list of node IDs to filter by")),
		),
		s.handleGetDevResources,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_create_dev_resource",
			mcp.WithDescription("Create a dev resource (link) attached to a node in a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("node_id", mcp.Required(), mcp.Description("The node ID to attach the resource to")),
			mcp.WithString("name", mcp.Required(), mcp.Description("Resource name")),
			mcp.WithString("url", mcp.Required(), mcp.Description("Resource URL")),
		),
		s.handleCreateDevResource,
	)

	s.mcp.AddTool(
		mcp.NewTool("figma_delete_dev_resource",
			mcp.WithDescription("Delete a dev resource from a Figma file."),
			mcp.WithString("file_key", mcp.Required(), mcp.Description("The Figma file key")),
			mcp.WithString("dev_resource_id", mcp.Required(), mcp.Description("The dev resource ID to delete")),
		),
		s.handleDeleteDevResource,
	)
}

func (s *Server) Serve() error {
	return server.ServeStdio(s.mcp)
}

func main() {
	token := os.Getenv("FIGMA_ACCESS_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "Error: FIGMA_ACCESS_TOKEN environment variable is required")
		os.Exit(1)
	}

	s := NewServer(token)
	if err := s.Serve(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
