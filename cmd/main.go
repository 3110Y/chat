package main

import (
	"fmt"
	"log"
	"net/http"
)

var a = 3

func main() {
	//databaseUser := os.Getenv("DATABASE_USER")
	//databasePassword := os.Getenv("DATABASE_PASSWORD")
	//databaseHost := os.Getenv("DATABASE_HOST")
	//databasePort := os.Getenv("DATABASE_PORT")
	//databaseDB := os.Getenv("DATABASE_DB")
	//_, err := sqlx.Connect(
	//	"postgres",
	//	fmt.Sprintf(
	//		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
	//		databaseUser,
	//		databasePassword,
	//		databaseHost,
	//		databasePort,
	//		databaseDB,
	//	),
	//)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	a++
	fmt.Println(a)
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handlerTwo)
	http.HandleFunc("/about$", handlerThree)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	a++
	fmt.Println(a)
	fmt.Fprintf(w, "URL.Path = %q 404 handler\n", r.URL.Path)
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q 200 handlerTwo\n", r.URL.Path)
}

func handlerThree(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q 200 handlerThree\n", r.URL.Path)
}
