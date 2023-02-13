package models

// Course struct
type Course struct {
	ID          string `json:"id"`
	CreatorId   string `json:"creator_id"`
	Title       string `json:"title"`
	SubTitle    string `json:"subtitle"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Price       string `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

// Courses struct
type Courses struct {
	Courses []Course `json:"courses"`
}
