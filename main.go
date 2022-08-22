package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"portfolio/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init(){

	err := godotenv.Load()
	if(err != nil){
		panic("error loading env")
	}

	connectionString := "postgresql://root:12345@127.0.0.1:5800/root?sslmode=disable"

	conn, err := sql.Open("postgres",connectionString)
	if(err != nil){
		panic(err)
	}
	if err = conn.Ping(); err != nil{
		panic(err)
	}

	DB = conn
}

func main(){
	defer DB.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	  }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/record", func(echo echo.Context) error{
		return routes.RecordUser(echo, DB)
	})
	e.GET("/users", func(echo echo.Context) error{
		return routes.GetUsers(echo, DB)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HOSTPORT"))))
}