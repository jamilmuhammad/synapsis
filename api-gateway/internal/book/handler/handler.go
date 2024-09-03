package handler

import (
	"context"
	"encoding/json"
	"lib"
	"net/http"
	"strconv"

	interfaces "api-gateway/internal/book/interface"
	"api-gateway/internal/models"
	"book-service/bookpb/book"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type bookHandler struct {
	r        *mux.Router
	validate *validator.Validate
	repoBook interfaces.BookInterfaceUseCase
}

func NewBookHandler(repoBook interfaces.BookInterfaceUseCase, r *mux.Router) *bookHandler {
	return &bookHandler{repoBook: repoBook, r: r, validate: validator.New()}
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetAllBooks")
	defer sp.Finish()

	page := r.URL.Query().Get("page")

	limit := r.URL.Query().Get("limit")

	payload := &models.GetAllBooksRequest{
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

	payloadPb := &book.GetAllBooksRequest{
		Page:  int32(pagePb),
		Limit: int32(limitPb),
	}

	result, err := h.repoBook.GetAllBooks(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "CreateBook")
	defer sp.Finish()

	var payload models.CreateBookRequest

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

	dopTimestampProto := timestamppb.New(payload.DateOfPublication.Time)

	payloadPb := &book.CreateBookRequest{
		Title:             payload.Title,
		Isbn:              payload.Isbn,
		DateOfPublication: dopTimestampProto,
		AuthorId:          payload.AuthorId,
		CategoryId:        payload.CategoryId,
	}

	result, err := h.repoBook.CreateBook(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *bookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetBookById")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.GetDetailBookRequest{
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

	payloadPb := &book.GetDetailBookRequest{
		Id: idPb,
	}

	result, err := h.repoBook.GetBookById(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrNotFound(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "UpdateBook")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	var payload = &models.UpdateBookRequest{}

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

	dopTimestampProto := timestamppb.New(payload.DateOfPublication.Time)

	payloadPb := &book.UpdateBookRequest{
		Id:                idPb,
		Title:             payload.Title,
		Isbn:              payload.Isbn,
		DateOfPublication: dopTimestampProto,
		AuthorId:          payload.AuthorId,
		CategoryId:        payload.CategoryId,
	}

	result, err := h.repoBook.UpdateBook(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "DeleteBook")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.DeleteBookRequest{
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

	payloadPb := &book.DeleteBookRequest{
		Id: int64(idPb),
	}

	result, err := h.repoBook.DeleteBook(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)
}
