package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func isSrcRemote(src string) bool {
	u, err := url.Parse(src)
	Check(err)

	return u.Scheme == "http" || u.Scheme == "https"
}

func readRemoteFile(src string) string {
	resp, err := http.Get(src)
	defer resp.Body.Close()
	Check(err)
	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	return string(body)
}

func readFile(src string) string {
	contents, err := ioutil.ReadFile(src)
	Check(err)

	return string(contents)
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

func main() {
	var txt string
	src := os.Args[1]

	if isSrcRemote(src) {
		txt = readRemoteFile(src)
	} else {
		txt = readFile(src)
	}

	fmt.Printf("%d %d %d %s", lines(txt), words(txt), chars(txt), src)
}
