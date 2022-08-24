package entity

type Player struct {
	UserID string `bson:"user_id,omitempty" binding:"required"`
	Score  int    `bson:"score,omitempty" binding:"required,gte=0"`
}
