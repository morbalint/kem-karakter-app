package dtos

type Character struct {
	Name       string `json:"name"`
	Class      string `json:"class"`
	Level      int    `json:"level"`
	Race       string `json:"race"`
	Alignment  string `json:"alignment"`
	Gender     string `json:"gender"`
	AgeInYears int    `json:"ageInYears"`
	Religion   string `json:"religion"`

	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`

	Endurance     int `json:"endurance"`
	ReflexBase    int `json:"reflexBase"`
	WillPowerBase int `json:"willPowerBase"`

	HPMax     int `json:"HPMax"`
	HPCurrent int `json:"HPCurrent"`

	ACArmor  int `json:"ACArmor"`
	ACShield int `json:"ACShield"`
	ACOther  int `json:"ACOther"`

	Initiative     int `json:"initiative"`
	MovementInFeet int `json:"movementInFeet"`
}
