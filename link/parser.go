package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link struct containing information about a parsed link
type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

// Extract returns a list of parsed from a HTML document
func Extract(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("parsing reader: %w", err)
	}

	anchorTags := searchAnchorTags(doc)

	links := make([]Link, 0, len(anchorTags))

	for _, anchorTag := range anchorTags {
		link := newLink(anchorTag)

		links = append(links, link)
	}

	return links, nil
}

func searchAnchorTags(node *html.Node) []*html.Node {
	anchorTags := make([]*html.Node, 0)

	if node.Type == html.ElementNode && node.Data == "a" {
		anchorTags = append(anchorTags, node)
	}

	if node.FirstChild != nil {
		anchorTags = append(anchorTags, searchAnchorTags(node.FirstChild)...)
	}

	if node.NextSibling != nil {
		anchorTags = append(anchorTags, searchAnchorTags(node.NextSibling)...)
	}

	return anchorTags
}

func newLink(node *html.Node) Link {
	href := getHref(node.Attr)
	textParts := extractTextParts(node)

	return Link{
		Href: href,
		Text: strings.Join(textParts, " "),
	}
}

func extractTextParts(node *html.Node) []string {
	textParts := make([]string, 0)

	if node.Type == html.TextNode {
		text := strings.TrimSpace(node.Data)
		if text != "" {
			textParts = append(textParts, text)
		}
	}

	if node.FirstChild != nil {
		textParts = append(textParts, extractTextParts(node.FirstChild)...)
	}

	if node.NextSibling != nil && node.Data != "a" {
		textParts = append(textParts, extractTextParts(node.NextSibling)...)
	}

	return textParts
}

func getHref(attributes []html.Attribute) string {
	for _, attr := range attributes {
		if attr.Key == "href" {
			return attr.Val
		}
	}

	return ""
}
