package admin

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"

    "github.com/deatil/doak-cms/app/model"
)

/**
 * 首页
 *
 * @create 2022-6-19
 * @author deatil
 */
type Index struct{
    Base
}

func (this *Index) Index(ctx *fiber.Ctx) error {
    // 文章总数
    artTotal, _ := db.Engine().Count(new(model.Art))

    // 分类总数
    cateTotal, _ := db.Engine().Count(new(model.Cate))

    // 标签总数
    tagTotal, _ := db.Engine().Count(new(model.Tag))

    // 账号总数
    userTotal, _ := db.Engine().Count(new(model.User))

    // 最新 10 条文章
    artList := make([]model.Art, 0)
    db.Engine().
        Limit(10, 0).
        Desc("add_time").
        Desc("id").
        Find(&artList)

    return this.View(ctx, "index/index", fiber.Map{
        "artTotal": artTotal,
        "cateTotal": cateTotal,
        "tagTotal": tagTotal,
        "userTotal": userTotal,

        "artList": artList,
    })
}
