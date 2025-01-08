package models

type TugasPanitia struct {
	ID        uint    `gorm:"primaryKey"`
	Deskripsi string  `gorm:"not null"`
	PanitiaID uint    `gorm:"not null;index;column:id_panitia"`           // Menyesuaikan dengan kolom di DB
	Panitia   Panitia `gorm:"foreignKey:PanitiaID;references:id_panitia"` // Menyesuaikan dengan kolom di DB
}
