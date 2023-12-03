package service

import "github.com/google/uuid"

type Object struct {
	ObjectName string    `json:"objectName,omitempty"`
	ObjectId   uuid.UUID `json:"objectId"`
}

type Relation struct {
	Type    string
	Object1 Object
	Object2 Object
}
