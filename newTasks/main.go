package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type User struct {
	ID    int
	Name  string `unpack:"-"`
	Login string
	Flags int
}

// формат - json: название, тип. omitempty - если поле пустое - то не нужно его писать при запаковке. "-" - вообще не нужно запаковывать(распаковывать)

func UnpackReflect(structs interface{}, data []byte) error {
	r := bytes.NewReader(data)
	val := reflect.ValueOf(structs).Elem()

	for i := 0; i < val.NumField(); i++ {
		valField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Tag.Get("unpack") == "-" {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			var value uint32
			binary.Read(r, binary.LittleEndian, &value)
			valField.Set(reflect.ValueOf(int(value)))
		case reflect.String:
			var len uint32
			binary.Read(r, binary.LittleEndian, &len)
			dataRaw := make([]byte, len)
			binary.Read(r, binary.LittleEndian, &dataRaw)
			valField.SetString(string(dataRaw))
		default:
			return fmt.Errorf("Bad type: %v for field %v", typeField.Type.Kind(), typeField.Name)
		}
	}
	return nil
}

func main() {
	data := []byte{
		128, 36, 17, 0,
		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,
		16, 0, 0, 0,
	}

	u := new(User)
	err := UnpackReflect(u, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", u)
}
