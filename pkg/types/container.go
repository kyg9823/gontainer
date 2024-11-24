package types

type ContainerInfo struct {
	ID     string `json:"id"`
	LongID string `json:"longId"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status string `json:"status"`
}
