package entity

import "time"

type GoMocktest struct {
	ID             int       `gorm:"column:id;type:int;not null;primaryKey;autoIncrement" json:"id"`
	IdentityNumber string    `gorm:"column:identity_number;type:varchar(256);default:null" json:"identity_number"`
	FullName       string    `gorm:"column:full_name;type:varchar(256);default:null" json:"full_name"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;not null;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;not null;autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (g *GoMocktest) TableName() string {
	return "go_mock_test"
}
