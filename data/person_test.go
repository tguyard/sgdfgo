package data

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	assert := assert.New(t)

	_, err := getClient().Database("sgdfgo").RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
	)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	p := Person{
		ID:         1,
		Contacts:   []Contact{},
		Diplomas:   []Diploma{},
		Formations: []Formation{},
		Function:   "101",
		Identity: Contact{
			Address:   "a",
			Firstname: "1",
			Name:      "2",
			Type:      Father,
			Tel:       "0",
			Email:     "11",
		},
	}

	err = p.Save()
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	s1, err := NewPersonById(1)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(p.ID, s1.ID)
	assert.Equal(p.Function, s1.Function)
	assert.Equal(p.Identity.Address, s1.Identity.Address)
}
