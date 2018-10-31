package data

import "time"

var DiplomaIds = map[int]string{
	185: "2 Buchettes",
	612: "3 Buchettes",
	613: "4 Buchettes",
	619: "AFPS",
	73:  "BAFA",
	607: "BAFA Qualification Voile",
	140: "BAFD",
	618: "BP JEPS",
	615: "Brevet d'Etat (activités marines)",
	617: "Brevet Professionnel (Skipper, marine marchande, .",
	616: "Carte Mer",
	609: "Certificat Radio Restreint CRR",
	606: "Chef de Flottille",
	605: "Chef de quart",
	626: "DE JEPS",
	628: "DEUG STAPS",
	624: "Directeur nautique",
	614: "DUT animation socioculturelle ",
	630: "Licence STAPS",
	622: "Médaille Argent de la JS",
	621: "Médaille Bronze de la JS",
	623: "Médaille Or de la JS",
	629: "Nœud de tisserand",
	604: "Patron d'embarcation",
	610: "Permis Côtier",
	625: "Permis E ",
	627: "Permis fluvial",
	611: "Permis Hauturier",
	603: "PSC1 Prevention et Secours Civiques de niveau 1",
	620: "Spirale AMGE",
	608: "Surveillant de Baignade",
}

var DiplomaNames = map[string]int{
	"2 Buchettes": 185,
	"3 Buchettes": 612,
	"4 Buchettes": 613,
	"AFPS":        619,
	"BAFA":        73,
	"BAFA Qualification Voile": 607,
	"BAFD":    140,
	"BP JEPS": 618,
	"Brevet d'Etat (activités marines)":                  615,
	"Brevet Professionnel (Skipper, marine marchande, .": 617,
	"Carte Mer":                                       616,
	"Certificat Radio Restreint CRR":                  609,
	"Chef de Flottille":                               606,
	"Chef de quart":                                   605,
	"DE JEPS":                                         626,
	"DEUG STAPS":                                      628,
	"Directeur nautique":                              624,
	"DUT animation socioculturelle ":                  614,
	"Licence STAPS":                                   630,
	"Médaille Argent de la JS":                        622,
	"Médaille Bronze de la JS":                        621,
	"Médaille Or de la JS":                            623,
	"Nœud de tisserand":                               629,
	"Patron d'embarcation":                            604,
	"Permis Côtier":                                   610,
	"Permis E ":                                       625,
	"Permis fluvial":                                  627,
	"Permis Hauturier":                                611,
	"PSC1 Prevention et Secours Civiques de niveau 1": 603,
	"Spirale AMGE":                                    620,
	"Surveillant de Baignade":                         608,
}

type Diploma struct {
	ID   int
	Ref  string
	Date time.Time
}

func (q Diploma) GetName() string {
	return DiplomaIds[q.ID]
}

func (q Diploma) GetType() string {
	if q.ID == 619 || q.ID == 603 {
		return "health"
	}
	if q.ID == 73 || q.ID == 607 || q.ID == 614 {
		return "annimation"
	}
	if q.ID == 140 || q.ID == 618 || q.ID == 626 {
		return "direction"
	}
	return "none"
}
