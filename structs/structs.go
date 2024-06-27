package structs

type Messages struct {
	Message string
	Type    string
	Peer    string
}

type UploadResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}
