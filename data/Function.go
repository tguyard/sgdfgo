package data

var FunctionIds = map[string][2]string{

	// CODE     NOM MASCULIN                       NOM FEMININ
	"170":  [2]string{"FARFADET", "FARFADET"},
	"110":  [2]string{"LOUVETEAU", "JEANNETTE"},
	"110M": [2]string{"MOUSSAILLON", "MOUSSAILLON"},
	"120":  [2]string{"SCOUT", "GUIDE"},
	"120M": [2]string{"MOUSSE", "MOUSSE"},
	"130":  [2]string{"PIONNIER", "CARAVELLE"},
	"130M": [2]string{"MARIN", "CARAVELLE MARINE"},
	"140":  [2]string{"COMPAGNON", "COMPAGNON"},
	"140M": [2]string{"COMPAGNON MARIN", "COMPAGNON MARIN"},
	"141":  [2]string{"COMPAGNON 3EME TEMPS", "COMPAGNON 3EME TEMPS"},

	"210":  [2]string{"RESPONSABLE D'UNITE LOUVETEAU JEANNETTE", "RESPONSABLE D'UNITE LOUVETEAU JEANNETTE"},
	"210M": [2]string{"RESPONSABLE D'UNITE MOUSSAILLON", "RESPONSABLE D'UNITE MOUSSAILLON"},
	"213":  [2]string{"CHEF LOUVETEAU JEANNETTE", "CHEFTAINE LOUVETEAU JEANNETTE"},
	"213M": [2]string{"CHEF MOUSSAILLON", "CHEFTAINE MOUSSAILLON"},
	"220":  [2]string{"RESPONSABLE D'UNITE SCOUT GUIDE", "RESPONSABLE D'UNITE SCOUT GUIDE"},
	"220M": [2]string{"RESPONSABLE D'UNITE MOUSSE", "RESPONSABLE D'UNITE MOUSSE"},
	"223":  [2]string{"CHEF SCOUT GUIDE", "CHEFTAINE SCOUT GUIDE"},
	"223M": [2]string{"CHEF MOUSSE", "CHEFTAINE MOUSSE"},
	"230":  [2]string{"RESPONSABLE D'UNITE PIONNIER CARAVELLE", "RESPONSABLE D'UNITE PIONNIER CARAVELLE"},
	"230M": [2]string{"RESPONSABLE D'UNITE FLOTILLE - CARAVELLE MARINE", "RESPONSABLE D'UNITE FLOTILLE - CARAVELLE MARINE"},
	"233":  [2]string{"CHEF PIONNIER CARAVELLE", "CHEFTAINE PIONNIER CARAVELLE"},
	"233M": [2]string{"CHEF FLOTILLE", "CHEFTAINE CARAVELLE MARINE"},

	"240": [2]string{"ACCOMPAGNATEUR COMPAGNON", "ACCOMPAGNATRICE COMPAGNON"},
	"270": [2]string{"RESPONSABLE FARFADET", "RESPONSABLE FARFADET"},
	"271": [2]string{"PARENT ANIMATEUR FARFADET", "PARENT ANIMATEUR FARFADET"},
	"293": [2]string{"RESPONSABLE VENT DU LARGE", "RESPONSABLE VENT DU LARGE"},

	"300":  [2]string{"RESPONSABLE DE GROUPE", "RESPONSABLE DE GROUPE"},
	"301":  [2]string{"RESPONSABLE DE GROUPE ADJOINT", "RESPONSABLE DE GROUPE ADJOINTE"},
	"302":  [2]string{"AUMONIER DE GROUPE", "AUMONIER DE GROUPE"},
	"307":  [2]string{"SECRETAIRE DE GROUPE", "SECRETAIRE DE GROUPE"},
	"309":  [2]string{"TRESORIER DE GROUPE", "TRESORIERE DE GROUPE"},
	"316":  [2]string{"ANIMATEUR DE VIE CHRETIENNE", "ANIMATRICE DE VIE CHRETIENNE"},
	"330":  [2]string{"CHARGE DE MISSION DU GROUPE", "CHARGEE DE MISSION DU GROUPE"},
	"330M": [2]string{"REFERENT TECHNIQUE MARIN", "REFERENTE TECHNIQUE MARIN"},
	"380":  [2]string{"RESPONSABLE LOCAL DEVELOPPEMENT ET RESEAUX", "RESPONSABLE LOCAL DEVELOPPEMENT ET RESEAUX"},

	"391": [2]string{"FF TRESORIER DE GROUPE", "TRESORIERE DE GROUPE"},
	"332": [2]string{"REPRESENTANT ASSOCIATIF", "REPRESENTANT ASSOCIATIF"},
	"399": [2]string{"FF RESPONSABLE DE GROUPE", "FF RESPONSABLE DE GROUPE"},

	// TERRITOIRE

	"500": [2]string{"DELEGUE TERRITORIAL", "DELEGUEE TERRITORIALE"},
	"501": [2]string{"DELEGUE TERRITORIAL ADJOINT", "DELEGUEE TERRITORIALE ADJOINTE"},
	"502": [2]string{"AUMONIER TERRITORIAL", "AUMONIER TERRITORIAL"},
	"503": [2]string{"RESPONSABLE POLE PEDAGOGIQUE", "RESPONSABLE POLE PEDAGOGIQUE"},
	"504": [2]string{"RESPONSABLE POLE DEVELOPPEMENT", "RESPONSABLE POLE DEVELOPPEMENT"},
	"505": [2]string{"RESPONSABLE POLE ADMINISTRATIF ET FINANCIER", "RESPONSABLE POLE ADMINISTRATIF ET FINANCIER"},

	"598": [2]string{"FF DELEGUE TERRITORIAL", "FF DELEGUEE TERRITORIALE"},
	"626": [2]string{"CORRESPONDANT HANDICAP", "CORRESPONDANT HANDICAP"},
	"631": [2]string{"MEDIATEUR", "MEDIATRICE"},
	"632": [2]string{"REPRESENTANT ASSOCIATIF", "REPRESENTANTE ASSOCIATIF"},
	"670": [2]string{"SECRETAIRE TERRITORIAL", "SECRETAIRE TERRITORIALE"},
	"691": [2]string{"FF TRESORIER TERRITORIAL", "FF TRESORIERE TERRITORIAL"},

	"600": [2]string{"ACCOMPAGNATEUR PEDAGOGIQUE", "ACCOMPAGNATRICE PEDAGOGIQUE"},
	"609": [2]string{"ACCOMPAGNATEUR DES RESPONSABLES DE GROUPE", "ACCOMPAGNATEUR DES RESPONSABLES DE GROUPE"},

	"610": [2]string{"EQUIPIER TERRITORIAL", "EQUIPIERE TERRITORIALE"},
	"616": [2]string{"ANIMATEUR DE VIE CHRETIENNE", "ANIMATRICE DE VIE CHRETIENNE"},
	"620": [2]string{"CHARGE MISSION DEVELOPPEMENT", "CHARGEE MISSION DEVELOPPEMENT"},
	"621": [2]string{"COORDINATEUR DEVELOPPEMENT ET RESEAUX", "COORDINATRICE DEVELOPPEMENT ET RESEAUX"},

	"622": [2]string{"OUVREUR DE GROUPE", "OUVREUR DE GROUPE"},
	"624": [2]string{"RESPONSABLE PARTENARIATS FINANCIERS", "RESPONSABLE PARTENARIATS FINANCIERS"},

	"625": [2]string{"RESPONSABLE COMMUNICATION", "RESPONSABLE COMMUNICATION"},
	"627": [2]string{"RESPONSABLE SCOUTISME EN QUARTIER", "RESPONSABLE SCOUTISME EN QUARTIER"},

	"630": [2]string{"CHARGE MISSION TERRITORIAL", "CHARGEE MISSION TERRITORIALE"},
	"650": [2]string{"RESPONSABLE BOUTIQUE SGDF", "RESPONSABLE BOUTIQUE SGDF"},
	"659": [2]string{"CHARGE DE MISSION RASSEMBLEMENT", "CHARGE DE MISSION RASSEMBLEMENT"},
	"660": [2]string{"CHARGE MISSION ADMINISTRATIF", "CHARGEE MISSION ADMINISTRATIF"},
	"661": [2]string{"GESTIONNAIRE MATERIEL ET LOGISTIQUE", "GESTIONNAIRE MATERIEL ET LOGISTIQUE"},

	"664": [2]string{"EQUIPIER DE BASE", "EQUIPIERE DE BASE"},
	"665": [2]string{"RESPONSABLE DE BASE", "RESPONSABLE DE BASE"},
	"690": [2]string{"TRESORIER TERRITORIAL", "TRESORIERE TERRITORIALE"},
}

