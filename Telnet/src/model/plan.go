package model

type Plan struct {
	Id       string  `gorm:"column:Id"`
	Duration int     `gorm:"column:Duration"`
	Cost     float32 `gorm:"column:Cost"`
	Sms      int     `gorm:"column:Sms"`
	Talktime string  `gorm:"column:Talktime"`
	Internet string  `gorm:"column:Internet"`
}

func (Plan) TableName() string {
	return "telnet.plan"
}
