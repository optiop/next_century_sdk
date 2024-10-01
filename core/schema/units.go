package schema

import "time"

type Unit struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	BuildingID string    `json:"building_id"`
	PropertyID int       `json:"property_id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
