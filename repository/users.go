package repository

import (
	"context"
	"fmt"
	"os"
	"realstate/db"
	"realstate/models"

	ippanel "github.com/ippanel/go-rest-sdk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	Update(user *models.User) error
	GetById(id primitive.ObjectID) (user *models.User, err error)
	GetByUserName(username string) (user *models.User, err error)
	GetByMobile(mobile string) (user *models.User, err error)
	GetAll() (users []models.User, err error)
	SendSms(mobile string, veryfiycode string) (int64, error)
	Delete(id primitive.ObjectID) error
}

type usersRepository struct {
	c *mongo.Collection
}

func (r *usersRepository) SendSms(mobile string, veryfiycode string) (int64, error) {

	apiKey := os.Getenv("API_KEY")

	sms := ippanel.New(apiKey)
	patternValues := map[string]string{
		"verification-code": veryfiycode}

	bulkid, err := sms.SendPattern("g0eepccptg", "+983000505", mobile, patternValues)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return bulkid, nil
}
func NewUsersRepository(DB *mongo.Client) UsersRepository {

	return &usersRepository{db.GetCollection(DB, UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {

	_, err := r.c.InsertOne(context.TODO(), &user)
	return err
}

func (r *usersRepository) Update(user *models.User) error {
	_, err := r.c.UpdateOne(context.TODO(), bson.M{"_id": user.Id}, &user)
	return err
}

func (r *usersRepository) GetByMobile(mobile string) (user *models.User, err error) {

	err = r.c.FindOne(context.TODO(), bson.M{"mobile": mobile}).Decode(&user)
	return user, err
}

func (r *usersRepository) GetById(id primitive.ObjectID) (user *models.User, err error) {

	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *usersRepository) GetByUserName(uasername string) (user *models.User, err error) {

	err = r.c.FindOne(context.TODO(), bson.M{"user_name": uasername}).Decode(&user)
	return user, err
}

func (r *usersRepository) GetAll() (users []models.User, err error) {

	result, err := r.c.Find(context.TODO(), bson.M{})
	if err != nil {
		return make([]models.User, 0), err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var singleUser models.User
		if err = result.Decode(&singleUser); err != nil {
			return make([]models.User, 0), err
		}

		users = append(users, singleUser)
	}
	if users == nil {
		users = make([]models.User, 0)

	}
	return users, err
}

func (r *usersRepository) Delete(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err

}
