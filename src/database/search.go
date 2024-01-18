package database

import (
	"fmt"
	"math"
)

func SearchByTime(timefrom int64, timeto int64, limit int, offset int) ([]Photo, error) {

	fmt.Printf("Searching from %v to %v using limit %v and offset %v\n", timefrom, timeto, limit, offset)
	db := GetDB()

	photos := []Photo{}
	if timeto == 0 {
		timeto = math.MaxInt64
	}

	rows, err := db.Query(
		"SELECT uuid, filename, filepath, unixtime FROM 'files' WHERE unixtime > ? and unixtime < ? ORDER BY unixtime LIMIT ? OFFSET ?",
		timefrom, timeto, limit, offset)
	for rows.Next() {
		var photo Photo
		if err = rows.Scan(&photo.UUID, &photo.Filename, &photo.Filepath, &photo.Unixtime); err != nil {
			return photos, err
		}
		photos = append(photos, photo)
	}

	return photos, nil
}
