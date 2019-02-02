package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markstgodard/linkerd2bug/baz/hello"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bar handler, calling baz:9002..!")

	// call baz rpc
	rpcClient := hello.NewBazServiceProtobufClient("http://baz:9002", &http.Client{})

	req := &hello.SayHelloRequest{
		Msg: "hello",
	}
	resp, err := rpcClient.Hello(r.Context(), req)
	if err != nil {
		http.Error(w, "error occurred", 500)
		return
	}

	w.Write([]byte(resp.GetMsg()))
}

func main() {
	http.HandleFunc("/", barHandler)

	fmt.Println("Listening..")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
