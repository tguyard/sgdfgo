package data

import (
	"context"
	"time"

	"github.com/imdario/mergo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
)

type ContactType int

type Inscription struct {
	Start time.Time
	End   time.Time
	Type  string
}

type Adhesion struct {
	Start  time.Time
	End    time.Time
	Status int
}

type JeunesseSport struct {
	Function, Diploma, Details, Qualite string
	LastModif                           time.Time
}

const (
	Member ContactType = 0
	Mother ContactType = iota
	Father ContactType = iota
	Other  ContactType = iota
)

type Contact struct {
	Name, Firstname, Address string
	Tel, Email               string
	Type                     ContactType
	Birthdate                time.Time
	BirthLocation            string
	Profession               string
}

type Person struct {
	ID            int `bson:"_id"`
	ScrapDate     time.Time
	Structure     int
	Function      string
	Identity      Contact
	NbAllocataire string

	RegimeGeneral, RegimeMSA, RegimeEtranger                 bool
	AutorisationInterventionChirurgicale, DroitImage         bool
	AssuranceResponsabiliteCivile, AutoriseMailInfoMouvement bool
	AutoriseMailInfoExterne                                  bool

	Contacts []Contact

	Diplomas   []Diploma
	Formations []Formation

	Qualif        Qualification
	Inscrtiption  Inscription
	Adhesion      Adhesion
	JeunesseSport JeunesseSport
}

func (p Person) Save() error {
	old, err := NewPersonByID(p.ID)
	if err != nil {
		err = mergo.Merge(&p, old)
		if err != nil {
			return err
		}
	}

	data := map[string]interface{}{
		"$set": &p,
	}
	_, err = db("persons").UpdateOne(context.Background(),
		bson.NewDocument(
			bson.EC.Int64("_id", int64(p.ID)),
		), data, updateopt.Upsert(true))
	return err
}

// mongo: no documents in result

func NewPersonByID(id int) (Person, error) {
	result := db("persons").FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("_id", int32(id)),
		))

	s := Person{}
	err := result.Decode(&s)
	return s, err
}
