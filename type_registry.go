package typeregistry

import "reflect"

var typeRegistry = make(map[string]reflect.Type)

func init() {

}

// AddType add type to registry center
func AddType(i interface{}) string {
	var key string
	tpe := reflect.TypeOf(i)
	switch tpe.Kind() {
	// case reflect.Ptr:
	// 	key = reflect.ValueOf(i).Type().Name()
	default:
		key = tpe.String()
	}

	typeRegistry[key] = tpe

	return key
}

// CleanRegistry clean registed types
func CleanRegistry() {
	if len(typeRegistry) > 0 {
		typeRegistry = make(map[string]reflect.Type)
	}
}

// RegistryLen return count of registied type
func RegistryLen() int {
	return len(typeRegistry)
}

// Make create type by key
func Make(key string) interface{} {
	if tpe, ok := typeRegistry[key]; ok {
		return reflect.New(tpe).Elem().Interface()
	}

	return nil
}
