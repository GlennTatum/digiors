package models

type Registration struct {
	Model
	DateBegin     string     `json:"dateBegin"`
	DateBegin_fmt string     `json:"dateBegin_fmt"`
	DateEnd       string     `json:"dateEnd"`
	DateEnd_fmt   string     `json:"dateEnd_fmt"`
	Steward       string     `json:"steward"`
	Users         []AnonUser `gorm:"foreignKey:RegistrationRefer" json:"users"`
	StationRefer  string
}
