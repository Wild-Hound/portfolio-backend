package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type user struct{
	ID int
	User_name string `json:"name"`
	Score int `json:"score"`
	Country string `json:"country"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Created_at string `json:"created_at"`
}

func GetUsers(router echo.Context, DB *sql.DB) error {
	query := `SELECT * FROM users`
	rows, err := DB.Query(query)
	if(err != nil){
		fmt.Println("Error executing DB query")
	}
	defer rows.Close()

	users := make([]user, 0)

	for rows.Next(){
		user := user{}
		err := rows.Scan(&user.ID, &user.User_name, &user.Score, &user.Country, &user.Lat, &user.Lon, &user.Created_at)
		if(err != nil){
			panic(err)
		}

		users = append(users, user)
	}

	return router.JSON(http.StatusOK, users)
}

func GetUser(router echo.Context, DB *sql.DB, userId int) error{
	query := fmt.Sprintf(`SELECT * FROM users WHERE id=%d`, userId)
	rows, err := DB.Query(query)
	if(err != nil){
		fmt.Println("Error executing DB query")
	}
	defer rows.Close()

	users := make([]user, 0)

	for rows.Next(){
		user := user{}
		err := rows.Scan(&user.ID, &user.User_name, &user.Score, &user.Country, &user.Lat, &user.Lon, &user.Created_at)
		if(err != nil){
			panic(err)
		}

		users = append(users, user)
	}

	return router.JSON(http.StatusOK, users)
}