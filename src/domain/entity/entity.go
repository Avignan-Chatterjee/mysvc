package entity

import "fmt"

type Entity struct {
	Id    int64  `json:"id"`
	Value string `json:"value"`
}

func (e *Entity) Verify() (bool, error) {
	if len(e.Value) == 0 {
		return false, fmt.Errorf("Invalid entity")
	}
	return true, nil
}
