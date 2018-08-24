package markdown

import "github.com/russross/blackfriday"

func MdToHTML(s string) string {
	formatted := blackfriday.MarkdownCommon([]byte(s))
	return string(formatted)
}
