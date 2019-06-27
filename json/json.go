package json

import (
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


func (this O) Override(elements ...string) O {
	o := make(O)
	for k, v := range this {
		o[k] = v
	}
	
	first := elements[0]

	if(len(elements) == 2) {
		o[first] = elements[1]
		return o
	}

	x :=  o[first].(O).Override(elements[1:]...)
	o[first] = x

	return o

}

type A []interface{}

func (this A) Print() {
	println("A")
}

func (this A) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}
