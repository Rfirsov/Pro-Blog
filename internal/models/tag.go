package models

type Tag struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:80;unique;not null" json:"name"`
	Posts []Post `gorm:"many2many:post_tags;constraint:OnDelete:CASCADE" json:"-"`
}
