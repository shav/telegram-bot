package serialization

import (
	"encoding/json"
	"github.com/mailru/easyjson"
)

// JsonSerializer выполняет сериализацию объектов в формат json.
type JsonSerializer struct {
}

// NewJsonSerializer создаёт новый json-сериализатор.
func NewJsonSerializer() *JsonSerializer {
	return &JsonSerializer{}
}

// Marshal сериализует object в формат json.
func (s *JsonSerializer) Marshal(object any) ([]byte, error) {
	if easyObj, ok := object.(easyjson.Marshaler); ok {
		return easyjson.Marshal(easyObj)
	}
	return json.Marshal(object)
}

// Unmarshal выполняет десериализацию сериализованного объекта serializedObj из формата json.
func (s *JsonSerializer) Unmarshal(serializedObj []byte, object any) error {
	if easyObj, ok := object.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(serializedObj, easyObj)
	}
	return json.Unmarshal(serializedObj, object)
}
