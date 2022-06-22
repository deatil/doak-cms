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
    arts := make([]model.Art, 0)
    db.Engine().
        Limit(listRows, start).
        Where("FIND_IN_SET(?, `tags`)", tag).
        Where("status = ?", 1).
        Desc("is_top").
        Desc("add_time").
        Find(&arts)

    // 分类列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Desc("sort").
        Asc("id").
        Find(&cates)

    newCates := make(map[int64]model.Cate)
    if len(cates) > 0 {
        for _, cv := range cates {
            newCates[cv.Id] = cv
        }
    }

    newArts := make([]model.ArtCatename, 0)
    if len(arts) > 0 {
        for _, v := range arts {
            cateName := ""
            cateSlug := ""
            if newCate, ok := newCates[v.CateId]; ok {
                cateName = newCate.Name
                cateSlug = newCate.Slug
            }

            newArts = append(newArts, model.ArtCatename{
                Art: v,
                CateName: cateName,
                CateSlug: cateSlug,
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
