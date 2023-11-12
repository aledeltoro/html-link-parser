package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Parser struct to handle parsing of links
type Parser struct {
	doc   *html.Node
	links []Link
}

// Link struct containing information about a parsed link
type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

// NewParser creates a parser instance for the passed HTML document
func NewParser(reader io.Reader) (*Parser, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return &Parser{
		doc:   doc,
		links: make([]Link, 0),
	}, nil
}

// ExtractLinks returns a list of parsed from a HTML document
func (p *Parser) ExtractLinks() []Link {
	links := make([]Link, 0)
	stack := make([]*html.Node, 0)

	stack = append([]*html.Node{p.doc}, stack...)

	for len(stack) != 0 {
		currentNode := stack[0]

		if len(stack) > 1 {
			stack = stack[1:]
		} else {
			stack = stack[:0]
		}

		if currentNode.Type == html.ElementNode && currentNode.Data == "a" {
			link := newLink(currentNode)

			links = append(links, link)
		}

		if currentNode.NextSibling != nil {
			stack = append([]*html.Node{currentNode.NextSibling}, stack...)
		}

		if currentNode.FirstChild != nil {
			stack = append([]*html.Node{currentNode.FirstChild}, stack...)
		}
	}

	return links
}

func (p *Parser) search2(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		link := newLink(node)

		p.links = append(p.links, link)
	}

	if node.FirstChild != nil {
		p.search2(node.FirstChild)
	}

	if node.NextSibling != nil {
		p.search2(node.NextSibling)
	}
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
	stack := make([]*html.Node, 0)

	stack = append([]*html.Node{node}, stack...)

	for len(stack) != 0 {
		currentNode := stack[0]

		if len(stack) > 1 {
			stack = stack[1:]
		} else {
			stack = stack[:0]
		}

		if currentNode.Type == html.TextNode {
			text := strings.TrimSpace(currentNode.Data)
			if text != "" {
				textParts = append(textParts, text)
			}
		}

		if currentNode.NextSibling != nil && currentNode.Data != "a" {
			stack = append([]*html.Node{currentNode.NextSibling}, stack...)
		}

		if currentNode.FirstChild != nil {
			stack = append([]*html.Node{currentNode.FirstChild}, stack...)
		}
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
