package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Error via API
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		ctx.Logger().Error(report)
		ctx.JSON(report.Code, report)
	}

	// Error Via PAGE
	// e.HTTPErrorHandler = func(err error, ctx echo.Context) {
	// 	report, ok := err.(*echo.HTTPError)
	// 	if !ok {
	// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}

	// 	errPage := fmt.Sprintf("%d.html", report.Code)
	// 	if err := ctx.File(errPage); err != nil {
	// 		ctx.HTML(report.Code, "Errrrorrrrr")
	// 	}
	// }

	e.POST("/users", func(ctx echo.Context) error {
		u := new(User)
		if err := ctx.Bind(u); err != nil {
			return err
		}
		if err := ctx.Validate(u); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, true)
	})

	// routes here
	e.Logger.Fatal(e.Start(":9000"))
}
