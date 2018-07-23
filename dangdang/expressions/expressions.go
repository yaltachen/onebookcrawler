package expressions

import "regexp"

var (
	PriceReg      = regexp.MustCompile(`<div class="price_m" id='original-price'>.+\n.+<\/span>(.+)<\/div>`)
	FormatReg     = regexp.MustCompile(`开 本：(\d+)开`)
	PaperReg      = regexp.MustCompile(`纸 张：(.+)<\/li><li>包`)
	PackReg       = regexp.MustCompile(`包 装：(.+)<\/li><li>是`)
	SuitReg       = regexp.MustCompile(`是否套装：(.+)<\/li><li>国`)
	IsbnReg       = regexp.MustCompile(`国际标准书号ISBN：(\d+)`)
	PressReg      = regexp.MustCompile(`出版社：([^。]+)。`)
	BookNameReg   = regexp.MustCompile(`<h1 title="([^"]+)">`)
	BriefIntroReg = regexp.MustCompile(`<span class="head_title_name" [^"]+">([^<]+)<\/span>`)
	AuthorReg     = regexp.MustCompile(`作者：(.+)，出版社`)

	ProductIdReg    = regexp.MustCompile(`"productId":"(\d+)"`)
	TemplateTypeReg = regexp.MustCompile(`"template":"(.+)","productType":`)
	DescribeMapReg  = regexp.MustCompile(`"describeMap":"(.+)","categoryId"`)
	ShopIdReg       = regexp.MustCompile(`"shopId":"(.+)","isCatalog"`)
	CategoryPathReg = regexp.MustCompile(`"categoryPath":"(.+)","describeMap"`)
)
