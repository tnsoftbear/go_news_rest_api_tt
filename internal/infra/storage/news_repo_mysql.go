package storage

import (
	"fmt"
	"log"
	"strings"

	"frr-news/internal/core/domain/model"

	"gopkg.in/reform.v1"
)

type NewsRepositoryMysql struct {
	db *reform.DB
}

func NewNewsRepositoryMysql(reformDB *reform.DB) *NewsRepositoryMysql {
	return &NewsRepositoryMysql{
		db: reformDB,
	}
}

func (r *NewsRepositoryMysql) LoadPagenated(page, perPage int) []*model.News {
	offset := (page - 1) * perPage
	tail := fmt.Sprintf(" ORDER BY id LIMIT %s, %s", r.db.Placeholder(1), r.db.Placeholder(2))
	rows, err := r.db.SelectRows(model.NewsTable, tail, offset, perPage)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []*model.News
	for {
		var news model.News
		if err = r.db.NextRow(&news, rows); err != nil {
			break
		}
		news.Categories = r.LoadCategoryIDs(news.ID)
		results = append(results, &news)
	}
	if err != reform.ErrNoRows {
		log.Print(err)
		return nil;
	}

	return results
}

func (r *NewsRepositoryMysql) LoadCategoryIDs(newsID int64) []int64 {
	tail := fmt.Sprintf("WHERE NewsId = %s", r.db.Placeholder(1))
	rows, err := r.db.SelectRows(model.NewsCategoryView, tail, newsID)
	if err != nil && err != reform.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	var categoryIDs []int64
	var nc model.NewsCategory
	for {
		err = r.db.NextRow(&nc, rows)
		if err == reform.ErrNoRows {
			break
		}
		categoryIDs = append(categoryIDs, nc.CategoryID)
	}
	return categoryIDs
}

// Return nil, when record absent
func (r *NewsRepositoryMysql) FindByID(id int64) *model.News {
	var newsModel model.News
	err := r.db.FindByPrimaryKeyTo(&newsModel, id)
	if err == reform.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return &newsModel
}

// Я заменил вызов r.db.Save(news) на r.db.Update(news) и r.db.Insert(news),
// потому что при сохранении существующей сущности,
// в случае полного совпадения входных данных с состоянием сущности в БД,
// после Update вызывается Insert на том же Id сущности, что приводит к SQL ошибке дублирующего INSERT.
func (r *NewsRepositoryMysql) Save(news *model.News) *model.News {
	if news.HasPK() {
		err := r.db.Update(news)
		if err != nil && err != reform.ErrNoRows {
			log.Fatalf("Error in r.db.Update(): %v", err)
		}
		return news
	}

	err := r.db.Insert(news)
	if err != nil && err != reform.ErrNoRows {
		log.Fatalf("Error in r.db.Insert(): %v", err)
	}
	return news
}

func (r *NewsRepositoryMysql) AssignCategory(newsID, catID int64) {
	tail := fmt.Sprintf("WHERE NewsId = %s AND CategoryId = %s", r.db.Placeholder(1), r.db.Placeholder(2))
	cnt, err := r.db.Count(model.NewsCategoryView, tail, newsID, catID)
	if err != nil && err != reform.ErrNoRows {
		log.Fatal(err)
	}
	if cnt > 0 {
		log.Printf("News (%d) to Category (%d) association already exists", newsID, catID)
		return
	}

	nc := &model.NewsCategory{
		NewsID:     newsID,
		CategoryID: catID,
	}
	err = r.db.Insert(nc)
	if err != nil && err != reform.ErrNoRows {
		log.Fatalf("Error in r.db.Insert(): %v", err)
	}
}

func (r *NewsRepositoryMysql) UnassignCategories(newsID int64, skipIDs []int64) {
	categoryIdCond := ""
	if len(skipIDs) > 0 {
		skipIDList := strings.Trim(strings.ReplaceAll(fmt.Sprint(skipIDs), " ", ","), "[]")
		categoryIdCond = fmt.Sprintf(" AND CategoryId NOT IN (%s)", skipIDList)
	}
	tail := fmt.Sprintf("WHERE NewsId = %d%s", newsID, categoryIdCond)
	cnt, err := r.db.DeleteFrom(model.NewsCategoryView, tail)
	if err != nil && err != reform.ErrNoRows {
		log.Fatal(err)
	}
	log.Printf("%d categories are unassigned from news record (%d)", cnt, newsID)
}

func (r *NewsRepositoryMysql) DeleteNewsById(newsID int64) *model.News {
	news := r.FindByID(newsID)
	if news == nil {
		return nil
	}
	return r.DeleteNews(news)
}

func (r *NewsRepositoryMysql) DeleteNews(news *model.News) *model.News {
	err := r.db.Delete(news)
	if err != nil && err != reform.ErrNoRows {
		log.Fatal(err)
	}
	r.UnassignCategories(news.ID, nil)
	return news
}
