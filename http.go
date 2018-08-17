//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//type Hello struct{}
//
//func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello!")
//}
//
//func main() {
//	var h Hello
//	err := http.ListenAndServe("localhost:4000", h)
//	if err != nil {
//		log.Fatal(err)
//	}
//}


package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

type Hello struct{}

func (v String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, v)
}

func (v Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, v)
}

func (v Struct) String() string{
	return fmt.Sprintf("{Greeting:%v, Punct:%v, Who:%v}", v.Greeting, v.Punct, v.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}