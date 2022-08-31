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
	country := router.FormValue("country")
	score, err1 := strconv.Atoi(router.FormValue("user_score"))
	latitude, err2 := strconv.ParseFloat(router.FormValue("latitude"), 64)
	longitude, err3 := strconv.ParseFloat(router.FormValue("longitude"), 64)
	
	if (err1 != nil || err2 != nil || err3 != nil){
		return router.JSON(http.StatusBadRequest, map[string]string{"status":"error"})
	}

	query := fmt.Sprintf(`INSERT INTO users(user_name, score, country, lat, lon) VALUES('%s', %d, '%s', %f, %f)`, userName, score, country, latitude, longitude)
	_, err := DB.Exec(query)
	if(err != nil){
		fmt.Println("Error executing DB query", err)
	}

	return router.JSON(http.StatusOK, map[string]int{userName:score})
}