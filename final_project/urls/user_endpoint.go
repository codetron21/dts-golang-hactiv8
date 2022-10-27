package urls

const (
	users               = "/users"
	POST_USERS_REGISTER = users + "/register"
	POST_USERS_LOGIN    = users + "/login"
	PATH_USER_ID        = "userId"
	PUT_USER            = users + "/:" + PATH_USER_ID
	DELETE_USER         = users + "/:" + PATH_USER_ID
)
