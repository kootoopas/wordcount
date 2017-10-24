package main

import (
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"os"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func chars(txt string) int {
	return len(txt)
}

func words(txt string) int {
	return len(strings.Split(txt, " "))
}

func lines(txt string) int {
	return len(strings.Split(txt, "\n"))
}

func isSrcRemote(src string) bool {
	u, err := url.Parse(src)
	Check(err)

	return u.Scheme == "http" || u.Scheme == "https"
}

func readRemoteFile(src string) string {
	req, err := http.Get(src)
	defer req.Body.Close()
	Check(err)
	body, err := ioutil.ReadAll(req.Body)
	Check(err)
	return string(body)
}

func readFile(src string) string {
	contents, err := ioutil.ReadFile(src)
	Check(err)
	return string(contents)
}

func main() {
	var txt string
	src := os.Args[1]

	if isSrcRemote(src) {
		txt = readRemoteFile(src)
	} else {
		txt = readFile(src)
 	}

	fmt.Printf("%d %d %d %s", chars(txt), words(txt), lines(txt), src)
}
