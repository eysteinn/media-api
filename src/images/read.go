package images

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

/*
type Photo struct {
	filepath string
}

func ReadFile(filepath string) (Photo, error) {
	photo := Photo{}
	file, err := os.Open(filepath)
	if err != nil {
		return photo, err
	}
	defer file.Close()
}

func (p *Photo) ExtractTakenTime() {
	file *os.File
}*/

func ReadImageTime(filepath string) (time.Time, error) {
	taken := time.Time{}
	file, err := os.Open(filepath)
	if err != nil {
		return taken, err
	}
	defer file.Close()

	// Decode the Exif data
	x, err := exif.Decode(file)
	if err != nil {
		return taken, err
	}
	taken, err = x.DateTime()
	return taken, err
}

func ReadImageSHA256(filepath string) (string, error) {
	hashString := ""
	file, err := os.Open(filepath)
	if err != nil {
		return hashString, err
	}
	defer file.Close()

	// Calculate the SHA256 hash
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashString, err
	}

	// Get the SHA256 hash sum as a byte slice
	hashSum := hash.Sum(nil)

	// Convert the byte slice to a hex string
	hashString = hex.EncodeToString(hashSum)
	return hashString, nil
}


