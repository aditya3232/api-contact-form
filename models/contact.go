// Package models defines the data models for the API Contact Form application.
//
// It includes the Contact struct, which represents a contact message submitted
// through the API. The struct is configured for use with GORM, an ORM library
// for Go, to handle database interactions.
//
// Author: Adit
package models

import (
	"time"
)

// Contact represents a contact message submitted through the API.
// ID is the unique identifier for each contact message.
// FullName is the name of the person submitting the contact message.
// Email is the email address of the person submitting the contact message.
// Phone is the phone number of the person submitting the contact message.
// Message is the content of the contact message.
// CreatedAt records the timestamp when the contact message was created.
// UpdatedAt records the timestamp when the contact message was last updated.
// DeletedAt records the timestamp when the contact message was deleted. This field is indexed to optimize deletion queries.
type Contact struct {
	ID        uint       `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	FullName  string     `gorm:"column:full_name;type:VARCHAR(100);not null"`
	Email     string     `gorm:"column:email_address;type:VARCHAR(100);not null"`
	Phone     string     `gorm:"column:phone_number;type:VARCHAR(20);not null"`
	Message   string     `gorm:"column:message_text;type:TEXT;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;type:DATETIME;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:DATETIME;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:DATETIME;index"`
}

// TableName specifies the table name for the Contact model in the database.
func (Contact) TableName() string {
	return "contact_messages"
}
