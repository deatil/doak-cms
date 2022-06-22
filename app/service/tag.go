package service

import (
    "context"

    "github.com/spf13/cast"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/rpc"

    "github.com/deatil/doak-cms/app/model"
)

/**
 * 标签
 *
 * @create 2022-6-21
 * @author deatil
 */
type Tag struct{}

// 返回列表
func (this *Tag) List(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
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
    list := make([]model.Tag, 0)
    err := db.Engine().
        Limit(limit, start).
        Where(where).
        OrderBy(orderby).
        Find(&list)

    // 总数
    total, _ := db.Engine().
        Where(where).
        Count(new(model.Tag))

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "total": total,
        "list": list,
        "error": errData,
    }

    return nil
}

// 返回详情
func (this *Tag) Info(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    // 条件
    where := cast.ToString(argsData["where"])
    if where == "" {
        reply.Body = map[string]any{
            "error": "where empty",
        }

        return nil
    }

    // 信息
    var data model.Tag
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
func (this *Tag) Add(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
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
        Table(new(model.Tag)).
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
func (this *Tag) Edit(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
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
        Table(new(model.Tag)).
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
func (this *Tag) Delete(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
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
        Delete(new(model.Tag))

    errData := ""
    if err != nil {
        errData = err.Error()
    }

    reply.Body = map[string]any{
        "error": errData,
    }

    return nil
}
