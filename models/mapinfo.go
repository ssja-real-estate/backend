package models

type mapInfo struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Langitude float64 `json:"langitude" bson:"langitude"`
	Zoom      float32 `json:"zoom" bson:"zoom"`
}
