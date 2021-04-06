package apis

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"rafaignacio.com/auth/src/pkg/apis/models"
	"rafaignacio.com/auth/src/pkg/apis/repos"
	"rafaignacio.com/auth/src/pkg/userInfo"
)

func AddNewUser(ctx *fiber.Ctx) error {
	var (
		u models.NewUserModel
	)

	err := ctx.BodyParser(&u)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errs := u.Validate()

	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	repo, err := repos.NewUserRepository()

	if err != nil {
		return err
	}

	r, err := userInfo.NewUserInfo(userInfo.ProviderType(u.ProviderType), u.ProviderValue, u.Password, repo)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ctx.Append("location", fmt.Sprintf("/api/users/%s", r.ID.String()))

	return ctx.SendStatus(fiber.StatusCreated)
}
