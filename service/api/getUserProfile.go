package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var profile Profile
	targetUserID, err := strconv.Atoi(ps.ByName("user_id"))

	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	askingUserID := ctx.UserID

	banCheck, err := rt.db.IsBanned(askingUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}
	if banCheck {
		dbUser, err := rt.db.GetUserByID(targetUserID)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error getting user profile")
			http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
			return
		}
		profile.User.ApiConversion(dbUser)
		profile.IsBanned = true
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(profile); err != nil {
			ctx.Logger.WithError(err).Error("Error marshaling json")
			http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	banCheck, err = rt.db.IsBanned(targetUserID, askingUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}
	if banCheck {
		dbUser, err := rt.db.GetUserByID(targetUserID)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error getting user profile")
			http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
			return
		}
		profile.User.ApiConversion(dbUser)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(profile); err != nil {
			ctx.Logger.WithError(err).Error("Error marshaling json")
			http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	dbUserProfile, err := rt.db.GetUserProfile(targetUserID, askingUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting user profile")
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}

	profile.ApiConversion(dbUserProfile)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		ctx.Logger.WithError(err).Error("Error marshaling json")
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}

}
