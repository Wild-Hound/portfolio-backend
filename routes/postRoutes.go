package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

func RecordUser(router echo.Context, DB *sql.DB) error{
	userName := router.FormValue("user_name")
	score, err := strconv.Atoi(router.FormValue("user_score"))
	if (err != nil){
		return router.JSON(http.StatusBadRequest, map[string]string{"status":"error"})
	}

	query := fmt.Sprintf(`INSERT INTO scors(name, score) VALUES('%s', %d)`, userName, score)
	_, err = DB.Exec(query)
	if(err != nil){
		fmt.Println("Error executing DB query", err)
	}

	return router.JSON(http.StatusOK, map[string]int{userName:score})
}