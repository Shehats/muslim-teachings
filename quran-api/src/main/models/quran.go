package models

// Quran data model to get from database [Original model]
type Quran struct {
	ID         int `gorm:"type:int(11);primary_key"`
	DatabaseID int `gorm:"type:smallint(6)"`
	SuraID     int `gorm:"type:int(11)"`
	VerseID    int `gorm:"type:int(11)"`
	AyahText   string
}
