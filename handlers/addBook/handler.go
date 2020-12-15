package addBook

import (
	"challenge-SmartMEI/controller"
	. "challenge-SmartMEI/controller/dto"
	"encoding/json"

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
	var book AddBookInput
	err := json.Unmarshal(request.Body, &book)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": "invalid body"})
	}
	err = Validate(book)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": err.Error()})
	}
	bookAdded, err := c.ctl.AddBookToMyCollection(request.Params["userid"], book)
	if err != nil {
		return httping.InternalServerError("Error to add book")
	}
	return httping.OK(bookAdded)
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
