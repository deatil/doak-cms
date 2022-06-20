package admin

import (
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
)

/**
 * 网站设置
 *
 * @create 2022-6-19
 * @author deatil
 */
type Setting struct{
    Base
}

// 页面
func (this *Setting) Index(ctx *fiber.Ctx) error {
    configs := make([]model.Config, 0)
    db.Engine().Find(&configs)

    settings := make(map[string]string)
    if len(configs) > 0 {
        for _, v := range configs {
            settings[v.Key] = v.Value
        }
    }

    return this.View(ctx, "setting/index", fiber.Map{
        "settings": settings,
    })
}

// 保存
func (this *Setting) Save(ctx *fiber.Ctx) error {
    name := cast.ToString(ctx.FormValue("website_name"))
    keywords := cast.ToString(ctx.FormValue("website_keywords"))
    description := cast.ToString(ctx.FormValue("website_description"))
    copyright := cast.ToString(ctx.FormValue("website_copyright"))
    status := cast.ToString(ctx.FormValue("website_status"))
    beian := cast.ToString(ctx.FormValue("website_beian"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "name": name,
            "status": status,
        },
        map[string]string{
            "name": "required",
            "status": "required",
        },
        map[string]string{
            "name.required": "网站名称不能为空",
            "status.required": "网站状态不能为空",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    // 需要更新的数据
    data := map[string]any{
        "website_name": name,
        "website_keywords": keywords,
        "website_description": description,
        "website_copyright": copyright,
        "website_status": status,
        "website_beian": beian,
    }

    // 更新
    for k, v := range data {
        db.Engine().
            Table(new(model.Config)).
            Where("`key` = ?", k).
            Update(map[string]any{
                "value": v,
            })
    }

    return http.Success(ctx, "更新成功", "")
}
