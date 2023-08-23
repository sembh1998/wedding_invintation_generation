package bootstrap

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
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

//go:embed assets
var assetsFolder embed.FS

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

func ExtractAssetsFolder() error {
	// Create the target directory if it doesn't exist
	targetDir := "assets"
	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		return err
	}
	// Extract the embedded content to the target directory
	err = fs.WalkDir(assetsFolder, "assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel("assets", path)
		if err != nil {
			return err
		}

		// Create the directory structure in the target directory
		targetPath := filepath.Join(targetDir, relPath)
		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// Check if the file already exists in the target directory
		if _, err := os.Stat(targetPath); err == nil {
			// File already exists, skip it
			return nil
		} else if !os.IsNotExist(err) {
			// An error occurred while checking the file, return the error
			return err
		}

		// Copy the file from embedded FS to target directory
		data, err := assetsFolder.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, 0644)
	})

	if err != nil {
		fmt.Println("Error extracting content:", err)
		return err
	}

	fmt.Println("Content extracted successfully.")
	return nil
}
