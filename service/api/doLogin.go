package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User

	//Read the body and decode it
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request"+err.Error(), http.StatusBadRequest)
		return
	}

	//Check if the username is valid
	if !user.IsValid() {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	//Check if the username exist
	exist, err := rt.db.UsernameExist(user.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't check if the username is already taken")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*If the username is not taken, create the user
	else do the login*/
	if !exist {
		user, err = rt.CreateUser(user)
		if err != nil {
			ctx.Logger.WithError(err).Error("Can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		dbUser, err := rt.db.GetUserByUsername(user.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("Can't get the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.ApiConversion(dbUser)
		w.WriteHeader(http.StatusOK)
	}

	authUser := AuthUser{user, user.ID}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(authUser); err != nil {
		ctx.Logger.WithError(err).Error("Can't encode the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
