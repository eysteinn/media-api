package server

/*
import (
	"io"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cache"

	"github.com/nfnt/resize"
)

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
