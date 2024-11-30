package models

import "time"

// BlockedIPsStruct - BlockedIPsStruct
type BlockedIPsStruct struct {
	IP        string    `json:"ip" bson:"ip"`
	Path      string    `json:"path" bson:"path"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

// PathsStruct - PathsStruct
type PathsStruct struct {
	Path         string    `json:"path" bson:"path"`
	Qty          int       `json:"qty" bson:"qty"`
	Timestamp    time.Time `json:"timestamp" bson:"timestamp"`
	IsProhibited bool      `json:"isProhibited" bson:"isProhibited"`
}
