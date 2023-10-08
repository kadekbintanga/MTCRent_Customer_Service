package Resizing

import (
	"fmt"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"github.com/gocraft/work"
	"github.com/mitchellh/mapstructure"
	"image"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type ResizeImagePayload struct {
	Path    string `json:"path"`
	Quality int    `json:"quality"`
	MaxSize uint   `json:"maxSize"`
}

func (c *Resizing) Image(job *work.Job) error {
	payload := ResizeImagePayload{}
	mapstructure.Decode(job.Args, &payload)

	filePath, destination, filename := c.SetFilePath(payload.Path)
	if len(destination) == 0 {
		xtremelog.Debug("Resize: Destination does not exists")
		return nil
	} else {
		_, err := os.Stat(destination)
		if err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(destination, 0777)
			} else {
				xtremelog.Error("Resize: Destination invalid!!")
				return err
			}
		}
	}

	if !isImage(filename) {
		xtremelog.Debug("Resize: File is not image!!")
		return nil
	}

	quality := 50
	if payload.Quality > 0 {
		quality = payload.Quality
	}

	maxSize := payload.MaxSize
	if maxSize == 0 {
		maxSize = 1000
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
		xtremelog.Error("Resize: Unable to open file")
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		xtremelog.Error("Resize: Unable to decode file")
		return err
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dx()

	var newWidth, newHeight int
	if width > height {
		newWidth = int(maxSize)
		newHeight = int(float64(maxSize) / float64(width) * float64(height))
	} else {
		newHeight = int(maxSize)
		newWidth = int(float64(maxSize) / float64(height) * float64(width))
	}

	img = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)

	temporary, err := os.Create(destination + fmt.Sprintf("tmp_resize_%s", filename))
	if err != nil {
		log.Println(err)
		xtremelog.Error("Resize: Unable to create new file")
		return err
	}
	defer temporary.Close()

	if err := webp.Encode(temporary, img, &webp.Options{Quality: float32(quality)}); err != nil {
		xtremelog.Error("Resize: Unable to generate webp file")
		return err
	}

	temporary.Seek(0, 0)
	newFile, err := os.Create(destination + filename)
	if err != nil {
		xtremelog.Error(fmt.Sprintf("Resize: Unable to create new file: %s", err))
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, temporary)
	if err != nil {
		xtremelog.Error(fmt.Sprintf("Resize: Unable to save new file: %s", err))
		return err
	}

	// Hapus temporary file
	err = os.Remove(temporary.Name())
	if err != nil {
		xtremelog.Error(fmt.Sprintf("Resize: Unable to remove temporary file: %s", err))
		return err
	}

	return nil
}

func isImage(filename string) bool {
	extension := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(extension)
	log.Println(mimeType)

	return strings.HasPrefix(mimeType, "image/")
}
