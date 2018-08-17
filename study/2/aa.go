//package main
//
//func main() {
//	//x := -1
//	//if x > 0 {
//	//	println("+x")
//	//} else if x < 0 {
//	//	println("-x")
//	//} else {
//	//	println("0")
//	//}
//
//	//switch {
//	//case x > 0:
//	//	println("+")
//	//case x < 0:
//	//	println("-")
//	//default:
//	//	println("x")
//	//}
//	for i := 0; i < 5; i++{
//		println(i)
//	}
//}
//
//
//
//

//package main
//
//import (
//	"os"
//	"text/template"
//	"fmt"
//	"net"
//)
//
//func main() {
//	tEmpty := template.New("template test")
//	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
//	tEmpty.Execute(os.Stdout, nil)
//
//	tWithValue := template.New("template test")
//	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
//	tWithValue.Execute(os.Stdout, nil)
//
//	tIfElse := template.New("template test")
//	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
//	tIfElse.Execute(os.Stdout, nil)
//}

//package main
//import (
//"net"
//"os"
//"fmt"
//)
//func main() {
//	if len(os.Args) != 2 {
//		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
//		os.Exit(1)
//	}
//	name := os.Args[1]
//	addr := net.ParseIP(name)
//	if addr == nil {
//		fmt.Println("Invalid address")
//	} else {
//		fmt.Println("The address is ", addr.String())
//	}
//	os.Exit(0)
//}

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	ch := make(chan int)
	defer close(ch)
	for i := 100; i < 1000; i++ {
		go calc(ch, i)
	}
	for v := range ch{
		fmt.Print(v)
		fmt.Print(" ")
	}
}

func calc(ch chan int, i int) {
	str_i := strconv.Itoa(i)
	list_i := strings.Split(str_i, "")
	a, _ := strconv.Atoi(list_i[0])
	b, _ := strconv.Atoi(list_i[1])
	c, _ := strconv.Atoi(list_i[2])
	if i == (a*a*a + b*b*b + c*c*c) {
		ch <- i
	}
}

//func main(){
//	var word string = "hello world"
//	list_word := strings.Split(word, " ")
//	for i:=len(list_word) - 1; i >=0; i--{
//		fmt.Print(list_word[i], " ")
//	}
//}
//package main
//
//import (
//"fmt"
//"strconv"
//"strings"
//)
//
//
//
//func main() {
//	ch := make(chan int)
//	//var wg sync.WaitGroup
//	defer close(ch)
//	for i := 100; i < 1000; i++ {
//		//wg.Add(1)
//		go calc(ch, i)
//	}
//	//wg.Wait()
//	x := 0
//	for {
//		select {
//		case v := <-ch:
//			fmt.Print(v)
//			fmt.Print(" ")
//		default:
//			x = 1
//		}
//		if x == 1{
//			break
//		}
//	}
//}
//
//func calc(ch chan int, i int) {
//	str_i := strconv.Itoa(i)
//	list_i := strings.Split(str_i, "")
//	a, _ := strconv.Atoi(list_i[0])
//	b, _ := strconv.Atoi(list_i[1])
//	c, _ := strconv.Atoi(list_i[2])
//	if i == (a*a*a + b*b*b + c*c*c) {
//		ch <- i
//	}
//}