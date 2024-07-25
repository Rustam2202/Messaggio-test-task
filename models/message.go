package models

import (
    "time"
)

type Message struct {
    ID        int       `json:"id" gorm:"primary_key"`
    Content   string    `json:"content"`
    Processed bool      `json:"processed"`
    CreatedAt time.Time `json:"created_at"`
}