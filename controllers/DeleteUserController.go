package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ID nya salah BOS!",
		})
	}
	user, err := database.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data has been DELETED",
		"data":    user,
	})
}