var FunctionSmallIds = map[string]string{

	"170":  "jeune",
	"110":  "jeune",
	"110M": "jeune",
	"120":  "jeune",
	"120M": "jeune",
	"130":  "jeune",
	"130M": "jeune",
	"140":  "compagnon",
	"140M": "compagnon",
	"141":  "compagnon",

	"210":  "chef",
	"210M": "chef",
	"213":  "chef",
	"213M": "chef",
	"220":  "chef",
	"220M": "chef",
	"223":  "chef",
	"223M": "chef",
	"230":  "chef",
	"230M": "chef",
	"233":  "chef",
	"233M": "chef",

	"240": "AC",
	"270": "responsable farfa",
	"271": "parent annimateur farfa",
	"293": "responsable vent du large",

	"300":  "RG",
	"301":  "RGA",
	"302":  "aumonier",
	"307":  "secretaire",
	"309":  "tresorier",
	"316":  "animateur de vie chretienne",
	"330":  "chargé de mission",
	"330M": "referent technique marin",
	"380":  "responsable dev",

	"391": "tresorier",
	"332": "representant associatif",
	"399": "RG",

	// TERRITOIRE

	"500": "DT",
	"501": "DTA",
	"502": "aumonier",
	"503": "RPP",
	"504": "RPDEV",
	"505": "RPAF",

	"598": "DT",
	"626": "correspondant handicap",
	"631": "mediateur",
	"632": "representant associatif",
	"670": "secretaire",
	"691": "tresorier",

	"600": "AP",
	"609": "accompagnateur RG",

	"610": "équipier territorial",
	"616": "animateur de vie chretienne",
	"620": "chargé mission developpement",
	"621": "coordinateur developpement et reseaux",

	"622": "ouvreur de groupe",
	"624": "responsable partenariats financiers",

	"625": "responsable com",
	"627": "responsable scoutisme en quartier",

	"630": "chargé mission",
	"650": "responsable boutique",
	"659": "chargé de mission rassemblement",
	"660": "chargé mission administratif",
	"661": "responsable matos",

	"664": "équipier de base",
	"665": "responsable de base",
	"690": "tresorier",
}

const (
	FuncNameSmall  int = 0
	FuncNameMale   int = 1
	FuncNameFemale int = 2
)

func GetFunctionName(code string, funcNameType int) string {
	if funcNameType == 0 {
		return FunctionSmallIds[code]
	}
	if funcNameType == 2 {
		return FunctionIds[code][1]
	}
	return FunctionIds[code][0]
}
