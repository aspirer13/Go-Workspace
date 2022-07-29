package model

type Consumer struct {
	Name        string `gorm:"column:Name"`
	Id          int    `gorm:"column:Id"`
	Age         int    `gorm:"column:Age"`
	PhoneNumber string `gorm:"column:Phone_number"`
	PlanId      string `gorm:"column:Plan_id"`
}

func (Consumer) TableName() string {
	return "telnet.consumer"
}
