package getUser

import (
	"challenge-SmartMEI/controller"

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
	user, err := c.ctl.GetUserDetails(request.Params["userid"])
	if err != nil {
		return httping.InternalServerError("Error to get user")
	}
	return httping.OK(user)
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
