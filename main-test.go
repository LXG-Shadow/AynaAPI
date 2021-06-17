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
	//fmt.Println(novel.GetDataByProvider(&novel.BiqugeProvider,"https://www.biquge.com.cn/book/32135/"))
	//fmt.Println(novel.SearchByProvider(&novel.BiqugeBProvider,"猫腻").Data)
	//fmt.Printf("%q",novel.SearchByProvider(&novel.BiqugeCProvider,"好色小姨").Data["result"])
	//fmt.Printf("%q",novel.GetData("http://www.liquge.com/book/228/").Data["abstraction"])
	//fmt.Printf("%q",novel.GetData("http://www.liquge.com/book/228/538726.html").Data["content"])
	//fmt.Println(novel.GetDataByProvider(&novel.BiqugeProvider,"https://www.biquge.com.cn/book/32135/1214264.html"))
	//fmt.Println(novel.GetDataByProvider(&novel.BiqugeBProvider,"https://www.biquwx.la/106_106507/25237418.html"))
	//fmt.Println(novel.GetDataByProvider(&novel.BiqugeBProvider,"https://www.biquwx.la/106_106507"))
	a := teststruct{A: "a"}
	b := []teststruct{}
	c := []teststruct{}
	b = append(b, a)
	c = append(c, a)
	a.A = "c"
	fmt.Println(a)
	fmt.Println(b[0], c[0])
}
