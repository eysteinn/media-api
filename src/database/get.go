package database

import (
	"fmt"
	"net/url"
	"strings"
)

type Photo struct {
	UUID         string `json:"uuid"`
	Filepath     string `json:"filepath"`
	Filename     string `json:"filename"`
	Unixtime     int64  `json:"unixtime"`
	Sha256Hash   string `json:"sha256hash"`
	ThumbnailUrl string `json:"thumbnail_url"`
	ImageUrl     string `json:"image_url"`
}

func (p *Photo) FillURLs(host string) error {
	var err error
	/*p.ThumbnailUrl, err = url.JoinPath(host, "/api/v1/photos/thumbnails/", p.UUID)
	if err != nil {
		return err
	}*/
	if !strings.HasPrefix(strings.ToLower(host), "http") {
		host = "http://" + host
	}

	p.ThumbnailUrl, err = url.JoinPath(host, "/api/v1/photos/", p.UUID, "/thumbnail")
	//p.ThumbnailUrl = host + "/api/v1/photos/thumbnails/" + p.UUID
	if err != nil {
		return err
	}
	//p.ImageUrl, err = url.JoinPath(host, "/api/v1/photos/thumbnails/", p.UUID)
	//p.ImageUrl = host + "/api/v1/photos/" + p.UUID
	p.ImageUrl, err = url.JoinPath(host, "api/v1/photos", p.UUID, "image")
	return err
}

func GetByUUID(uuid string) (Photo, error) {

	db := GetDB()

	photo := Photo{}
	err := db.QueryRow("SELECT uuid, filename, filepath, unixtime FROM 'files' WHERE uuid = ?", uuid).
		Scan(&photo.UUID, &photo.Filename, &photo.Filepath, &photo.Unixtime)
	if err != nil {
		return photo, err
	}
	fmt.Println(photo)
	return photo, nil
}
