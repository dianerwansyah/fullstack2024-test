package model

type Client struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	IsProject    string  `json:"is_project"`
	SelfCapture  string  `json:"self_capture"`
	ClientPrefix string  `json:"client_prefix"`
	ClientLogo   string  `json:"client_logo"`
	Address      string  `json:"address"`
	PhoneNumber  string  `json:"phone_number"`
	City         string  `json:"city"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at"`
}
