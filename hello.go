//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main(){
//	fmt.Println( 1 << 8)
//	fmt.Println( 1 << 8)
//}

//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	fmt.Print("Go runs on ")
//	switch os := runtime.GOOS; os {
//	case "darwin":
//		fmt.Println("OS X.")
//	case "linux":
//		fmt.Println("Linux.")
//	default:
//		// freebsd, openbsd,
//		// plan9, windows...
//		fmt.Printf("%s.", os)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("counting")
//
//	for i := 0; i < 10; i++ {
//		defer fmt.Println(i)
//	}
//
//	fmt.Println("done")
//}

//package main
//
//import "fmt"
//
//func main() {
//	i, j := 42, 2701
//
//	p := &i         // point to i
//	fmt.Println(*p) // read i through the pointer
//	*p = 21         // set i through the pointer
//	fmt.Println(i)  // see the new value of i
//
//	p = &j         // point to j
//	*p = *p / 37   // divide j through the pointer
//	fmt.Println(j) // see the new value of j
//}



//package main
//
//import "fmt"
//
//type Vertex struct {
//	X, Y int
//}
//
//var (
//	v1 = Vertex{1, 2}  // 类型为 Vertex
//	v2 = Vertex{X: 1}  // Y:0 被省略
//	v3 = Vertex{}      // X:0 和 Y:0
//	p  = &Vertex{1, 2} // 类型为 *Vertex
//)
//
//func main() {
//	fmt.Println(v1, p, v2, v3)
//}

//package main
//
//import "fmt"
//
//func main() {
//	a:= []string{"hello", "world"}
//	a[0] = "Hello"
//	a[1] = "World"
//	fmt.Println(a[0], a[1])
//	fmt.Println(a)
//}

//package main
//
//import "fmt"
//
//func main() {
//	s := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println("s ==", s)
//
//	for i := 0; i < len(s); i++ {
//		fmt.Printf("s[%d] == %d\n", i, s[i])
//	}
//}



//package main
//
//import "fmt"
//
//func main() {
//	a := make([]int, 5)
//	printSlice("a", a)
//	b := make([]int, 0, 5)
//	printSlice("b", b)
//	c := b[:2]
//	printSlice("c", c)
//	d := c[2:5]
//	printSlice("d", d)
//}
//
//func printSlice(s string, x []int) {
//	fmt.Printf("%s len=%d cap=%d %v\n",
//		s, len(x), cap(x), x)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var z []int
//	fmt.Println(z, len(z), cap(z))
//	if z == nil {
//		fmt.Println("nil!")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a []int
//	printSlice("a", a)
//
//	// append works on nil slices.
//	a = append(a, 0)
//	printSlice("a", a)
//
//	// the slice grows as needed.
//	a = append(a, 1)
//	printSlice("a", a)
//
//	// we can add more than one element at a time.
//	a = append(a, 2, 3, 4)
//	printSlice("a", a)
//}
//
//func printSlice(s string, x []int) {
//	fmt.Printf("%s len=%d cap=%d %v\n",
//		s, len(x), cap(x), x)
//}

//package main
//
//import "fmt"
//
//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
//
//func main() {
//	for i, v := range pow {
//		fmt.Printf("2**%d = %d\n", i, v)
//	}
//}


//实现 Pic 。它返回一个长度为 dy 的 slice，其中每个元素是一个长度为 dx 且元素类型为8位无符号整数的 slice。当你运行这个程序时，
// 它会将每个整数作为对应像素的灰度值（好吧，其实是蓝度）并显示这个 slice 所对应的图像。
//
//计算每个像素的灰度值的方法由你决定；几个有意思的选择包括 (x+y)/2、x*y 和 x^y 。
//
//（需要使用循环来分配 [][]uint8 中的每个 []uint8 。）
//
//（使用 uint8(intValue) 来在类型之间进行转换。）
//package main
//
//import "golang.org/x/tour/pic"
//
//func Pic(dx, dy int) [][]uint8 {
//	a := make([][]uint8, dx)
//	for x := range a{
//		b := make([]uint8, dy)
//		for y := range b{
//			b[y] = uint8(x*y - 1)
//		}
//		a[x] = b
//	}
//	return a
//}
//
//func main() {
//	pic.Show(Pic)
//}


