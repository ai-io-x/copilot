package model

import (
    "gorm.io/gorm"
)

type Bookmark struct {
    ID          uint   `gorm:"primaryKey"`
    WebsiteName string `gorm:"not null"`
    Link        string `gorm:"not null"`
    Description string
    Tags        []Tag   `gorm:"many2many:bookmark_tags;"`
    Category    Category `gorm:"foreignKey:CategoryID"`
    CategoryID  uint
}

type Tag struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null;unique"`
}

type Category struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null;unique"`
}

func Migrate(db *gorm.DB) error {
    err := db.AutoMigrate(&Bookmark{}, &Tag{}, &Category{})
    return err
}