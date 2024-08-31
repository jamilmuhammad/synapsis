package usecase

import (
	"context"
	"lib"
	"log"

	"category-service/internal/interfaces"
	category_proto "category-service/categorypb"
)

type CategoryUseCase struct {
	repoCategory interfaces.CategoryInterfaceUseCase
}

func NewUseCase(repoCategory interfaces.CategoryInterfaceUseCase) *CategoryUseCase {
	return &CategoryUseCase{
		repoCategory,
	}
}

func (uc *CategoryUseCase) GetAllCategories(ctx context.Context, payload *category_proto.GetAllCategoriesRequest) (*category_proto.GetAllCategoriesResponse, error) {
	result, err := uc.repoCategory.GetAllCategories(ctx, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, payload *category_proto.CreateCategoryRequest) (*category_proto.Category, error) {

	result, err := uc.repoCategory.CreateCategory(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *CategoryUseCase) GetCategoryById(ctx context.Context, payload *category_proto.GetDetailCategoryRequest) (*category_proto.GetCategoryResponse, error) {

	result, err := uc.repoCategory.GetCategoryById(ctx, payload)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return &category_proto.GetCategoryResponse{}, lib.NewErrNotFound("category not found")
	}

	return result, nil
}

func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, payload *category_proto.UpdateCategoryRequest) (*category_proto.Category, error) {

	result, err := uc.repoCategory.UpdateCategory(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, id *category_proto.DeleteCategoryRequest) (*category_proto.DeleteCategoryResponse, error) {

	result, err := uc.repoCategory.DeleteCategory(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil

}
