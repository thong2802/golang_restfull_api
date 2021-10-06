package entity 

import (
    "fmt"
)
 
type User struct {
   ID       uint64 `gorm:"primary_key:auto_increment" json:"id`
   Name     string `gorm:"type:"varchar(255)" json:"name"`
   Email	string `gorm:"uniqueIndex; type:"varchar(255)" json:"email"`
   Pass		string `gorm:"->;<-;not null" json:"-"`
   Token	string `gorm:"-" json:"token, omitempty"`

}
 
