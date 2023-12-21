package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"Spazio-D/wasa-project/service/database"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	commentID, err := strconv.Atoi(ps.ByName("comment_id"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	var comment database.Comment
	comment, err = rt.db.GetCommentByID(commentID, commentedID, postID)
	if err == sql.ErrNoRows {
		http.Error(w, "Comment not exist", http.StatusBadRequest)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Error getting comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	askingUserID := ctx.UserID

	if askingUserID != commentedID && askingUserID != comment.User.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = rt.db.DeleteComment(commentedID, postID, commentID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
