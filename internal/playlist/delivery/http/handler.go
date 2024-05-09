package http

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuanhoang/music-library/internal/playlist/usecase"
	pkgErrors "github.com/xuanhoang/music-library/pkg/errors"
	"github.com/xuanhoang/music-library/pkg/jwt"
	"github.com/xuanhoang/music-library/pkg/paginator"
	"github.com/xuanhoang/music-library/pkg/response"
)

// @Summary Create Track
// @Schemes
// @Description Create Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param title formData string true "Title of the track"
// @Param artist formData string true "Artist of the track"
// @Param album formData string true "Album of the track"
// @Param genre formData string true "Genre of the track"
// @Param release_year formData int true "Release Year of the track"
// @Param duration formData int true "Duration of the track in seconds"
// @Param mp3_file formData file true "MP3 file of the track"
// @Produce json
// @Tags Music Track
// @Accept json
// @Produce json
// @Success 200 {object} detailTrackRes
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/music-tracks [post]
func (h handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req createMusicTrackReq
	if err := c.ShouldBind(&req); err != nil {
		h.l.Warnf(ctx, "music.http.handler.Create.ShouldBind: %v", err)
		response.Error(c, errInvalidFormData)
		return
	}

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "music.http.handler.Create.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	o, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "music.http.handler.Create.Create: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newDetailTrackRes(o))
}

// @Summary Update Track
// @Schemes
// @Description Update Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param title formData string true "Title of the track"
// @Param artist formData string true "Artist of the track"
// @Param album formData string true "Album of the track"
// @Param genre formData string true "Genre of the track"
// @Param release_year formData int true "Release Year of the track"
// @Param duration formData int true "Duration of the track in seconds"
// @Param mp3_file formData file true "MP3 file of the track"
// @Produce json
// @Tags Music Track
// @Accept json
// @Produce json
// @Success 200 {object} detailTrackRes
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/music-tracks/{id} [PUT]
func (h handler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var req updateMusicTrackReq
	if err := c.ShouldBind(&req); err != nil {
		h.l.Warnf(ctx, "music.http.handler.Update.ShouldBind: %v", err)
		response.Error(c, errInvalidFormData)
		return
	}
	req.ID = c.Param("id")

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "music.http.handler.Update.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	o, err := h.uc.Update(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "music.http.handler.Update.Update: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newDetailTrackRes(o))
}

// @Summary Delete Track
// @Schemes
// @Description Delete Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Produce json
// @Tags Music Track
// @Accept json
// @Produce json
// @Success 200 {object} detailTrackRes
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/music-tracks/{id} [DELETE]
func (h handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "music.http.handler.Update.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	err := h.uc.Delete(ctx, sc, c.Param("id"))
	if err != nil {
		h.l.Warnf(ctx, "music.http.handler.Update.Update: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, nil)
}

// @Summary Detail Track
// @Schemes
// @Description Detail Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Produce json
// @Tags Music Track
// @Accept json
// @Produce json
// @Success 200 {object} detailTrackRes
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/music-tracks/{id} [GET]
func (h handler) Detail(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "music.http.handler.Update.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	track, err := h.uc.Detail(ctx, sc, c.Param("id"))
	if err != nil {
		h.l.Warnf(ctx, "music.http.handler.Update.Update: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newDetailTrackRes(track))
}

// @Summary List Track
// @Schemes
// @Description List Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Param artist query string false "artist"
// @Param title query string false "title"
// @Param album query string false "album"
// @Produce json
// @Tags Music Track
// @Accept json
// @Produce json
// @Success 200 {object} detailTrackRes
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/music-tracks [GET]
func (h handler) List(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "music.http.handler.Update.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	var req listMusicTrackReq
	if err := c.ShouldBind(&req); err != nil {
		h.l.Warnf(ctx, "music.http.handler.List.ShouldBind: %v", err)
		response.Error(c, errInvalidFormData)
		return
	}

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Warn(ctx, "service.http.Get: invalid request")
		response.Error(c, pkgErrors.NewBadRequestHTTPError())
		return
	}
	pagQuery.Adjust()

	track, err := h.uc.List(ctx, sc, usecase.ListInput{
		Filter: usecase.Filter{
			Title:  req.Title,
			Artist: req.Artist,
			Album:  req.Album,
		},
		PaginatorQuery: pagQuery,
	})
	if err != nil {
		h.l.Warnf(ctx, "music.http.handler.Update.Update: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newListTrackResp(track))
}

// @Summary Get File
// @Schemes
// @Description Get File
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Produce json
// @Tags Music Track
// @Accept json
// @Produce json
// @Success 200 {object}
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/music-tracks/{id}/file [GET]
func (h handler) GetFile(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "music.http.handler.Update.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	f, err := h.uc.GetFile(ctx, sc, c.Param("id"))
	if err != nil {
		h.l.Warnf(ctx, "music.http.handler.Update.Update: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}
	fname := f.Name()
	ext := strings.Split(fname, ".")
	c.Header("Content-Disposition", "inline; filename="+fname)
	c.Header("Content-Type", ContentType(ext[len(ext)-1]))
	c.File(f.Name())
}

func ContentType(extension string) string {
	switch extension {
	case "jpg", "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	case "pdf":
		return "application/pdf"
	case "doc":
		return "application/msword"
	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case "xls":
		return "application/vnd.ms-excel"
	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case "ppt":
		return "application/vnd.ms-powerpoint"
	case "pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case "mp4":
		return "video/mp4"
	case "mp3":
		return "audio/mpeg"
	case "wav":
		return "audio/wav"
	case "zip":
		return "application/zip"
	case "rar":
		return "application/x-rar-compressed"
	case "7z":
		return "application/x-7z-compressed"
	case "txt":
		return "text/plain"
	case "csv":
		return "text/csv"
	case "json":
		return "application/json"
	case "xml":
		return "application/xml"
	case "html":
		return "text/html"
	case "css":
		return "text/css"
	}

	return "application/octet-stream"
}