//package main
//
//import "fmt"
//
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m map[string]Vertex
//
//func main() {
//	m = make(map[string]Vertex)
//	m["Bell Labs"] = Vertex{
//		40.68433, -74.39967,
//	}
//	fmt.Println(m["Bell Labs"])
//}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m = map[string]Vertex{
//	"Bell Labs": Vertex{
//		40.68433, -74.39967,
//	},
//	"Google": Vertex{
//		37.42202, -122.08408,
//	},
//}
//
//func main() {
//	fmt.Println(m)
//}

//package main
//
//import "fmt"
//
//func main() {
//	m := make(map[string]int)
//
//	m["Answer"] = 42
//	fmt.Println("The value:", m["Answer"])
//
//	m["Answer"] = 0
//	fmt.Println("The value:", m["Answer"])
//
//	v, ok := m["Answer"]
//	fmt.Println("The value:", v, "Present?", ok)
//
//	delete(m, "Answer")
//	fmt.Println("The value:", m["Answer"])
//
//	v, ok = m["Answer"]
//	fmt.Println("The value:", v, "Present?", ok)
//}


//package main
//
//import (
//	"strings"
//	"fmt"
//	"math"
//)
//
//func WordCount(s string) map[string]int {
//	str_list := strings.Fields(s)
//	ret := make(map[string]int)
//	for _,v := range str_list{
//		ret[v] += 1
//	}
//	return ret
//}
//
//func main() {
//	fmt.Println(WordCount("i am ha am ha h"))
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func compute(fn func(float64, float64) float64) float64 {
//	return fn(3, 4)
//}
//
//func main() {
//	hypot := func(x, y float64) float64 {
//		return math.Sqrt(x*x + y*y)
//	}
//	fmt.Println(hypot(5, 12))
//	fmt.Println(compute(hypot))
//	fmt.Println(compute(math.Pow))
//}


//package main
//
//import "fmt"
//
//func adder() func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
//
//func main() {
//	pos, neg := adder(), adder()
//	for i := 0; i < 10; i++ {
//		fmt.Println(
//			pos(i),
//			neg(-2*i),
//		)
//	}
//}



//package main
//
//import "fmt"
//
//// fibonacci 函数会返回一个返回 int 的函数。
//func fibonacci() func() int {
//	a,b := 0,1
//	return func() int {
//		tmp := b
//		b = a+b
//		a = tmp
//		return b
//	}
//}
//
//func main() {
//	f := fibonacci()
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}


//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := &Vertex{3, 4}
//	fmt.Println(v.Abs())
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//func main() {
//	f := MyFloat(-math.Sqrt2)
//	fmt.Println(f.Abs())
//}


//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := &Vertex{3, 4}
//	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
//	v.Scale(5)
//	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
//}


//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//type Reader interface {
//	Read(b []byte) (n int, err error)
//}
//
//type Writer interface {
//	Write(b []byte) (n int, err error)
//}
//
//type ReadWriter interface {
//	Reader
//	Writer
//}
//
//func main() {
//	var w ReadWriter
//
//	// os.Stdout 实现了 Writer
//	w = os.Stdout
//
//	fmt.Fprintf(w, "hello, writer\n")
//}


