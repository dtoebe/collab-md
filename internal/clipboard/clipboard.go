package clipboard

import (
	"github.com/atotto/clipboard"
)

func WriteToClipboard(s string) error {
	return clipboard.WriteAll(s)
}
