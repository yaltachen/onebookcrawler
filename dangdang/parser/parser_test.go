package parser

import (
	"crypto/sha1"
	"fmt"
	"testing"

	"github.com/yaltachen/onebookcrawler/fetcher"
	"github.com/yaltachen/onebookcrawler/models"
)

func TestParseBasicInfo(t *testing.T) {
	// t.SkipNow()
	tests := []struct {
		url       string
		basicInfo *models.BookBasicInfo
		detailUrl *models.BookDetailUrl
	}{
		{"http://product.dangdang.com/22610008.html", &models.BookBasicInfo{BookName: "计算机网络（第5版）", ISBN: "9787302274629", Paper: "胶版纸", Suit: "否", Price: 89.5, Format: "16", Pack: "平装", Press: "清华大学出版社", AuthorName: "（美）特南鲍姆", Category: "计算机/网络", BriefIntro: "AndrewS．Tanenbaum国内外使用*广泛、*权威的计算机网络经典教材。", BookPicURL: "http://img3m8.ddimg.cn/91/11/22610008-1_w_3.jpg"}, &models.BookDetailUrl{ProductId: "22610008", TemplateType: "publish", DescribeMap: "0100002835:1", ShopId: "0", CategoryPath: "01.54.07.05.00.00"}},
		{"http://product.dangdang.com/23578344.html", &models.BookBasicInfo{BookName: "计算机网络:自顶向下方法（原书第6版）", ISBN: "9787111453789", Paper: "胶版纸", Suit: "否", Price: 79, Format: "16", Pack: "平装-胶订", Press: "机械工业出版社", AuthorName: "（美）库罗斯", Category: "计算机/网络", BriefIntro: "", BookPicURL: "http://img3m4.ddimg.cn/9/20/23578344-1_w_1.jpg"}, &models.BookDetailUrl{ProductId: "23578344", TemplateType: "publish", DescribeMap: "0100002835:1", ShopId: "0", CategoryPath: "01.54.07.05.00.00"}},
		{"http://product.dangdang.com/22911745.html", &models.BookBasicInfo{BookName: "少年Pi的奇幻漂流（插图珍藏版）(央视朗读者王耀庆倾情朗读）", ISBN: "9787544731706", Paper: "胶版纸", Suit: "否", Price: 35, Format: "16", Pack: "平装", Press: "译林出版社", AuthorName: "(加)扬.马特尔", Category: "小说", BriefIntro: "", BookPicURL: "http://img3m5.ddimg.cn/76/13/22911745-1_w_1.jpg"}, &models.BookDetailUrl{ProductId: "22911745", TemplateType: "publish", DescribeMap: "", ShopId: "0", CategoryPath: "01.03.56.03.00.00"}},
		{"http://product.dangdang.com/463785.html", &models.BookBasicInfo{BookName: "时间简史（插图本）（央视《朗读者》推荐）", ISBN: "9787535732309", Paper: "胶版纸", Suit: "否", Price: 45, Format: "16", Pack: "平装-胶订", Press: "湖南科技出版社", BriefIntro: "时间只留简史  世间再无霍金", AuthorName: "史蒂芬·霍金", Category: "科普读物", BookPicURL: "http://img3m5.ddimg.cn/69/27/463785-1_w_1.jpg"}, &models.BookDetailUrl{ProductId: "463785", TemplateType: "publish", DescribeMap: "0100002835:1", ShopId: "0", CategoryPath: "01.52.01.00.00.00"}},
	}

	for _, test := range tests {
		body, _ := fetcher.Fetch(test.url)
		basicInfo, detailUrl := ParseBasicInfo(body)
		if *basicInfo != *test.basicInfo {
			t.Errorf("Expected bookBasicInfo %#v, but got %#v", *test.basicInfo, *basicInfo)
		}
		if *detailUrl != *test.detailUrl {
			t.Errorf("Expected bookDetailUrl %#v, but got %#v", *test.detailUrl, *detailUrl)
		}
	}
}

