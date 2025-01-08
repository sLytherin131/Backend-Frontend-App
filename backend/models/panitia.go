package models

type Panitia struct {
	IdPanitia uint   `gorm:"primaryKey"`
	Nama      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Jabatan   string
}
