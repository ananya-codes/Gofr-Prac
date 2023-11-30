package main

import(
	"fmt"
	"net/http"
)
func eventHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to goFr's event!"))
}
func main() {
	fmt.Println("heyoo")
	http.HandleFunc("/event", eventHandler)

	fmt.Println("server is runnning on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server encountered error %s\n", err)
	}

}
