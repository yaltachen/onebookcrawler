package parser

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/yaltachen/onebookcrawler/dangdang/expressions"
	"github.com/yaltachen/onebookcrawler/models"

	"github.com/PuerkitoBio/goquery"
)

func ParseBasicInfo(body []byte) (*models.BookBasicInfo, *models.BookDetailUrl) {
	basicInfo := &models.BookBasicInfo{}

	// get book basic infos through regular expression
	basicInfo.BookName = extractString(body, expressions.BookNameReg)
	basicInfo.ISBN = extractString(body, expressions.IsbnReg)
	basicInfo.Format = extractString(body, expressions.FormatReg)
	basicInfo.Pack = extractString(body, expressions.PackReg)
	basicInfo.Press = extractString(body, expressions.PressReg)
	basicInfo.Suit = extractString(body, expressions.SuitReg)
	basicInfo.Paper = extractString(body, expressions.PaperReg)
	basicInfo.AuthorName = extractString(body, expressions.AuthorReg)
	price, err := strconv.ParseFloat(strings.TrimSpace(extractString(body, expressions.PriceReg)), 64)
	if err != nil {
		log.Printf("Parse book's price error. Book's price: %s, error: %v",
			extractString(body, expressions.PriceReg), err)
		basicInfo.Price = 0
	}
	basicInfo.Price = price

	// get book basic infos through dom
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Printf("Can not generate book:%s's basic info document", basicInfo.BookName)
	}
	basicInfo.BriefIntro = strings.TrimSpace(doc.Find(".head_title_name").Text())
	doc.Find(".breadcrumb").Find("a").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			basicInfo.Category = s.Text()
			return
		}
	})
	if picURL, exist := doc.Find("#main-img-slider").Find("a").Attr("data-imghref"); exist {
		basicInfo.BookPicURL = picURL
	}

	// get book detail request url params through regular expression
	bookDetailUrl := &models.BookDetailUrl{}
	bookDetailUrl.ProductId = extractString(body, expressions.ProductIdReg)
	bookDetailUrl.TemplateType = extractString(body, expressions.TemplateTypeReg)
	bookDetailUrl.DescribeMap = extractString(body, expressions.DescribeMapReg)
	bookDetailUrl.ShopId = extractString(body, expressions.ShopIdReg)
	bookDetailUrl.CategoryPath = extractString(body, expressions.CategoryPathReg)

	return basicInfo, bookDetailUrl
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
