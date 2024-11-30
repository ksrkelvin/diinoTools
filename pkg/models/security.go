package models

import "time"

// ProhibitedPathsStruct - ProhibitedPathsStruct
type ProhibitedPathsStruct struct {
	Path      string    `json:"path" bson:"path"`
	Qtd       int       `json:"qtd" bson:"qtd"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

// BlockedIPsStruct - BlockedIPsStruct
type BlockedIPsStruct struct {
	IP        string    `json:"ip" bson:"ip"`
	Path      string    `json:"path" bson:"path"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
