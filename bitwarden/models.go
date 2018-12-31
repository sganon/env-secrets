package bitwarden

// FolderID is a bw folder ID
type FolderID string

// FolderSearch is the response of a `bw list folders`
type FolderSearch struct {
	ID   FolderID
	Name string
}

// ItemID is bw item ID
type ItemID string

// Item represents results from `bw list items`
type Item struct {
	ID    ItemID
	Name  string
	Notes string
}
