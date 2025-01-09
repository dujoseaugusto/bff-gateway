package models

import "encoding/json"

type User struct {
	User json.RawMessage `json:"user-content"`
}
