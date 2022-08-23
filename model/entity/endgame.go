package entity

type Endgame struct {
	Players []Player `bson:"tags,omitempty"`
}
