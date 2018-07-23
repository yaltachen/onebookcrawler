package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yaltachen/onebookcrawler/dangdang/parser"
	"github.com/yaltachen/onebookcrawler/fetcher"
	"github.com/yaltachen/onebookcrawler/models"
	"github.com/yaltachen/onebookcrawler/wordutil"

	"github.com/julienschmidt/httprouter"
)

func registerRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/book/:url", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(p.ByName("url")))
		basicURL, err := ioutil.ReadAll(decoder)
		if err != nil {
			fmt.Println(err)
			return
		}
		body, _ := fetcher.Fetch(string(basicURL))
		basicInfo, detailURL := parser.ParseBasicInfo(body)
		detailBody, _ := fetcher.Fetch(detailURL.GetRequestURL())
		detailInfo := parser.ParseBookDetail(basicInfo.BookName, detailBody)
		book := &models.Book{*basicInfo, *detailInfo}
		doc, err := wordutil.WriteDoc(book)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Error")
			return
		}
		w.Header().Set("Content-type", "application/msword")
		w.Header().Set("Content-Disposition", "attachment;fileName="+book.BookName+".doc")
		w.WriteHeader(http.StatusOK)
		err = doc.Save(w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Error")
			return
		}
		return
	})
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w, nil)
	})
	return router
}

func main() {
	r := registerRouter()
	server := http.Server{Addr: ":9000", Handler: r}
	server.ListenAndServe()
}
