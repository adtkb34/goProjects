package account

func AllModels() []interface{} {
	models := []interface{}{
		&UserInfo12{},
		&UserInfo17{},
	}
	return models
}

type UserInfo17 struct {
	ID   uint   `gorm:"size:10"`
	Name string `gorm:"size:10"`
}

type UserInfo14 struct {
	ID   uint   `gorm:"size:10"`
	Name string `gorm:"size:10"`
}

type UserInfo12 struct {
	ID    uint   `gorm:"size:10"`
	Name  string `gorm:"size:10"`
	Name2 string `gorm:"size:10"`
}
type UserInfo13 struct {
	ID   uint   `gorm:"size:10"`
	Name string `gorm:"size:10"`
}
