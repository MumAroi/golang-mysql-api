package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title     string `gorm:"type:varchar(255);not null;unique" json:"title"`
	Image     string `gorm:"type:text;not null" json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (i *Image) Prepare() {
	i.ID = 0
	i.Title = html.EscapeString(strings.TrimSpace(i.Title))
	i.Image = html.EscapeString(strings.TrimSpace(i.Image))
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
}

func (i *Image) Validate() error {

	if i.Title == "" {
		return errors.New("Required Title")
	}
	if i.Image == "" {
		return errors.New("Required Image")
	}
	return nil
}

func (i *Image) SaveImage(db *gorm.DB) (*Image, error) {
	var err error
	err = db.Debug().Model(&Image{}).Create(&i).Error
	if err != nil {
		return &Image{}, err
	}
	return i, nil
}

func (i *Image) FindAllImage(db *gorm.DB) (*[]Image, error) {
	var err error
	images := []Image{}
	err = db.Debug().Model(&Image{}).Limit(100).Find(&images).Error
	if err != nil {
		return &[]Image{}, err
	}

	return &images, nil
}

func (i *Image) FindImageByID(db *gorm.DB, pid uint64) (*Image, error) {
	var err error
	err = db.Debug().Model(&Image{}).Where("id = ?", pid).Take(&i).Error
	if err != nil {
		return &Image{}, err
	}
	return i, nil
}

func (i *Image) UpdateAImage(db *gorm.DB, pid uint64) (*Image, error) {

	var err error
	db = db.Debug().Model(&Image{}).Where("id = ?", pid).Take(&Image{}).UpdateColumns(
		map[string]interface{}{
			"title":      i.Title,
			"image":      i.Image,
			"updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Image{}).Where("id = ?", pid).Take(&i).Error
	if err != nil {
		return &Image{}, err
	}
	return i, nil
}

func (i *Image) DeleteAImage(db *gorm.DB, pid uint64) (int64, error) {

	var err error

	db = db.Debug().Model(&Image{}).Where("id = ?", pid).Take(&Image{}).Delete(&Image{})
	err = db.Error
	if db.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("Image not found")
		}

		return 0, db.Error
	}
	return db.RowsAffected, nil
}
