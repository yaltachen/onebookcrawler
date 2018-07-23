package wordutil

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/yaltachen/onebookcrawler/models"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
)

func WriteDoc(b *models.Book) (*document.Document, error) {
	doc := document.New()

	// file, _ := os.Create("test.doc")
	table := doc.AddTable()
	writeRowText(table, "书名", b.BookName)
	writeRowText(table, "分类", b.Category)
	writeRowText(table, "作者", b.AuthorName)
	writeRowText(table, "出版社", b.Press)
	writeRowText(table, "出版时间", "")
	writeRowText(table, "原价", strconv.FormatFloat(b.Price, 'g', -1, 64))
	writeRowText(table, "折扣", "请填写折扣")
	writeRowText(table, "印刷时间", "")
	writeRowText(table, "开本", b.Format)
	writeRowText(table, "纸张", b.Paper)
	writeRowText(table, "包装", b.Pack)
	writeRowText(table, "套装", b.Suit)
	writeRowText(table, "ISBN", b.ISBN)
	writeRowText(table, "编辑推荐", b.EditorRecommend)
	writeRowText(table, "内容简介", b.ContentIntro)
	writeRowText(table, "作者简介", b.AuthorIntro)
	writeRowText(table, "封面", "")

	doc, err := writeRowImage(doc, table, "大图", b.BookPicURL)
	fmt.Println(b.BookPicURL)
	if err != nil {
		return nil, err
	}
	doc, err = writeRowImage(doc, table, "详情图", b.DetailPicURL)
	fmt.Println(b.DetailPicURL)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func writeRowText(table document.Table, key, value string) {
	row := table.AddRow()
	cell := row.AddCell()
	cell.AddParagraph().AddRun().AddText(key)
	cell = row.AddCell()
	cell.AddParagraph().AddRun().AddText(value)
}

// gooxml can not insert jpg to word(only support png)
func writeRowImage(doc *document.Document, table document.Table, key, url string) (*document.Document, error) {
	row := table.AddRow()
	cell := row.AddCell()
	cell.AddParagraph().AddRun().AddText(key)
	cell = row.AddCell()
	if url == "" {
		cell.AddParagraph().AddRun().AddText("")
	} else {
		client := http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			cell.AddParagraph().AddRun().AddText("")
		}
		defer resp.Body.Close()

		imgJpg, err := jpeg.Decode(resp.Body)
		if err != nil {
			return nil, err
		}
		buffer := bytes.NewBuffer([]byte{})
		err = png.Encode(buffer, imgJpg)
		all, _ := ioutil.ReadAll(buffer)
		ioutil.WriteFile(key, all, 0666)
		image, err := common.ImageFromFile(key)
		if err != nil {
			log.Printf("Can not load image from file error: %v", err)
			return nil, err
		}
		imageRef, err := doc.AddImage(image)
		if err != nil {
			return nil, err
		}
		inl, err := cell.AddParagraph().AddRun().AddDrawingInline(imageRef)
		inl.SetSize(1*measurement.Inch, 1*measurement.Inch)
		if err != nil {
			return nil, err
		}
	}
	return doc, nil
}
