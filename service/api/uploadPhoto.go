package api

import (
	"Spazio-D/wasa-project/service/api/reqcontext"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	if userID != ctx.UserID {
		http.Error(w, UnauthorizedError, http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(30000000)
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, BadRequestError+err.Error(), http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't read the file")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" {
		http.Error(w, "Bad Request, wrong file format "+err.Error(), http.StatusBadRequest)
		return
	}

	defer func() { err = file.Close() }()

	dbUser, err := rt.db.GetUserByID(userID)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "User not exist", http.StatusBadRequest)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Can't get the user")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	var user User
	user.ApiConversion(dbUser)

	var post = Post{
		User: user,
	}

	dbPost := post.DatabaseConversion()
	dbPost, err = rt.db.CreatePost(dbPost, data)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't create the post")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	err = post.ApiConversion(dbPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't convert the post")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		ctx.Logger.WithError(err).Error("Can't encode the post")
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
}
