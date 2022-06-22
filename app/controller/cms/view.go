package cms

import (
    "strings"
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
)

/**
 * 详情
 *
 * @create 2022-6-19
 * @author deatil
 */
type View struct{
    Base
}

func (this *View) Index(ctx *fiber.Ctx) error {
    id := cast.ToString(ctx.Params("id"))
    if id == "" {
        return response.CmsErrorRender(ctx, "数据不存在")
    }

    // 文章信息
    var data model.Art
    _, err := db.Engine().
        Where("uuid = ?", id).
        Get(&data)
    if err != nil || data.Id == 0 {
        return response.CmsErrorRender(ctx, "数据不存在")
    }

    // 标签
    tags := strings.Split(data.Tags, ",")

    // 分类信息
    var cate model.Cate
    db.Engine().
        Where("id = ?", data.CateId).
        Get(&cate)

    // 添加浏览量
    db.Engine().
        Table(new(model.Art)).
        Where("id = ?", data.Id).
        Update(map[string]any{
            "views": data.Views + 1,
        })

    return this.View(ctx, "view", fiber.Map{
        "id": id,
        "cate": cate,
        "data": data,
        "tags": tags,
    })
}
