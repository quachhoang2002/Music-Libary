package http

import (
	"github.com/quachhoang2002/Music-Library/internal/music/usecase"
	pkgErrors "github.com/quachhoang2002/Music-Library/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(0001, "Wrong pagination query")
	errInvalidBody          = pkgErrors.NewHTTPError(0002, "Invalid body")
	errInvalidFormData      = pkgErrors.NewHTTPError(0003, "Invalid form data")
	errInvalidParamsQuery   = pkgErrors.NewHTTPError(0004, "Invalid params query")
	errInvalidValidation    = pkgErrors.NewHTTPError(0005, "Invalid validation")
	errUnauthorized         = pkgErrors.NewHTTPError(0006, "Unauthorized")
)

// map usecase error to http error
var mapError = map[error]*pkgErrors.HTTPError{
	usecase.ErrMusicTrackNotFound: pkgErrors.NewHTTPError(0007, "Music track not found"),
}
