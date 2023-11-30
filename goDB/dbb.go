package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)

func conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/Students")
	if err != nil {
		fmt.Println("error connecting to database", err)
		return nil, err
	}
	return db, nil
}

type data struct {
	Enroll int    `json:"enrollmentNumber"`
	Name   string `json:"name"`
}

func main() {
	db, err := conn()
	if err != nil {
		fmt.Println("Couldn't initialize db")
		return
	}

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		var d data

		reqBody, _ := io.ReadAll(r.Body)

		err = json.Unmarshal(reqBody, &d)
		
		if err != nil {
			fmt.Println("Unmarshall error", err)
		}

		_, err := db.Exec(`INSERT INTO students (enrollment_number,name) values (?,?)`, d.Enroll, d.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Student added successfully!"))

	})

	fmt.Println("server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server encountered an arror: #{err}\n")
	}
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// )
// func conn() (*sql.DB, error ){
// 	db, err := sql.Open("mysql", "")
// 	if err != nil {}
// }
// func eventHandler(w http.ResponseWriter, r *http.Request){
// 	if r.Method != "GET" {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Welcome to goFr's event!"))
// }
// func main() {
// 	fmt.Println("heyoo")
// 	http.HandleFunc("/event", eventHandler)

// 	fmt.Println("server is runnning on http://localhost:8080")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Printf("server encountered error %s\n", err)
// 	}

// }
