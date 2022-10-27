package urls

const (
	comments         = "/comments"
	POST_COMMENTS    = comments
	GET_COMMENTS     = comments
	PATH_COMMENTS_ID = "commentId"
	PUT_COMMENT      = comments + "/:" + PATH_COMMENTS_ID
	DELETE_COMMENT   = comments + "/:" + PATH_COMMENTS_ID
)
