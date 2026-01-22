package handler

import "time"

// CategoryCreateRequest represents the payload for creating a category
type CategoryCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CategoryUpdateRequest represents the payload for updating a category
type CategoryUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CategoryResponse represents the API response for a category
type CategoryResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
