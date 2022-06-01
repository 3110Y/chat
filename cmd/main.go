package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	log.Print("URL.Path5 = %q\n", r.URL.Path)
}
