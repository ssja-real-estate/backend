package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/security"
	"realstate/util"
	"strconv"
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
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
	VeryfiyMobile(ctx *fiber.Ctx) error
	Changepassword(ctx *fiber.Ctx) error
}

type authController struct {
	usersRepo repository.UsersRepository
}

func NewAuthController(usersRepo repository.UsersRepository) AuthController {
	return &authController{usersRepo}
}

// Verify ... Verify user
// @Summary Verify User
// @Description Signin
// @Tags User
// @Success 200 {object} models.User
// @Param mobile path string true "Item mobile"
// @Param verify_code path string true "Item verify_code"
// @Failure 404 {object} object
// @Router /verify [Get]
func (c *authController) VeryfiyMobile(ctx *fiber.Ctx) error {

	mobile := ctx.Query("mobile")
	veryfiyCode := ctx.Query("verify_code")
	exists, err := c.usersRepo.GetByMobile(mobile)
	if err != nil || exists.Mobile != mobile {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNotMobile)
	}

	if exists.VerifyCode != veryfiyCode {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrVeryfiyCodeNotValid)
	}

	exists.Verify = true

	err = c.usersRepo.Update(exists)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrSignup)
	}
	token, err := security.NewToken(exists.Id.Hex())
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(

			fiber.Map{
				"user":  exists,
				"token": fmt.Sprintf("Bearer %s", token),
			})

}
func (c *authController) SignUp(ctx *fiber.Ctx) error {
	var newUser models.User
	err := ctx.BodyParser(&newUser)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	exists, err := c.usersRepo.GetByMobile(newUser.Mobile)

	if err == mgo.ErrNotFound {
		if strings.TrimSpace(newUser.Mobile) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyMobile))
		}
		newUser.Password, err = security.EncryptPassword(newUser.Password)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		newUser.CreatedAt = time.Now()
		if newUser.Role == 0 {
			newUser.Role = 3
		}
		newUser.UpdatedAt = newUser.CreatedAt
		newUser.Verify = false
		newUser.VerifyCode = strconv.FormatInt(int64(rand.Intn(89000)+10000), 10)

		newUser.Id = bson.NewObjectId()
		err = c.usersRepo.Save(&newUser)

		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}

		_, err = c.usersRepo.SendSms(newUser.Mobile, newUser.VerifyCode)
		if err != nil {
			ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}

		return ctx.Status(http.StatusOK).JSON(util.SuccessSendSms)

	}
	if !exists.Verify {
		exists.Password, err = security.EncryptPassword(newUser.Password)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		exists.UpdatedAt = time.Now()
		exists.Role = 3
		exists.VerifyCode = strconv.FormatInt(int64(rand.Intn(89000)+10000), 10)

		err = c.usersRepo.Update(exists)

		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}

		_, err = c.usersRepo.SendSms(exists.Mobile, exists.VerifyCode)
		if err != nil {
			ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}

		return ctx.Status(http.StatusOK).JSON(util.SuccessSendSms)
	}

	if exists != nil {
		err = util.ErrMobileAlreadyExists
	}

	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

// Signin ... Login in web api
// @Summary Signin
// @Description Signin
// @Tags User
// @Success 200 {object} models.User
// @Failure 404 {object} object
// @Router /signin [post]
func (c *authController) SignIn(ctx *fiber.Ctx) error {

	var input models.User
	var user *models.User
	err := ctx.BodyParser(&input)
    
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
    
	user , err = c.usersRepo.GetByMobile(input.Mobile)
   
	if !user.Verify {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrNotVerifyed))
	}

	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Mobile, err.Error())
		return ctx.Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrInvalidCredentials))
	}

	err = security.VerifyPassword(user.Password, input.Password)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Name, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Name, err.Error())
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
// @Router /user/id [get]
func (c *authController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNotFound)
	}

	user, err := c.usersRepo.GetById(id)
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
// @Router /users [get]
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
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNotFound)
	}
	var update models.User
	err := ctx.BodyParser(&update)
	if update.Role < 1 && update.Role > 3 {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrBadRole))
	}
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	exists, err := c.usersRepo.GetByMobile(update.Mobile)
	if err == mgo.ErrNotFound || exists.Id.Hex() == id {
		user, err := c.usersRepo.GetById(id)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		if update.Name != "" {
			user.Name = update.Name
		}
		if update.Role != 0 {
			user.Role = update.Role
		}

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
		err = util.ErrMobileAlreadyExists
	}

	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

func (c *authController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNotFound)
	}
	err := c.usersRepo.Delete(id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	ctx.Set("Entity", id)
	return ctx.
		Status(http.StatusOK).
		JSON(util.SuccessDelete)
	// return ctx.SendStatus(http.StatusNoContent)
}

// ChangePassword ... Change login password
// @Summary change login password
// @Description change password for login
// @Tags User
// @Success 200 {string} success change password
// @Param current_password path string true "Item current_password"
// @Param new_password path string true "Item new_password"
// @Param confirm_password path string true "Item confirm_password"
// @Failure 404 {object} object
// @Router /changepassword [post]
func (c *authController) Changepassword(ctx *fiber.Ctx) error {
	currentpassword := ctx.Query("current_password")
	newpassowrd := ctx.Query("new_password")
	confirm_password := ctx.Query("confirm_password")
	fmt.Println(currentpassword, "-", newpassowrd, "-", confirm_password)

	if newpassowrd != confirm_password {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNotCompatablePassword)
	}
	id, err := security.GetUserByToken(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(util.ErrInvalidAuthToken)
	}
	fmt.Println(id)

	user, err := c.usersRepo.GetById(id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrInvalidAuthToken)
	}
	fmt.Print(user.Name)

	err = security.VerifyPassword(user.Password, currentpassword)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrNoMatchPassword)
	}
	fmt.Println(user.Password)

	user.Password, err = security.EncryptPassword(newpassowrd)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	fmt.Println("user update")
	err = c.usersRepo.Update(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErroNotUserUpdate)
	}
	fmt.Println("user update done")

	return ctx.Status(http.StatusOK).JSON(util.SuccessUpdate)

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
