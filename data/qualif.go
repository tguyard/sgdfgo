package data

import "time"

var QualifNames = map[string]int{
	"Animateur SF (CAFASF)":          2,
	"Directeur SF (CAFDSF)":          1,
	"Responsable Unité SF (CAFRUSF)": 3,
}

var QualifIds = map[int]string{
	2: "Animateur SF (CAFASF)",
	1: "Directeur SF (CAFDSF)",
	3: "Responsable Unité SF (CAFRUSF)",
}

type Qualification struct {
	ID          int
	StartDate   time.Time
	EndDate     time.Time
	Unlimited   bool
	IsTitulaire bool
}

func (q Qualification) GetName() string {
	return QualifIds[q.ID]
}
