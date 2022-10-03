package controller

import (
	"encoding/xml"
	"try_xml/database"
	"try_xml/model"

	"github.com/gofiber/fiber/v2"
)

func GetAccount(c *fiber.Ctx) error {

	dbModel := []model.User{}

	database.DBConn.Debug().Table("users").Find(&dbModel)

	responseBody := []model.ResponseBody{}

	for _, user := range dbModel {
		ul := UserList(user)
		responseBody = append(responseBody, ul)
	}

	m, err := xml.MarshalIndent(responseBody, "", "")

	if err != nil {
		panic(err)
	}
	return c.Send(m)
}

func UserList(u model.User) model.ResponseBody {
	return model.ResponseBody{
		UserID:   u.UserID,
		Fullname: u.Fullname,
		Username: u.Username,
		Password: u.Password,
	}
}
