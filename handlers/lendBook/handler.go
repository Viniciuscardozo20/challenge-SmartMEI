package lendBook

import (
	"challenge-SmartMEI/controller"
	. "challenge-SmartMEI/controller/dto"
	"encoding/json"
	"strconv"

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
	var lendBook LendBookInput
	err := json.Unmarshal(request.Body, &lendBook)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": "invalid body"})
	}
	err = Validate(lendBook)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": err.Error()})
	}
	v, err := strconv.Atoi(request.Params["userid"])
	if err != nil {
		return httping.BadRequest(map[string]string{"body": err.Error()})
	}
	lendedBook, err := c.ctl.LendBook(v, lendBook)
	if err != nil {
		return httping.InternalServerError("Error to lend book")
	}
	return httping.OK(lendedBook)
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
