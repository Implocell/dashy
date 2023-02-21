package meme

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type MemeService struct {
	db MemeRepository
}

func NewMemeService(db MemeRepository) *MemeService {
	return &MemeService{db: db}
}

func (s *MemeService) GetMemeByID(c echo.Context) error {
	// test this by telling echo it should cancel immediately?
	ctx, cancelFunc := context.WithCancel(c.Request().Context())
	defer cancelFunc()

	id := c.Param("id")

	if id == "" {
		return fmt.Errorf("failed to extract id from url")
	}

	memeSerializable, err := s.db.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(200, memeSerializable)
}
