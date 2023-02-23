package meme

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/implocell/dashy/core/meme/internal"
	"github.com/labstack/echo/v4"
)

type MemeService struct {
	db      MemeRepository
	storage MemeStorage
}

type PostMemeRequest struct {
	Text string `json:"text"`
}

func NewMemeService(db MemeRepository, storage MemeStorage) *MemeService {
	return &MemeService{db: db, storage: storage}
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
	ctx, cancelFunc := context.WithTimeout(c.Request().Context(), time.Minute*2)
	defer cancelFunc()

	var memeRequest PostMemeRequest
	err := json.NewDecoder(c.Request().Body).Decode(&memeRequest)
	if err != nil {
		return err
	}

	englishPoem, err := internal.GetPoem(memeRequest.Text)
	if err != nil {
		return err
	}

	imgUrl, err := internal.GetImageFromPoem(englishPoem)
	if err != nil {
		return err
	}

	meme := NewMeme(imgUrl, englishPoem, Generated)

	serializableMeme := meme.AsSerializable()

	if err := s.db.Create(ctx, &serializableMeme); err != nil {
		return c.JSON(500, err.Error())
	}

	res, err := http.Get(imgUrl)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	b := []byte{}
	n, err := res.Body.Read(b)
	if err != nil {
		c.JSON(500, err.Error())
	}
	log.Printf("Got %d bytes from image", n)

	buf := bytes.NewBuffer(b)

	err = s.storage.Upload(ctx, buf.Bytes())
	if err != nil {
		c.JSON(500, err.Error())
	}

	return c.JSON(200, imgUrl)
}
