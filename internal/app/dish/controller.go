package dish

import (
	"dancing-pony/internal/app/rating"
	"dancing-pony/internal/common/jwt"
	"dancing-pony/internal/common/response"
	"dancing-pony/internal/common/utils"
	"dancing-pony/internal/platform/paginator"
	"dancing-pony/internal/platform/validator"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type dishController struct {
	service       Service
	ratingService rating.Service
}

func NewController(service Service, ratingService rating.Service) *dishController {
	return &dishController{service, ratingService}
}

// Show all dishes resource handler
// @Description Get all dishes and browse through them.
// @Summary get all dishes
// @Tags Dishes
// @Accept json
// @Produce json
// @Success 200 {object} response.PaginatedResponse{data=[]Dish,meta=paginator.Paginator,links=paginator.PageLink}
// @Success 400 {object} response.ErrorResponse{}
// @Success 401  {object} response.UnauthenticatedResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Security ApiKeyAuth
// @Router /v1/dishes [get]
func (r *dishController) Browse(c *fiber.Ctx) error {
	limit := utils.ParseInt(c.Query(paginator.PerPageVar), paginator.DefaultPageSize)
	page := utils.ParseInt(c.Query(paginator.PageVar), 0)

	pagination, err := r.service.ListDishes(c.Context(), page, limit)

	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL(), c.Path())

	return c.
		Status(fiber.StatusOK).
		JSON(response.NewPaginatedResponse(pagination, fullURL, limit))
}

// @Description View a single dish by ID
// @Summary get dish by a given ID
// @Tags Dishes
// @Accept json
// @Produce json
// @Param id path string true "Dish ID"
// @Success 200 {object} response.JsonResponse{data=DishResourceResponse}
// @Success 400 {object} response.ErrorResponse{}
// @Success 401  {object} response.UnauthenticatedResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Security ApiKeyAuth
// @Router /v1/dishes/{id} [get]
func (r *dishController) Read(c *fiber.Ctx) error {
	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	dish, err := r.service.ViewDish(c.Context(), id)

	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	ratings, err := r.ratingService.FindDishRating(c.Context(), dish.Id)

	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	var ratingSlice = []RatingResource{}
	for _, v := range *ratings {
		data := RatingResource{
			UserId: v.UserId,
			Rating: int(v.Rate),
		}

		ratingSlice = append(ratingSlice, data)
	}

	dishResource := DishResourceResponse{Dish: *dish, Ratings: ratingSlice}

	return c.JSON(response.NewJsonResponse(dishResource))
}

// @Description Create a new dish.
// @Summary create a new dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Param request body CreateDishRequest  true "query params"
// @Success 201 {object} response.JsonResponse{data=Dish}
// @Success 400 {object} response.ErrorResponse{}
// @Success 401  {object} response.UnauthenticatedResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Security ApiKeyAuth
// @Router /v1/dishes [post]
func (r *dishController) Add(c *fiber.Ctx) error {
	var request CreateDishRequest

	if err := c.BodyParser(&request); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	validator := validator.NewValidator()

	if err := validator.Validate(request); err != nil {
		return validator.JsonResponse(c, err)
	}

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(response.NewUnauthenticatedResponse())
	}

	request.UserId = int(token.Identifier.(float64))

	dish, err := r.service.CreateDish(c.Context(), request)

	if err != nil {
		code := fiber.StatusInternalServerError

		if err == ValidationNameAlreadyExist {
			code = fiber.StatusBadRequest
		}

		return c.
			Status(code).
			JSON(response.NewErrorResponse(err))
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(response.NewJsonResponse(dish))
}

// @Description Update an existing dish by given ID
// @Summary update dish by ID
// @Tags Dishes
// @Accept json
// @Produce json
// @Param request body UpdateDishRequest  true "query params"
// @Param id path string true "Dish ID"
// @Success 200 {object} response.JsonResponse{data=Dish}
// @Success 400 {object} response.ErrorResponse{}
// @Success 401  {object} response.UnauthenticatedResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Security ApiKeyAuth
// @Router /v1/dishes/{id} [put]
func (r *dishController) Edit(c *fiber.Ctx) error {
	var request UpdateDishRequest

	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		response.NewErrorResponse(err)
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	if err := c.BodyParser(&request); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(response.NewUnauthenticatedResponse())
	}

	request.UserId = int(token.Identifier.(float64))

	dish, err := r.service.UpdateDish(c.Context(), request, id)

	if err != nil {
		code := fiber.StatusInternalServerError

		if err == ErrorResourceNotFound {
			code = fiber.StatusNotFound
		}

		return c.
			Status(code).
			JSON(response.NewErrorResponse(err))
	}

	return c.JSON(response.NewJsonResponse(dish))
}

// @Description Delete dish by given ID.
// @Summary delete book by given ID
// @Tags Dishes
// @Accept json
// @Produce json
// @Param id path string true "Dish ID"
// @Success 204 {} status "No Content"
// @Success 400 {object} response.ErrorResponse{}
// @Success 401  {object} response.UnauthenticatedResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Security ApiKeyAuth
// @Router /v1/dishes/{id} [delete]
func (r *dishController) Delete(c *fiber.Ctx) error {
	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   jwt.ErrorUnAuthenticated,
		})
	}

	userId := int(token.Identifier.(float64))

	if err := r.service.DeleteDish(c.Context(), id, userId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})
}
