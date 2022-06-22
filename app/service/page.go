package service

import (
    "context"

    "github.com/spf13/cast"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/rpc"

    "github.com/deatil/doak-cms/app/model"
)

/**
 * 单页
 *
 * @create 2022-6-21
 * @author deatil
 */
type Page struct{}

// 返回列表
func (this *Page) List(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    // 每页数量
    limit := cast.ToInt(argsData["limit"])
    if limit < 1 {
        limit = 5
    }

    // 开始位置
    start := cast.ToInt(argsData["start"])
    if start < 0 {
        start = 0
    }

    // 条件
    where := cast.ToString(argsData["where"])

    // 排序
    orderby := cast.ToString(argsData["orderby"])
    if where == "" {
        orderby = "id desc"
    }

    // 列表
    cates := make([]model.Page, 0)
    err := db.Engine().
        Limit(limit, start).
        Where(where).
        OrderBy(orderby).
        Find(&cates)

    // 总数
    total, _ := db.Engine().
        Where(where).
        Count(new(model.Page))

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "total": total,
        "list": cates,
        "error": errData,
    }

    return nil
}

// 返回详情
func (this *Page) Info(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    // 条件
    where := cast.ToString(argsData["where"])
    if where == "" {
        reply.Body = map[string]any{
            "error": "where empty",
        }

        return nil
    }

    // 分类信息
    var data model.Page
    _, err := db.Engine().
        Where(where).
        Get(&data)

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "data": data,
        "error": errData,
    }

    return nil
}

// 添加
func (this *Page) Add(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    // 添加的数据
    data := cast.ToStringMap(argsData["data"])
    if len(data) == 0 {
        reply.Body = map[string]any{
            "error": "data empty",
        }

        return nil
    }

    _, err := db.Engine().
        Table(new(model.Page)).
        Insert(data)

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "error": errData,
    }

    return nil
}

// 编辑
func (this *Page) Edit(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    // 条件
    where := cast.ToString(argsData["where"])
    if where == "" {
        reply.Body = map[string]any{
            "error": "where empty",
        }

        return nil
    }

    // 添加的数据
    data := cast.ToStringMap(argsData["data"])
    if len(data) == 0 {
        reply.Body = map[string]any{
            "error": "data empty",
        }

        return nil
    }

    _, err := db.Engine().
        Table(new(model.Page)).
        Where(where).
        Update(data)

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "error": errData,
    }

    return nil
}

// 删除
func (this *Page) Delete(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    // 条件
    where := cast.ToString(argsData["where"])
    if where == "" {
        reply.Body = map[string]any{
            "error": "where empty",
        }

        return nil
    }

    _, err := db.Engine().
        Where(where).
        Delete(new(model.Page))

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "error": errData,
    }

    return nil
}
