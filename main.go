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

type user struct{
	id int
	name string
	score int
}

func init(){

	err := godotenv.Load()
	if(err != nil){
		fmt.Println("error loading env")
	}

	connectionString := "postgresql://root:12345@127.0.0.1:5800/root?sslmode=disable"

	conn, err := sql.Open("postgres",connectionString)
	if(err != nil){
		panic(err)
	}
	if err = conn.Ping(); err != nil{
		panic(err)
	}
	defer conn.Close()

	query := `SELECT * FROM "scors"`
	rows, err := conn.Query(query)
	if(err != nil){
		fmt.Println("Error executing DB query")
	}
	defer rows.Close()

	users := make([]user, 0)

	for rows.Next(){
		user := user{}
		err := rows.Scan(&user.id, &user.name, &user.score)
		if(err != nil){
			panic(err)
		}

		users = append(users, user)
		for _, ux := range(users){
			fmt.Println(ux)
		}
	}
	

}

func main(){
	
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	  }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/record", routes.RecordUser)
	e.GET("/users", routes.GetUsers)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HOSTPORT"))))
}