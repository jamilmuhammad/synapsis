package delivery

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sharing_vasion_indonesia/api_gateway/internal/article/interfaces"
	"sharing_vasion_indonesia/api_gateway/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type articleDelivery struct {
	r              *mux.Router
	validate       *validator.Validate
	articleUseCase interfaces.ArticleUseCase
}

func NewDelivery(articleUseCase interfaces.ArticleUseCase, r *mux.Router) *articleDelivery {
	return &articleDelivery{articleUseCase: articleUseCase, r: r, validate: validator.New()}
}

func (d *articleDelivery) GetArticles(w http.ResponseWriter, r *http.Request) {
	limit := mux.Vars(r)["limit"]
	offset := mux.Vars(r)["offset"]

	ctx := context.Background()
	result, err := d.articleUseCase.GetArticles(ctx, limit, offset)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resByte, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resByte)
}

func (d *articleDelivery) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateArticleRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	err = d.validate.StructCtx(ctx, payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = d.articleUseCase.CreateArticle(ctx, payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}

func (d *articleDelivery) GetArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ctx := context.Background()
	result, err := d.articleUseCase.GetArticle(ctx, id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resByte, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resByte)
}

func (d *articleDelivery) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var payload models.UpdateArticleRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	payload.ID = id

	err = d.validate.StructCtx(ctx, payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = d.articleUseCase.UpdateArticle(ctx, payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}

func (d *articleDelivery) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ctx := context.Background()
	err := d.articleUseCase.DeleteArticle(ctx, id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
