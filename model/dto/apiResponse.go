package dto

// - status (string): İşlem sonuç durumu. success
// - timestamp (string): İşlem zaman bilgisi
// - result (object): İşlem sonucunda dönen veriler
type ApiResponse struct {
	Status    string      `bson:"status,omitempty"`
	TimeStamp string      `bson:"timestamp,omitempty"`
	Result    interface{} `bson:"result,omitempty"`
}
