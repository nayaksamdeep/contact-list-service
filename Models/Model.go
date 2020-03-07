package Models

type Contact struct {
	ID uint            `json:"id"`
	Name string       `json:"name"`
	Designation string `json:"designation"`
	Company string `json:"company"`
}

func (b *Contact) TableName() string {
	return "contact"
}

