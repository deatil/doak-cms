package model

// 分类
type Cate struct {
    Id      int64 `xorm:"int(11)"`
    Pid     int64 `xorm:"int(11)"`
    Name    string `xorm:"varchar(50)"`
    Slug    string `xorm:"varchar(50)"`
    Desc    string `xorm:"varchar(200)"`
    Sort    int `xorm:"int(5)"`
    Status  int `xorm:"tinyint(1)"`
    AddTime int64 `xorm:"created"`
    AddIp   string `xorm:"varchar(50)"`
}

// 分类树
type CateTree struct {
    Id       int64
    Pid      int64
    Name     string
    Slug     string
    Desc     string
    Sort     int
    Status   int
    AddTime  int64
    AddIp    string
    Children []CateTree
}

// 转换为树形结构
func ToCateTree(data []Cate, pid int64) []CateTree {
    var list []CateTree

    for _, item := range data {
        if item.Pid == pid {
            newItem := CateTree{
                Id: item.Id,
                Pid: item.Pid,
                Name: item.Name,
                Slug: item.Slug,
                Desc: item.Desc,
                Sort: item.Sort,
                AddTime: item.AddTime,
                AddIp: item.AddIp,
                Children: ToCateTree(data, item.Id),
            }

            list = append(list, newItem)
        }
    }

    return list
}
