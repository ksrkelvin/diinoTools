package models

import "time"

// ProhibitedPathsStruct - ProhibitedPathsStruct
type ProhibitedPathsStruct struct {
	Path      string    `json:"path" bson:"path"`
	Qty       int       `json:"qty" bson:"qty"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

// BlockedIPsStruct - BlockedIPsStruct
type BlockedIPsStruct struct {
	IP        string    `json:"ip" bson:"ip"`
	Path      string    `json:"path" bson:"path"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

// PathsStruct - PathsStruct
type PathsStruct struct {
	Path      string    `json:"path" bson:"path"`
	Qty       int       `json:"qty" bson:"qty"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	IsDanger  bool      `json:"isDanger" bson:"isDanger"`
}
