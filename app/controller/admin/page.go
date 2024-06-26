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

    // 状态
    status := cast.ToString(ctx.Query("status", ""))

    // 列表
    arts := make([]model.Page, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("title like ?", "%" + keywords + "%").
        Desc("add_time").
        Desc("id")
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
    if status != "" {
        countdb = countdb.Where("status = ?", status)
    }

    total, _ := countdb.Count(new(model.Page))

    // url 链接信息
    uri := ctx.Request().URI()
    parameters, _ := url.ParseQuery(uri.QueryArgs().String())
    pageHtml := page.
        Paginate(listRows, int(total), string(uri.Path()), parameters).
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
        return http.Error(ctx, errs.One())
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    // 单页信息
    var data model.Page
    has, _ := db.Engine().
        Where("slug = ?", slug).
        Get(&data)
    if has {
        return http.Error(ctx, "添加失败, 单页标识[" + slug + "]已经存在")
    }

    // 当前账号
    userId := appAuth.GetUserInfo(ctx).Id

    // 模板列表
    tpls := this.ListTplFiles("page")
    tpl := "page"
    if len(tpls) > 0 {
        tpl = tpls[0]
    }

    _, err := db.Engine().Insert(&model.Page{
        UserId: userId,
        Slug: slug,
        Title: title,
        Content: "",
        Tpl: tpl,
        Status: newStatus,
        AddIp: ctx.IP(),
    })
    if err != nil {
        return http.Error(ctx, "添加失败")
    }

    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Page) Edit(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 单页信息
    var data model.Page
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil || data.Id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 模板列表
    tpls := this.ListTplFiles("page")

    return this.View(ctx, "page/edit", fiber.Map{
        "id": id,
        "data": data,
        "tpls": tpls,
    })
}

// 编辑保存
func (this *Page) EditSave(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, "编辑失败")
    }

    slug := cast.ToString(ctx.FormValue("slug"))
    title := cast.ToString(ctx.FormValue("title"))
    keywords := cast.ToString(ctx.FormValue("keywords"))
    description := cast.ToString(ctx.FormValue("description"))
    content := cast.ToString(ctx.FormValue("content"))
    tpl := cast.ToString(ctx.FormValue("tpl", ""))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "slug": slug,
            "title": title,
            "content": content,
            "tpl": tpl,
            "status": status,
        },
        map[string]string{
            "slug": "required|minLen:3|isAlphaDash",
            "title": "required",
            "content": "required",
            "tpl": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "slug.required": "单页标识不能为空",
            "slug.minLen": "单页标识不能少于3个字符",
            "slug.isAlphaDash": "单页标识错误",
            "title.required": "单页标题不能为空",
            "content.required": "单页内容不能为空",
            "tpl.required": "单页模板不能为空",
            "status.required": "单页状态不能为空",
            "status.in": "单页状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    // 单页信息
    var data model.Page
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, "单页不存在")
    }

    // 判断是否存在
    var slugData model.Page
    db.Engine().
        Where("slug = ?", slug).
        Get(&slugData)
    if slugData.Id > 0 && slugData.Id != id {
        return http.Error(ctx, "编辑失败, 单页标识[" + slug + "]已经存在")
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
            "tpl": tpl,
            "status": newStatus,
        })
    if err != nil {
        return http.Error(ctx, "编辑失败")
    }

    return http.Success(ctx, "编辑成功", "")
}

// 删除
func (this *Page) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, "删除失败")
    }

    // 单页信息
    var data model.Page
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, "单页不存在")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Page))
    if err != nil {
        return http.Error(ctx, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
