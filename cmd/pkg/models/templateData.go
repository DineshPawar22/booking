package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FlotMap   map[string]float32
	Data      map[string]interface{}
	CSEFToken string
	Flash     string
	Warning   string
	Error     string
}
