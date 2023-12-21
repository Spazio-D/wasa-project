package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	likedUserID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("post_id"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	likerUserID, err := strconv.Atoi(ps.ByName("liker_id"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID
	if userID != likerUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	likeCheck, err := rt.db.IsLiked(postID, likedUserID, likerUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking like")
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}

	if !likeCheck {
		http.Error(w, "Not liked", http.StatusBadRequest)
		return
	}

	banCheck1, err := rt.db.IsBanned(likerUserID, likedUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	banCheck2, err := rt.db.IsBanned(likedUserID, likerUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if banCheck1 || banCheck2 {
		http.Error(w, "Forbidden ", http.StatusForbidden)
		return
	}

	err = rt.db.DeleteLike(likerUserID, likedUserID, postID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting like")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
