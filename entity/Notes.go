package entity

import "time"

type Notes struct {
	Title string	`bson:"title"`
	Note string		`bson:"note"`
	CreatedDate time.Time	`bson:"created_date"`
	UpdatedDate time.Time 	`bson:"updated_date"`
}
