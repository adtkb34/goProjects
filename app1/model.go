package app1

func AllModels() []interface{} {
	models := []interface{}{
		&Example{},
	}
	return models
}

type Example struct {
	ID uint `gorm:"size:10"`
}
