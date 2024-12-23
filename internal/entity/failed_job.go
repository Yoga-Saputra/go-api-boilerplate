package entity

import "time"

// FailedJob defines the structure of failed job's data
type FailedJob struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Connection string    `gorm:"type:text;not null" json:"connection"`
	Queue      string    `gorm:"type:text;not null" json:"queue"`
	Payload    string    `gorm:"type:text;not null" json:"payload"`
	Exception  string    `gorm:"type:text;not null" json:"exception"`
	FailedAt   time.Time `gorm:"index" json:"failedAt"`
}

// TableName return the name of failed job's table on the data source
func (a FailedJob) TableName() string {
	return "FailedJobs"
}
