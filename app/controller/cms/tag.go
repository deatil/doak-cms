package cms

import (
    "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/page"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
)

/**
 * 标签文章列表
 *
 * @create 2022-6-19
 * @author deatil
 */
type Tag struct{
    Base
}

func (this *Tag) Index(ctx *fiber.Ctx) error {
    tagName := cast.ToString(ctx.Params("tag"))
    if tagName == "" {
        return response.CmsErrorRender(ctx, "标签错误")
    }

    // url.QueryEscape(str)
    tag, _ := url.QueryUnescape(tagName)

    // 当前页码
    currentPage := cast.ToInt(ctx.Query("page", "1"))
    if currentPage < 1 {
        currentPage = 1
    }

    // 每页数量
    listRows := 5
    start := (currentPage - 1) * listRows

    // 文章列表
    arts := make([]model.ArtCate, 0)
    db.Engine().
        Table(new(model.Art)).Alias("a").
        Join("LEFT", []any{new(model.Cate), "c"}, "c.id = a.cate_id").
        Limit(listRows, start).
        Where("FIND_IN_SET(?, `a`.`tags`)", tag).
        Where("a.status = ?", 1).
        Desc("a.is_top").
        Desc("a.add_time").
        Find(&arts)

    newArts := make([]model.ArtCatename, 0)
    if len(arts) > 0 {
        for _, a := range arts {
            newArts = append(newArts, model.ArtCatename{
                Art: a.Art,
                CateName: a.Cate.Name,
                CateSlug: a.Cate.Slug,
            })
        }
    }

    // 总数
    total, _ := db.Engine().
        Where("status = ?", 1).
        Where("FIND_IN_SET(?, `tags`)", tag).
        Count(new(model.Art))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "tag", fiber.Map{
        "tag": tag,
        "arts": newArts,
        "total": total,
        "pageHtml": pageHtml,
    })
}
