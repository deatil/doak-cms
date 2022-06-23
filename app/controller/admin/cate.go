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

/**
 * 分类
 *
 * @create 2022-6-19
 * @author deatil
 */
type Cate struct{
    Base
}

// 列表
func (this *Cate) Index(ctx *fiber.Ctx) error {
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

    // 总数
    countdb := db.Engine().
        Where("name like ?", "%" + keywords + "%")
    if status != "" {
        countdb = countdb.Where("status = ?", status)
    }

    total, _ := countdb.Count(new(model.Cate))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "cate/index", fiber.Map{
        "keywords": keywords,
        "status": status,
        "total": total,
        "list": cates,
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
            "slug": "required|minLen:3|isAlphaDash",
            "status": "required|in:y,n",
        },
        map[string]string{
            "name.required": "分类名称不能为空",
            "name.minLen": "分类名称不能少于1位",
            "slug.required": "分类标识不能为空",
            "slug.minLen": "分类标识不能少于3位",
            "slug.isAlphaDash": "分类标识错误",
            "status.required": "分类状态不能为空",
            "status.in": "分类状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    // 分类信息
    var data model.Cate
    has, _ := db.Engine().
        Where("slug = ?", slug).
        Get(&data)
    if has {
        return http.Error(ctx, 1, "添加失败, [" + slug + "] 标识已经存在")
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    _, err := db.Engine().Insert(&model.Cate{
        Pid: 0,
        Name: name,
        Slug: slug,
        Desc: desc,
        Sort: 100,
        Status: newStatus,
        AddIp: ctx.IP(),
    })
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
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
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
            "slug": "required|minLen:3|isAlphaDash",
            "sort": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "name.required": "分类名称不能为空",
            "name.minLen": "分类名称不能少于1位",
            "slug.required": "分类标识不能为空",
            "slug.minLen": "分类标识不能少于3位",
            "slug.isAlphaDash": "分类标识错误",
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

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    // 分类信息
    var data model.Cate
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, 1, "分类不存在")
    }

    // 判断是否重复
    var slugData model.Cate
    db.Engine().
        Where("slug = ?", slug).
        Get(&slugData)
    if slugData.Id > 0 && slugData.Id != id {
        return http.Error(ctx, 1, "编辑失败, 分类标识[" + slug + "]已经存在")
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
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, 1, "分类不存在")
    }

    // 分类信息
    var pidData model.Cate
    db.Engine().
        Where("pid = ?", id).
        Get(&pidData)
    if pidData.Id > 0 {
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
