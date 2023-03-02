package meme

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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

type UploadMemeRequest struct {
	Url  string `json:"url"`
	Name string `json:"name"`
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

func (s *MemeService) GetMemes(c echo.Context) error {
	ctx, cancelFunc := context.WithCancel(c.Request().Context())
	defer cancelFunc()

	serializableMemes, err := s.db.GetAll(ctx)
	if err != nil {
		return err
	}

	return c.JSON(200, serializableMemes)
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
	id := randStringBytes(10)
	localImgUrl, err := s.uploadMeme(ctx, imgUrl, id)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	meme := NewMeme(localImgUrl, englishPoem, Generated)

	serializableMeme := meme.AsSerializable()

	if err := s.db.Create(ctx, &serializableMeme); err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, imgUrl)
}

func (s *MemeService) UploadMeme(c echo.Context) error {
	ctx, cancelFunc := context.WithCancel(c.Request().Context())
	defer cancelFunc()

	var urlRequest UploadMemeRequest
	err := json.NewDecoder(c.Request().Body).Decode(&urlRequest)
	if err != nil {
		return err
	}

	url, err := s.uploadMeme(ctx, urlRequest.Url, urlRequest.Name)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, url)
}

func (s *MemeService) uploadMeme(ctx context.Context, url string, fileName string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf := bytes.NewBuffer([]byte{})

	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return "", err
	}

	imageUrl, err := s.storage.Upload(ctx, buf, fileName)
	if err != nil {
		return "", err
	}

	return imageUrl, nil
}

func seed() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	seed()
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
