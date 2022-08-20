package routes

import (
	// "fmt"
	// "math/rand"
	"net/http"
	// "os"
	// "strconv"
	// "time"

	"github.com/labstack/echo/v4"
)

var users = make(map[string]int)

func GetUsers(router echo.Context) error {

	return router.JSON(http.StatusOK, users)
}