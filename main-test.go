package main

import "fmt"

type teststruct struct {
	A string
	B string
	C string
}

func test(x *teststruct) {
	(*x) = teststruct{A: "2"}
}

func main() {
	//gredis.Initialize()
	//fmt.Println(gredis.IsOnline())
	//a := teststruct{
	//	A: "1",
	//	B: "2",
	//	C: "3",
	//}
	//ok, err := gredis.SetData("miao",a,0)
	//fmt.Println(ok,err)
	//fmt.Println(gredis.Set("miao","123123",0))
	//fmt.Println(gredis.GetString("miao"))
	//fmt.Println(gredis.Delete("miao"))
	//fmt.Println(gredis.GetString("miao"))
	//var a teststruct
	//fmt.Println(gredis.GetString("miao"))
	//ok:=gredis.GetData("miao",&a)
	//fmt.Println(ok,a.A,a.B,a.C)
	a := &teststruct{A: "1"}
	test(a)
	fmt.Println(a)
}
