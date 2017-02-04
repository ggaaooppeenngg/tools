# callgraph

这个版本是基于官方的版本改写的，主要是基于 AST 到 SSA 的表示然后导出调用关系。
能够指定一个函数作为 root 生成调用图，并且指定展开的递归层数。

## 安装方式

目前包目录是没有改动的，还是基于 golang.org/x/tools 目录，不能用 go get 获取，所以要基于 GOPATH 创建 $GOPATH/src/golang.org/x
```
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/ggaaooppeenngg/tools.git
cd tools
go install ./cmd/callgraph
```
## 使用

比如生成 callgraph 本身的调用关系，entry 要带上完整的路径，不然无法分别来自其他包的函数。
```
cd $GOPATH/src/golang.org/x/tools/cmd/callgraph
callgraph --level=2 --entry "golang.org/x/tools/cmd/callgraph.main" .
```
如果想要生成函数调用的图而不是文本格式可以借助 graphviz
```
cd $GOPATH/src/golang.org/x/tools/cmd/callgraph
callgraph --format graphviz --level 2 --entry "golang.org/x/tools/cmd/callgraph.main" . > callgraph.dot
```
MAC 安装 graphviz 
```
brew install graphviz
```
UBUNTU 安装 graphviz
```
sudo apt-get install graphviz
```
再执行下面的命令就能得到图片格式的函数调用图了
```
dot -Tpng callgraph.dot -o callgraph.png
```

## call graph 的调用算法

1.  static calls only (unsound)
2.  cha         (Class Hierarchy Analysis)
3.  rta         (Rapid Type Analysis)
4.  pta         (inclusion-based Points-To Analysis)

每个算法的精确度依次增加，消耗时间也会增加，默认使用的是 rta。

## 标示某个函数是 go 出去的
