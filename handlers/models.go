package handlers

type ErrorModel struct {
	Status     string `json:"status" example:"error"`
	StatusCode int    `json:"status_code" example:"500"`
	Message    string `json:"message" example:"internal server error"`
}

type SigninInput struct {
	Username string `json:"username" example:"harshtalks"`
	Password string `json:"password" password:"xxxxx"`
}

type DocBusinessModel struct {
	BusinessID  string                 `json:"business_id"`
	Name        string                 `json:"name"`
	Address     string                 `json:"address"`
	City        string                 `json:"city"`
	State       string                 `json:"state"`
	PostalCode  string                 `json:"postal_code"`
	Latitude    float64                `json:"latitude"`
	Longitude   float64                `json:"longitude"`
	Stars       float64                `json:"stars"`
	ReviewCount int                    `json:"review_count"`
	IsOpen      int                    `json:"is_open"`
	Attributes  map[string]interface{} `gorm:"type:jsonb" json:"attributes"`
	Categories  string                 `json:"categories"`
	Hours       map[string]interface{} `gorm:"type:jsonb" json:"hours"`
	UserID      int                    `json:"userID" gorm:"not null"`
	GeoHash     string                 `json:"geoHash" gorm:"not null"`
}
