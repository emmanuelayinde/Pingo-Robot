package models

import (
    "gorm.io/gorm"
)

type Monitor struct {
    gorm.Model
    URL         string `json:"url"`
    Name        string `json:"name"`
    Interval    int    `json:"interval"` // in minutes
    LastChecked string `json:"last_checked"`
    Status      string `json:"status"`
}
