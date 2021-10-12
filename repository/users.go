package repository

import (
	"fmt"
	"math/rand"
	"realstate/db"
	"realstate/models"
	"time"

	ippanel "github.com/ippanel/go-rest-sdk"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	Verify(mobile string) (int64, error)
	Update(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByUserName(username string) (user *models.User, err error)
	GetByMobile(mobile string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Delete(id string) error
}

type usersRepository struct {
	c *mgo.Collection
}

func sendsms(mobile string, veryfiycode int) (int64, error) {
	apiKey := "xa6nMhNisMZP92-0giaTIJeFQz0VIm6o7UQTbYK2L7Q="
	sms := ippanel.New(apiKey)
	mobiles := make([]string, 1)
	mobiles = append(mobiles, "989147256898")
	fmt.Println(mobiles)
	bulkid, err := sms.Send("+983000505", []string{"989147256898"}, "پیام تستی ارسال شد")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return bulkid, nil
}
func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{conn.DB().C(UsersCollection)}
}

func (r *usersRepository) Verify(mobile string) (int64, error) {
	var user models.User
	user.Mobile = mobile
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.VerifyCode = rand.Int63n(99000)

	sendsms(mobile, int(user.VerifyCode))
	return user.VerifyCode, r.c.Insert()
}
func (r *usersRepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *usersRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *usersRepository) GetByMobile(mobile string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"mobile": mobile}).One(&user)
	return user, err
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) GetByUserName(uasername string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"user_name": uasername}).One(&user)
	return user, err
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&users)
	if users == nil {
		users = make([]*models.User, 0)

	}
	return users, err
}

func (r *usersRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
