package models

type Station struct {
	Model
	Image         string         `json:"image"`
	Name          string         `json:"name"`
	Geo           string         `json:"geo"`
	Slug          string         `json:"slug"`
	Status        string         `json:"status"`
	Description   string         `json:"description"`
	Registrations []Registration `gorm:"foreignKey:StationRefer" json:"registrations"`
}

// TODO format slug on BeforeCreate()
