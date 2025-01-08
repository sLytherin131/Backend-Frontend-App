package models

type JadwalWebinar struct {
	ID         uint   `gorm:"primaryKey"`
	Topik      string `gorm:"not null"`
	Tanggal    string `gorm:"not null"`
	Waktu      string `gorm:"not null"`
	Platform   string
	LinkWebinar string
	PanitiaID  uint `gorm:"not null"`
}
