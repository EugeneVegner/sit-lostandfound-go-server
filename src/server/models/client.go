package models

type Client struct {
	Version         string `json:"version" valid:"required`
	Platform        string `json:"platform" valid:"length(1|1),required"`
}


