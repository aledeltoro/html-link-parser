package link

import (
	"io"

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
	p.search(p.doc)

	return p.links
}

func (p *Parser) search(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		link := newLink(node)

		p.links = append(p.links, link)
	}

	if node.FirstChild != nil {
		p.search(node.FirstChild)
	}

	if node.NextSibling != nil {
		p.search(node.NextSibling)
	}
}

func newLink(node *html.Node) Link {
	href := getHref(node.Attr)
	text := ""

	if node.FirstChild != nil {
		text = node.FirstChild.Data
	}

	return Link{
		Href: href,
		Text: text,
	}
}

func getHref(attributes []html.Attribute) string {
	for _, attr := range attributes {
		if attr.Key == "href" {
			return attr.Val
		}
	}

	return ""
}
