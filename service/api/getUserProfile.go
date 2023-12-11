package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	targetUserID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	askingUserID := ctx.UserID

	banCheck, err := rt.db.IsBanned(askingUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	if banCheck {
		http.Error(w, "Bad Request, you have banned this user", http.StatusBadRequest)
		return
	}

	banCheck, err = rt.db.IsBanned(targetUserID, askingUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	if banCheck {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	dbUserProfile, err := rt.db.GetUserProfile(targetUserID, askingUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting user profile")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	var profile Profile
	profile.ApiConversion(dbUserProfile)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		ctx.Logger.WithError(err).Error("Error marshaling json")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

}
