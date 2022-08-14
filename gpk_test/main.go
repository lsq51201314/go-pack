package main

import (
	"fmt"

	gpk "github.com/lsq51201314/go-pack"
)

func main() {
	var g gpk.Object
	defer g.Close()
	// //创建的同时也是导入，可以省略下面的Load
	// if err := g.CreateFromFolder(
	// 	"D:/UserData/Desktop/test1",
	// 	"D:/UserData/Desktop/test1.gpk",
	// 	"123456",
	// 	process,
	// ); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //此处测试导入
	// if err := g.Load("D:/UserData/Desktop/test1.gpk", "123456"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //数量
	// fmt.Println(g.GetTotal())
	// //列表
	// fmt.Println(g.GetList())
	// //读取
	// if data, err := g.GetData("dir2/dir3/0_0.jpg"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else if err := os.WriteFile("D:/UserData/Desktop/0_0.png", data, 0777); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //删除
	// g.Delete("dir2/dir3/0_0.jpg")
	// if err := g.Save(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //读取
	// if data, err := g.GetData("dir2/dir3/0_0.jpg"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else if err := os.WriteFile("D:/UserData/Desktop/0_0_test.png", data, 0777); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //添加或更新
	// if err := g.AddOrUpdate("test.txt", []byte("hello world")); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if err := g.Save(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //读取
	// if data, err := g.GetData("test.txt"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else if err := os.WriteFile("D:/UserData/Desktop/test.txt", data, 0777); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//创建一个空文件
	if err := g.Create("D:/UserData/Desktop/test3.gpk", "123456"); err != nil {
		fmt.Println(err)
		return
	}
	if err := g.Merge("D:/UserData/Desktop/test1.gpk", "123456", process); err != nil {
		fmt.Println(err)
		return
	}
	//重新载入试试
	if err := g.Load("D:/UserData/Desktop/test3.gpk", "123456"); err != nil {
		fmt.Println(err)
		return
	}
	//数量
	fmt.Println(g.GetTotal())
	//列表
	fmt.Println(g.GetList())
}

func process(file string, current, count int) {
	fmt.Println(file, current, count)
}
