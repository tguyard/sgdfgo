package serv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sgdfgo/data"
	"github.com/sgdfgo/user"
	"github.com/sgdfgo/utils"
)

type errorMsg struct {
	message string
}
type ConnectionRqt struct {
	Login, Password string
}

func connect(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if r.Method != "POST" {
		w.WriteHeader(405)
		encoder.Encode(errorMsg{"Method Not Allowed"})
		return
	}

	var request ConnectionRqt
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(&errorMsg{"Wrong request"})
		return
	}

	user, err := user.Connect(request.Login, request.Password)
	if err != nil {
		w.WriteHeader(403)
		encoder.Encode(&errorMsg{"Connexion error"})
		return
	}

	structure, err := data.NewStructureByID(user.StructureID)
	if err != nil {
		w.WriteHeader(202)
		encoder.Encode(&errorMsg{"Accepted, but retry later (we need to update data, it may take a long time)"})
		return
	}

	if structure.TypeName() != "Territoire" {
		w.WriteHeader(503)
		encoder.Encode(&errorMsg{"On accepte pour l'instant que les membres d'un territoire."})
		return
	}

	t := time.Now().Format("20060102150405")
	data, err := json.Marshal(map[string]interface{}{
		"l": user.ID,
		"p": request.Password,
		"s": user.StructureID,
	})
	if err != nil {
		w.WriteHeader(500)
		encoder.Encode(&errorMsg{"Connexion error"})
		return
	}
	edata, err := utils.EncryptAes(string(data), Secret)
	if err != nil {
		w.WriteHeader(500)
		encoder.Encode(&errorMsg{"Connexion error"})
		return
	}

	signature := utils.HmacSign([]byte(t+edata), Secret)

	value := fmt.Sprintf("t=%s&data=%s&s=%s",
		url.QueryEscape(t),
		url.QueryEscape(edata),
		url.QueryEscape(signature),
	)

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   value,
		Expires: time.Now().Add(365 * 24 * time.Hour),
	})

	encoder.Encode(user)

	// message := r.URL.Path
	// message = strings.TrimPrefix(message, "/")
	// message = "Hello " + message

}
