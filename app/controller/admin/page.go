package admin

import (
    "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/log"
    "github.com/deatil/doak-cms/pkg/page"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
    appAuth "github.com/deatil/doak-cms/app/auth"
)

/**
 * 单页
 *
 * @create 2022-6-19
 * @author deatil
 */
type Page struct{
    Base
}

// 列表
func (this *Page) Index(ctx *fiber.Ctx) error {
    // 当前页码
    currentPage := cast.ToInt(ctx.Query("page", "1"))
    if currentPage < 1 {
        currentPage = 1
    }

    // 每页数量
    listRows := 5
    start := (currentPage - 1) * listRows

    // 搜索关键字
    keywords := cast.ToString(ctx.Query("keywords", ""))

    // 分类
    cateid := cast.ToInt(ctx.Query("cateid"))

    // 状态
    status := cast.ToString(ctx.Query("status", ""))

    // 列表
    arts := make([]model.Page, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("title like ?", "%" + keywords + "%").
        Desc("add_time").
        Desc("id")
    if cateid != 0 {
        modeldb = modeldb.Where("cate_id = ?", cateid)
    }
    if status != "" {
        modeldb = modeldb.Where("status = ?", status)
    }

    err := modeldb.Find(&arts)
    if err != nil {
        log.Log().Error(err.Error())
    }

    // 总数
    countdb := db.Engine().
        Where("title like ?", "%" + keywords + "%")
    if cateid != 0 {
        countdb = countdb.Where("cate_id = ?", cateid)
    }
    if status != "" {
        countdb = countdb.Where("status = ?", status)
    }

    total, _ := countdb.Count(new(model.Page))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "page/index", fiber.Map{
        "keywords": keywords,
        "status": status,

        "total": total,
        "list": arts,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 添加
func (this *Page) Add(ctx *fiber.Ctx) error {
    return this.View(ctx, "page/add", fiber.Map{})
}

// 添加保存
func (this *Page) AddSave(ctx *fiber.Ctx) error {
    slug := cast.ToString(ctx.FormValue("slug"))
    title := cast.ToString(ctx.FormValue("title"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "slug": slug,
            "title": title,
            "status": status,
        },
        map[string]string{
            "slug": "required|minLen:3|isAlphaDash",
            "title": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "slug.required": "单页标识不能为空",
            "slug.minLen": "单页标识不能少于3个字符",
            "slug.isAlphaDash": "单页标识错误",
            "title.required": "单页标题不能为空",
            "status.required": "单页状态不能为空",
            "status.in": "单页状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    // 当前账号
    userId := appAuth.GetUserInfo(ctx).Id

    _, err := db.Engine().Insert(&model.Page{
        UserId: userId,
        Slug: slug,
        Title: title,
        Content: "",
        Status: newStatus,
        AddIp: ctx.IP(),
    })
    if err != nil {
        return http.Error(ctx, 1, "添加失败")
    }

    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Page) Edit(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 分类信息
    var data model.Page
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    return this.View(ctx, "page/edit", fiber.Map{
        "id": id,
        "data": data,
    })
}

// 编辑保存
func (this *Page) EditSave(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "编辑失败")
    }

    slug := cast.ToString(ctx.FormValue("slug"))
    title := cast.ToString(ctx.FormValue("title"))
    keywords := cast.ToString(ctx.FormValue("keywords"))
    description := cast.ToString(ctx.FormValue("description"))
    content := cast.ToString(ctx.FormValue("content"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "slug": slug,
            "title": title,
            "content": content,
            "status": status,
        },
        map[string]string{
            "slug": "required|minLen:3|isAlphaDash",
            "title": "required",
            "content": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "slug.required": "单页标识不能为空",
            "slug.minLen": "单页标识不能少于3个字符",
            "slug.isAlphaDash": "单页标识错误",
            "title.required": "单页标题不能为空",
            "content.required": "单页内容不能为空",
            "status.required": "单页状态不能为空",
            "status.in": "单页状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    _, err := db.Engine().
        Table(new(model.Page)).
        Where("id = ?", id).
        Update(map[string]any{
            "slug": slug,
            "title": title,
            "keywords": keywords,
            "description": description,
            "content": content,
            "status": newStatus,
        })
    if err != nil {
        return http.Error(ctx, 1, "编辑失败")
    }

    return http.Success(ctx, "编辑成功", "")
}

// 删除
func (this *Page) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "删除失败")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Page))
    if err != nil {
        return http.Error(ctx, 1, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
