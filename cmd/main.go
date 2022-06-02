package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Fatal(r.Run(":3000"))
}
