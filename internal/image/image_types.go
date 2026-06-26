package image

import "mime/multipart"

type ResizeImageRequest struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	Width  int                   `form:"width" binding:"required, min=1"`
	Height int                   `form:"height" binding:"required, min=1"`
}
