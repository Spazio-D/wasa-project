package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := r.URL.Query().Get("username")
	if !regexp.MustCompile(`^\w{1,16}$`).MatchString(username) {
		http.Error(w, BadRequestError, http.StatusBadRequest)
		return
	}

	dbUsers, err := rt.db.SearchUsers(ctx.UserID, username)
	if err != nil {
		ctx.Logger.Error("Error during search ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	users := make([]User, len(dbUsers))
	for i, dbUser := range dbUsers {
		var user User
		user.ApiConversion(dbUser)
		if err != nil {
			ctx.Logger.Error("Error during conversion ", err)
			http.Error(w, InternalServerError, http.StatusInternalServerError)
			return
		}
		users[i] = user
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		ctx.Logger.Error("Error during encoding ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

}
