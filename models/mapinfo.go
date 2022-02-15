package models

type mapInfo struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	Zoom      float32 `json:"zoom" bson:"zoom"`
}
