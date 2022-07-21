package admin

import (
    "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/log"
    "github.com/deatil/doak-cms/pkg/time"
    "github.com/deatil/doak-cms/pkg/page"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/utils"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/auth"
    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
    appAuth "github.com/deatil/doak-cms/app/auth"
)

/**
 * 文章
 *
 * @create 2022-6-19
 * @author deatil
 */
type Art struct{
    Base
}

// 列表
func (this *Art) Index(ctx *fiber.Ctx) error {
    // 当前页码
    currentPage := cast.ToInt(ctx.Query("page", "1"))
    if currentPage < 1 {
        currentPage = 1
    }

    // 每页数量
    listRows := 10
    start := (currentPage - 1) * listRows

    // 搜索关键字
    keywords := cast.ToString(ctx.Query("keywords", ""))

    // 分类
    cateid := cast.ToInt64(ctx.Query("cateid"))

    // 开始时间
    startTime := cast.ToString(ctx.Query("start_time", ""))

    // 结束时间
    endTime := cast.ToString(ctx.Query("end_time", ""))

    // 置顶
    istop := cast.ToString(ctx.Query("istop", ""))

    // 状态
    status := cast.ToString(ctx.Query("status", ""))

    // 列表
    arts := make([]model.Art, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("title like ? or uuid like ?", "%" + keywords + "%", "%" + keywords + "%")

    if cateid != 0 {
        modeldb = modeldb.Where("cate_id = ?", cateid)
    }

    if startTime != "" {
        startTimeObj := time.CreateFromFormat(time.DefaultFormat, startTime)

        modeldb = modeldb.Where("add_time >= ?", startTimeObj.Timestamp())
    }
    if endTime != "" {
        endTimeObj := time.CreateFromFormat(time.DefaultFormat, endTime)

        modeldb = modeldb.Where("add_time <= ?", endTimeObj.Timestamp())
    }

    if istop != "" {
        modeldb = modeldb.Where("is_top = ?", istop)
    }
    if status != "" {
        modeldb = modeldb.Where("status = ?", status)
    }

    // 非管理员
    if !auth.IsAdmin(ctx) {
        modeldb = modeldb.Where("user_id = ?", auth.GetUserId(ctx))
    }

    err := modeldb.
        Desc("add_time").
        Desc("id").
        Find(&arts)
    if err != nil {
        log.Log().Error(err.Error())
    }

    // 总数
    countdb := db.Engine().
        Where("title like ? or uuid like ?", "%" + keywords + "%", "%" + keywords + "%")
    if cateid != 0 {
        countdb = countdb.Where("cate_id = ?", cateid)
    }

    if startTime != "" {
        startTimeObj := time.CreateFromFormat(time.DefaultFormat, startTime)

        countdb = countdb.Where("add_time >= ?", startTimeObj.Timestamp())
    }
    if endTime != "" {
        endTimeObj := time.CreateFromFormat(time.DefaultFormat, endTime)

        countdb = countdb.Where("add_time <= ?", endTimeObj.Timestamp())
    }

    if istop != "" {
        countdb = countdb.Where("is_top = ?", istop)
    }
    if status != "" {
        countdb = countdb.Where("status = ?", status)
    }

    // 非管理员
    if !auth.IsAdmin(ctx) {
        countdb = countdb.Where("user_id = ?", auth.GetUserId(ctx))
    }

    total, _ := countdb.Count(new(model.Art))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    // 分类列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Desc("sort").
        Asc("id").
        Find(&cates)
    newCates := make(map[int64]string)
    if len(cates) > 0 {
        for _, cate := range cates {
            newCates[cate.Id] = cate.Name
        }
    }

    return this.View(ctx, "art/index", fiber.Map{
        "keywords": keywords,
        "cateid": cateid,
        "istop": istop,
        "status": status,

        "start_time": startTime,
        "end_time": endTime,

        "cates": newCates,

        "total": total,
        "list": arts,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 添加
func (this *Art) Add(ctx *fiber.Ctx) error {
    // 分类列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Desc("sort").
        Asc("id").
        Find(&cates)

    // 转为树结构
    newCates := model.ToCateTree(cates, 0)

    return this.View(ctx, "art/add", fiber.Map{
        "cates": newCates,
    })
}

// 添加保存
func (this *Art) AddSave(ctx *fiber.Ctx) error {
    cateId := cast.ToInt64(ctx.FormValue("cate_id"))
    title := cast.ToString(ctx.FormValue("title"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "cate_id": cateId,
            "title": title,
            "status": status,
        },
        map[string]string{
            "cate_id": "required",
            "title": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "cate_id.required": "文章分类不能为空",
            "title.required": "文章标题不能为空",
            "status.required": "文章状态不能为空",
            "status.in": "文章状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    // 当前账号
    userId := appAuth.GetUserInfo(ctx).Id
    nickname := appAuth.GetUserInfo(ctx).Nickname

    _, err := db.Engine().Insert(&model.Art{
        Uuid: utils.Uniqueid(),
        CateId: cateId,
        UserId: userId,
        Title: title,
        Content: "",
        From: nickname,
        IsTop: 0,
        Status: newStatus,
        AddIp: ctx.IP(),
    })
    if err != nil {
        return http.Error(ctx, "添加失败")
    }

    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Art) Edit(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 文章信息
    var data model.Art
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 非管理员
    if !auth.IsAdmin(ctx) && data.UserId != auth.GetUserId(ctx) {
        return response.AdminErrorRender(ctx, "你不能修改该文章")
    }

    // 分类列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Desc("sort").
        Asc("id").
        Find(&cates)

    // 转为树结构
    newCates := model.ToCateTree(cates, 0)

    return this.View(ctx, "art/edit", fiber.Map{
        "id": id,
        "data": data,
        "cates": newCates,
    })
}

// 编辑保存
func (this *Art) EditSave(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, "编辑失败")
    }

    cateId := cast.ToInt64(ctx.FormValue("cate_id"))
    title := cast.ToString(ctx.FormValue("title"))
    keywords := cast.ToString(ctx.FormValue("keywords"))
    description := cast.ToString(ctx.FormValue("description"))
    cover := cast.ToString(ctx.FormValue("cover"))
    content := cast.ToString(ctx.FormValue("content"))
    tags := cast.ToString(ctx.FormValue("tags"))
    from := cast.ToString(ctx.FormValue("from"))
    addTime := cast.ToString(ctx.FormValue("add_time"))
    isTop := cast.ToString(ctx.FormValue("is_top"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "cate_id": cateId,
            "title": title,
            "content": content,
            "add_time": addTime,
            "status": status,
        },
        map[string]string{
            "cate_id": "required",
            "title": "required",
            "content": "required",
            "add_time": "required",
            "status": "required|in:y,n",
        },
        map[string]string{
            "cate_id.required": "文章分类不能为空",
            "title.required": "文章标题不能为空",
            "content.required": "文章内容不能为空",
            "add_time.required": "文章添加时间不能为空",
            "status.required": "文章状态不能为空",
            "status.in": "文章状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    if cateId < 0 {
        return http.Error(ctx, "分类选择错误")
    }

    newIsTop := 0
    if isTop == "y" {
        newIsTop = 1
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    addTimeTimestamp := time.CreateFromFormat(time.DefaultFormat, addTime).Timestamp()

    // 文章信息
    var data model.Art
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, "文章不存在")
    }

    // 非管理员
    if !auth.IsAdmin(ctx) && data.UserId != auth.GetUserId(ctx) {
        return http.Error(ctx, "你不能修改该文章")
    }

    _, err := db.Engine().
        Table(new(model.Art)).
        Where("id = ?", id).
        Update(map[string]any{
            "cate_id": cateId,
            "title": title,
            "keywords": keywords,
            "description": description,
            "cover": cover,
            "content": content,
            "tags": tags,
            "from": from,
            "add_time": addTimeTimestamp,
            "is_top": newIsTop,
            "status": newStatus,
        })
    if err != nil {
        return http.Error(ctx, "编辑失败")
    }

    // engine.In("column", []int{1, 2, 3}).Find()

    return http.Success(ctx, "编辑成功", "")
}

// 删除
func (this *Art) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, "删除失败")
    }

    // 文章信息
    var data model.Art
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, "文章不存在")
    }

    // 非管理员
    if !auth.IsAdmin(ctx) && data.UserId != auth.GetUserId(ctx) {
        return http.Error(ctx, "你不能修改该文章")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Art))
    if err != nil {
        return http.Error(ctx, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
