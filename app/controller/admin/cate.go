package admin

import (
    "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/page"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
)

// 分类
type Cate struct{
    Base
}

// 列表
func (this *Cate) Index(ctx *fiber.Ctx) error {
    // 总数
    cate := new(model.Cate)
    total, _ := db.
        Engine().
        Where("id >?", 1).
        Count(cate)

    // 当前页码
    currentPage := cast.ToInt(ctx.Query("page", "1"))
    if currentPage < 1 {
        currentPage = 1
    }

    // 每页数量
    listRows := 5
    start := (currentPage - 1) * listRows

    // 列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Limit(listRows, start).
        Asc("id").
        Find(&cates)

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "cate/index", fiber.Map{
        "total": total,
        "list": cates,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 添加
func (this *Cate) Add(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "cate/add", fiber.Map{
        "Title": data,
    })
}

// 添加保存
func (this *Cate) AddSave(ctx *fiber.Ctx) error {
    name := cast.ToString(ctx.FormValue("name"))
    slug := cast.ToString(ctx.FormValue("slug"))
    desc := cast.ToString(ctx.FormValue("desc"))
    status := cast.ToInt(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "name": name,
            "slug": slug,
            "status": status,
        },
        map[string]string{
            "name": "required|minLen:3",
            "slug": "required|minLen:3",
            "status": "required",
        },
        map[string]string{
            "name.required": "分类名称不能为空",
            "name.minLen": "分类名称不能少于3位",
            "slug.required": "分类标志不能为空",
            "slug.minLen": "分类标志不能少于3位",
            "status.required": "分类状态不能为空",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    cate := new(model.Cate)
    cate.Name = name
    cate.Slug = slug
    cate.Desc = desc
    cate.Status = status
    cate.AddIp = ctx.IP()
    _, err := db.Engine().Insert(cate)
    if err != nil {
        return http.Error(ctx, 1, "添加失败")
    }

    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Cate) Edit(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "cate/edit", fiber.Map{
        "Title": data,
    })
}

// 编辑保存
func (this *Cate) EditSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "保存成功", "")
}

// 删除
func (this *Cate) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "删除失败")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Cate))
    if err != nil {
        return http.Error(ctx, 1, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
