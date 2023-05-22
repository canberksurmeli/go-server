package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const port = 8080

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello, from client!\n")
	io.WriteString(w, "Hello, from server! It is "+time.Now().Format(time.RFC850)+"\n")
}

func startHttpServer() {
	http.HandleFunc("/", getHello)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("ListenAndServe: %s", err)
		os.Exit(1)
	}
}
