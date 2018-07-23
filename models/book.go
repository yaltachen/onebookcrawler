package models

type BookBasicInfo struct {
	BookName   string  `json:"book_name,omitempty"`    // 书名
	ISBN       string  `json:"isbn,omitempty"`         // ISBN
	Paper      string  `json:"paper,omitempty"`        // 纸张
	Suit       string  `json:"suit,omitempty"`         // 是否为套装
	Price      float64 `json:"price,omitempty"`        // 定价
	Format     string  `json:"format,omitempty"`       // 开本
	Pack       string  `json:"pack,omitempty"`         // 装订
	Press      string  `json:"press,omitempty"`        // 出版社
	BriefIntro string  `json:"brief_intro,omitempty"`  // 一句话简介
	AuthorName string  `json:"author_name,omitempty"`  // 作者
	Category   string  `json:"category,omitempty"`     // 分类
	BookPicURL string  `json:"book_pic_url,omitempty"` // 图书封面URL
}

type BookDetail struct {
	DetailPicURL    string `json:"detail_pic,omitempty"`       // 详情图URL
	EditorRecommend string `json:"editor_recommend,omitempty"` // 编辑推荐
	ContentIntro    string `json:"content_intro,omitempty"`    // 内容简介
	AuthorIntro     string `json:"author_intro,omitempty"`     // 作者简介
}

type BookDetailUrl struct {
	ProductId    string
	TemplateType string
	DescribeMap  string
	ShopId       string
	CategoryPath string
}

func (b *BookDetailUrl) GetRequestURL() string {
	return "http://product.dangdang.com/index.php?r=callback/detail&productId=" + b.ProductId + "&templateType=" + b.TemplateType + "&describeMap=" + b.DescribeMap + "&shopId=" + b.ShopId + "&categoryPath=" + b.CategoryPath
}

type BookDetailResp struct {
	Data *struct {
		Html             string `json:"html,omitempty"`
		NavigationLabels []*struct {
			Index      string `json:"index,omitempty"`
			DdName     string `json:"ddName,omitempty"`
			ColumnName string `json:"columnName,omitempty"`
		} `json:"navigationLabels,omitempty"`
	}
	Elapse   float64 `json:"elapse,omitempty"`
	ErrMsg   string  `json:"errMsg,omitempty"`
	Location string  `json:"location,omitempty"`
}

type Book struct {
	BookBasicInfo
	BookDetail
}
