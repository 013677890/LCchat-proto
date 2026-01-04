package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker! Current time: %s", time.Now().Format(time.RFC3339))
	})

	fmt.Println("Server starting on :8080...")
	// 这是一个阻塞调用，会防止 main 函数退出，从而保持容器运行
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