//package main
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) String() string {
//	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
//}
//
//func main() {
//	a := Person{"Arthur Dent", 42}
//	z := Person{"Zaphod Beeblebrox", 9001}
//	fmt.Println(a, z)
//}


//package main
//
//import (
//	"fmt"
//)
//
//type IPAddr [4]byte
//
//// TODO: Add a "String() string" method to IPAddr.
//func (v IPAddr) String() string{
//	return fmt.Sprintf("%v.%v.%v.%v", v[0],v[1],v[2],v[3])
//}
//
//
//func main() {
//	addrs := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for n, a := range addrs {
//		fmt.Printf("%v: %v\n", n, a)
//	}
//}


//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//type MyError struct {
//	When time.Time
//	What string
//}
//
//func (e *MyError) Error() string {
//	return fmt.Sprintf("at %v, %s",
//		e.When, e.What)
//}
//
//func run() error {
//	return &MyError{
//		time.Now(),
//		"it didn't work",
//	}
//}
//
//func main() {
//	if err := run(); err != nil {
//		fmt.Println(err)
//	}
//}

//package main
//
//import (
//    "fmt"
//    "math"
//)
//// 定义类型
//type ErrNegativeSqrt float64
//
//// 重写Error()
//func (e ErrNegativeSqrt) Error() string {
//    return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
//}
//
//func Sqrt(x float64) (float64, error) {
//    if x < 0 {
//        return 0, ErrNegativeSqrt(x)
//    }
//    return math.Sqrt(x), nil
//}
//
//func main() {
////通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil，    来进行错误处理。
////这里只是简单的打印
//    fmt.Println(Sqrt(2))
//    fmt.Println(Sqrt(-2))
//}


//package main
//
//import (
//	"fmt"
//	"io"
//	"strings"
//)
//
//func main() {
//	r := strings.NewReader("Hello, Reader!")
//
//	b := make([]byte, 8)
//	for {
//		n, err := r.Read(b)
//		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
//		fmt.Printf("b[:n] = %q\n", b[:n])
//		if err == io.EOF {
//			break
//		}
//	}
//}

//package main
//
//import "golang.org/x/tour/reader"
//
//type MyReader struct{}
//
//// TODO: Add a Read([]byte) (int, error) method to MyReader.
//func (v MyReader) Read(p []byte) (n int, err error){
//	p[0] = 'A'
//    return 1, nil
//}
//
//func main() {
//	reader.Validate(MyReader{})
//}


//package main
//
//import (
//	"io"
//	"os"
//	"strings"
//	//"fmt"
//)
//
//type rot13Reader struct {
//	r io.Reader
//}
//
//func (v rot13Reader) Read(p []byte) (n int, err error){
//	old := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
//	new_ := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
//	n, e := v.r.Read(p)
//	for i,v := range p{
//		index := strings.Index(old, string(v))
//		if index == -1{
//			p[i] = []byte(" ")[0]
//		}else{
//			p[i] = new_[index]
//		}
//	}
//	return n, e
//}
//
//func main() {
//	s := strings.NewReader("Lbh penpxrq gur pbqr!")
//	r := rot13Reader{s}
//	io.Copy(os.Stdout, &r)
//}


//package main
//
//import (
//	"fmt"
//	//"image"
//)
//
//func main() {
//	//m := image.NewRGBA(image.Rect(0, 0, 100, 100))
//	//fmt.Println(m.Bounds())
//	//fmt.Println(m.At(1, 1).RGBA())
//	var x, y int
//	fmt.Println(&x == &x, &x == &y, &x == nil)
//}


//package main
//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//	"io"
//)
//
//type IPAddr [4]byte
//
//func (ip IPAddr) String() string {
//	var ret [] string
//	for _, value := range ip{
//		ret = append(ret, strconv.Itoa(int(value)))
//	}
//	return strings.Join(ret, ".")
//}
//func main() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for name, ip := range hosts {
//		fmt.Printf("%v: %v\n", name, ip)
//	}
//}



//package main
//
//import "golang.org/x/tour/pic"
//import "image/color"
//import "image"
//
//type Image struct{}
//
//func (m *Image) ColorModel() color.Model {
//	return color.RGBAModel
//}
//func (m *Image) Bounds() image.Rectangle {
//	return image.Rect(0, 0, 100, 100)
//}
//func (m *Image) At(x, y int) color.Color{
//	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
//}
//
//func main() {
//	m := Image{}
//	pic.ShowImage(&m)
//}


//package main
//
//import "golang.org/x/tour/tree"
//import "fmt"
//
//// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
//func Walk(t *tree.Tree, ch chan int) {
//	sendvalue(t, ch)
//	close(ch)
//}
//func sendvalue(t *tree.Tree, ch chan int){
//	if t.Left != nil{
//		sendvalue(t.Left, ch)
//	}
//	ch <- t.Value
//	if t.Right != nil{
//		sendvalue(t.Right, ch)
//	}
//}
//
//// Same 检测树 t1 和 t2 是否含有相同的值。
//func Same(t1, t2 *tree.Tree) bool{
//	ch1, ch2 := make(chan int), make(chan int)
//	go Walk(t1, ch1)
//	go Walk(t2, ch2)
//	for i := range ch1{
//		j := <- ch2
//		if i != j{
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	fmt.Println(Same(tree.New(2), tree.New(1)))
//}


package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type History struct {
	m map[string] int
	s sync.Mutex
}

func (h History) get(key string) int{
	h.s.Lock()
	defer h.s.Unlock()
	return h.m[key]
}

func (h History) set(key string) {
	h.s.Lock()
	defer h.s.Unlock()
	h.m[key] += 1
}


// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// 下面并没有实现上面两种情况：
	defer i.Done()
	if depth <= 0 {
		return
	}
	if h.get(url) > 0 {
		return
	}

	h.set(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		i.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return
}

var (
	i sync.WaitGroup
	h = History{m: make(map[string] int)}
)
func main() {
	i.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	i.Wait()
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
