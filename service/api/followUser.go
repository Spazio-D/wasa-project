package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	followerID, err := strconv.Atoi(ps.ByName("follower_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	followedID, err := strconv.Atoi(ps.ByName("followed_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	if followerID != ctx.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if followerID == followedID {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	followCheck, err := rt.db.IsFollowing(followerID, followedID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if user is following")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	if followCheck {
		http.Error(w, "Bad Request, already following", http.StatusBadRequest)
		return
	}

	banCheck1, err := rt.db.IsBanned(followerID, followedID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	banCheck2, err := rt.db.IsBanned(followedID, followerID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking for ban")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	if banCheck1 || banCheck2 {
		http.Error(w, "Bad Request, user already banned", http.StatusBadRequest)
		return
	}

	err = rt.db.CreateFollow(followerID, followedID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating follow")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
