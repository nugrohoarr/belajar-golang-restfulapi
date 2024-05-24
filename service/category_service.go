package service

import (
	"context"
	"restfullapi/golang/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delate(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error)
	FindAll(ctx context.Context) []web.CategoryResponse
}
