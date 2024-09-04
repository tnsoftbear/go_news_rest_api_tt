package repository

import "frr-news/internal/core/domain/model"

type NewsRepository interface {
	AssignCategory(newsID, catID int64)
	DeleteNews(news *model.News) *model.News
	DeleteNewsById(newsID int64) *model.News
	FindByID(id int64) *model.News
	LoadCategoryIDs(newsID int64) []int64
	LoadPagenated(page, perPage int) []*model.News
	Save(*model.News) *model.News
	UnassignCategories(newsID int64, skipIDs []int64)
}
