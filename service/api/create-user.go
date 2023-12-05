package api

func (rt *_router) CreateUser(user User) (User, error) {

	dbUser, err := rt.db.CreateUser(user.DatabaseConversion())
	if err != nil {
		return user, err
	}

	user.ApiConversion(dbUser)

	return user, nil
}
