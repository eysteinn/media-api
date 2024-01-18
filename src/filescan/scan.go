package filescan

import (
	"fmt"
	"path/filepath"

	"github.com/eysteinn/media-api/src/database"
	"github.com/eysteinn/media-api/src/images"
)

func Scan(folderPath string) error {
	//folderPath := "photos"
	fmt.Println("Searching: ", folderPath)

	files, err := filepath.Glob(filepath.Join(folderPath, "*.jpg"))
	if err != nil {
		return err
	}

	for _, fileName := range files {
		fmt.Println("Filename: ", fileName)
		/*file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()

		// Decode the Exif data
		x, err := exif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		taken, err := x.DateTime()
		if err != nil {
			log.Println("No date/time found in Exif")
		} else {
			fmt.Println("Exif timestamp:", taken)
		}*/

		t, err := images.ReadImageTime(fileName)
		if err != nil {
			return err
		}
		fmt.Println("Time: ", t)

		hash, err := images.ReadImageSHA256(fileName)
		if err != nil {
			return err
		}
		fmt.Println("Hash: ", hash)

		///name := path.Base(fileName)
		err = database.InsertPhoto(fileName, int32(t.Unix()), hash)
		if err != nil {
			//return err
			fmt.Println(err)
		}
		fmt.Println("\tTaken: ", t.String())
	}
	return nil
}
