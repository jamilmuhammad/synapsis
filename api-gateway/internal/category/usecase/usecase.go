package usecase

import (
	"context"

	category_proto "category-service/categorypb"
)

type categoryUseCase struct {
	categoryClient category_proto.CategoryServiceClient
}

func NewCategoryUseCase(categoryClient category_proto.CategoryServiceClient) *categoryUseCase {
	return &categoryUseCase{categoryClient}
}

func (u *categoryUseCase) GetAllCategories(ctx context.Context, payload *category_proto.GetAllCategoriesRequest) (*category_proto.GetAllCategoriesResponse, error) {

	result, err := u.categoryClient.GetAllCategories(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *categoryUseCase) CreateCategory(ctx context.Context, payload *category_proto.CreateCategoryRequest) (*category_proto.Category, error) {

	result, err := u.categoryClient.CreateCategory(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *categoryUseCase) GetCategoryById(ctx context.Context, payload *category_proto.GetDetailCategoryRequest) (*category_proto.GetCategoryResponse, error) {

	result, err := uc.categoryClient.GetCategoryById(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *categoryUseCase) UpdateCategory(ctx context.Context, payload *category_proto.UpdateCategoryRequest) (*category_proto.Category, error) {

	result, err := u.categoryClient.UpdateCategory(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *categoryUseCase) DeleteCategory(ctx context.Context, payload *category_proto.DeleteCategoryRequest) (*category_proto.DeleteCategoryResponse, error) {

	result, err := u.categoryClient.DeleteCategory(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}
