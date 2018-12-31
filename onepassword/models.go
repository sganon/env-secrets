package onepassword

// ItemID represents a one password item id
type ItemID string

// Item is op item
type Item struct {
	ID       ItemID `json:"uuid"`
	Overview struct {
		Tags  []string `json:"tags"`
		Title string   `json:"title"`
	} `json:"overview"`
	Details struct {
		Password string `json:"password"`
	} `json:"details"`
}
