package interfaces

import (
	"context"
	
	category_proto "category-service/categorypb"
)

type CategoryInterfaceUseCase interface {
	GetAllCategories(ctx context.Context, payload *category_proto.GetAllCategoriesRequest) (*category_proto.GetAllCategoriesResponse, error)
	CreateCategory(ctx context.Context, post *category_proto.CreateCategoryRequest) (*category_proto.Category, error)
	UpdateCategory(ctx context.Context, post *category_proto.UpdateCategoryRequest) (*category_proto.Category, error)
	GetCategoryById(ctx context.Context, payload *category_proto.GetDetailCategoryRequest) (*category_proto.GetCategoryResponse, error)
	DeleteCategory(ctx context.Context, id *category_proto.DeleteCategoryRequest) (*category_proto.DeleteCategoryResponse, error)
}
