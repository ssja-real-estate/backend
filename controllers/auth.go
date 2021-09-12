package controllers

import (
	"fmt"
	"log"
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/security"
	"realstate/util"
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/asaskevich/govalidator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	GetUsers(ctx *fiber.Ctx) error
	PutUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type authController struct {
	usersRepo repository.UsersRepository
}

func NewAuthController(usersRepo repository.UsersRepository) AuthController {
	return &authController{usersRepo}
}

func (c *authController) SignUp(ctx *fiber.Ctx) error {
	var newUser models.User
	err := ctx.BodyParser(&newUser)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	// newUser.Username = newUser.Username
	// if !govalidator.IsEmail(newUser.Email) {
	// 	return ctx.
	// 		Status(http.StatusBadRequest).
	// 		JSON(util.NewJError(util.ErrInvalidEmail))
	// }
	exists, err := c.usersRepo.GetByUserName(newUser.UserName)
	if err == mgo.ErrNotFound {
		if strings.TrimSpace(newUser.UserName) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyName))
		}
		newUser.Password, err = security.EncryptPassword(newUser.Password)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = newUser.CreatedAt
		newUser.Id = bson.NewObjectId()
		err = c.usersRepo.Save(&newUser)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		return ctx.
			Status(http.StatusCreated).
			JSON(newUser)
	}

	if exists != nil {
		err = util.ErrEmailAlreadyExists
	}

	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

func (c *authController) SignIn(ctx *fiber.Ctx) error {
	var input models.User
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	user, err := c.usersRepo.GetByUserName(input.UserName)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.UserName, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	err = security.VerifyPassword(user.Password, input.Password)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.UserName, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.UserName, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"user":  user,
			"token": fmt.Sprintf("Bearer %s", token),
		})
}

// GetUser ... Get user by id
// @Summary Get user by id
// @Description get user by id
// @Tags User
// @Success 200 {object} models.User
// @Param id path int true "Item ID"
// @Failure 404 {object} object
// @Router /id [get]
func (c *authController) GetUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	user, err := c.usersRepo.GetById(payload.Id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(user)
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags User
// @Success 200 {array} models.User
// @Failure 404 {object} object
// @Router / [get]
func (c *authController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.usersRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(users)
}

func (c *authController) PutUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	var update models.User
	err = ctx.BodyParser(&update)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	if !govalidator.IsEmail(update.UserName) {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrInvalidEmail))
	}
	exists, err := c.usersRepo.GetByUserName(update.UserName)
	if err == mgo.ErrNotFound || exists.Id.Hex() == payload.Id {
		user, err := c.usersRepo.GetById(payload.Id)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		user.UserName = update.UserName
		user.UpdatedAt = time.Now()
		err = c.usersRepo.Update(user)
		if err != nil {
			return ctx.
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
		}
		return ctx.
			Status(http.StatusOK).
			JSON(user)
	}

	if exists != nil {
		err = util.ErrEmailAlreadyExists
	}

	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

func (c *authController) DeleteUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	err = c.usersRepo.Delete(payload.Id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	ctx.Set("Entity", payload.Id)
	return ctx.
		Status(http.StatusOK).
		JSON(util.ErrEmailAlreadyExists)
	// return ctx.SendStatus(http.StatusNoContent)
}

func AuthRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return nil, util.ErrUnauthorized
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		return nil, err
	}
	if payload.Id != id || payload.Issuer != id {
		return nil, util.ErrUnauthorized
	}
	return payload, nil
}
