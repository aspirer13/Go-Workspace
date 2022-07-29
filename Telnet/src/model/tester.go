package model

type Tester struct {
	Status bool
}

// sets return value as table name
// by default, schema is set to public
func (Tester) TableName() string {
	return "telnet.Tester"
}
