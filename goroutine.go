//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	go say("world")
//	say("hello")
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func sum(a []int, c chan int) {
//	sum := 0
//	for _, v := range a {
//		sum += v
//	}
//	c <- sum // 将和送入 c
//}
//
//func main() {
//	a := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(a[:len(a)/2], c)
//	x := <-c
//	time.Sleep(time.Second)
//	go sum(a[len(a)/2:], c)
//	// 从 c 中获取
//	var y int
//	go get(c, &y)
//	time.Sleep(time.Second * 2)
//	fmt.Println(x, y, x+y)
//}
//func get(ch chan int, y *int){
//	*y = <-ch
//	fmt.Println(y)
//	fmt.Println(666)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch := make(chan int, 2)
//	ch <- 1
//	go get(ch)
//	go get(ch)
//	go get(ch)
//	ch <- 2
//	time.Sleep(time.Second)
//}
//
//func get(ch chan int){
//	fmt.Println(<-ch)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func main() {
//	c := make(chan int, 10)
//	go fibonacci(cap(c), c)
//	for i := range c {
//		fmt.Println(i)
//	}
//}

//package main
//
//import "fmt"
//
//func fibonacci(c, quit chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			fmt.Println(111)
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		default:
//			fmt.Println("wait")
//		}
//	}
//}
//
//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	go func() {
//		for i := 0; i < 2; i++ {
//			fmt.Println(<-c)
//		}
//		quit <- 0
//	}()
//	fibonacci(c, quit)
//}


//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	tick := time.Tick(100 * time.Millisecond)
//	boom := time.After(500 * time.Millisecond)
//	for {
//		select {
//		case <-tick:
//			fmt.Println("tick.")
//		case <-boom:
//			fmt.Println("BOOM!")
//			return
//		default:
//			fmt.Println("    .")
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
//}


//1. 实现 Walk 函数。
//
//2. 测试 Walk 函数。
//
//函数 tree.New(k) 构造了一个随机结构的二叉树，保存了值 k，2k，3k，...，10k。 创建一个新的 channel ch 并且对其进行步进：
//
//go Walk(tree.New(1), ch)
//然后从 channel 中读取并且打印 10 个值。应当是值 1，2，3，...，10。
//
//3. 用 Walk 实现 Same 函数来检测是否 t1 和 t2 存储了相同的值。
//
//4. 测试 Same 函数。
//
//Same(tree.New(1), tr

//package main
//
//import (
//	"golang.org/x/tour/tree"
//	"fmt"
//)
//
//
//// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
//func Walk(t *tree.Tree, ch chan int){
//	SendValue(t, ch)
//	close(ch)
//}
//
//func SendValue(t *tree.Tree,ch chan int){
//	if t != nil{
//		SendValue(t.Left, ch)
//		ch <- t.Value
//		SendValue(t.Right, ch)
//	}
//}
//
//// Same 检测树 t1 和 t2 是否含有相同的值。
//func Same(t1, t2 *tree.Tree) bool{
//	cha := make(chan int)
//	chb := make(chan int)
//	go Walk(t1, cha)
//	go Walk(t2, chb)
//	for i := range cha{
//		if i != <- chb{
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	ret:=Same(tree.New(1), tree.New(1))
//	fmt.Println(ret)
//	ret =Same(tree.New(2), tree.New(2))
//	fmt.Println(ret)
//
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//// SafeCounter 的并发使用是安全的。
//type SafeCounter struct {
//	v   map[string]int
//	mux sync.Mutex
//}
//
//// Inc 增加给定 key 的计数器的值。
//func (c *SafeCounter) Inc(key string) {
//	c.mux.Lock()
//	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
//	c.v[key]++
//	c.mux.Unlock()
//}
//
//// Value 返回给定 key 的计数器的当前值。
//func (c *SafeCounter) Value(key string) int {
//	c.mux.Lock()
//	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
//	defer c.mux.Unlock()
//	return c.v[key]
//}
//
//func main() {
//	c := SafeCounter{v: make(map[string]int)}
//	for i := 0; i < 1000; i++ {
//		go c.Inc("somekey")
//	}
//
//	time.Sleep(time.Second)
//	fmt.Println(c.Value("somekey"))
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

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
        // 下面并没有实现上面两种情况：
	defer i.Done()
	if depth <= 0 {
		return
	}
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

func main() {
	i.Add(1)
	Crawl("http://golang.org/", 4, fetcher)
	i.Wait()
}

var (
fetched = Fetched{make(map[string]bool), sync.Mutex{}}
i sync.WaitGroup
)

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	fetched.mux.Lock()
	defer fetched.mux.Unlock()
	if fetched.find[url]{
		return "", nil, fmt.Errorf("have found: %s", url)
	}
	fetched.find[url] = true
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type Fetched struct{
	find map[string] bool
	mux sync.Mutex
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

