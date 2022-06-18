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

    // 状态
    status := cast.ToString(ctx.Query("status", ""))

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

    // 列表
    cates := make([]model.Cate, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("name like ?", "%" + keywords + "%").
        Desc("sort").
        Asc("id")
    if status != "" {
        modeldb = modeldb.Where("status = ?", status)
    }

    modeldb.Find(&cates)

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
        "keywords": keywords,
        "status": status,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 添加
func (this *Cate) Add(ctx *fiber.Ctx) error {
    return this.View(ctx, "cate/add", fiber.Map{})
}

// 添加保存
func (this *Cate) AddSave(ctx *fiber.Ctx) error {
    name := cast.ToString(ctx.FormValue("name"))
    slug := cast.ToString(ctx.FormValue("slug"))
    desc := cast.ToString(ctx.FormValue("desc"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "name": name,
            "slug": slug,
            "status": status,
        },
        map[string]string{
            "name": "required|minLen:1",
            "slug": "required|minLen:3",
            "status": "required|in:y,n",
        },
        map[string]string{
            "name.required": "分类名称不能为空",
            "name.minLen": "分类名称不能少于1位",
            "slug.required": "分类标志不能为空",
            "slug.minLen": "分类标志不能少于3位",
            "status.required": "分类状态不能为空",
            "status.in": "分类状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    // 分类信息
    var data model.Cate
    db.Engine().
        Where("slug = ?", slug).
        Get(&data)
    if data.Id > 0 {
        return http.Error(ctx, 1, "添加失败, [" + slug + "] 标识已经存在")
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    cate := new(model.Cate)
    cate.Pid = 0
    cate.Name = name
    cate.Slug = slug
    cate.Desc = desc
    cate.Sort = 100
    cate.Status = newStatus
    cate.AddIp = ctx.IP()
    _, err := db.Engine().Insert(cate)
    if err != nil {
        return http.Error(ctx, 1, "添加失败")
    }

    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Cate) Edit(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 分类信息
    var data model.Cate
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 分类列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Desc("sort").
        Asc("id").
        Find(&cates)

    // 转为树结构
    newCates := model.ToCateTree(cates, 0)

    return this.View(ctx, "cate/edit", fiber.Map{
        "id": id,
        "data": data,
        "cates": newCates,
    })
}

// 编辑保存
func (this *Cate) EditSave(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "编辑失败")
    }

    pid := cast.ToInt64(ctx.FormValue("pid"))
    name := cast.ToString(ctx.FormValue("name"))
    slug := cast.ToString(ctx.FormValue("slug"))
    desc := cast.ToString(ctx.FormValue("desc"))
    sort := cast.ToInt(ctx.FormValue("sort"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "name": name,
            "slug": slug,
            "sort": sort,
            "status": status,
        },
        map[string]string{
            "name": "required|minLen:1",
            "slug": "required|minLen:3",
            "sort": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "name.required": "分类名称不能为空",
            "name.minLen": "分类名称不能少于1位",
            "slug.required": "分类标志不能为空",
            "slug.minLen": "分类标志不能少于3位",
            "sort.required": "分类排序不能为空",
            "status.required": "分类状态不能为空",
            "status.in": "分类状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    if pid < 0 {
        return http.Error(ctx, 1, "父级ID错误")
    }

    // 分类信息
    var data model.Cate
    db.Engine().
        Where("slug = ?", slug).
        Get(&data)
    if data.Id > 0 && data.Id != id {
        return http.Error(ctx, 1, "添加失败, [" + slug + "] 标识已经存在")
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    _, err := db.Engine().
        Table(new(model.Cate)).
        Where("id = ?", id).
        Update(map[string]any{
            "pid": pid,
            "name": name,
            "slug": slug,
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
func (this *Cate) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, 1, "删除失败")
    }

    // 分类信息
    var data model.Cate
    db.Engine().
        Where("pid = ?", id).
        Get(&data)
    if data.Id > 0 {
        return http.Error(ctx, 1, "当前分类有子级，不能被删除")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Cate))
    if err != nil {
        return http.Error(ctx, 1, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
