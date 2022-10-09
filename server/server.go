package server

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/dpk2012/url-shortener/model"
	"github.com/dpk2012/url-shortener/utils"
	"github.com/gofiber/fiber/v2"
)

func redirect(c *fiber.Ctx) error {

	shortUrl := c.Params("redirect")
	url, err := model.FindByUrl(shortUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find url in db " + err.Error(),
		})
	}
	// grab the status
	// url.Clicked += 1
	// err = model.UpdateUrl(url)

	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(url.LongUrl, fiber.StatusTemporaryRedirect)
}

func getAllRedirects(c *fiber.Ctx) error {
	urls, err := model.GetAllUrls()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all url links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(urls)
}

func getUrl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id " + err.Error(),
		})
	}

	url, err := model.GetUrl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not retrive url from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(url)
}

func createUrl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var url model.Url
	err := c.BodyParser(&url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if !govalidator.IsURL(url.LongUrl) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid url ",
		})
	}

	if url.ShortUrl == "" {
		url.ShortUrl = utils.RandomURL(5)
		err = model.CreateUrl(url)
		for err != nil {
			url.ShortUrl = utils.RandomURL(5)
			err = model.CreateUrl(url)
		}
	} else {
		err = model.CreateUrl(url)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "url already exist in db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(url)
}

func updateUrl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var url model.Url

	err := c.BodyParser(&url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	err = model.UpdateUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update url link in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func deleteUrl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteUrl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "url deleted",
	})
}

func SetupAndListen() {
	app := fiber.New()

	app.Get("/api", getAllRedirects)
	app.Get("/api/:id", getUrl)
	app.Post("/api", createUrl)
	app.Patch("/api", updateUrl)
	app.Delete("/api/:id", deleteUrl)
	app.Get("/:redirect", redirect)
	app.Listen((":3000"))

}
