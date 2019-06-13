package json

import(
	"encoding/json"
)

type Object interface {
	ToJson() []byte
}

type O map[string]interface{}

func (this O) Print() {
  println("O")
}

func (this O) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}

type A []interface{}

func (this A) Print() {
  println("A")
}

func (this A) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}
