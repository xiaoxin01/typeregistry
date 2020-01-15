package typeregistry

import "reflect"

var typeRegistry = make(map[string]reflect.Type)

func init() {

}

// AddTypes add multiple types to registry center
func AddTypes(interfaces []interface{}) {
	for _, i := range interfaces {
		AddType(i)
	}
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
// If type is pointer, Make will create an object, point to it and return none null pointer
func Make(key string) interface{} {
	var value reflect.Value
	if tpe, ok := typeRegistry[key]; ok {
		value = reflect.New(tpe).Elem()
		switch tpe.Kind() {
		case reflect.Ptr:
			tValue := reflect.New(tpe.Elem()).Elem().Addr()
			value.Set(tValue)
		default:
		}
		return value.Interface()
	}

	return nil
}

// MakeSlice create slice type by key
func MakeSlice(key string) interface{} {
	var value reflect.Value
	if tpe, ok := typeRegistry[key]; ok {
		value = reflect.MakeSlice(reflect.SliceOf(tpe), 0, 0)
		// switch tpe.Kind() {
		// case reflect.Ptr:
		// 	tValue := reflect.New(tpe.Elem()).Elem().Addr()
		// 	value.Set(tValue)
		// default:
		// }
		return value.Interface()
	}

	return nil
}
