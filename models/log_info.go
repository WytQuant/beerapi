package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogInfo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	RequestAt  string             `bson:"request_at,omitempty"`
	StatusCode int                `bson:"statusCode,omitempty"`
	Path       string             `bson:"path,omitempty"`
	Method     string             `bson:"method,omitempty"`
	ClientAddr string             `bson:"client_addr,omitempty"`
}
