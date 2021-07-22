package controllers

import (
	"net/http"
	"project/lib/database"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id ngga valid bos!",
		})
	}
	user, err := database.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "alhamdulillah",
		"user":    user,
	})

}
