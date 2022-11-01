package models

type LogInfo struct {
	RequestAt  string `bson:"request_at,omitempty"`
	StatusCode int    `bson:"statusCode,omitempty"`
	Path       string `bson:"path,omitempty"`
	Method     string `bson:"method,omitempty"`
	ClientAddr string `bson:"client_addr,omitempty"`
}
