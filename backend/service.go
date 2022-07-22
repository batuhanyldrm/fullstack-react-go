package main

import (
	"strings"

	"example.com/greetings/models"
	"github.com/google/uuid"
)

type Service struct {
	Repository *Repository
}

func NewService(Repository *Repository) Service {
	return Service{
		Repository: Repository,
	}
}

func (service *Service) GetBlogs() ([]models.Blog, error) {

	blogs, err := service.Repository.GetBlogs()

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (service *Service) PostBlogs(blog models.Blog) error {

	blog.ID = GenerateUUID(8)

	err := service.Repository.PostBlogs(blog)

	if err != nil {
		return err
	}

	return nil
}

func (service *Service) DeleteBlogs(blogId string) error {

	err := service.Repository.DeleteBlogs(blogId)

	if err != nil {
		return err
	}

	return nil
}

func (service *Service) UpdateBlogs(blog models.BlogDTO, ID string) (models.Blog, error) {

	updatedBlog, err := service.Repository.UpdateBlogs(blog, ID)

	if err != nil {
		return models.Blog{}, err
	}

	return updatedBlog, nil
}

func (service *Service) GetBlog(ID string) (models.Blog, error) {

	updatedBlog, err := service.Repository.GetBlog(ID)

	if err != nil {
		return models.Blog{}, err
	}

	return updatedBlog, nil
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}
