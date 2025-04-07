package writer

import (
	"strings"
	"unicode/utf8"
)

type Config struct {
	width   int
	padding int
	align   Alignment
	end     End
}

func (c Config) innerWidth() int {
	return c.width - 2*c.padding
}

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

type End int

const (
	Overflow End = iota
	Wrap
)

type Writer struct {
	config  Config
	builder strings.Builder
}

type Opt func(*Config)

func WithAlignment(align Alignment) Opt {
	return func(config *Config) {
		config.align = align
	}
}

func WithEnd(end End) Opt {
	return func(config *Config) {
		config.end = end
	}
}

func WithFullWidth() Opt {
	return func(config *Config) {
		config.padding = 0
	}
}

func New(opts ...Opt) *Writer {
	c := Config{
		width:   80,
		padding: 10,
		align:   AlignLeft,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return &Writer{
		config: c,
	}
}

func (w *Writer) Write(s string, opts ...Opt) {
	c := w.config
	for _, opt := range opts {
		opt(&c)
	}

	var length = utf8.RuneCountInString

	blockWidth := 0
	for _, line := range strings.Split(s, "\n") {
		blockWidth = max(blockWidth, length(line))
	}

	padding := strings.Repeat(" ", c.padding)

	var innerPaddingLen int
	switch c.align {
	case AlignLeft:
		innerPaddingLen = 0
	case AlignCenter:
		innerPaddingLen = (c.innerWidth() - blockWidth) / 2
	case AlignRight:
		innerPaddingLen = c.innerWidth() - blockWidth
	}
	innerPadding := strings.Repeat(" ", innerPaddingLen)

	for _, line := range strings.Split(s, "\n") {
		if c.end == Wrap && innerPaddingLen+length(line) > c.innerWidth() {
			w.writeWrappedLine(line, c.innerWidth()-innerPaddingLen, padding, innerPadding)
			continue
		}

		out := padding + innerPadding + line
		w.builder.WriteString(out + "\n")
	}
}

// writeWrappedLine handles text wrapping without recursive calls
func (w *Writer) writeWrappedLine(line string, width int, padding, innerPadding string) {
	var length = utf8.RuneCountInString
	remaining := line

	for length(remaining) > 0 {
		chunkSize := min(width, length(remaining))
		chunk := remaining[:chunkSize]

		out := padding + innerPadding + chunk
		w.builder.WriteString(out + "\n")

		remaining = remaining[chunkSize:]
	}
}

func (w *Writer) Hr() {
	s := strings.Repeat("-", w.config.innerWidth())
	w.Write(s)
}

func (w *Writer) Tr() {
	s := strings.Repeat("=", w.config.width)
	w.Write(s, WithFullWidth())
}

func (w *Writer) String() string {
	return w.builder.String()
}
