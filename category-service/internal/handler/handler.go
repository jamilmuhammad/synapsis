package handler

import (
	"context"
	"log"

	category_proto "category-service/categorypb"
	"category-service/internal/interfaces"
)

type CategoryHandler struct {
	category_proto.UnimplementedCategoryServiceServer
	repoCategory interfaces.CategoryInterfaceHandler
}

func NewHandler(repoCategory interfaces.CategoryInterfaceHandler) *CategoryHandler {
	return &CategoryHandler{repoCategory: repoCategory}
}

func (uc *CategoryHandler) GetAllCategories(ctx context.Context, payload *category_proto.GetAllCategoriesRequest) (*category_proto.GetAllCategoriesResponse, error) {
	result, err := uc.repoCategory.GetAllCategories(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *CategoryHandler) CreateCategory(ctx context.Context, payload *category_proto.CreateCategoryRequest) (*category_proto.Category, error) {

	result, err := uc.repoCategory.CreateCategory(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *CategoryHandler) GetCategoryById(ctx context.Context, payload *category_proto.GetDetailCategoryRequest) (*category_proto.GetCategoryResponse, error) {

	result, err := uc.repoCategory.GetCategoryById(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *CategoryHandler) UpdateCategory(ctx context.Context, payload *category_proto.UpdateCategoryRequest) (*category_proto.Category, error) {

	result, err := uc.repoCategory.UpdateCategory(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *CategoryHandler) DeleteCategory(ctx context.Context, id *category_proto.DeleteCategoryRequest) (*category_proto.DeleteCategoryResponse, error) {

	result, err := uc.repoCategory.DeleteCategory(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}
