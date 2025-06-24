package filemime

import (
	"mime"

	"github.com/gabriel-vasile/mimetype"
)

// default provided mime type for checks. Add more using the set function
var validMimeType map[string]struct{} = map[string]struct{}{
	"image/jpeg":      {},
	"image/png":       {},
	"application/pdf": {},
	"image/heic":      {},
	"image/webp":      {},
}

func SetMimeTypeValidation(mimetype string) {
	validMimeType[mimetype] = struct{}{}
}

// IsValidFileMIMEType checks whether the MIME type for property Body is a valid MIME Type
/*
	By default, these mimes types are the one to be check on this function.
	Other than this, this function will return false.
	You can also add more mimetype for check using the SetMimeTypeValidation function. Be cautious, that the set function needs to be called first before anything else, to avoid race cond
*/
func IsValidFileMIMEType(body []byte) bool {
	mtype := mimetype.Detect(body)
	_, found := validMimeType[mtype.String()]
	return found
}

func GetFileMIMEType(body []byte) ([]string, error) {
	return mime.ExtensionsByType(mimetype.Detect(body).String())
}
