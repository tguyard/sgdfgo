package data

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
)

type Scrap struct {
	ID          objectid.ObjectID `bson:"_id"`
	Date        time.Time
	ByUser      int
	OnStructure int
	Status      string
	Err         string
}

func StartScrap(person Person) *Scrap {
	return &Scrap{
		ID:          objectid.New(),
		Date:        time.Now(),
		ByUser:      person.ID,
		OnStructure: person.Structure,
		Status:      "loading",
	}
}

func (s *Scrap) Save() error {
	_, err := db("scraps").InsertOne(context.Background(), s)
	return err
}

func (s *Scrap) SetFinished() error {
	s.Status = "finished"
	_, err := db("scraps").UpdateOne(context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", s.ID),
		), bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("Status", s.Status),
			),
		), updateopt.Upsert(true))
	return err
}
func (s *Scrap) SetError(err error) error {
	s.Err = err.Error()
	s.Status = "error"

	_, err = db("scraps").UpdateOne(context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", s.ID),
		), bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("Status", s.Status),
				bson.EC.String("Err", s.Err),
			),
		), updateopt.Upsert(true))
	return err
}
