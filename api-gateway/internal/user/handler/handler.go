package handler

import (
	"api-gateway/internal/models"
	interfaces "api-gateway/internal/user/interface"
	"context"
	"encoding/json"
	"lib"
	"log"
	"net/http"
	"strconv"
	"user-service/userpb"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type userHandler struct {
	r        *mux.Router
	validate *validator.Validate
	repoUser interfaces.UserInterfaceUseCase
}

func NewUserHandler(repoUser interfaces.UserInterfaceUseCase, r *mux.Router) *userHandler {
	return &userHandler{repoUser: repoUser, r: r, validate: validator.New()}
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetAllUsers")
	defer sp.Finish()

	page := r.URL.Query().Get("page")

	limit := r.URL.Query().Get("limit")

	payload := &models.GetAllUsersRequest{
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

	payloadPb := &user.GetAllUsersRequest{
		Page:  int32(pagePb),
		Limit: int32(limitPb),
	}

	result, err := h.repoUser.GetAllUsers(ctx, payloadPb)

	if err != nil {
		log.Println(err)
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "CreateUser")
	defer sp.Finish()

	var payload models.CreateUserRequest

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

	payloadPb := &user.CreateUserRequest{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     payload.Role,
		Status:   payload.Status,
	}

	result, err := h.repoUser.CreateUser(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, err, nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "GetUserById")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.GetDetailUserRequest{
		ID: id,
	}

	log.Println(payload, "payload")

	lib.LogRequest(sp, payload)

	ctx := context.Background()

	err := h.validate.StructCtx(ctx, payload)

	if err != nil {
		log.Println(err, "validate struct")
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	idPb, err := strconv.ParseInt(id, 10, 64)
	log.Println(idPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	payloadPb := &user.GetDetailUserRequest{
		Id: idPb,
	}

	result, err := h.repoUser.GetUserById(ctx, payloadPb)
	log.Println(result, "result")

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrNotFound(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "UpdateUser")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	var payload = &models.UpdateUserRequest{}

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

	// idPb, err := strconv.Atoi(mux.Vars(r)["id"])
	idPb, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	payloadPb := &user.UpdateUserRequest{
		Id:       int64(idPb),
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     payload.Role,
		Status:   payload.Status,
	}

	result, err := h.repoUser.UpdateUser(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)

}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "DeleteUser")
	defer sp.Finish()

	id := mux.Vars(r)["id"]

	payload := &models.DeleteUserRequest{
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

	payloadPb := &user.DeleteUserRequest{
		Id: int64(idPb),
	}

	result, err := h.repoUser.DeleteUser(ctx, payloadPb)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "Login")
	defer sp.Finish()

	var payload models.LoginUserRequest

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

	payloadPb := &user.LoginRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}

	result, err := h.repoUser.Login(ctx, payloadPb)

	log.Println(result)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)
}

func (h *userHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	sp := lib.CreateRootSpan(r, "RefreshToken")
	defer sp.Finish()

	var payload models.RefreshTokenRequest

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

	payloadPb := &user.RefreshTokenRequest{
		RefreshToken: payload.RefreshToken,
	}

	result, err := h.repoUser.RefreshToken(ctx, payloadPb)

	log.Println(result)

	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	lib.WriteResponse(sp, w, nil, result)
}
