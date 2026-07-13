package page

import (
    "fmt"
    "sort"
    "math"
    "strconv"
    "net/url"
)

func Paginate(
    listRows int,
    total int,
    urlPath string,
    parameters url.Values,
) *Pagination {
    return New().Paginate(listRows, total, urlPath, parameters)
}

// 分页
type Pagination struct {
    // 当前页
    CurrentPage int
    // 最后一页
    LastPage int
    // 数据总数
    Total int
    // 每页数量
    ListRows int
    // 当前开始位置
    NowStart int
    // 是否有下一页
    HasMore bool
    // 链接
    UrlPath string
    // url中参数
    Parameters url.Values
    // 生成链接
    PageHtml string
}

// 构造函数
func New() *Pagination {
    return &Pagination{}
}

// 每页数量
// listRows 每页数量
// total    总数量,
// urlPath  链接,
// config 配置参数
//   page:当前页,
//   path:url路径,
//   query:url额外参数,
//   fragment:url锚点,
func (this *Pagination) Paginate(
    listRows int,
    total int,
    urlPath string,
    parameters url.Values,
) *Pagination {
    this.UrlPath = urlPath
    this.Parameters = parameters

    // 当前页
    var page int
    pageStr := this.Parameters.Get("page")
    if pageStr == "" {
        page = 1
    } else {
        page, _ = strconv.Atoi(pageStr)
    }

    if page < 1 {
        page = 1
    }

    this.CurrentPage = page
    this.ListRows = listRows

    this.Total = total
    this.LastPage = int(math.Ceil(float64(this.Total) / float64(this.ListRows)))
    this.HasMore = this.CurrentPage < this.LastPage

    // 开始位置
    this.NowStart = (this.CurrentPage - 1) * this.ListRows

    this.PageHtml = this.render()

    return this
}

// 渲染分页html
func (this *Pagination) render() string {
    if !this.hasPages() {
        return ""
    }

    return fmt.Sprintf(
        `<ul class="pagination">%s %s %s</ul>`,
        this.getPreviousButton(),
        this.getLinks(),
        this.getNextButton(),
    )
}

// 上一页按钮
func (this *Pagination) getPreviousButton() string {
    text := `<span aria-hidden="true">&laquo;</span>`
    if this.CurrentPage <= 1 {
        return this.getDisabledTextWrapper(text)
    }

    url := this.url(this.CurrentPage - 1)
    return this.getPageLinkWrapper(url, text)
}

// 下一页按钮
func (this *Pagination) getNextButton() string {
    text := `<span aria-hidden="true">&raquo;</span>`
    if !this.HasMore {
        return this.getDisabledTextWrapper(text)
    }

    url := this.url(this.CurrentPage + 1)
    return this.getPageLinkWrapper(url, text)
}

// 页码按钮
func (this *Pagination) getLinks() string {
    block := map[string]map[int]string{
        "first":  nil,
        "slider": nil,
        "last":   nil,
    }

    side := 3
    window := side * 2

    // 判断
    switch {
        case this.LastPage < window+6:
            block["first"] = this.getUrlRange(1, this.LastPage)
        case this.CurrentPage <= window:
            block["first"] = this.getUrlRange(1, window+2)
            block["last"] = this.getUrlRange(this.LastPage-1, this.LastPage)
        case this.CurrentPage > (this.LastPage - window):
            block["first"] = this.getUrlRange(1, 2)
            block["last"] = this.getUrlRange(this.LastPage-(window+2), this.LastPage)
        default:
            block["first"] = this.getUrlRange(1, 2)
            block["slider"] = this.getUrlRange(this.CurrentPage-side, this.CurrentPage+side)
            block["last"] = this.getUrlRange(this.LastPage-1, this.LastPage)
    }

    html := ""
    if len(block["first"]) > 0 {
        html += this.getUrlLinks(block["first"])
    }

    if len(block["slider"]) > 0 {
        html += this.getDots()
        html += this.getUrlLinks(block["slider"])
    }

    if len(block["last"]) > 0 {
        html += this.getDots()
        html += this.getUrlLinks(block["last"])
    }

    return html

}

// 创建一组分页链接
func (this *Pagination) getUrlRange(start, end int) map[int]string {
    urls := map[int]string{}
    for page := start; page <= end; page++ {
        urls[page] = this.url(page)
    }

    return urls
}

// 获取页码对应的链接
func (this *Pagination) url(page int) string {
    urlPath := this.UrlPath

    parameters := this.Parameters
    if len(parameters) > 0 {
        // 复制值
        parameters.Set("page", strconv.Itoa(page))
        urlStr := parameters.Encode()
        return urlPath + "?" + urlStr
    }

    return urlPath + "?page=" + strconv.Itoa(page)
}

// 生成一个可点击的按钮
func (this *Pagination) getAvailablePageWrapper(url string, page string) string {
    return `<li class="page-item"><a class="page-link" href="` + url + `">` + page + `</a></li>`
}

// 生成一个禁用的按钮
func (this *Pagination) getDisabledTextWrapper(text string) string {
    return `<li class="page-item disabled"><a class="page-link">` + text + `</a></li>`
}

// 生成一个激活的按钮
func (this *Pagination) getActivePageWrapper(text string) string {
    return `<li class="page-item active"><a class="page-link">` + text + `</a></li>`
}

// 生成省略号按钮
func (this *Pagination) getDots() string {
    return this.getDisabledTextWrapper("...")
}

// 批量生成页码按钮
func (this *Pagination) getUrlLinks(urls map[int]string) string {
    html := ""
    var sortKeys []int
    for page, _ := range urls {
        sortKeys = append(sortKeys, page)
    }

    sort.Ints(sortKeys)
    for _, page := range sortKeys {
        html += this.getPageLinkWrapper(urls[page], page)
    }

    return html
}

// 生成普通页码按钮
func (this *Pagination) getPageLinkWrapper(url string, page interface{}) string {
    pageInt, ok := page.(int)
    if ok {
        if this.CurrentPage == pageInt {
            return this.getActivePageWrapper(strconv.Itoa(pageInt))
        }
        return this.getAvailablePageWrapper(url, strconv.Itoa(pageInt))
    }

    pageStr := page.(string)
    return this.getAvailablePageWrapper(url, pageStr)
}

// 数据是否足够分页
func (this *Pagination) hasPages() bool {
    return !(1 == this.CurrentPage && !this.HasMore)
}
