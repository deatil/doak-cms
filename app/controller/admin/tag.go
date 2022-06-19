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
    "github.com/deatil/doak-cms/app/response"
)

// 标签
type Tag struct{
    Base
}

// 页面
func (this *Tag) Index(ctx *fiber.Ctx) error {
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
    cates := make([]model.Tag, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("name like ?", "%" + keywords + "%").
        Desc("sort").
        Asc("id")
    if status != "" {
        modeldb = modeldb.Where("status = ?", status)
    }

    modeldb.Find(&cates)

    // 总数
    countdb := db.Engine().
        Where("name like ?", "%" + keywords + "%")
    if status != "" {
        countdb = countdb.Where("status = ?", status)
    }

    total, _ := countdb.Count(new(model.Tag))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "tag/index", fiber.Map{
        "keywords": keywords,
        "status": status,
        "total": total,
        "list": cates,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 编辑
func (this *Tag) Edit(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 标签信息
    var data model.Tag
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    return this.View(ctx, "tag/edit", fiber.Map{
        "id": id,
        "data": data,
    })
}

// 编辑保存
func (this *Tag) EditSave(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "编辑失败")
    }

    name := cast.ToString(ctx.FormValue("name"))
    desc := cast.ToString(ctx.FormValue("desc"))
    sort := cast.ToInt(ctx.FormValue("sort"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "name": name,
            "sort": sort,
            "status": status,
        },
        map[string]string{
            "name": "required|minLen:1",
            "sort": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "name.required": "标签名称不能为空",
            "name.minLen": "标签名称不能少于1位",
            "sort.required": "标签排序不能为空",
            "status.required": "标签状态不能为空",
            "status.in": "标签状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    // 标签信息
    var data model.Tag
    db.Engine().
        Where("name = ?", name).
        Get(&data)
    if data.Id > 0 && data.Id != id {
        return http.Error(ctx, 1, "编辑失败, 标签[" + name + "]已经存在")
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    _, err := db.Engine().
        Table(new(model.Tag)).
        Where("id = ?", id).
        Update(map[string]any{
            "name": name,
            "desc": desc,
            "sort": sort,
            "status": newStatus,
        })
    if err != nil {
        return http.Error(ctx, 1, "编辑失败")
    }

    return http.Success(ctx, "编辑成功", "")
}

// 删除
func (this *Tag) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "删除失败")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Tag))
    if err != nil {
        return http.Error(ctx, 1, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
