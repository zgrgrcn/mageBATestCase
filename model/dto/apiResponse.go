package dto

type ApiResponse struct {
	Status    string      `bson:"status,omitempty"`
	TimeStamp string      `bson:"timestamp,omitempty"`
	Result    interface{} `bson:"result,omitempty"`
}
