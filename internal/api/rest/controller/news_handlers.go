package controller

import (
	"fmt"
	"log"
	"strconv"

	"frr-news/internal/core/domain/model"
	"frr-news/internal/core/domain/repository"

	"github.com/gofiber/fiber/v2"
)

///////////////////////////////////////////////////////////////////////////////

type GetNewsListResponse struct {
	Success bool          `json:"Success"`
	News    []*model.News `json:"News,omitempty"`
}

// GetNewsList godoc
// @Security	 ApiKeyAuth
// @Summary      News List
// @Description  Retrieve news list at some page
// @Tags         news
// @Produce      json
// @Param		 page		query	int		false	"Show page number (def: 1)"
// @Param		 per-page	query	int		false	"Records per page (def: 10)"
// @Success      200  {object}  controller.GetNewsListResponse
// @Failure      500  {object}  error
// @Router       /list [get]
func GetNewsList(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per-page", 10)
	news := repo.LoadPagenated(page, perPage)
	return ctx.JSON(&GetNewsListResponse{
		Success: len(news) > 0,
		News:    news,
	})
}

///////////////////////////////////////////////////////////////////////////////

type PostNewsEditByIdRequest struct {
	ID         int64   `json:"Id"`
	Title      string  `json:"Title"`
	Content    string  `json:"Content"`
	Categories []int64 `json:"Categories"`
}

type PostNewsEditByIdResponse struct {
	Success    bool       `json:"Success"`
	Message    string     `json:"Message"`
	News       model.News `json:"News,omitempty"`
	Categories []int64    `json:"Categories,omitempty"`
}

// PostNewsEditById godoc
// @Security	 ApiKeyAuth
// @Summary      Edit News
// @Description  Modify the existing News record
// @Tags         news
// @Accept		 json
// @Produce      json
// @Param		 Id			path	int		true	"News record ID"
// @Param		 req		body	controller.PostNewsEditByIdRequest	true	"News record data"
// @Success      200  		{object}  controller.PostNewsEditByIdResponse
// @Failure		 404
// @Router       /edit/:Id [post]
func PostNewsEditById(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	var req PostNewsEditByIdRequest
	err := ctx.BodyParser(&req)
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

	if newsID != req.ID {
		message := fmt.Sprintf("News ID in route (ID: %d) does not match News record ID in request body (ID: %d)", newsID, req.ID)
		log.Println(message)
		return ctx.JSON(&PostNewsEditByIdResponse{
			Success: false,
			Message: message,
		})
	}

	newsModel := repo.FindByID(req.ID)
	if newsModel == nil {
		message := fmt.Sprintf("Cannot find News record (ID: %d)", req.ID)
		return ctx.JSON(&PostNewsEditByIdResponse{
			Success: false,
			Message: message,
		})
	}

	if req.Title != "" {
		newsModel.Title = req.Title
	}
	if req.Content != "" {
		newsModel.Content = req.Content
	}

	repo.Save(newsModel)

	for _, catID := range req.Categories {
		repo.AssignCategory(newsID, catID)
	}
	repo.UnassignCategories(newsID, req.Categories)

	return ctx.JSON(&PostNewsEditByIdResponse{
		Success:    true,
		Message:    fmt.Sprintf("News updated (ID: %d)", newsID),
		News:       *newsModel,
		Categories: repo.LoadCategoryIDs(newsID),
	})
}

///////////////////////////////////////////////////////////////////////////////

type PostNewsAddRequest struct {
	ID         int64   `json:"Id"`
	Title      string  `json:"Title"`
	Content    string  `json:"Content"`
	Categories []int64 `json:"Categories"`
}

type PostNewsAddResponse struct {
	Success    bool        `json:"Success"`
	Message    string      `json:"Message"`
	News       *model.News `json:"News"`
	Categories []int64     `json:"Categories,omitempty"`
}

// PostNewsAdd godoc
// @Security	 ApiKeyAuth
// @Summary      Add News
// @Description  Add a News record
// @Tags         news
// @Accept		 json
// @Produce      json
// @Param		 req		body	controller.PostNewsAddRequest	true	"News record data"
// @Success      200  		{object}  controller.PostNewsAddResponse
// @Failure		 404
// @Router       /add [post]
func PostNewsAdd(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	var req PostNewsAddRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		log.Printf("BodyParser error: %v\n", err)
		return ctx.SendStatus(404)
	}

	newsModel := &model.News{
		Title:   req.Title,
		Content: req.Content,
	}
	repo.Save(newsModel)
	for _, catID := range req.Categories {
		repo.AssignCategory(newsModel.ID, catID)
	}

	return ctx.JSON(&PostNewsAddResponse{
		Success:    true,
		Message:    fmt.Sprintf("News added (ID: %d)", newsModel.ID),
		News:       newsModel,
		Categories: repo.LoadCategoryIDs(newsModel.ID),
	})
}

///////////////////////////////////////////////////////////////////////////////

type DeleteNewsByIdResponse struct {
	Success bool
	Message string
}

// DeleteNewsById godoc
// @Security	 ApiKeyAuth
// @Summary      Delete News
// @Description  Delete News record by ID
// @Tags         news
// @Accept		 json
// @Produce      json
// @Param		 NewsId		path		int		true	"News record ID"
// @Success      200  		{object}  controller.DeleteNewsByIdResponse
// @Failure		 404
// @Router       /:NewsId [delete]
func DeleteNewsById(ctx *fiber.Ctx, repo repository.NewsRepository) error {
	newsID, err := strconv.Atoi(ctx.Params("NewsId"))
	if err != nil {
		log.Printf("Route param (NewsId) parsing error: %v\n", err)
		return ctx.SendStatus(404)
	}
	deletedNews := repo.DeleteNewsById(int64(newsID))
	if deletedNews != nil {
		return ctx.JSON(DeleteNewsByIdResponse{Success: true, Message: fmt.Sprintf("News record (ID: %d) is deleted", deletedNews.ID)})
	} else {
		return ctx.JSON(DeleteNewsByIdResponse{Success: false, Message: fmt.Sprintf("News record (ID: %d) not found", newsID)})
	}
}

///////////////////////////////////////////////////////////////////////////////

type PostNewsAddCategoryResponse struct {
	Success bool
	Message string
}

// PostNewsAddCategory godoc
// @Security	 ApiKeyAuth
// @Summary      Assign Category
// @Description  Assign category to some news record
// @Tags         news
// @Accept		 json
// @Produce      json
// @Param		 NewsId		path		int		true	"News record ID"
// @Param		 CatId		path		int		true	"Category ID"
// @Success      200  		{object}  controller.PostNewsAddCategoryResponse
// @Failure		 404
// @Router       /add-category/:NewsId/:CatId [post]
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

	return ctx.JSON(PostNewsAddCategoryResponse{Success: true, Message: fmt.Sprintf("Category (ID: %d) assigned to news record (ID: %d)", catID, newsID)})
}
