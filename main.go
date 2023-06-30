package main

import (
	"fmt"
	"net/http"

	"github.com/justhands13/learn_go/server"
	r "github.com/justhands13/reverse_int/v2"
	s "github.com/justhands13/reverse_string/v2"
	g "github.com/justhands13/web_scrapper/v2/greeting"
	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	fmt.Println("hello world")
	//func from different module and package
	fmt.Println(g.Greet())

	fmt.Println(tinytime.Now())

	//from different packages
	fmt.Println(s.ReverseString("string desrever"))
	fmt.Println(r.ReverseInt(12938))

	srv := server.NewServer()

	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("service is running"))
	})

	http.ListenAndServe(":8080", srv)

}
