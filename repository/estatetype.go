package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const estatetypeCollection = "estatetype"

type EstateTypeRepository interface {
	SaveEstateType(estatetype *models.EstateType) error
	UpdateEstateType(estatetype *models.EstateType, estateid primitive.ObjectID) error
	GetEstateTypeById(id primitive.ObjectID) (estatetype *models.EstateType, err error)
	GetEstateTypeByName(name string) (estatetype *models.EstateType, err error)
	GetEstateTypeAll() ([]models.EstateType, error)
	DeleteEstateType(id primitive.ObjectID) error
}

type estateTypeRepository struct {
	c *mongo.Collection
}

func NewEstateTypesRepository(Db *mongo.Client) EstateTypeRepository {
	return &estateTypeRepository{db.GetCollection(Db, estatetypeCollection)}
}

func (r *estateTypeRepository) SaveEstateType(estatetype *models.EstateType) error {
	_, err := r.c.InsertOne(context.TODO(), estatetype)
	return err
}
func (r *estateTypeRepository) UpdateEstateType(estatetype *models.EstateType, estatetypeid primitive.ObjectID) error {
	filter := bson.M{"_id": estatetypeid}

	// فقط فیلدهایی را برای آپدیت ارسال می‌کنیم که قابل تغییر هستند
	// این کار از آپدیت ناخواسته فیلدهایی مانند _id یا CreatedAt جلوگیری می‌کند
	updatePayload := bson.M{
		"name":  estatetype.Name,
		"order": estatetype.Order,
		// هر فیلد دیگری که باید آپدیت شود را اینجا اضافه کنید
	}

	update := bson.M{"$set": updatePayload}

	result, err := r.c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments // یا یک خطای سفارشی
	}

	return nil
}
func (r *estateTypeRepository) GetEstateTypeById(id primitive.ObjectID) (estatetype *models.EstateType, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&estatetype)
	return estatetype, err
}

func (r *estateTypeRepository) GetEstateTypeByName(name string) (estatetype *models.EstateType, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"name": name}).Decode(&estatetype)
	return estatetype, err
}
func (r *estateTypeRepository) GetEstateTypeAll() ([]models.EstateType, error) {

	var estattypes []models.EstateType

	// مرحله ۱: ساخت یک آپشن برای Find
	findOptions := options.Find()

	// مرحله ۲: تنظیم مرتب‌سازی بر اساس فیلد "order" به صورت صعودی
	findOptions.SetSort(bson.D{{"order", 1}}) // عدد 1 برای ترتیب صعودی (A-Z, 1-10)

	// مرحله ۳: ارسال آپشن به عنوان پارامتر سوم به متد Find
	result, err := r.c.Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return make([]models.EstateType, 0), err
	}

	defer result.Close(context.TODO())

	for result.Next(context.TODO()) {
		var estatetype models.EstateType
		if err = result.Decode(&estatetype); err != nil {
			return make([]models.EstateType, 0), err
		}
		estattypes = append(estattypes, estatetype)
	}

	if estattypes == nil {
		estattypes = make([]models.EstateType, 0)
	}

	return estattypes, err
}
func (r *estateTypeRepository) DeleteEstateType(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
