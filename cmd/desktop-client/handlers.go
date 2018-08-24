package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

var (
	host  = "127.0.0.1"
	port  = "3080"
	proto = "http"
)

func createDoc(title, body string) ([]byte, error) {
	log.Println("sending document")
	c := &http.Client{}
	vals := url.Values{}
	vals.Set("title", title)
	vals.Set("body", body)
	fmt.Printf("VALS: %v\n", vals.Encode())
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s://%s/documents/create/", proto, net.JoinHostPort(host, port)),
		strings.NewReader(vals.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Connection-Type", "direct")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Println("response received")

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Println("body:", string(b))
	return b, nil
}
