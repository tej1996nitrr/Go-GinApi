package controller

import (
	"GoGinWebServices/entity"
	service "GoGinWebServices/services"
	"GoGinWebServices/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}
type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func new(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil

}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service}
}
