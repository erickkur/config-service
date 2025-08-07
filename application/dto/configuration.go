package dto

type CreateConfigurationRequest struct {
	Name    string `json:"name"`
	Version int32  `json:"version"`
	Data    string `json:"data"`
}
