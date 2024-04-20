package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SliderRepository interface {
	CreateSlider(models.Slider) error
	GetSliders() ([]models.Slider, error)
	DeleteSlider(primitive.ObjectID) error
	Get(primitive.ObjectID) (*models.Slider, error)
}

const slidercollection = "slider"

type sliderRepository struct {
	c *mongo.Collection
}

// CreateSlider implements SliderRepository.
func (r sliderRepository) CreateSlider(slider models.Slider) error {
	_, err := r.c.InsertOne(context.TODO(), slider)
	return err
}
func (r *sliderRepository) Get(id primitive.ObjectID) (slider *models.Slider, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&slider)
	return slider, err
}

// DeleteSlider implements SliderRepository.
func (r sliderRepository) DeleteSlider(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// GetSliders implements SliderRepository.
func (r sliderRepository) GetSliders() ([]models.Slider, error) {
	result, err := r.c.Find(context.TODO(), bson.M{})
	var sliders []models.Slider
	if err != nil {
		return make([]models.Slider, 0), err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var slider models.Slider
		if err = result.Decode(&slider); err != nil {
			return make([]models.Slider, 0), err
		}

		sliders = append(sliders, slider)
	}
	if sliders == nil {
		sliders = make([]models.Slider, 0)
	}

	return sliders, err
}

func NewSliderRepository(DB *mongo.Client) SliderRepository {
	return &sliderRepository{db.GetCollection(db.DB, slidercollection)}
}
