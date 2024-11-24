package types

type ImageInfo struct {
	ID         string `json:"id"`
	Digest     string `json:"digest"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Path       string `json:"path"`
	Size       int64  `json:"size"`
}
