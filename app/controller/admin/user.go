package admin

import (
    "net/url"

    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/auth"
    "github.com/deatil/doak-cms/pkg/page"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
    appAuth "github.com/deatil/doak-cms/app/auth"
)

/**
 * 账号
 *
 * @create 2022-6-19
 * @author deatil
 */
type User struct{
    Base
}

// 登录
func (this *User) Index(ctx *fiber.Ctx) error {
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
    cates := make([]model.User, 0)
    modeldb := db.Engine().
        Limit(listRows, start).
        Where("username like ?", "%" + keywords + "%").
        Asc("id")
    if status != "" {
        modeldb = modeldb.Where("status = ?", status)
    }

    modeldb.Find(&cates)

    // 总数
    countdb := db.Engine().
        Where("username like ?", "%" + keywords + "%")
    if status != "" {
        countdb = countdb.Where("status = ?", status)
    }

    total, _ := countdb.Count(new(model.User))

    // url 链接信息
    uri := ctx.Request().URI()
    parameters, _ := url.ParseQuery(uri.QueryArgs().String())
    pageHtml := page.
        Paginate(listRows, int(total), string(uri.Path()), parameters).
        PageHtml

    return this.View(ctx, "user/index", fiber.Map{
        "keywords": keywords,
        "status": status,
        "total": total,
        "list": cates,
        "currentPage": currentPage,
        "pageHtml": pageHtml,
    })
}

// 添加
func (this *User) Add(ctx *fiber.Ctx) error {
    return this.View(ctx, "user/add", fiber.Map{})
}

// 添加保存
func (this *User) AddSave(ctx *fiber.Ctx) error {
    username := cast.ToString(ctx.FormValue("username"))
    nickname := cast.ToString(ctx.FormValue("nickname"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "username": username,
            "nickname": nickname,
        },
        map[string]string{
            "username": "required|minLen:1",
            "nickname": "required|minLen:2",
        },
        map[string]string{
            "username.required": "账号名称不能为空",
            "username.minLen": "账号名称不能少于1位",
            "nickname.required": "账号昵称不能为空",
            "nickname.minLen": "账号昵称不能少于3位",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    // 账号信息
    var data model.User
    has, _ := db.Engine().
        Where("username = ?", username).
        Get(&data)
    if has {
        return http.Error(ctx, "添加失败, 账号[" + username + "]已经存在")
    }

    _, err := db.Engine().Insert(&model.User{
        Username: username,
        Password: "",
        Nickname: nickname,
        Sign: "",
        Status: 0,
        AddIp: ctx.IP(),
    })
    if err != nil {
        return http.Error(ctx, "添加失败")
    }

    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *User) Edit(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    // 账号信息
    var data model.User
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil || data.Id == 0 {
        return response.AdminErrorRender(ctx, "数据不存在")
    }

    return this.View(ctx, "user/edit", fiber.Map{
        "id": id,
        "data": data,
    })
}

// 编辑保存
func (this *User) EditSave(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, "编辑失败")
    }

    // 当前账号
    userId := appAuth.GetUserInfo(ctx).Id
    if id == userId {
        return http.Error(ctx, "你不能编辑自己的账号")
    }

    username := cast.ToString(ctx.FormValue("username"))
    password := cast.ToString(ctx.FormValue("password"))
    nickname := cast.ToString(ctx.FormValue("nickname"))
    sign := cast.ToString(ctx.FormValue("sign"))
    status := cast.ToString(ctx.FormValue("status"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "username": username,
            "nickname": nickname,
            "status": status,
        },
        map[string]string{
            "username": "required|minLen:1",
            "nickname": "required|minLen:2",
            "status": "required|in:y,n",
        },
        map[string]string{
            "username.required": "账号名称不能为空",
            "username.minLen": "账号名称不能少于1位",
            "nickname.required": "账号昵称不能为空",
            "nickname.minLen": "账号昵称不能少于3位",
            "status.required": "账号状态不能为空",
            "status.in": "账号状态信息错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    // 账号信息
    var data model.User
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, "账号不存在")
    }

    // 账号信息
    var usernameData model.User
    db.Engine().
        Where("username = ?", username).
        Get(&usernameData)
    if usernameData.Id > 0 && usernameData.Id != id {
        return http.Error(ctx, "编辑失败, 账号[" + username + "]已经存在")
    }

    newStatus := 0
    if status == "y" {
        newStatus = 1
    }

    updateData := map[string]any{
        "username": username,
        "nickname": nickname,
        "sign": sign,
        "status": newStatus,
    }

    if password != "" {
        updateData["password"] = auth.Hash(password)
    }

    _, err := db.Engine().
        Table(new(model.User)).
        Where("id = ?", id).
        Update(updateData)
    if err != nil {
        return http.Error(ctx, "编辑失败")
    }

    return http.Success(ctx, "编辑成功", "")
}

// 删除
func (this *User) Delete(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return http.Error(ctx, "删除失败")
    }

    // 账号信息
    var data model.User
    has, _ := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if !has {
        return http.Error(ctx, "账号不存在")
    }

    _, err := db.Engine().
        Where("id = ?", id).
        Delete(new(model.User))
    if err != nil {
        return http.Error(ctx, "删除失败")
    }

    return http.Success(ctx, "删除成功", "")
}
