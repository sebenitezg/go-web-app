package models

// Holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // Generic data type. When we don't know what type of data we are going to send to the template
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
