package parser

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/yaltachen/onebookcrawler/models"

	"github.com/PuerkitoBio/goquery"
)

func ParseBookDetail(bookName string, body []byte) *models.BookDetail {
	bookDetailResp := &models.BookDetailResp{}
	err := json.Unmarshal(body, bookDetailResp)
	if err != nil {
		log.Printf("Can not unmarshal book:%s detail response\n", bookName)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bookDetailResp.Data.Html))

	if err != nil {
		log.Printf("Can not generate book:%s's detail info document\n", bookName)
	}

	bookDetail := &models.BookDetail{}
	// 详情图
	if src, exist := doc.Find("#feature").Find(".descrip").Find("img").Attr("src"); exist {
		bookDetail.DetailPicURL = src
	}

	// 作者简介
	s := doc.Find("#authorIntroduction").Find(".descrip")
	for len(s.Children().Nodes) != 0 && s.Text() != "" {
		bookDetail.AuthorIntro += s.Text() + "\r\n"
		s = s.NextAll()
	}
	// 内容简介
	s = doc.Find("#content").Find(".descrip")
	for ; len(s.Children().Nodes) != 0 && s.Text() != ""; s = s.NextAll() {
		bookDetail.ContentIntro += s.Text() + "\r\n"
	}
	// 编辑推荐
	s = doc.Find("#abstract").Find(".descrip")
	for ; len(s.Children().Nodes) != 0 && s.Text() != ""; s = s.NextAll() {
		bookDetail.EditorRecommend += s.Text() + "\r\n"
	}

	return bookDetail
}
