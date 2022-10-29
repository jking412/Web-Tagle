package user

type EmailStatus struct {
	Id         int    `gorm:"column:id;primaryKey;autoIncrement"`
	Email      string `gorm:"column:email;type:varchar(40)"`
	IsActivate bool   `gorm:"column:is_activate;type:tinyint(1);default:0;comment:是否激活"`
}

func (es *EmailStatus) TableName() string {
	return "email_info"
}
