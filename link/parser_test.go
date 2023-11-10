package link

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	c := require.New(t)

	var parser *Parser
	var file io.Reader
	var err error

	file, err = os.Open("samples/ex1.html")
	c.NoError(err)

	parser, err = NewParser(file)
	c.NoError(err)
	c.NotNil(parser)

	links := parser.ExtractLinks()
	fmt.Printf("%+v \n", links)
}
