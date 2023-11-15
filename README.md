# HTML Link Parser

Simple package to parse the hyperlink and text from the anchortags in a HTML document.

## Setup

```
go get github.com/aledeltoro/html-link-parser/link
```

## Example

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aledeltoro/html-link-parser/link"
)

func main() {
	rawDocument := `
	<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/other-page">A link to another page</a>
	  <a href="/dog">
	    <span>Something in a span</span>
	    Text not in a span
	    <b>Bold text!</b>
	  </a>
	</body>
	</html>
	`

	reader := strings.NewReader(rawDocument)

	extractedLinks, err := link.Extract(reader)
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(extractedLinks)
}
```
