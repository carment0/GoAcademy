package model // keep logic in one package

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model // inheriting all model methods from ORM
	// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
	//    type User struct {
	//      gorm.Model
	//    }
	Name           string    `gorm:"type:varchar(100)"`
	Email          string    `gorm:"type:varchar(100);unique_index"`
	SessionToken   string    `gorm:"type:varchar(100);unique_index"`
	PasswordDigest []byte    `gorm:"type:bytea"`        // byte array
	Messages       []Message `gorm:"ForeignKey:UserID"` // example of has_many
}

func (u *User) ResetSessionToken() {
	if randStr, err := GenerateRandomString(20); err == nil {
		u.SessionToken = randStr
	}
}
