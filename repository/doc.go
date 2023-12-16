package repository

import (
	"context"
	"realstate/db"
	"realstate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const docCollection = "document"

type DocumentRepository interface {
	SaveDoc(estatetype *models.Document) error
	GetDocumentById(id primitive.ObjectID) (Document *models.Document, err error)
	GetDocumentAll() ([]models.Document, error)
	DeleteDocument(id primitive.ObjectID) error
}

type documentRepository struct {
	c *mongo.Collection
}

// DeleteDocument implements DocumentRepository.
func (r *documentRepository) DeleteDocument(id primitive.ObjectID) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err

}

// GetDocumentAll implements DocumentRepository.
func (r *documentRepository) GetDocumentAll() ([]models.Document, error) {
	result, err := r.c.Find(context.TODO(), bson.M{})
	var documents []models.Document
	if err != nil {
		return make([]models.Document, 0), err
	}

	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var document models.Document
		if err = result.Decode(&document); err != nil {
			return make([]models.Document, 0), err
		}

		documents = append(documents, document)
	}
	if documents == nil {
		documents = make([]models.Document, 0)
	}

	return documents, err

}

// GetDocumentById implements DocumentRepository.
func (r *documentRepository) GetDocumentById(id primitive.ObjectID) (Document *models.Document, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&Document)
	return Document, err
}

// SaceDoce implements DocumentRepository.
func (r *documentRepository) SaveDoc(document *models.Document) error {
	_, err := r.c.InsertOne(context.TODO(), &document)
	return err
}

func NewDocumentRepository(Db *mongo.Client) DocumentRepository {
	return &documentRepository{db.GetCollection(Db, docCollection)}
}
