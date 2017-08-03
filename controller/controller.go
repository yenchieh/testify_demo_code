package controller

import (
	"net/http"

	"strconv"

	"fmt"

	"github.com/labstack/echo"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email" form:"email"`
}

var users []User

func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}

	user.ID = int64(len(users) + 1)
	users = append(users, user)

	return c.JSON(http.StatusOK, user)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func ChargeUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.FormValue("user_id"))

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	var user *User
	for _, u := range users {
		if u.ID == int64(userID) {
			user = &u
			break
		}
	}

	if user == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "The user is not exist",
		})
	}

	gateway := NewVisaGateway(user)

	if ChargeCustomer(gateway) {
		c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, true)
}
