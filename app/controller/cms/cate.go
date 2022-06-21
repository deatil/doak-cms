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
 * 分类
 *
 * @create 2022-6-19
 * @author deatil
 */
type Cate struct{
    Base
}

func (this *Cate) Index(ctx *fiber.Ctx) error {
    slug := cast.ToString(ctx.Params("slug"))
    if slug == "" {
        return response.CmsErrorRender(ctx, "分类错误")
    }

    // 当前页码
    currentPage := cast.ToInt(ctx.Query("page", "1"))
    if currentPage < 1 {
        currentPage = 1
    }

    // 每页数量
    listRows := 5
    start := (currentPage - 1) * listRows

    // 分类信息
    var cate model.Cate
    _, err := db.Engine().
        Where("slug = ?", slug).
        Get(&cate)
    if err != nil || cate.Id == 0 {
        return response.CmsErrorRender(ctx, "分类错误")
    }

    // 文章列表
    arts := make([]model.Art, 0)
    db.Engine().
        Limit(listRows, start).
        Where("cate_id = ?", cate.Id).
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
        Where("cate_id = ?", cate.Id).
        Count(new(model.Art))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    return this.View(ctx, "cate", fiber.Map{
        "cate": cate,
        "arts": newArts,
        "total": total,
        "pageHtml": pageHtml,
    })
}
