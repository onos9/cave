package controller

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/cave/pkg/database"
	"github.com/gofiber/fiber/v2"
)

var file *File

const MAX_UPLOAD_SIZE = 1000 * 1024 * 1024 // 1GB

// Webhook is an anonymous struct for user controller
type File struct{}

type Progress struct {
	TotalSize int64
	BytesRead int64
}

func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

func (c *File) download(ctx *fiber.Ctx) error {

	filename := ctx.Params("filename")
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	// set the default MIME type to send
	mime := http.DetectContentType(buffer)
	fileSize := len(string(buffer))

	reader := bytes.NewReader(buffer)

	// Generate the server headers
	ctx.Request().Header.Set("Content-Type", mime)
	ctx.Request().Header.Set("Content-Disposition", "attachment; filename="+filename+"")
	ctx.Request().Header.Set("Expires", "0")
	ctx.Request().Header.Set("Content-Transfer-Encoding", "binary")
	ctx.Request().Header.Set("Content-Length", strconv.Itoa(fileSize))
	ctx.Request().Header.Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	_, err = io.Copy(ctx.Response().BodyWriter(), reader)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return nil
}


func (c *File) upload(ctx *fiber.Ctx) error {

	// Parse the multipart form:
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	// => *multipart.Form

	// Get all files from "documents" key:
	files := form.File["file"]
	// => []*multipart.file

	// Loop through files:
	for _, fileHeader := range files {
		fmt.Println(fileHeader.Filename, fileHeader.Size, fileHeader.Header["Content-Type"][0])
		// => "tutorial.pdf" 360641 "application/pdf"

		if fileHeader.Size > MAX_UPLOAD_SIZE {
			return ctx.Status(http.StatusOK).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Errorf("the uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename),
			})
		}

		file, err := fileHeader.Open()
		if err != nil {
			return ctx.Status(http.StatusOK).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			return ctx.Status(http.StatusOK).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		// filetype := http.DetectContentType(buff)
		// if filetype != "image/jpeg" && filetype != "image/png" {
		// 	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		// 		"success": false,
		// 		"error":   errors.New("the provided file format is not allowed. Please upload a JPEG or PNG image"),
		// 	})
		// }

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		userId := ctx.Query("userId")
		rdb := database.RedisClient(0)
		defer rdb.Close()

		id, err := rdb.Get(ctx.Context(), userId).Result()
		if err != nil {
			return err
		}

		err = os.MkdirAll("./uploads/"+id, os.ModePerm)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		//ext := filepath.Ext(fileHeader.Filename)
		f, err := os.Create(fmt.Sprintf("./uploads/%s/%s", id, fileHeader.Filename))
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		defer f.Close()

		pr := &Progress{
			TotalSize: fileHeader.Size,
		}

		_, err = io.Copy(f, io.TeeReader(file, pr))
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}
	}

	// return ctx.Status(http.StatusOK).JSON(fiber.Map{
	// 	"success": true,
	// })

	return nil
}
