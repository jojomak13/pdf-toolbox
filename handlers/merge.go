package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jojomak13/pdf-toolbox/core"
)

type MergeRequest struct {
	Urls []string `json:"urls"`
}

func Merge(c *fiber.Ctx) error {
	var req MergeRequest

	if err := c.BodyParser(&req); err != nil {
		return core.WithError(c, err.Error(), http.StatusBadRequest)
	}

	toolBox := core.NewToolBox(c.Locals("requestid").(string))

	if err := toolBox.Merge(req.Urls); err != nil {
		core.Logger.Println(err.Error())

		return core.WithError(c, err.Error(), http.StatusBadRequest)
	}

	return core.WithSuccess(c, "success", fiber.Map{})
}
