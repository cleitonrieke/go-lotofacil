package models

import "time"

type Convite struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	ChaveConvite string    `json:"chave_convite"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	IdUtilizado  bool      `json:"id_utilizado"`
}

func (c *Convite) TableName() string {
	// custom table name, this is default
	return "lotofacil.convites"
}
