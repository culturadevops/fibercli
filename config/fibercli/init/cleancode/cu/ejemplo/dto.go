package ejemplo

type Dto struct {
	Awsaccount_id uint   `gorm:"type:int;not null;"`
	Region        string `gorm:"type:varchar(100);not null;"`
}

func (Dto) TableName() string {
	return "ejemplo"
}
