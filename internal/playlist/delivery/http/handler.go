package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanhoang/music-library/internal/playlist/usecase"
	pkgErrors "github.com/xuanhoang/music-library/pkg/errors"
	"github.com/xuanhoang/music-library/pkg/jwt"
	"github.com/xuanhoang/music-library/pkg/paginator"
	"github.com/xuanhoang/music-library/pkg/response"
)

// @Summary Create Playlist
// @Schemes
// @Description Create Playlist
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param user_id path string true "user_id"
// @Param body body createPlaylistReq true "body"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object} itemPlaylistResp
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id} [POST]
func (h handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req createPlaylistReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.Create.ShouldBind: %v", err)
		response.Error(c, errInvalidFormData)
		return
	}

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.Create.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	o, err := h.uc.Create(ctx, sc, req.toInput(c.Param("user_id")))
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.Create.uc.Create: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newItemPlaylistResp(o))
}

// @Summary Update Playlist
// @Schemes
// @Description Update Playlist
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Param user_id path string true "user_id"
// @Param body body updatePlayListReq true "body"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object} itemPlaylistResp
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id}/{id} [PUT]
func (h handler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var req updatePlayListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.Update.ShouldBind: %v", err)
		response.Error(c, errInvalidFormData)
		return
	}

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.Update.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	o, err := h.uc.Update(ctx, sc, req.toInput(c.Param("id"), c.Param("user_id")))
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.Update.uc.Update: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newItemPlaylistResp(o))
}

// @Summary Delete Playlist
// @Schemes
// @Description Delete Playlist
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Param user_id path string true "user_id"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object} response.Resp
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id}/{id} [DELETE]
func (h handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.Delete.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	err := h.uc.Delete(ctx, sc, c.Param("id"))
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.Delete.uc.Delete: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, nil)
}

// @Summary Detail Playlist
// @Schemes
// @Description Detail Playlist
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object} detailPlaylistRes
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id}/{id} [GET]
func (h handler) Detail(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.Detail.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	playlist, tracks, err := h.uc.Detail(ctx, sc, c.Param("id"))
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.Detail.uc.Detail: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newDetailPlaylistRes(playlist, tracks))
}

// @Summary List Playlist
// @Schemes
// @Description List Playlist
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Param artist query string false "artist"
// @Param title query string false "title"
// @Param album query string false "album"
// @Param user_id path string true "user_id"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object} listPlaylistResp
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id} [GET]
func (h handler) List(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.List.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	var req listPlaylistReq
	if err := c.ShouldBind(&req); err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.List.ShouldBind: %v", err)
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

	playlists, err := h.uc.List(ctx, sc, usecase.ListInput{
		TrackFilter: usecase.TrackFilter{
			Title:  req.Title,
			Artist: req.Artist,
			Album:  req.Album,
		},
		UserID:         c.Param("user_id"),
		PaginatorQuery: pagQuery,
	})
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.List.uc.List: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, newPlaylistsResp(playlists))
}

// @Summary Playlist Add Track
// @Schemes
// @Description Playlist Add Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Param track_id path string true "track_id"
// @Param user_id path string true "user_id"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object}  response.Resp
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id}/{id}/tracks/{track_id} [POST]
func (h handler) AddTrack(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.AddTrack.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	err := h.uc.AddTrack(ctx, sc, c.Param("id"), c.Param("track_id"))
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.AddTrack.uc.AddTrack: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, nil)
}

// @Summary Playlist Remove Track
// @Schemes
// @Description Playlist Remove Track
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjYyYzc3MzU1MmRmMzZjMGJkMjUxNDdkIiwiZ3JvdXBfaWQiOiI2NGY0NDUyNTlkODNkM2JkZDg0ZGZjOWEiLCJncm91cF9yb2xlIjoiYWRtaW4iLCJleHBpcmVkX2F0IjoiMjAyNC0wNS0yN1QxMDo1Nzo0Ny40Mjg1NTgrMDc6MDAiLCJleHAiOjE3MTY3ODIyNjd9.pb6sLIq4F2FDXE2ASWacYRzI5qs1ae48_DeQ2b3jJLU)"
// @Param Language header string false "Language" default(en)
// @Param id path string true "id"
// @Param track_id path string true "track_id"
// @Param user_id path string true "user_id"
// @Produce json
// @Tags Playlist
// @Accept json
// @Produce json
// @Success 200 {object}  response.Resp
// @Failure 400 {object} response.Resp "Bad Request,Error..."
// @Router /api/v1/playlists/{user_id}/{id}/tracks/{track_id} [DELETE]
func (h handler) RemoveTrack(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "playlist.http.handler.RemoveTrack.GetPayloadFromContext: %v", errUnauthorized)
		// response.Error(c, errUnauthorized)
	}
	sc := jwt.NewScope(payload)

	err := h.uc.RemoveTrack(ctx, sc, c.Param("id"), c.Param("track_id"))
	if err != nil {
		h.l.Warnf(ctx, "playlist.http.handler.RemoveTrack.uc.RemoveTrack: %v", err)
		response.ErrorWithMap(c, err, mapError)
		return
	}

	response.OK(c, nil)
}
