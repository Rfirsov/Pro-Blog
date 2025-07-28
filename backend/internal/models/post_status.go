package models

type PostStatus struct {
	ID    uint   `gorm:"primaryKey" json:"id" example:"1"`
	Value string `gorm:"uniqueIndex;not null" json:"value" example:"draft"`
	Label string `gorm:"not null" json:"label" example:"Draft"`
}
