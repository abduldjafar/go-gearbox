package controller

import "github.com/gogearbox/gearbox"

type file struct{}

func (*file) Create() interface{} {
	return func(ctx gearbox.Context) {
		form, err := ctx.Context().Request.MultipartForm()
		if err != nil {
			ctx.SendJSON(map[string]interface{}{
				"error": err.Error(),
				"code":  500,
			})
		}

		fileHeader := form.File["file"][0]

		file, err := fileHeader.Open()
		if err != nil {
			ctx.SendJSON(map[string]interface{}{
				"error": err.Error(),
				"code":  500,
			})
		}
		_ = file
	}
}

func ImplFileController() FileController {
	return &file{}
}
