package images

import (
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

var (
	thumbnail_width  = 200
	thumbnail_height = 200
)

/*
func GenerateThumbnail(imagepath string) ([]byte, error) {

	buf, err := os.ReadFile(imagepath)
	if err != nil {
		return nil, err
	}
	image, err := vips.NewThumbnailFromBuffer(buf, thumbnail_width, thumbnail_height, vips.InterestingAll) //.InterestingCentre)
	if err != nil {
		return nil, err
	}
	log.Println("Creating thumbnail for " + imagepath)
	ep := vips.NewJpegExportParams()
	ep.StripMetadata = true
	ep.Quality = 75
	ep.Interlace = true
	ep.OptimizeCoding = true
	ep.SubsampleMode = vips.VipsForeignSubsampleAuto
	ep.TrellisQuant = true
	ep.OvershootDeringing = true
	ep.OptimizeScans = true
	ep.QuantTable = 3
	imageBytes, _, err := image.ExportJpeg(ep)

	return imageBytes, err
}*/

func GenerateThumbnail(imagePath string, thumbnailpath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}

	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// Extract thumbnail
	thumbnail := resize.Thumbnail(200, 200, img, resize.NearestNeighbor)

	out, err := os.Create(thumbnailpath)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, thumbnail, nil)
}

/*
func getThumbnail(cacheKey string, imagePath string) (io.ReadCloser, error) {
	// Check if the thumbnail is cached
	hit, err := cache.Get(cacheKey)
	if err != nil {
		return nil, err
	}

	if hit {
		// Use cached thumbnail
		return hit, nil
	}

	// Thumbnail is not cached, create it
	img, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}

	defer img.Close()

	// Extract thumbnail
	thumbnail, err := resize.Thumbnail(img, 200, 200, resize.NearestNeighbor)
	if err != nil {
		return nil, err
	}

	// Cache thumbnail for 10 minutes
	err = cache.Set(cacheKey, thumbnail, cacheDuration)
	if err != nil {
		return nil, err
	}

	return thumbnail, nil
}
*/
