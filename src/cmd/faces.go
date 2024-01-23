/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/Kagami/go-face"
	"github.com/eysteinn/media-api/src/database"
	"github.com/spf13/cobra"
)

var (
	modelsDir string
)

// facesCmd represents the faces command
var facesCmd = &cobra.Command{
	Use:   "faces",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := dofaces()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(facesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// facesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// facesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	facesCmd.Flags().StringVarP(&modelsDir, "model-dir", "m", "./faces/models/", "Models")
}

func dofaces() error {
	page := 0
	limit := 20

	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	for {
		fmt.Println("Loading page: ", page+1)
		offset := page * limit
		photos, err := database.SearchByTime(0, 0, limit, offset)
		if err != nil {
			return err
		}
		page++
		if len(photos) == 0 {
			break
		}

		for _, photo := range photos {
			faces, err := rec.RecognizeFile(photo.Filepath)
			if err != nil {
				return err
			}
			if faces == nil || len(faces) == 0 {
				log.Println("No faces found in: ", photo.Filepath)
				continue
			}
			log.Println("Found ", len(faces), " faces in ", photo.Filepath)

			for _, face := range faces {

				db := database.GetDB()
				insertSQL := `
				INSERT INTO faces (photo_uuid, descriptor) VALUES (?, ?)
			`
				b, err := structToBytes(face.Descriptor)
				if err != nil {
					return err
				}
				_, err = db.Exec(insertSQL, photo.UUID, b)
				if err != nil {
					return err
				}
			}

			fmt.Println(photo.Filepath)
		}
	}
	return nil
}

func structToBytes(data [128]float32) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
