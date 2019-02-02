package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markstgodard/linkerd2bug/baz/hello"
)

func main() {
	s := &hello.Service{}
	rpcHandler := hello.NewBazServiceServer(s, nil)

	fmt.Println("[baz] Listening...")
	log.Fatal(http.ListenAndServe(":9000", rpcHandler))
}
