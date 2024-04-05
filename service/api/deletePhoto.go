package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	UserID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("post_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	if UserID != ctx.UserID {
		http.Error(w, UnauthorizedError, http.StatusUnauthorized)
		return
	}

	err = rt.db.DeletePost(postID, UserID)
	if os.IsNotExist(err) {
		http.Error(w, "Post not exist", http.StatusBadRequest)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting the post")
		http.Error(w, InternalServerError+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
