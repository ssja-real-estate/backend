package controllers

import (
	"net/http"
	"realstate/models"
	"realstate/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2/bson"
)

type FormRegisterController interface {
	CreateFormRegister(ctx *fiber.Ctx) error
}

type formRegisterController struct {
	formregister repository.FormRegisterRepository
}

func NewFormRegisterController(formregisterrepo repository.FormRegisterRepository) FormRegisterController {
	return &formRegisterController{formregisterrepo}
}

func (r *formRegisterController) CreateFormRegister(ctx *fiber.Ctx) error {
	 var formregister models.FormRegister
	 err:=ctx.BodyParser(&formregister)
	 if err!=nil {
		 return ctx.Status(http.StatusBadRequest).JSON(err)
	 }
	 formregister.Id=bson.NewObjectId()
	 formregister.CreatedAt=time.Now()
	 formregister.UpdatedAt=time.Now()
	 formregister.Accept=false
	 
    
	 return nil
}
