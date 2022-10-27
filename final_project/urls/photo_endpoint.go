package urls

const (
	photos        = "/photos"
	POST_PHOTOS   = photos
	GET_PHOTOS    = photos
	PATH_PHOTO_ID = "photoId"
	PUT_PHOTO     = photos + "/:" + PATH_PHOTO_ID
	DELETE_PHOTO  = photos + "/:" + PATH_PHOTO_ID
)
