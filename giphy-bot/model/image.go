package model

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Image gorm model of a image
type Image struct {
	gorm.Model
	GiphyID  string
	ImageKey string
}

// Store ...
type Store struct {
	db *gorm.DB
}

// NewStore ...
func NewStore() *Store {
	return &Store{}
}

// Connect ...
func (store *Store) Connect() error {
	var err error
	store.db, err = gorm.Open(sqlite.Open("images.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	store.db.AutoMigrate(&Image{})
	return err
}

// Get ...
func (store *Store) Get(giphyID string) (Image, *gorm.DB) {
	var image Image
	db := store.db.Last(&image, "giphy_id = ?", giphyID)
	return image, db
}

// BeforeSave use hook to check params
func (img *Image) BeforeSave(tx *gorm.DB) (err error) {
	if len(img.GiphyID) == 0 || len(img.ImageKey) == 0 {
		err = errors.New("can't save invalid data")
	}
	return
}

// Create ...
func (store *Store) Create(giphyID, imageKey string) {
	store.db.Create(&Image{GiphyID: giphyID, ImageKey: imageKey})
}
