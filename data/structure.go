package data

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

var StructureNames = map[string]int{
	"Autres":                       1211,
	"Centre National":              1205,
	"Groupe":                       1200,
	"Membres associés local":       1201,
	"Membres associés National":    1206,
	"Membres associés Territorial": 1204,
	"Sommet":                       1207,
	"Territoire":                   1203,

	"Unité 11-14 ans":     1212,
	"Unité 14-17 ans":     1210,
	"Unité 17-20 ans":     1209,
	"Unité 8-11 ans":      1199,
	"Unité Farfadet":      1208,
	"Unité Vent du Large": 1202,
}

var StructureIDs = map[int]string{
	1211: "Autres",
	1205: "Centre National",
	1200: "Groupe",
	1201: "Membres associés local",
	1206: "Membres associés National",
	1204: "Membres associés Territorial",
	1207: "Sommet",
	1203: "Territoire",
	1212: "Unité 11-14 ans",
	1210: "Unité 14-17 ans",
	1209: "Unité 17-20 ans",
	1199: "Unité 8-11 ans",
	1208: "Unité Farfadet",
	1202: "Unité Vent du Large",
}

var SpecialityNames = map[string]int{
	"Marine":          622,
	"sans spécialité": 624,
	"Vent du Large":   623,
}

type Structure struct {
	ID        int `bson:"_id"`
	ScrapDate time.Time

	Name        string
	Type        int
	ParentID    int
	Speciality  int
	Email       string
	URL         string
	Tel         string
	Lat         float64
	Long        float64
	City        string
	Country     string
	Zip         string
	Description string
}

func (s Structure) Save() error {
	_, err := db("structures").InsertOne(context.Background(), &s)

	return err
}

func NewStructureByID(id int) (Structure, error) {
	result := db("structures").FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("_id", int32(id)),
		))

	s := Structure{}
	err := result.Decode(&s)
	return s, err
}

func (s Structure) TypeName() string {
	if s.IsNational() {
		return "National"
	}
	if s.IsTer() {
		return "Territoire"
	}
	if s.IsGroup() {
		return "Groupe"
	}
	if s.IsUnit() {
		return "Unité"
	}
	return "Autre"
}

// IsNational TODO
func (s Structure) IsNational() bool {
	return s.Type == 1205 || s.Type == 1206 || s.Type == 1207
}

// IsTer TODO
func (s Structure) IsTer() bool {
	return s.Type == 1204 || s.Type == 1203
}

// IsGroup TODO
func (s Structure) IsGroup() bool {
	return s.Type == 1200 || s.Type == 1201
}

// IsUnit TODO
func (s Structure) IsUnit() bool {
	return s.Type == 1212 || s.Type == 1210 || s.Type == 1209 || s.Type == 1199 || s.Type == 1208 || s.Type == 1202
}

// IsOther TODO
func (s Structure) IsOther() bool {
	return s.Type == 1211
}
