package models

import "time"

type Estate struct {
	Id               int16      `json:"id"`
	Price            float32    `json:"price"`
	CurrencyCode     string     `json:"currency_code"`
	Region           string     `json:"region"`
	City             *string    `json:"city"`
	Village          *string    `json:"village"`
	District         *string    `json:"district"`
	Type             string     `json:"type"`
	ConstructionType string     `json:"construction_type"`
	LandSize         *int64     `json:"land_size"`
	BuildingSize     *int64     `json:"building_size"`
	Rooms            *int       `json:"rooms"`
	Bathrooms        *int       `json:"bathrooms"`
	Floors           *int       `json:"floors"`
	FloorNumber      *int       `json:"floor_number"`
	Description      string     `json:"description"`
	ConstructionDate *time.Time `json:"construction_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// 'author' => new PublicUserResource($this->whenLoaded('author')),