func TestParseBookDetail(t *testing.T) {
	// t.SkipNow()
	tests := []struct {
		bookName      string
		bookDetailURL *models.BookDetailUrl
		bookDetail    *models.BookDetail
	}{
		{"计算机网络（第5版）", &models.BookDetailUrl{ProductId: "22610008", TemplateType: "publish", DescribeMap: "0100002835:1", ShopId: "0", CategoryPath: "01.54.07.05.00.00"}, &models.BookDetail{EditorRecommend: "2b8dfa099613e21512ea22592bc0916395db5332", ContentIntro: "977edf4641f1bac4a7ee7304ab4aeb16294c0352", AuthorIntro: "e112f510989573b387e5eddeb999c70eb2d0d77f", DetailPicURL: ""}},
		{"计算机网络:自顶向下方法", &models.BookDetailUrl{ProductId: "23578344", TemplateType: "publish", DescribeMap: "0100002835:1", ShopId: "0", CategoryPath: "01.54.07.05.00.00"}, &models.BookDetail{EditorRecommend: "f3452d5274b4220a2dfa262bce69bdec378693ff", ContentIntro: "d45376ed13ca1e268f8f6ae733c5ae3720b777d5", AuthorIntro: "221e61697cb56eb7848f1ca051ebfcfcfe41b996", DetailPicURL: ""}},
		{"少年Pi的奇幻漂流", &models.BookDetailUrl{ProductId: "22911745", TemplateType: "publish", DescribeMap: "", ShopId: "0", CategoryPath: "01.03.56.03.00.00"}, &models.BookDetail{EditorRecommend: "698e0c4727563dde552fc4c7b203e91adbb57088", ContentIntro: "3208666c1853657364984c34bcd592681f6bcee7", AuthorIntro: "b77fd12488e6cffe0cf90a51f3e365f9d4458763", DetailPicURL: ""}},
		{"肖生克的救赎", &models.BookDetailUrl{ProductId: "23824407", TemplateType: "publish", DescribeMap: "0100002831:1,0100002833:1", ShopId: "0", CategoryPath: "01.03.38.00.00.00"}, &models.BookDetail{EditorRecommend: "766185cc961b59532dd7454e2e61d638a1138066", ContentIntro: "e2b0d363f47a234610360399686d76f5c3cc8a8d", AuthorIntro: "15c834ffba3509eff0d647f6b5ee5eaa3c99425d"}},
		{"时间简史",
			&models.BookDetailUrl{ProductId: "463785", TemplateType: "publish", DescribeMap: "0100002835:1", ShopId: "0", CategoryPath: "01.52.01.00.00.00"},
			&models.BookDetail{AuthorIntro: "89907466d9a9ce126ec2ff9182ef0855bd60d6f6", ContentIntro: "918d9960dfa794e2134b3db0b296a10fe9dd5b99", EditorRecommend: "af80b87cccc9964d85736452c2b306f3f25944d1", DetailPicURL: "http://img59.ddimg.cn/99999990001575079.jpg"}},
	}

	for _, test := range tests {
		body, _ := fetcher.Fetch(test.bookDetailURL.GetRequestURL())
		bookDetail := ParseBookDetail(test.bookName, body)
		h := sha1.New()
		h.Write([]byte(bookDetail.AuthorIntro))
		a := fmt.Sprintf("%x", h.Sum(nil))
		if a != test.bookDetail.AuthorIntro {
			t.Errorf("Expected author intro: %s, but got: %s", test.bookDetail.AuthorIntro, a)
		}
		h.Write([]byte(bookDetail.ContentIntro))
		c := fmt.Sprintf("%x", h.Sum(nil))
		if c != test.bookDetail.ContentIntro {
			t.Errorf("Expected content intro: %s, but got: %s", test.bookDetail.ContentIntro, c)
		}
		h.Write([]byte(bookDetail.EditorRecommend))
		e := fmt.Sprintf("%x", h.Sum(nil))
		if e != test.bookDetail.EditorRecommend {
			t.Errorf("Expected editor recommend: %s, but got: %s", test.bookDetail.EditorRecommend, e)
		}
		if test.bookDetail.DetailPicURL != bookDetail.DetailPicURL {
			t.Errorf("Expected detail pic URL: %s, but got: %s", test.bookDetail.DetailPicURL, bookDetail.DetailPicURL)
		}
	}
}
