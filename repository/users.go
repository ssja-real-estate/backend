package repository

import (
	"fmt"
	"os"

	"realstate/db"
	"realstate/models"

	ippanel "github.com/ippanel/go-rest-sdk"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	Update(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByUserName(username string) (user *models.User, err error)
	GetByMobile(mobile string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	SendSms(mobile string, veryfiycode string) (int64, error)
	Delete(id string) error
}

type usersRepository struct {
	c *mgo.Collection
}

func (r *usersRepository) SendSms(mobile string, veryfiycode string) (int64, error) {
	apiKey := os.Getenv("SMS_KEY")
	smsPatern := os.Getenv("SMS_PATTERN")
	smsSendNumber := os.Getenv("SMS_SENDNUMBER")
	sms := ippanel.New(apiKey)
	fmt.Println("++++++++++++++++++++++++++++++")
	fmt.Println(apiKey)
	fmt.Println("++++++++++++++++++++++++++++++")
	patternValues := map[string]string{
		"verification-code": veryfiycode}

	bulkid, err := sms.SendPattern(smsPatern, smsSendNumber, mobile, patternValues)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return bulkid, nil
}
func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{conn.DB().C(UsersCollection)}
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
