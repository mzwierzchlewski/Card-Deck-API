package models

type Deck struct {
	ID        string `json:"deck_id" example:"e89efc4d-9294-443a-8a54-d78cd0e8a0c9"`
	Shuffled  bool   `json:"shuffled" example:"false"`
	Remaining int    `json:"remaining" example:"52"`
	Cards     []Card `json:"cards"`
}
