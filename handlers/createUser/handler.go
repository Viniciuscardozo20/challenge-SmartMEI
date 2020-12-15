package createUser

import (
	"challenge-SmartMEI/controller"
	. "challenge-SmartMEI/controller/dto"
	"encoding/json"
	"fmt"

	httping "github.com/ednailson/httping-go"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	ctl controller.Controller
}

func NewHandler(ctl controller.Controller) *Handler {
	return &Handler{ctl: ctl}
}

func (c *Handler) Handle(request httping.HttpRequest) httping.IResponse {
	var user CreateUserInput
	err := json.Unmarshal(request.Body, &user)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": "invalid body"})
	}
	err = Validate(user)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": err.Error()})
	}
	userCreated, err := c.ctl.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		return httping.InternalServerError("Error to create user")
	}
	return httping.OK(userCreated)
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
