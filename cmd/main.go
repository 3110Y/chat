package main

import (
	"fmt"
	"github.com/3110Y/chat/pkg/user"
	"github.com/3110Y/chat/pkg/user/usermigration"
	"github.com/3110Y/migrator"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseDB := os.Getenv("DATABASE_DB")
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			databaseUser,
			databasePassword,
			databaseHost,
			databasePort,
			databaseDB,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	m := migrator.NewMigrator(migrator.NewDbSqlx(db))
	err = m.AddMigration(usermigration.Migration20220606())
	if err != nil {
		log.Fatalln(err)
	}
	err = m.MigrationUp()
	if err != nil {
		log.Fatalln(err)
	}
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userRest := user.NewRest(userService)
	r.POST("/user", userRest.Create)
	r.PUT("/user", userRest.Update)
	r.DELETE("/user", userRest.Delete)
	r.GET("/user/:id", userRest.Item)
	r.GET("/user", userRest.List)
	log.Fatal(r.Run(":3000"))
}
