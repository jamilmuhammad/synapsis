package handler

import (
	"context"
	"encoding/json"
	"lib"
	"log"
	"net/http"
	"strconv"

	interfaces "api-gateway/internal/author/interface"
	"api-gateway/internal/models"
	"author-service/authorpb"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authorHandler struct {
	r          *mux.Router
	validate   *validator.Validate
	repoAuthor interfaces.AuthorInterfaceUseCase
}

func NewAuthorHandler(repoAuthor interfaces.AuthorInterfaceUseCase, r *mux.Router) *authorHandler {
	return &authorHandler{repoAuthor: repoAuthor, r: r, validate: validator.New()}
}

func (h *authorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetAllAuthors")
	defer sp.Finish()

	page := r.URL.Query().Get("page")

	limit := r.URL.Query().Get("limit")

	payload := &models.GetAllAuthorsRequest{
		Page:  page,
		Limit: limit,
	}

	lib.LogRequest(sp, payload)

	ctx := context.Background()

	err := h.validate.StructCtx(ctx, payload)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	pagePb, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	limitPb, err := strconv.ParseInt(limit, 10, 32)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	payloadPb := &author.GetAllAuthorsRequest{
		Page:  int32(pagePb),
		Limit: int32(limitPb),
	}

	result, err := h.repoAuthor.GetAllAuthors(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *authorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "CreateAuthor")
	defer sp.Finish()

	var payload models.CreateAuthorRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	lib.LogRequest(sp, payload)

	ctx := context.Background()

	err = h.validate.StructCtx(ctx, payload)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	dobTimestampProto := timestamppb.New(payload.DateOfBirth.Time)

	payloadPb := &author.CreateAuthorRequest{
		Name:        payload.Name,
		Bio:         payload.Bio,
		DateOfBirth: dobTimestampProto,
	}

	result, err := h.repoAuthor.CreateAuthor(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *authorHandler) GetAuthorById(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetAuthorById")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.GetDetailAuthorRequest{
		ID: id,
	}

	lib.LogRequest(sp, payload)

	ctx := context.Background()

	err := h.validate.StructCtx(ctx, payload)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	idPb, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	payloadPb := &author.GetDetailAuthorRequest{
		Id: idPb,
	}

	result, err := h.repoAuthor.GetAuthorById(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrNotFound(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *authorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "UpdateAuthor")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	var payload = &models.UpdateAuthorRequest{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	payload.ID = id

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	lib.LogRequest(sp, payload)

	ctx := context.Background()

	err = h.validate.StructCtx(ctx, payload)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	idPb, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	dobTimestampProto := timestamppb.New(payload.DateOfBirth.Time)

	payloadPb := &author.UpdateAuthorRequest{
		Id:          int64(idPb),
		Name:        payload.Name,
		Bio:         payload.Bio,
		DateOfBirth: dobTimestampProto,
	}

	result, err := h.repoAuthor.UpdateAuthor(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *authorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "DeleteAuthor")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.DeleteAuthorRequest{
		ID: id,
	}

	lib.LogRequest(sp, payload)

	ctx := context.Background()

	err := h.validate.StructCtx(ctx, payload)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	idPb, err := strconv.Atoi(id)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	payloadPb := &author.DeleteAuthorRequest{
		Id: int64(idPb),
	}

	result, err := h.repoAuthor.DeleteAuthor(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)
}
