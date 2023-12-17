package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	commentedID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("post_id"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	commenterID := ctx.UserID

	banCheck1, err := rt.db.IsBanned(commentedID, commenterID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	banCheck2, err := rt.db.IsBanned(commenterID, commentedID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if banCheck1 || banCheck2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var comment Comment

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	if !comment.IsValid() {
		http.Error(w, "Bad Request, the comment is invalid", http.StatusBadRequest)
		return
	}

	dbComment, err := rt.db.CreateComment(commenterID, commentedID, postID, comment.Text)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	comment.ApiConversion(dbComment)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error encoding comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
