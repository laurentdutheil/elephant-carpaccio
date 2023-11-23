package http_server

import (
	"elephant_carpaccio/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleApi(router *http.ServeMux, game *domain.Game) {
	router.Handle("/api/register", handleRegistration(game))
}

type RegistrationRequestBody struct {
	TeamName string `json:"teamName"`
	IP       string `json:"ip"`
}

func handleRegistration(game *domain.Game) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			decoder := json.NewDecoder(r.Body)
			var requestBody RegistrationRequestBody
			err := decoder.Decode(&requestBody)
			if err != nil || requestBody.TeamName == "" || requestBody.IP == "" {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = fmt.Fprint(w, "the body of your request don't respect {\"teamName\":\"<your team name>\", \"ip\":\"<your api address>\"}")
				return
			}

			existingTeam := game.FindTeamByName(requestBody.TeamName)
			if existingTeam != nil {
				existingTeam.SetIp(requestBody.IP)
				w.WriteHeader(http.StatusOK)
			} else {
				game.Register(requestBody.TeamName, requestBody.IP)
				w.WriteHeader(http.StatusCreated)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, "please use PUT method to register/update your team")
		}
	})
}
