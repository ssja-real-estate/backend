package security

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewToken(t *testing.T) {
	id := primitive.NewObjectID()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	id := primitive.NewObjectID()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, id.Hex(), payload.Id)
	assert.Equal(t, id.Hex(), payload.Issuer)
	assert.Equal(t, time.Now().Year(), time.Unix(payload.IssuedAt, 0).Year())
	assert.Equal(t, time.Now().Month(), time.Unix(payload.IssuedAt, 0).Month())
	assert.Equal(t, time.Now().Day(), time.Unix(payload.IssuedAt, 0).Day())
}
