package service

import (
    "context"

    "github.com/spf13/cast"

    "github.com/deatil/doak-cms/pkg/rpc"
)

type Arith struct{}

func (t *Arith) Mul(ctx context.Context, args *rpc.Args, reply *rpc.Reply) error {
    argsData := cast.ToStringMap(args.Body)

    res := cast.ToInt(argsData["A"]) * cast.ToInt(argsData["B"]) * 6

    reply.Body = map[string]any{
        "data": "成功",
        "res": res,
    }

    return nil
}
