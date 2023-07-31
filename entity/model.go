package entity

type ApiKey struct {
	Id    uint64 `json:"-" gorm:"primary_key;auto_increment" `
	Owner string `json:"owner" gorm:"type:varchar(32)"`
	Key   []byte `json:"key" gorm:"type:binary(32)"`
}
