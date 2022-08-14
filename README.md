# go-pack
游戏资源包的创建、修改、读取等操作。\
注意：资源包为游戏等需要快速读取的场景使用，所以并未对文件进行分块操作，都是一整个读取压缩的，所以需要自行保证文件大小的合理或者自行更改代码进行分包压缩，再则不对重复文件做校验，不同路径的同一个文件将被认为是两个文件。

引入包：
```go
gpk "github.com/lsq51201314/go-pack"
```

创建资源包：\
gpk.Object.Create(创建空的资源包)\
gpk.Object.CreateFromFolder(从文件夹创建资源包)
```go
var g gpk.Object
defer g.Close()
if err := g.CreateFromFolder(
	"D:/UserData/Desktop/resource",
	"D:/UserData/Desktop/resource.gpk",
	"123456",
	process,
); err != nil {
	fmt.Println(err)
	return
}
```
用于显示处理进度的函数：
```go
func process(file string, current, count int) {
	fmt.Println(file, current, count)
}
```

载入资源包：
```go
if err:= g.Load("D:/UserData/Desktop/resource.gpk","123456");err!= nil {
	fmt.Println(err)
	return
}
```

获取资源包内文件数量：
```go
g.GetTotal()
```

获取资源包内文件列表：
```go
g.GetList()
```

读取资源包内文件：
```go
if data, err := g.GetData("image 6021.png"); err != nil {
	fmt.Println(err)
	return
} else if err := os.WriteFile("D:/UserData/Desktop/image 6021.png", data, 0777); err != nil {
	fmt.Println(err)
	return
}
``` 

删除资源包内文件：\
注意：为了快速处理，并没有实际删除文件，仅删除了文件列表内索引，资源包的大小并不会有变化。
```go
g.Delete("image 6021.png")
if err := g.Save(); err != nil {
	fmt.Println(err)
	return
}
``` 

添加或更新文件：\
注意：为了快速处理，并没有实际更新文件，不管是添加还是更新都是在文件的末尾操作，资源包的大小会随着操作不断变大。
```go
if err := g.AddOrUpdate("test.txt", []byte("hello world")); err != nil {
	fmt.Println(err)
	return
}
if err := g.Save(); err != nil {
	fmt.Println(err)
	return
}
``` 

重建文件：\
说明：当资源包不断的操作之后将产生很多的无用碎片导致资源包非常大，使用此方法重新生成一个资源包删除掉无用的碎片是最好的选择，此操作为漫长过程，需要用户等待。
```go
if err := g.NewFile("D:/UserData/Desktop/newRes.gpk", process); err != nil {
	fmt.Println(err)
	return
}
``` 

合并资源包：\
说明：合并来自另外一个资源包的资源。
```go
if err := g.Merge("D:/UserData/Desktop/update.gpk", "123456", process); err != nil {
	fmt.Println(err)
	return
}
``` 
