package handler

import (
	"net/http"
	"strconv"

	"github.com/ahmadeko2017/backed-golang-tugas/internal/entity"
	"github.com/ahmadeko2017/backed-golang-tugas/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with name and description
// @Tags categories
// @Accept json
// @Produce json
// @Param category body CategoryCreateRequest true "Category Data"
// @Success 201 {object} entity.Category
// @Failure 400 {object} map[string]string "Bad Request"
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := entity.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get a list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} entity.Category
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Get details of a specific category
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} entity.Category
// @Failure 404 {object} map[string]string "Not Found"
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category's name and description
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body CategoryUpdateRequest true "Category Data"
// @Success 200 {object} entity.Category
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := entity.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.UpdateCategory(uint(id), &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string "Message"
// @Failure 404 {object} map[string]string "Not Found"
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
