package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	// "io/ioutil"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type locationResponse struct{
	Status string
	Country string
	CountryCode string
	Region string
	RegionName string
	City string
	Zip string
	Lat float64
	Lon float64
	Timezone string
	Isp string
	Org string
	As string
	Query string
}

func RecordUser(router echo.Context, DB *sql.DB) error{
	userName := router.FormValue("user_name")
	country := router.FormValue("country")
	score, err1 := strconv.Atoi(router.FormValue("user_score"))
	latitude, err2 := strconv.ParseFloat(router.FormValue("latitude"), 64)
	longitude, err3 := strconv.ParseFloat(router.FormValue("longitude"), 64)

	if(country == "" || err2 != nil || err3 != nil){
		IPAddress := router.Request().Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = router.Request().Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = router.Request().RemoteAddr
		}
		var locationData locationResponse
		err := GetUserLoc(strings.Split(IPAddress, ":")[0], &locationData) //strings.Split(IPAddress, ":")[0]
		if err!= nil{
			fmt.Println(err)
		}
		country = locationData.Country
		latitude = locationData.Lat
		longitude = locationData.Lon
	}
	
	if (err1 != nil){
		return router.JSON(http.StatusBadRequest, map[string]string{"status":"error"})
	}

	query := fmt.Sprintf(`INSERT INTO users(user_name, score, country, lat, lon) VALUES('%s', %d, '%s', %f, %f)`, userName, score, country, latitude, longitude)
	_, err := DB.Exec(query)
	if(err != nil){
		fmt.Println("Error executing DB query", err)
	}

	return router.JSON(http.StatusOK, map[string]int{userName:score})
}

func GetUserLoc(ipAdd string, target interface{}) error {
	res, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ipAdd))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}