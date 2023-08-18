package bootstrap

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
)

//go:embed htmx/crudguests.html
var crudguestsFileByte []byte

//go:embed htmx/guest.html
var guestFileByte []byte

//go:embed htmx/login.html
var loginFileByte []byte

//go:embed htmx/gift-preferences.html
var giftPreferencesFileByte []byte

func ExtractEmbeddedFile() error {
	filenames := []string{"crudguests.html", "guest.html", "login.html", "gift-preferences.html"}
	// Create the fonts directory if it doesn't exist
	err := os.MkdirAll("htmx", 0755)
	if err != nil {
		return err
	}
	for _, filename := range filenames {
		filePath := filepath.Join("htmx", filename)

		// If the file already exists, don't overwrite it
		if _, err := os.Stat(filePath); err == nil {
			log.Printf("File %s already exists, skipping", filePath)
			continue
		}

		var fileByte []byte
		switch filename {
		case "crudguests.html":
			fileByte = crudguestsFileByte
		case "guest.html":
			fileByte = guestFileByte
		case "login.html":
			fileByte = loginFileByte
		case "gift-preferences.html":
			fileByte = giftPreferencesFileByte
		}

		// Write the file
		err := os.WriteFile(filePath, fileByte, 0644)
		if err != nil {
			return err
		}

	}

	return nil
}
