package link

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input string
	want  []Link
}

func TestExtractLinks(t *testing.T) {
	c := require.New(t)

	tests := []test{
		{input: "samples/ex1.html", want: []Link{
			{Href: "/other-page", Text: "A link to another page"},
			{Href: "/dog", Text: "Something in a span Text not in a span Bold text!"},
		}},
		{input: "samples/ex2.html", want: []Link{
			{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
			{Href: "https://github.com/gophercises", Text: "Gophercises is on Github !"},
		}},
		{input: "samples/ex3.html", want: []Link{
			{Href: "#", Text: "Login"},
			{Href: "/lost", Text: "Lost? Need help?"},
			{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
		}},
		{input: "samples/ex4.html", want: []Link{
			{Href: "/dog-cat", Text: "dog cat"},
		}},
		{input: "samples/ex5.html", want: []Link{
			{Href: "#", Text: "Something here"},
			{Href: "/dog", Text: "nested dog link"},
		}},
	}

	for _, test := range tests {
		file, err := os.Open(test.input)
		c.NoError(err)

		links, err := Extract(file)
		c.NoError(err)
		c.Equal(test.want, links)
	}
}
