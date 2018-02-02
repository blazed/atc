package teamserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"code.cloudfoundry.org/lager"

	"github.com/concourse/atc"
	"github.com/concourse/atc/api/auth"
	"github.com/concourse/atc/api/present"
)

func (s *Server) SetTeam(w http.ResponseWriter, r *http.Request) {
	hLog := s.logger.Session("set-team")

	hLog.Debug("setting-team")

	authTeam, authTeamFound := auth.GetTeam(r)
	if !authTeamFound {
		hLog.Error("failed-to-get-team-from-auth", errors.New("failed-to-get-team-from-auth"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	teamName := r.FormValue(":team_name")

	var atcTeam atc.Team
	err := json.NewDecoder(r.Body).Decode(&atcTeam)
	if err != nil {
		hLog.Error("malformed-request", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	atcTeam.Name = teamName
	if !authTeam.IsAdmin() && !authTeam.IsAuthorized(teamName) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	team, found, err := s.teamFactory.FindTeam(teamName)
	if err != nil {
		hLog.Error("failed-to-lookup-team", err, lager.Data{"teamName": teamName})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found {
		hLog.Debug("updating-credentials")
		err = team.UpdateProviderAuth(atcTeam.Auth)
		if err != nil {
			hLog.Error("failed-to-update-team", err, lager.Data{"teamName": teamName})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	} else if authTeam.IsAdmin() {
		hLog.Debug("creating team")

		team, err = s.teamFactory.CreateTeam(atcTeam)
		if err != nil {
			hLog.Error("failed-to-save-team", err, lager.Data{"teamName": teamName})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(present.Team(team))
	if err != nil {
		hLog.Error("failed-to-encode-team", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
