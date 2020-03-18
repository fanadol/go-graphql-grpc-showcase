package model

type User struct {
	ID      int32   `json:"id"`
	Name    string  `json:"name"`
	Friends []int32 `json:"friends"`
}
