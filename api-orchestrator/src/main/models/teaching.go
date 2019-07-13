package models

// Teaching is a struct to return from server
type Teaching struct {
	Slug         string   `json:"slug"`
	CategorySlug string   `json:"category_slug"`
	Text         []string `json:"text"`
	MediaBase64  []string `json:"media_base64"`
}
