package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"Spazio-D/wasa-project/service/database"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	commentedID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("post_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	commentID, err := strconv.Atoi(ps.ByName("comment_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	var comment database.Comment
	comment, err = rt.db.GetCommentByID(commentID, commentedID, postID)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Comment not exist", http.StatusBadRequest)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Error getting comment")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	askingUserID := ctx.UserID

	if askingUserID != commentedID && askingUserID != comment.User.ID {
		http.Error(w, UnauthorizedError, http.StatusUnauthorized)
		return
	}

	err = rt.db.DeleteComment(commentedID, postID, commentID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting comment")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
