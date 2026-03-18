package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	client := mcp.NewClient(&mcp.Implementation{Name: "bear-fried", Version: "v1.0.0"}, nil)

	transport := &mcp.CommandTransport{Command: exec.Command("../bear-server/weather")}
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	params := &mcp.CallToolParams{
		Name:      "get_forecast",
		Arguments: map[string]any{"latitude": 40.7128, "longitude": -74.0060},
	}
	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatal("CallTool failed: %v", err)
	}
	if res.IsError {
		log.Fatal("tool failed")
	}
	for _, c := range res.Content {
		log.Print(c.(*mcp.TextContent).Text)
	}
}
