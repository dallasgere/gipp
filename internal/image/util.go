package image

import "slices"

// Starter list that I need to explore more
var allowedImageTypes = []string{"jpeg", "jpg", "png", "tiff", "gif"}

// Helper to check if the image type is allowed to be processed by gipp
func IsImageTypeAllowed(imageType string) bool {
	return slices.Contains(allowedImageTypes, imageType)
}
