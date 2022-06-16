package rpc

// 接收参数数据
type Args struct {
    Body map[string]any
}

// 回复
type Reply struct {
    Body map[string]any
}
