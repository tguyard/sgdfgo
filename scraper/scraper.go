package scraper

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/sgdfgo/data"
)

// Scraper TODO
type Scraper struct {
	baseURL                   string
	c                         *colly.Collector
	structureURL, adherentURL string
	callbackError             error
}

// New TODO
func New() *Scraper {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; rv:31.0) Gecko/20100101 Firefox/61.0"),
		colly.AllowedDomains("intranet.sgdf.fr"),
		// colly.CacheDir("/tmp/sgdfgo-cache"),
	)
	c.SetRequestTimeout(180 * time.Second)
	c.WithTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   180 * time.Second,
			KeepAlive: 180 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       180 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	return &Scraper{
		baseURL: "https://intranet.sgdf.fr",
		c:       c,
	}
}

func (s *Scraper) consumeErr() error {
	err := s.callbackError
	s.callbackError = nil
	return err
}

// Connect TODO
func (s *Scraper) Connect(login, password string) (int, error) {
	if s.callbackError != nil {
		return -1, s.consumeErr()
	}
	startURL := s.baseURL + "/Specialisation/Sgdf/Default.aspx"
	structureID := -1

	s.c.OnHTML("#ctl00__hlVoirFicheStructure[href]", func(e *colly.HTMLElement) {
		s.structureURL = e.Attr("href")
	})
	s.c.OnHTML("#ctl00__hlVoirFicherAdherent[href]", func(e *colly.HTMLElement) {
		s.adherentURL = e.Attr("href")
	})

	s.c.OnHTML("#ctl00__ddDelegations", func(e *colly.HTMLElement) {
		option := e.ChildText("option[selected=selected]")
		re := regexp.MustCompile("@ *([0-9]+)$")
		if matches := re.FindStringSubmatch(option); len(matches) == 2 {
			structureID, _ = strconv.Atoi(matches[1])
		}
	})

	s.c.OnHTML("body", func(e *colly.HTMLElement) {
		if strings.Contains(e.DOM.Text(), "Erreur inconnue") {
			s.callbackError = errors.New("Intranet ERROR")
		}
	})

	//
	// Start Crawling
	//
	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Connexion", r.URL)
	})

	err := s.c.Post(startURL, map[string]string{
		"__VIEWSTATE":                               "/wEPZwUPOGQ2M2E4YzIwZjI0OWQ249N596drjy0+MDT+x2hPLl4PC9o=",
		"__VIEWSTATEGENERATOR":                      "F4403698",
		"__EVENTTARGET":                             "",
		"__EVENTARGUMENT":                           "",
		"__EVENTVALIDATION":                         "/wEdAAUQ/jH9gFZneLTWQDe6b/Lg/55zOzzEBlNtbr7pJB5UvDGVEAOAsrBEy6P94sJdCUVzsej/WjmIl9BiFEsnoqSU5wxz7TGst0yyQ7xMskRui7rb/SOMwBgGGcHAd5TtCY22WJIY",
		"ctl00$MainContent$login":                   login,
		"ctl00$MainContent$password":                password,
		"ctl00$MainContent$_btnValider":             "Se+connecter",
		"ctl00$_hidReferenceStatistiqueUtilisation": "-1",
	})

	if structureID == -1 {
		return -1, errors.New("Could not find the structure ID")
	}
	if s.callbackError != nil {
		return -1, s.consumeErr()
	}

	return structureID, err
}

func (s Scraper) atoi(val string) int {
	ret, err := strconv.Atoi(val)
	if err != nil {
		s.callbackError = err
	}
	return ret
}
func (s Scraper) atof(val string) float64 {
	if val == "" {
		return 0.
	}
	// "Latitude : 48,8322°"
	regex := regexp.MustCompile("^.* ([0-9]+),([0-9]+)°$")
	matches := regex.FindStringSubmatch(val)
	if len(matches) == 3 {
		val = matches[1] + "." + matches[2]
	}

	ret, err := strconv.ParseFloat(val, 64)
	if err != nil {
		s.callbackError = err
	}
	return ret
}

