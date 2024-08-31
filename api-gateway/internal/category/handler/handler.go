package handler

import (
	"context"
	"encoding/json"
	"lib"
	"net/http"
	"strconv"

	interfaces "api-gateway/internal/category/interface"
	"api-gateway/internal/models"
	"category-service/categorypb"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type categoryHandler struct {
	r            *mux.Router
	validate     *validator.Validate
	repoCategory interfaces.CategoryInterfaceUseCase
}

func NewCategoryHandler(repoCategory interfaces.CategoryInterfaceUseCase, r *mux.Router) *categoryHandler {
	return &categoryHandler{repoCategory: repoCategory, r: r, validate: validator.New()}
}

func (h *categoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetAllCategories")
	defer sp.Finish()

	page := r.URL.Query().Get("page")

	limit := r.URL.Query().Get("limit")

	payload := &models.GetAllCategoriesRequest{
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

	payloadPb := &category.GetAllCategoriesRequest{
		Page:  int32(pagePb),
		Limit: int32(limitPb),
	}

	result, err := h.repoCategory.GetAllCategories(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "CreateCategory")
	defer sp.Finish()

	var payload models.CreateCategoryRequest

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

	payloadPb := &category.CreateCategoryRequest{
		Name: payload.Name,
	}

	result, err := h.repoCategory.CreateCategory(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *categoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetCategoryById")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.GetDetailCategoryRequest{
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

	payloadPb := &category.GetDetailCategoryRequest{
		Id: idPb,
	}

	result, err := h.repoCategory.GetCategoryById(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrNotFound(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "UpdateCategory")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	var payload = &models.UpdateCategoryRequest{}

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

	payloadPb := &category.UpdateCategoryRequest{
		Id:   int64(idPb),
		Name: payload.Name,
	}

	result, err := h.repoCategory.UpdateCategory(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "DeleteCategory")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.DeleteCategoryRequest{
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

	payloadPb := &category.DeleteCategoryRequest{
		Id: int64(idPb),
	}

	result, err := h.repoCategory.DeleteCategory(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)
}
