package models

// TemplateData = data I pass to the template I will render
type TemplateData struct {
	// Map of key(string) : value (string) pairs
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// Interface in go represents different kinds of data types
	Data       map[string]interface{}
	CSRFToken  string
	FlashMsg   string
	WarningMsg string
	ErrorMsg   string
}
