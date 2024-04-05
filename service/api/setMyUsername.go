package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	if userID != ctx.UserID {
		http.Error(w, UnauthorizedError, http.StatusUnauthorized)
		return
	}

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	if !user.IsValid() {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	exist, err := rt.db.UsernameExist(user.Username)
	if err != nil {
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}

	if exist {
		http.Error(w, "Username already exist", http.StatusBadRequest)
		return
	}

	if err := rt.db.ChangeUsername(user.Username, userID); err != nil {
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	if err := json.NewEncoder(w).Encode("Username changed"); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
