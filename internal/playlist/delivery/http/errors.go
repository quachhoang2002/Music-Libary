package http

import (
	"github.com/xuanhoang/music-library/internal/playlist/usecase"
	pkgErrors "github.com/xuanhoang/music-library/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(101, "Wrong pagination query")
	errInvalidBody          = pkgErrors.NewHTTPError(102, "Invalid body")
	errInvalidFormData      = pkgErrors.NewHTTPError(103, "Invalid form data")
	errInvalidParamsQuery   = pkgErrors.NewHTTPError(104, "Invalid params query")
	errInvalidValidation    = pkgErrors.NewHTTPError(105, "Invalid validation")
	errUnauthorized         = pkgErrors.NewHTTPError(106, "Unauthorized")
)

// map usecase error to http error
var mapError = map[error]*pkgErrors.HTTPError{
	usecase.ErrPlaylistNotFound: pkgErrors.NewHTTPError(107, "Playlist not found"),
	usecase.ErrTrackNotFound:    pkgErrors.NewHTTPError(108, "Music track not found"),
}
