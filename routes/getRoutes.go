package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type user struct{
	Id int
	Name string `json:"name"`
	Score int `json:"score"`
}

func GetUsers(router echo.Context, DB *sql.DB) error {
	query := `SELECT * FROM scors`
	rows, err := DB.Query(query)
	if(err != nil){
		fmt.Println("Error executing DB query")
	}
	defer rows.Close()

	users := make([]user, 0)

	for rows.Next(){
		user := user{}
		err := rows.Scan(&user.Id, &user.Name, &user.Score)
		if(err != nil){
			panic(err)
		}

		users = append(users, user)
	}

	fmt.Println(users)

	return router.JSON(http.StatusOK, users)
}