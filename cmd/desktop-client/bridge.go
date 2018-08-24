package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dtoebe/collab-md/internal/document"
	"github.com/dtoebe/collab-md/internal/utils"

	"github.com/dtoebe/collab-md/internal/clipboard"
	"github.com/dtoebe/collab-md/internal/markdown"
	"github.com/therecipe/qt/core"
)

type QmlBridge struct {
	core.QObject

	_ string `property:"mdareaText"`
	_ string `property:"rtareaText"`
	_ string `property:"mdCopyText"`
	_ string `property:"rtCopyText"`
	_ string `property:"connClick"`

	_ func() `constructor:"init"`
}

func (b *QmlBridge) init() {
	b.ConnectMdareaTextChanged(b.mdareaTextChanged)
	b.ConnectRtareaTextChanged(b.rtareaTextChanged)
	b.ConnectMdCopyTextChanged(b.mdCopyTextChanged)
	b.ConnectRtCopyTextChanged(b.rtCopyTextChanged)
	b.ConnectConnClickChanged(b.connClickChanged)

	b.SetMdareaText("# Markdown text")
	b.SetRtareaText(markdown.MdToHTML(b.MdareaText()))
}

func (b *QmlBridge) connClickChanged(s string) {
	log.Println("go: connClicked:", s)
	if s != "clicked" {
		return
	}
	res, err := createDoc("document-"+utils.GenRandStr(5), b.MdareaText())
	if err != nil {
		log.Println(err)
		return
	}

	var docData struct {
		Document document.Document `json:"document"`
		ID       string            `json:"id"`
		Metadata json.RawMessage   `json:"metadata,omitempty"`
		URL      string            `json:"url"`
	}

	err = json.Unmarshal(res, &docData)
	if err != nil {
		log.Println(err)
		return
	}

	b.SetMdareaText(fmt.Sprintf("TITLE: %s\n\nURL: %s\n\n%s",
		docData.Document.Title, docData.URL, docData.Document.Body))
}

func (b *QmlBridge) mdareaTextChanged(s string) {
	b.SetRtareaText(markdown.MdToHTML(s))
	log.Println("go: mdareaText:", s)
}

func (b *QmlBridge) rtareaTextChanged(s string) {
	log.Println("go: rtareaText:", s)
}

func (b *QmlBridge) mdCopyTextChanged(s string) {
	err := clipboard.WriteToClipboard(s)
	if err != nil {
		log.Println(err)
	}
}

func (b *QmlBridge) rtCopyTextChanged(s string) {
	formatted := markdown.MdToHTML(s)
	err := clipboard.WriteToClipboard(formatted)
	if err != nil {
		log.Println(err)
	}
}
