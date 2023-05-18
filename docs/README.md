## etcd 用法

```cmd
$ etcdctl put foo bar
OK
```

```cmd
$ etcdctl get foo
foo
bar
```

```cmd
#只是获取值
$ etcdctl get foo --print-value-only
bar
```

```cmd
$ etcdctl put foo1 bar1
$ etcdctl put foo2 bar2
$ etcdctl put foo3 bar3

#获取从foo到foo3的值，不包括foo3
$ etcdctl get foo foo3 --print-value-only 
bar
bar1
bar2

# 获取前缀为foo的值
$ etcdctl get --prefix foo --print-value-only
bar
bar1
bar2
bar3
```

#删除foo
$ etcdctl del foo
1

#删除foo到foo2，不包括foo2
$ etcdctl del foo foo2
1
#删除key前缀为foo的
$ etcdctl del --prefix foo
2


============

#监视foo单个key
$ etcdctl watch foo
#另一个控制台执行： etcdctl put foo bar
PUT
foo
bar

#同时监视多个值
$ etcdctl watch -i
$ watch foo
$ watch zoo
# 另一个控制台执行: etcdctl put foo bar
PUT
foo
bar
# 另一个控制台执行: etcdctl put zoo val
PUT
zoo
val

#监视foo前缀的key
$ etcdctl watch --prefix foo
#另一个控制台执行： etcdctl put foo1 bar1
PUT
foo1
bar1
#另一个控制台执行： etcdctl put fooz1 barz1
PUT
fooz1
barz1
