package main

import (
	"fmt"
	"os"
	"path/filepath"

	"dbut.sh/content"
)

//go:generate go run .

type TextContentGenerator func() string

func main() {
	pages := map[string]TextContentGenerator{
		"index":      content.Index,
		"swarm-gpus": content.GpuSwarm,
	}

	outputDir := "../../pages"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	for page, generator := range pages {
		outputPath := filepath.Join(outputDir, page+".md")
		content := generator()

		fmt.Printf("Generating %s...\n", outputPath)
		err := os.WriteFile(outputPath, []byte(content), 0666)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", outputPath, err)
			os.Exit(1)
		}
	}

	fmt.Println("Content generation complete")
}
