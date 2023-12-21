package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"Spazio-D/wasa-project/service/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	askingUserID := ctx.UserID

	banCheck1, err := rt.db.IsBanned(askingUserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if user is banned")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}
	banCheck2, err := rt.db.IsBanned(userID, askingUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if user is banned")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	if banCheck1 || banCheck2 {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var user database.User
	user, err = rt.db.GetUserByID(userID)
	if err == sql.ErrNoRows {
		http.Error(w, "User not exist", http.StatusBadRequest)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Error getting user")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	dbPosts, err := rt.db.GetPosts(askingUserID, user, 0, 1000)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting posts")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	posts := make([]Post, len(dbPosts))

	for i, dbPost := range dbPosts {
		var post Post
		err = post.ApiConversion(dbPost)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error converting post")
			http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
			return
		}
		
		posts[i] = post
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
}
