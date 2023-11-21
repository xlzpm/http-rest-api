package model

import "time"

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	IsAlive   bool      `json:"is_alive"`
	CreatedAT time.Time `json:"created_at"`
}
