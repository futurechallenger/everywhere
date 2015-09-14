package config

/*
  ErrorJson
*/
type TemplateJson struct {
	State   string      `json:"state"`   // 0: correct, other values means there're something wrong.
	Message string      `json:"message"` // message: success or error
	Data    interface{} `json:"data"`    // if success, contains a dic or array, or null.
}
