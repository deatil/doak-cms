package cms

import (
    "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/page"
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

// 详情
func (this *Cate) Index(ctx *fiber.Ctx) error {
    slug := cast.ToString(ctx.Params("slug"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "slug": slug,
        },
        map[string]string{
            "slug": "required|minLen:3|isAlphaDash",
        },
        map[string]string{
            "slug.required": "分类标识不能为空",
            "slug.minLen": "分类标识不能少于3位",
            "slug.isAlphaDash": "分类标识错误",
        },
    )

    if (errs != nil) {
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
    arts := make([]model.ArtCate, 0)
    db.Engine().
        Table(new(model.Art)).Alias("a").
        Join("LEFT", []any{new(model.Cate), "c"}, "c.id = a.cate_id").
        Limit(listRows, start).
        Where("a.cate_id = ?", cate.Id).
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
        Where("cate_id = ?", cate.Id).
        Count(new(model.Art))

    // url 链接信息
    urlPath := string(ctx.Request().URI().Path())
    urlQuery := ctx.Request().URI().QueryArgs().String()
    parameters, _ := url.ParseQuery(urlQuery)
    pageHtml := page.New().
        Paginate(listRows, int(total), urlPath, parameters).
        PageHtml

    // 模板
    tpl := cate.Tpl
    if tpl == "" {
        return response.CmsErrorRender(ctx, "分类页面不存在")
    }

    return this.View(ctx, tpl, fiber.Map{
        "cate": cate,
        "arts": newArts,
        "total": total,
        "pageHtml": pageHtml,
    })
}
