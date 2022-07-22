package main

import (
	"example.com/greetings/models"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) Api {
	return Api{
		Service: service,
	}
}

func (api *Api) GetBlogsHandler(c *fiber.Ctx) error {

	blogs, err := api.Service.GetBlogs()

	switch err {
	case nil:
		c.JSON(blogs)
		c.Status(fiber.StatusOK)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) PostBlogsHandler(c *fiber.Ctx) error {

	createBlogs := models.Blog{}
	err := c.BodyParser(&createBlogs)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}
	err = api.Service.PostBlogs(createBlogs)

	switch err {
	case nil:
		c.Status(fiber.StatusCreated)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) DeleteBlogsHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	err := api.Service.DeleteBlogs(ID)
	switch err {
	case nil:
		c.Status(fiber.StatusNoContent)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) UpdateBlogsHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	blog := models.BlogDTO{}
	err := c.BodyParser(&blog)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	updatedBlog, err := api.Service.UpdateBlogs(blog, ID)

	switch err {
	case nil:
		//response
		c.JSON(updatedBlog)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) GetBlogHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	blog, err := api.Service.GetBlog(ID)

	switch err {
	case nil:
		//response
		c.JSON(blog)
		c.Status(fiber.StatusOK)
	/* case BlogNotFound:
	c.Status(fiber.StatusNotFound) */
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}
