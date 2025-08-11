package models

type Customer struct {
    ID            int64  `json:"id"`
    Code          string `json:"code"`
    FullName      string `json:"full_name"`
    Phone         string `json:"phone"`
    Email         string `json:"email"`
    Address       string `json:"address"`
    DateOfBirth   string `json:"date_of_birth"` // YYYY-MM-DD
    DriverLicense string `json:"driver_license"`
    CreatedAt     string `json:"created_at"`
    UpdatedAt     string `json:"updated_at"`
}
