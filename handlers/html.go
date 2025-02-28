package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jojomak13/pdf-toolbox/core"
)

type HTMLRequest struct {
	FilePath    string `json:"file_path"`
	HTMLContent string `json:"html_content"`
}

func HTML(c *fiber.Ctx) error {
	if c.Get("FILE-PATH", "") == "" {
		return core.WithError(c, "File-PATH header cannot be empty", http.StatusBadRequest)
	}

	toolBox := core.NewToolBox(c.Locals("requestid").(string))

	if err := toolBox.HTML(string(c.Body())); err != nil {
		core.Logger.Println(err.Error())

		return core.WithError(c, err.Error(), http.StatusBadRequest)
	}

	url, err := toolBox.Upload(c.Get("FILE-PATH"))
	if err != nil {
		core.Logger.Println(err.Error())

		return core.WithError(c, err.Error(), http.StatusBadRequest)
	}

	if err = toolBox.Clean(); err != nil {
		core.Logger.Println(err.Error())

		return core.WithError(c, err.Error(), http.StatusBadRequest)
	}

	return core.WithSuccess(c, "success", fiber.Map{
		"url": url,
	})
}
