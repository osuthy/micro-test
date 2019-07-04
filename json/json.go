package json

import (
	"encoding/json"
	// "fmt"
	// "reflect"
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


func (this O) OverrideByStrings(elements ...string) O {
	o := make(O)
	for k, v := range this {
		o[k] = v
	}
	
	first := elements[0]

	if(len(elements) == 2) {
		o[first] = elements[1]
		return o
	}

	x :=  o[first].(O).OverrideByStrings(elements[1:]...)
	o[first] = x

	return o

}


func (this O) OverrideByObject(elements ...string) []string {
	o := make(O)
	for k, v := range this {
		o[k] = v
	}

	r := []string{}
	r = append(r, elements...)
	
	for key, obj := range this {

		if _, ok := obj.(O); ok {
			r = append(r, key)
			obj.(O).OverrideByObject(r...)
		}

		if _, ok := obj.(string); ok {
			r = append(r, key)
			r = append(r, obj.(string))
			p("--------------")
			p(r)
			return r
		}
	}

	p("--------------!!!!")
	p(this)
	return []string{}

}


func (this O) Override(elements ...interface{}) O {
	if _, ok := elements[0].(string); ok {
		r := []string{}
		for _, e := range elements {
			r = append(r, e.(string))
		}
		return this.OverrideByStrings(r...)
	}

	if _, ok := elements[0].(O); ok {
		// for _, o := range elements {
		// 	a := o.(O).OverrideByObject()
		// 	p("1kaettekita")
		// 	p(a)
		// }
		p(elements[0])
		elements[0].(O).OverrideByObject()
		p("asdfharhfliawruhg")
		// return this.OverrideByObject()
	}

	return O{}
}



type A []interface{}

func (this A) Print() {
	println("A")
}

func (this A) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}
