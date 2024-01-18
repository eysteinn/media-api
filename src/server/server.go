package server

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/eysteinn/media-api/src/database"
	"github.com/eysteinn/media-api/src/images"
	"github.com/eysteinn/media-api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const cacheDuration = 10 * time.Minute

func servePhoto(c *fiber.Ctx) error {
	photouuid := c.Params("photoid")
	if photouuid == "" {
		return fmt.Errorf("Missing parameter 'photoid'")
	}
	photo, err := database.GetByUUID(photouuid)
	if err != nil {
		return err
	}

	filepath := photo.Filepath
	fmt.Println("Reading file: ", filepath)
	img, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	fmt.Println("Sending image")
	err = c.Send(img)
	if err != nil {
		return err
	}

	fmt.Println("Done serving image")
	return nil
}
func serveThumbnail(c *fiber.Ctx) error {
	photouuid := c.Params("photoid")
	if photouuid == "" {
		return fmt.Errorf("Missing parameter 'photoid'")
	}
	err := os.MkdirAll("thumbnails", os.ModePerm)
	if err != nil {
		return err
	}

	thumbname := path.Join("thumbnails", photouuid+".jpg")
	fmt.Println("Searching for: ", thumbname)

	if !utils.FileExists(thumbname) {
		fmt.Println("Generating thumbnail")
		photo, err := database.GetByUUID(photouuid)
		if err != nil {
			return err
		}

		fmt.Println("Creating thumbnail")
		err = images.GenerateThumbnail(photo.Filepath, thumbname)
		if err != nil {
			return err
		}
	}

	fmt.Println("Reading file: ", thumbname)
	img, err := os.ReadFile(thumbname)
	if err != nil {
		return err
	}

	fmt.Println("Sending image")
	err = c.Send(img)
	if err != nil {
		return err
	}

	fmt.Println("Done serving thumbnail")
	//images.GenerateThumbnail()
	/*database.GetByUUID(photouuid)
	images.GenerateThumbnail())*/
	/*// Generate cache key
	cacheKey := fmt.Sprintf("%s-thumbnail", imagePath)

	// Get thumbnail
	thumbnail, err := images.(cacheKey, filepath.Join(imagesPath, imagePath))
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error retrieving thumbnail: " + err.Error())
	}

	return c.Send(thumbnail)*/
	return nil
}

type SearchResult struct {
	Photos []database.Photo `json:"photos"`
	/*Query   struct {
		Abc string
	}*/
}

func parseStringToUnixtime(timestr string) (int64, error) {
	layout := "20060102T1504"

	timefrom, err := strconv.Atoi(timestr)
	if err != nil {
		parsedTime, err := time.Parse(layout, timestr)
		if err != nil {
			fmt.Println("Error:", err)
			return 0, err
		}
		timefrom = int(parsedTime.Unix())
	}
	return int64(timefrom), nil
}
func queryPhoto(c *fiber.Ctx) error {
	photouuid := c.Params("photoid")
	photo, err := database.GetByUUID(photouuid)
	if err != nil {
		return err
	}
	err = photo.FillURLs(c.Hostname())
	if err != nil {
		return err
	}

	return c.JSON(photo)

}
func queryPhotos(c *fiber.Ctx) error {

	fmt.Println("Query: ", c.Queries())
	timefrom_str := c.Query("timefrom", "0")
	timeto_str := c.Query("timeto", "0")

	page := c.QueryInt("page", 0)
	limit := c.QueryInt("limit", 20)

	timefrom, err := parseStringToUnixtime(timefrom_str)
	fmt.Println("TimeFrom: ", timefrom)

	timeto, err := parseStringToUnixtime(timeto_str)
	fmt.Println("TimeTo: ", timeto)
	if timeto == 0 {
		timeto = math.MaxInt64
	}

	fmt.Println("Page: ", page)
	fmt.Println("limit: ", limit)
	offset := page * limit
	photos, err := database.SearchByTime(int64(timefrom), int64(timeto), limit, offset)
	if err != nil {
		return err
	}
	for idx, _ := range photos {
		err = photos[idx].FillURLs(c.Hostname())
		if err != nil {
			return err
		}
	}

	res := SearchResult{}
	res.Photos = photos
	return c.JSON(res)
}

func Serve(port string) error {
	log.Println("Starting image server on port " + port)
	app := fiber.New()

	app.Use(cors.New())
	//app.Use(cache.New())

	//app.Get("/api/v1/photos/:photoid", serveImage)
	app.Get("/api/v1/photos/:photoid/image", servePhoto)
	app.Get("/api/v1/photos/:photoid/thumbnail", serveThumbnail)
	app.Get("/api/v1/photos/:photoid", queryPhoto)
	app.Get("/api/v1/photos", queryPhotos)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendString("404 Not Found")
	})

	err := app.Listen("0.0.0.0:" + port)
	if err != nil {
		return err
	}
	return nil
}
