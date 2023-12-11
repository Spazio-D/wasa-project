package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	targetUserID, err := strconv.Atoi(ps.ByName("target_user_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	if userID != ctx.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if targetUserID == userID {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	checkBan, err := rt.db.IsBanned(userID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	if !checkBan {
		http.Error(w, "Bad Request, user not banned", http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteBan(userID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting ban")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
