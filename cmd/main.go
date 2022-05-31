package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	appPort := os.Getenv("APP_PORT")
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
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe("localhost:"+appPort, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hi!"))
}
