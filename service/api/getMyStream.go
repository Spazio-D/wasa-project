package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if userID != ctx.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dbStream, err := rt.db.GetStream(userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting stream")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	posts := make([]Post, len(dbStream))

	for i, dbPost := range dbStream {
		var post Post
		err = post.ApiConversion(dbPost)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error converting post")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		posts[i] = post
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
