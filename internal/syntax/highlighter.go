package syntax

import (
	"encoding/json"
	"regexp"
	"strings"

	"TUI_HTTP/internal/styles"

	"github.com/charmbracelet/lipgloss"
)

type Highlighter struct {
	jsonKeyStyle    lipgloss.Style
	jsonStringStyle lipgloss.Style
	jsonNumberStyle lipgloss.Style
	jsonBoolStyle   lipgloss.Style
	jsonNullStyle   lipgloss.Style
	htmlTagStyle    lipgloss.Style
	htmlAttrStyle   lipgloss.Style
	xmlTagStyle     lipgloss.Style
}

func NewHighlighter() *Highlighter {
	return &Highlighter{
		jsonKeyStyle:    lipgloss.NewStyle().Foreground(styles.Blue).Bold(true),
		jsonStringStyle: lipgloss.NewStyle().Foreground(styles.Green),
		jsonNumberStyle: lipgloss.NewStyle().Foreground(styles.Purple),
		jsonBoolStyle:   lipgloss.NewStyle().Foreground(styles.Orange),
		jsonNullStyle:   lipgloss.NewStyle().Foreground(styles.DarkGray),
		htmlTagStyle:    lipgloss.NewStyle().Foreground(styles.HotPink),
		htmlAttrStyle:   lipgloss.NewStyle().Foreground(styles.Blue),
		xmlTagStyle:     lipgloss.NewStyle().Foreground(styles.Purple),
	}
}

func (h *Highlighter) Highlight(content, contentType string) string {
	if content == "" {
		return content
	}

	if contentType == "" {
		contentType = h.detectContentType(content)
	}

	switch {
	case strings.Contains(contentType, "json"):
		return h.highlightJSON(content)
	case strings.Contains(contentType, "html"):
		return h.highlightHTML(content)
	case strings.Contains(contentType, "xml"):
		return h.highlightXML(content)
	default:
		return content
	}
}

func (h *Highlighter) detectContentType(content string) string {
	trimmed := strings.TrimSpace(content)

	if (strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}")) ||
		(strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]")) {

		var js json.RawMessage
		if json.Unmarshal([]byte(content), &js) == nil {
			return "application/json"
		}
	}

	if strings.HasPrefix(trimmed, "<") && strings.HasSuffix(trimmed, ">") {
		if strings.Contains(strings.ToLower(content), "<html") {
			return "text/html"
		}
		return "application/xml"
	}

	return "text/plain"
}

func (h *Highlighter) highlightJSON(content string) string {

	var obj interface{}
	if err := json.Unmarshal([]byte(content), &obj); err == nil {
		if formatted, err := json.MarshalIndent(obj, "", "  "); err == nil {
			content = string(formatted)
		}
	}

	patterns := []struct {
		regex *regexp.Regexp
		style lipgloss.Style
	}{

		{regexp.MustCompile(`"([^"\\]|\\.)*":`), h.jsonKeyStyle},

		{regexp.MustCompile(`:\s*"([^"\\]|\\.)*"`), h.jsonStringStyle},

		{regexp.MustCompile(`:\s*-?\d+\.?\d*([eE][+-]?\d+)?`), h.jsonNumberStyle},

		{regexp.MustCompile(`:\s*(true|false)`), h.jsonBoolStyle},

		{regexp.MustCompile(`:\s*null`), h.jsonNullStyle},
	}

	result := content
	for _, pattern := range patterns {
		result = pattern.regex.ReplaceAllStringFunc(result, func(match string) string {
			if strings.Contains(match, ":") {

				parts := strings.SplitN(match, ":", 2)
				if len(parts) == 2 {
					if strings.Contains(match, `"`) && !strings.Contains(parts[1], `"`) {

						return h.jsonKeyStyle.Render(parts[0]) + ":" + parts[1]
					}
					return parts[0] + ":" + pattern.style.Render(parts[1])
				}
			}
			return pattern.style.Render(match)
		})
	}

	return result
}

func (h *Highlighter) highlightHTML(content string) string {

	tagRegex := regexp.MustCompile(`<[^>]+>`)
	content = tagRegex.ReplaceAllStringFunc(content, func(match string) string {
		return h.htmlTagStyle.Render(match)
	})

	attrRegex := regexp.MustCompile(`(\w+)=("[^"]*"|'[^']*')`)
	content = attrRegex.ReplaceAllStringFunc(content, func(match string) string {
		return h.htmlAttrStyle.Render(match)
	})

	return content
}

func (h *Highlighter) highlightXML(content string) string {

	tagRegex := regexp.MustCompile(`<[^>]+>`)
	content = tagRegex.ReplaceAllStringFunc(content, func(match string) string {
		return h.xmlTagStyle.Render(match)
	})

	return content
}
