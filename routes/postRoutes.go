package routes

import (
	// "fmt"
	// "math/rand"
	"net/http"
	// "os"
	"strconv"
	// "time"

	"github.com/labstack/echo/v4"
)

func RecordUser(router echo.Context) error{

	userName := router.FormValue("user_name")
	score, err := strconv.Atoi(router.FormValue("user_score"))

	if (err != nil){
		return router.JSON(http.StatusBadRequest, map[string]string{"status":"error"})
	}

	users[userName] = score

	return router.JSON(http.StatusOK, map[string]int{userName:score})
}