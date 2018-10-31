package data

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/stretchr/testify/assert"
)

func TestStructure(t *testing.T) {
	databaseName = "sgdftest"
	assert := assert.New(t)

	_, err := getClient().Database(databaseName).RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
	)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	s := Structure{
		Email:      "email",
		ID:         1,
		Name:       "name",
		ParentID:   2,
		Speciality: 3,
		Type:       4,
		URL:        "url",
	}

	err = s.Save()
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	s1, err := NewStructureByID(1)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(s.Email, s1.Email)
	assert.Equal(s.ID, s1.ID)
	assert.Equal(s.Name, s1.Name)
	assert.Equal(s.ParentID, s1.ParentID)
	assert.Equal(s.Speciality, s1.Speciality)
	assert.Equal(s.Type, s1.Type)
	assert.Equal(s.URL, s1.URL)
}
