package admin

import (
    gourl "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/page"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/utils"

    "github.com/deatil/doak-cms/app/url"
    "github.com/deatil/doak-cms/app/model"
)

/**
 * 附件
 *
 * @create 2022-6-19
 * @author deatil
 */
type Attach struct{
    Base
}

// 页面
func (this *Attach) Index(ctx *fiber.Ctx) error {
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
    cates := make([]model.Attach, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("name like ?", "%" + keywords + "%").
        Desc("add_time")
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

    total, _ := countdb.Count(new(model.Attach))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := gourl.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "attach/index", fiber.Map{
        "keywords": keywords,
        "status": status,
        "total": total,
        "list": cates,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 删除
func (this *Attach) Delete(ctx *fiber.Ctx) error {
    id := cast.ToString(ctx.Params("id"))
    if id == "" {
        return http.Error(ctx, 1, "删除失败")
    }

    // 附件信息
    var data model.Attach
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, 1, "附件不存在")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.Attach))
    if err != nil {
        return http.Error(ctx, 1, "删除失败")
    }

    // 删除原始文件
    newPath := url.AttachPath(data.Path)
    if utils.FileExists(newPath) {
        utils.FileDelete(newPath)
    }

    return http.Success(ctx, "删除成功", "")
}

// 下载
func (this *Attach) Download(ctx *fiber.Ctx) error {
    id := cast.ToString(ctx.Params("id"))
    if id == "" {
        return ctx.SendString("附件不存在")
    }

    // 附件信息
    var data model.Attach
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil || data.Id == "" {
        return ctx.SendString("附件不存在")
    }

    newPath := url.AttachPath(data.Path)
    if !utils.FileExists(newPath) {
        return ctx.SendString("附件不存在")
    }

    return ctx.Download(newPath, data.Name);
}