func (s *Scraper) ScrapStructures() error {
	if s.callbackError != nil {
		return s.consumeErr()
	}
	return s.visitStructurePage(s.structureURL)
}

func (s *Scraper) visitStructurePage(url string) error {
	if s.callbackError != nil {
		return s.consumeErr()
	}

	c := s.c.Clone()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		postData := make(map[string]string)
		e.ForEach("form#aspnetForm input[type=hidden]", func(index int, e *colly.HTMLElement) {
			postData[e.Attr("name")] = e.Attr("value")
		})

		parentID := e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabHierarche__gvParents tr.ligne1 #ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabHierarche__gvParents_ctl02__hlStructure")

		var children []*colly.HTMLElement
		e.ForEach("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabHierarche__gvEnfants tr:not(.entete)", func(index int, e *colly.HTMLElement) {
			children = append(children, e)
		})

		if intPID, err := strconv.Atoi(parentID); err != nil {
			s.callbackError = err
		} else if err = s.parseStructureDetails(url, postData, intPID, children); err != nil {
			s.callbackError = err
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Structure", r.URL)
	})

	err := c.Visit(s.baseURL + "/" + url)
	if s.callbackError != nil {
		return s.consumeErr()
	}
	return err
}

func (s *Scraper) parseStructureDetails(url string, basePost map[string]string, parentID int, children []*colly.HTMLElement) error {
	if s.callbackError != nil {
		return s.consumeErr()
	}

	c := s.c.Clone()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblNom"))

		structure := data.Structure{
			ID:          s.atoi(e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblCodeStructure")),
			ScrapDate:   time.Now(),
			Email:       (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblCourrier")),
			Name:        (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblNom")),
			ParentID:    parentID,
			Speciality:  0,
			Type:        data.StructureNames[e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblType")],
			URL:         (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__hlSiteWeb")),
			Lat:         s.atof(e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblCoordonneesGPSLatitude")),
			Long:        s.atof(e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblCoordonneesGPSLongitude")),
			City:        (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__resumeAdresse__lbVille")),
			Country:     (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblNom")),
			Description: (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblNom")),
			Tel:         (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__lblTelephone")),
			Zip:         (e.ChildText("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeStructure__tabResume__resume__resumeAdresse__lbCodePostal")),
		}
		if err := structure.Save(); err != nil {
			s.callbackError = err
		}

		for _, child := range children {
			if strings.Contains(child.ChildText("td:last-child"), "Aucune structure") {
				continue
			}
			childURL := child.ChildAttr("td:first-child a", "href")
			if childURL == "" {
				continue
			}

			if structure.Type == data.StructureNames["Groupe"] {
				stype := -1
				childName := child.ChildText("td:last-child")
				if strings.Contains(childName, "LOUVETEAUX") || strings.Contains(childName, "JEANNETTES") {
					stype = 1199
				}
				if strings.Contains(childName, "SCOUTS") || strings.Contains(childName, "GUIDES") {
					stype = 1212
				}
				if strings.Contains(childName, "PIONNIERS") || strings.Contains(childName, "CARAVELLES") {
					stype = data.StructureNames["Groupe"]
				}
				if strings.Contains(childName, "FARFADET") {
					stype = 1208
				}
				if strings.Contains(childName, "VENT") && strings.Contains(childName, "LARGE") {
					stype = 1202
				}
				if strings.Contains(childName, "COMPAGNONS") {
					stype = 1209
				}
				if stype != -1 {
					childStruct := data.Structure{
						ID:         s.atoi(child.ChildText("td:first-child a")),
						ScrapDate:  time.Now(),
						Name:       childName,
						ParentID:   structure.ID,
						Speciality: 0,
						Type:       stype,
					}
					err := childStruct.Save()
					if err != nil {
						s.callbackError = err
					}
					continue
				}
			}
			if err := s.visitStructurePage("Specialisation/Sgdf/structures/" + childURL); err != nil {
				s.callbackError = err
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("StructureDetails", r.URL)
	})

	basePost["__EVENTTARGET"] = "ctl00$ctl00$MainContent$TabsContent$TabContainerResumeStructure"
	basePost["__EVENTARGUMENT"] = "activeTabChanged:1"

	err := c.Post(s.baseURL+"/"+url, basePost)
	if s.callbackError != nil {
		return s.consumeErr()
	}
	return err
}

// ScrapExport TODO
func (s *Scraper) ScrapExport() error {
	if s.callbackError != nil {
		return s.consumeErr()
	}

	exported, err := s.Export()
	if err != nil {
		return err
	}

	indexes := make(map[string]int)

	for j, cell := range exported[0] {
		indexes[cell] = j
	}
	get := func(name string, row []string) string {
		return row[indexes[name]]
	}
	getInt := func(name string, row []string) int {
		val := row[indexes[name]]
		if val == "" {
			return 0
		}
		intval, err := strconv.Atoi(val)
		if err != nil {
			s.callbackError = err
		}
		return intval
	}
	getTime := func(name string, row []string) time.Time {
		val := row[indexes[name]]
		if val == "" {
			return time.Time{}
		}
		res, err := time.Parse("02/01/2006", val)
		if err == nil {
			return res
		}
		return time.Time{}
	}

	getContact := func(name string, row []string, ctype data.ContactType) data.Contact {
		email := get(name+".CourrielPersonnel", row)
		if email == "" {
			email = get(name+".CourrielProfessionnel", row)
		}
		tel := get(name+".TelephonePortable1", row)
		if tel == "" {
			tel = get(name+".TelephonePortable2", row)
		}
		if tel == "" {
			tel = get(name+".TelephoneDomicile", row)
		}
		if tel == "" {
			tel = get(name+".TelephoneBureau", row)
		}
		return data.Contact{
			Address:   strings.Trim(get(name+".Adresse.Ligne1", row)+"\n"+get(name+".Adresse.Ligne2", row)+"\n"+get(name+".Adresse.Ligne3", row), "\n "),
			Firstname: get(name+".Prenom", row),
			Name:      get(name+".Nom", row),
			Type:      ctype,
			Tel:       tel,
			Email:     email,

			Birthdate:     getTime(name+".DateNaissance", row),
			BirthLocation: get(name+"PaysNaissance.PaysLib", row),
			Profession:    get(name+".Profession", row),
		}
	}

	oui := "Oui"
	formables := make([]data.Person, 0, 32)
	for i, row := range exported {
		if i == 0 {
			continue
		}

		p := data.Person{
			ID:        getInt("Individu.CodeAdherent", row),
			ScrapDate: time.Now(),
			Structure: getInt("Structure.CodeStructure", row),
			Function:  get("Fonction.Code", row),
			Identity:  getContact("Individu", row, data.Member),

			NbAllocataire: get("Individu.NumeroAllocataire", row),

			RegimeGeneral:                        get("Individu.RegimeGeneral", row) == oui,
			RegimeMSA:                            get("Individu.RegimeMSA", row) == oui,
			RegimeEtranger:                       get("Individu.RegimeEtranger", row) == oui,
			AutorisationInterventionChirurgicale: get("Individu.AutorisationInterventionChirurgicale", row) == oui,
			DroitImage:                           get("Individu.DroitImage", row) == oui,
			AssuranceResponsabiliteCivile:        get("Individu.AssuranceResponsabiliteCivile", row) == oui,
			AutoriseMailInfoMouvement:            get("Individu.AutoriseMailInfoMouvement", row) == oui,
			AutoriseMailInfoExterne:              get("Individu.RegimeGeneralAutoriseMailInfoExterne", row) == oui,

			Contacts: make([]data.Contact, 0),
			Inscrtiption: data.Inscription{
				Start: getTime("Inscriptions.DateDebut", row),
				End:   getTime("Inscriptions.DateFin", row),
				Type:  get("Inscriptions.Type", row),
			},
			Adhesion: data.Adhesion{
				Start:  getTime("Adhesions.DateDebut", row),
				End:    getTime("Adhesions.DateFin", row),
				Status: getInt("Adhesions.Statut", row),
			},
			JeunesseSport: data.JeunesseSport{
				Details:   get("IntervenantJS.DiplomeDetailJS", row),
				Diploma:   get("IntervenantJS.DiplomeJS", row),
				Function:  get("IntervenantJS.FonctionJS", row),
				Qualite:   get("IntervenantJS.QualiteJS", row),
				LastModif: getTime("IntervenantJS.DerniereModification", row),
			},
			Qualif: data.Qualification{
				ID:          data.QualifNames[get("QualificationsQualificationJeunesseSports.Libelle", row)],
				EndDate:     getTime("Qualifications.DateFinValidite", row),
				IsTitulaire: get("Qualifications.EstTitulaire", row) == oui,
				Unlimited:   get("Qualifications.DateFinValidite", row) == "",
			},
		}
		if get("Pere.Nom", row) != "" {
			p.Contacts = append(p.Contacts, getContact("Pere", row, data.Father))
		}
		if get("Mere.Nom", row) != "" {
			p.Contacts = append(p.Contacts, getContact("Mere", row, data.Mother))
		}
		err := p.Save()
		if err != nil {
			return err
		}

		if p.Function != "170" && p.Function != "110" &&
			p.Function != "110M" && p.Function != "120" &&
			p.Function != "120M" && p.Function != "130" {
			formables = append(formables, p)
		}
	}

	for _, person := range formables {
		if err := s.AddFormation(person); err != nil {
			return err
		}
	}

	return s.consumeErr()
}

func (s *Scraper) AddFormation(person data.Person) error {
	if s.callbackError != nil {
		return s.consumeErr()
	}

	c := s.c.Clone()
	url := "/Specialisation/Sgdf/adherents/RechercherAdherent.aspx?code=" + strconv.Itoa(person.ID)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		postData := make(map[string]string)
		e.ForEach("form#aspnetForm input[type=hidden]", func(index int, e *colly.HTMLElement) {
			postData[e.Attr("name")] = e.Attr("value")
		})

		if err := s.doAddFormation(e.Request.URL.String(), postData, person); err != nil {
			s.callbackError = err
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("AddFormation", r.URL)
	})

	c.AllowURLRevisit = true

	err := c.Visit(s.baseURL + url)
	if s.callbackError != nil {
		return s.consumeErr()
	}
	return err
}

var typestr = "Type"

func (s *Scraper) doAddFormation(fullURL string, basePost map[string]string, person data.Person) error {
	if s.callbackError != nil {
		return s.consumeErr()
	}

	c := s.c.Clone()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		formation := make([]data.Formation, 0)
		e.ForEach("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeAdherent__tabFormations__formations__gvFormations__gvFormations tr", func(index int, tr *colly.HTMLElement) {
			if tr.ChildText("td:first-child") == typestr || tr.ChildText("th:first-child") == typestr {
				return
			}
			line := make([]string, 0)
			tr.ForEach("td", func(index int, td *colly.HTMLElement) {
				line = append(line, td.DOM.Text())
			})

			if len(line) < 3 {
				return
			}
			if strings.Trim(line[1], " \t") == "Stagiaire" {
				res, err := time.Parse("02/01/2006", line[2])
				if err != nil {
					return
				}
				formation = append(formation, data.Formation{
					Date: res,
					ID:   data.FormationNames[strings.Trim(line[0], " \t")],
				})
			}
		})

		diplomes := make([]data.Diploma, 0)
		e.ForEach("#ctl00_ctl00_MainContent_TabsContent_TabContainerResumeAdherent__tabFormations__formations__gvDiplomes__gvDiplomes tr", func(index int, tr *colly.HTMLElement) {
			if tr.ChildText("td:first-child") == typestr {
				return
			}
			line := make([]string, 0)
			tr.ForEach("td", func(index int, td *colly.HTMLElement) {
				line = append(line, td.DOM.Text())
			})

			if len(line) < 3 {
				return
			}
			res, err := time.Parse("02/01/2006", line[2])
			if err != nil {
				return
			}
			diplomes = append(diplomes, data.Diploma{
				Date: res,
				ID:   data.DiplomaNames[strings.Trim(line[0], " \t")],
				Ref:  strings.Trim(line[1], " \t"),
			})
		})

		person.Formations = formation
		person.Diplomas = diplomes
		s.callbackError = person.Save()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("FormationScp", r.URL)
	})
	basePost["ctl00$MainContent$_btnExporter.y"] = "13"
	basePost["ctl00$MainContent$_btnExporter.x"] = "65"
	basePost["ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent$_tabDeclarationTam$_hidAfficherEditeurDeclarationTam"] = "0"
	basePost["ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent$_tabResume$_resume$_cbAssuranceRc"] = "on"
	basePost["ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent$_tabResume$_resume$_cbAutorisationImage"] = "on"
	basePost["ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent$_tabResume$_resume$_cbAutorisationInterventionChirurgicale"] = "on"
	basePost["ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent$_tabResume$_resume$_modeleIndividu$_cbRegimeGeneral"] = "on"
	basePost["ctl00$ctl00$ScriptManager1"] = "ctl00$ctl00$_upMainContent|ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent"
	basePost["ctl00_ctl00_MainContent_TabsContent_TabContainerResumeAdherent_ClientState"] = "{\"ActiveTabIndex\":5,\"TabState\":[true,true,true,true,true,true,true,true,true,true,true]}"
	basePost["__EVENTARGUMENT"] = "activeTabChanged:5"
	basePost["__EVENTTARGET"] = "ctl00$ctl00$MainContent$TabsContent$TabContainerResumeAdherent"
	basePost["_eo_js_modules"] = ""
	basePost["_eo_obj_inst"] = ""

	err := c.Post(fullURL, basePost)
	if s.callbackError != nil {
		return s.consumeErr()
	}
	return err
}

// Export TODO
func (s *Scraper) Export() ([][]string, error) {
	c := s.c.Clone()
	url := "/Specialisation/Sgdf/adherents/ExtraireAdherents.aspx"
	var data [][]string
	if s.callbackError != nil {
		return data, s.consumeErr()
	}

	c.OnHTML("body", func(e *colly.HTMLElement) {
		postData := make(map[string]string)
		e.ForEach("form#aspnetForm input", func(index int, e *colly.HTMLElement) {
			postData[e.Attr("name")] = e.Attr("value")
		})

		data, s.callbackError = s.doExport(url, postData)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Export", r.URL)
	})

	err := c.Visit(s.baseURL + url)
	if s.callbackError != nil {
		return data, s.consumeErr()
	}
	return data, err
}

func (s *Scraper) doExport(url string, basePost map[string]string) ([][]string, error) {
	c := s.c.Clone()
	data := make([][]string, 0)
	if s.callbackError != nil {
		return data, s.consumeErr()
	}

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(index int, tr *colly.HTMLElement) {
			row := make([]string, 0)
			tr.ForEach("td", func(index int, td *colly.HTMLElement) {
				row = append(row, td.DOM.Text())
			})
			data = append(data, row)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("ExportDetails", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		r.Headers.Set("Content-Type", "text/html")
	})

	basePost["ctl00$MainContent$_btnExporter.y"] = "12"
	basePost["ctl00$MainContent$_btnExporter.x"] = "65"
	basePost["ctl00$MainContent$_cbExtraireIndividu"] = "on"
	basePost["ctl00$MainContent$_cbExtraireParents"] = "on"
	basePost["ctl00$MainContent$_cbExtraireInscription"] = "on"
	basePost["ctl00$MainContent$_cbExtraireAdhesion"] = "on"
	basePost["ctl00$MainContent$_cbExtraireJsInformations"] = "on"
	basePost["_eo_js_modules"] = ""
	basePost["_eo_obj_inst"] = ""

	err := c.Post(s.baseURL+url, basePost)
	return data, err
}
