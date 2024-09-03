package controller

import (
	"fmt"
	"log"
	"strconv"

	"fiber-reform-rest/internal/core/domain/model"
	"fiber-reform-rest/internal/core/domain/repository"

	"github.com/gofiber/fiber/v2"
)

func GetNewsList(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	type Response struct {
		Success bool          `json:"Success"`
		News    []*model.News `json:"News,omitempty"`
	}
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per-page", 10)
	news := repo.LoadPagenated(page, perPage)
	return ctx.JSON(&Response{
		Success: len(news) > 0,
		News:    news,
	})
}

func PostNewsEditById(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	type Request struct {
		ID         int64   `json:"Id"`
		Title      string  `json:"Title"`
		Content    string  `json:"Content"`
		Categories []int64 `json:"Categories"`
	}
	type Response struct {
		Success    bool       `json:"Success"`
		Message    string     `json:"Message"`
		News       model.News `json:"News,omitempty"`
		Categories []int64    `json:"Categories,omitempty"`
	}
	var newsRequest Request
	err := ctx.BodyParser(&newsRequest)
	if err != nil {
		log.Printf("BodyParser error: %v\n", err)
		return ctx.SendStatus(404)
	}

	id, err := strconv.Atoi(ctx.Params("Id"))
	if err != nil {
		log.Printf("Route param parsing error: %v\n", err)
		return ctx.SendStatus(404)
	}
	newsID := int64(id)

	if newsID != newsRequest.ID {
		message := fmt.Sprintf("News ID in route (ID: %d) does not match News record ID in request body (ID: %d)", newsID, newsRequest.ID)
		log.Println(message)
		return ctx.JSON(&Response{
			Success: false,
			Message: message,
		})
	}

	newsModel := repo.FindByID(newsRequest.ID)
	if newsModel == nil {
		message := fmt.Sprintf("Cannot find News record (ID: %d)", newsRequest.ID)
		return ctx.JSON(&Response{
			Success: false,
			Message: message,
		})
	}

	if newsRequest.Title != "" {
		newsModel.Title = newsRequest.Title
	}
	if newsRequest.Content != "" {
		newsModel.Content = newsRequest.Content
	}

	repo.Save(newsModel)

	for _, catID := range newsRequest.Categories {
		repo.AssignCategory(newsID, catID)
	}
	repo.UnassignCategories(newsID, newsRequest.Categories)

	return ctx.JSON(&Response{
		Success:    true,
		Message:    fmt.Sprintf("News updated (ID: %d)", newsID),
		News:       *newsModel,
		Categories: repo.LoadCategoryIDs(newsID),
	})
}

func PostNewsAdd(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	type Request struct {
		ID         int64   `json:"Id"`
		Title      string  `json:"Title"`
		Content    string  `json:"Content"`
		Categories []int64 `json:"Categories"`
	}
	type Response struct {
		Success    bool        `json:"Success"`
		Message    string      `json:"Message"`
		News       *model.News `json:"News"`
		Categories []int64     `json:"Categories,omitempty"`
	}
	var newsRequest Request
	err := ctx.BodyParser(&newsRequest)
	if err != nil {
		log.Printf("BodyParser error: %v\n", err)
		return ctx.SendStatus(404)
	}

	newsModel := &model.News{
		Title:   newsRequest.Title,
		Content: newsRequest.Content,
	}
	repo.Save(newsModel)
	for _, catID := range newsRequest.Categories {
		repo.AssignCategory(newsModel.ID, catID)
	}

	return ctx.JSON(&Response{
		Success:    true,
		Message:    fmt.Sprintf("News added (ID: %d)", newsModel.ID),
		News:       newsModel,
		Categories: repo.LoadCategoryIDs(newsModel.ID),
	})
}

func DeleteNewsById(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	newsID, err := strconv.Atoi(ctx.Params("NewsId"))
	if err != nil {
		log.Printf("Route param (NewsId) parsing error: %v\n", err)
		return ctx.SendStatus(404)
	}
	deletedNews := repo.DeleteNewsById(int64(newsID))
	if deletedNews != nil {
		return ctx.JSON(fiber.Map{"Success": true, "Message": fmt.Sprintf("News record (ID: %d) is deleted", deletedNews.ID)})
	} else {
		return ctx.JSON(fiber.Map{"Success": false, "Message": fmt.Sprintf("News record (ID: %d) not found", newsID)})
	}
}

func PostNewsAddCategory(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	newsID, err := strconv.Atoi(ctx.Params("NewsId"))
	if err != nil {
		log.Printf("Route param (NewsId) parsing error: %v\n", err)
		return ctx.SendStatus(404)
	}
	catID, err := strconv.Atoi(ctx.Params("CatId"))
	if err != nil {
		log.Printf("Route param (CatId) parsing error: %v\n", err)
		return ctx.SendStatus(404)
	}
	if repo.FindByID(int64(newsID)) == nil {
		return ctx.JSON(fiber.Map{"Success": false, "Message": fmt.Sprintf("Cannot find News record (ID: %d)", newsID)})
	}

	repo.AssignCategory(int64(newsID), int64(catID))

	return ctx.JSON(fiber.Map{"Success": true, "Message": fmt.Sprintf("Category (ID: %d) assigned to news record (ID: %d)", catID, newsID)})
}
