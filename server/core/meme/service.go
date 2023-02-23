package meme

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/implocell/dashy/core/meme/internal"
	"github.com/labstack/echo/v4"
)

type MemeService struct {
	db MemeRepository
}

type PostMemeRequest struct {
	Text string `json:"text"`
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

func (s *MemeService) GenerateMemeByText(c echo.Context) error {
	_, cancelFunc := context.WithTimeout(c.Request().Context(), time.Minute*2)
	defer cancelFunc()
	
	var memeRequest PostMemeRequest
	err := json.NewDecoder(c.Request().Body).Decode(&memeRequest)
	if err != nil {
		return err
	}

	fmt.Printf("Got meme text %s", memeRequest.Text)
	internal.GenerateMemeByText(memeRequest.Text)
	return nil
}

