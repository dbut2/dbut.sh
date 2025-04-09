package provider

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/glamour"

	"dbut.sh/pages"
)

type ContentProvider interface {
	GetContent(path string) ([]byte, error)
	GetContentType() string
}

type HTMLProvider struct{}

func (p *HTMLProvider) GetContent(path string) ([]byte, error) {
	if path == "" || path == "/" {
		path = "index"
	}
	path = strings.TrimPrefix(path, "/")
	return pages.HTML.ReadFile(path + ".html")
}

func (p *HTMLProvider) GetContentType() string {
	return "text/html;charset=UTF-8"
}

type MarkdownProvider struct {
	renderer *glamour.TermRenderer
}

func NewMarkdownProvider(width int) (*MarkdownProvider, error) {
	renderWidth := 80
	if width > 0 {
		renderWidth = width
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(renderWidth),
		glamour.WithEmoji(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create markdown renderer: %w", err)
	}

	return &MarkdownProvider{
		renderer: renderer,
	}, nil
}

func (p *MarkdownProvider) GetContent(path string) ([]byte, error) {
	if path == "" || path == "/" {
		path = "index"
	}
	path = strings.TrimPrefix(path, "/")

	content, err := pages.MD.ReadFile(path + ".md")
	if err != nil {
		return nil, fmt.Errorf("failed to read markdown file: %w", err)
	}

	if p.renderer != nil {
		return p.renderer.RenderBytes(content)
	}

	return content, nil
}

func (p *MarkdownProvider) GetContentType() string {
	return "text/plain;charset=UTF-8"
}

type RawMarkdownProvider struct{}

func (p *RawMarkdownProvider) GetContent(path string) ([]byte, error) {
	if path == "" || path == "/" {
		path = "index"
	}
	path = strings.TrimPrefix(path, "/")

	return pages.MD.ReadFile(path + ".md")
}

func (p *RawMarkdownProvider) GetContentType() string {
	return "text/plain;charset=UTF-8"
}

type VanityProvider struct {
	template string
}

func NewVanityProvider(template string) *VanityProvider {
	return &VanityProvider{
		template: template,
	}
}

func (p *VanityProvider) GetContent(_ string) ([]byte, error) {
	return []byte(p.template), nil
}

func (p *VanityProvider) GetContentType() string {
	return "text/html;charset=UTF-8"
}

func WriteContent(w io.Writer, provider ContentProvider, path string) error {
	content, err := provider.GetContent(path)
	if err != nil {
		return err
	}

	_, err = w.Write(content)
	return err
}
