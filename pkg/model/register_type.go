package model

type RegisterRequest struct {
	Addons []Addon
}

type Addon struct {
	Name       string
	Properties map[string]string
}
