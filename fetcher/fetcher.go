package fetcher

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var RESPONSESTATUSCODENOT200 = errors.New("Response status code is not 200")

// send http request
// return response's body
func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Client get error: %v", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Printf("%s response status is not 200", url)
		return nil, RESPONSESTATUSCODENOT200
	}
	defer resp.Body.Close()
	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	reader := transform.NewReader(
		bodyReader, e.NewDecoder())
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

// auto check html charset
// default return utf-8
func determinEncoding(r *bufio.Reader) encoding.Encoding {
	// read 1024 bytes without advancing the reader
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Peek error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
