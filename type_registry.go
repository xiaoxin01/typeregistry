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

// AddTypesWithKey add type to registry center with key generator func
func AddTypesWithKey(interfaces []interface{}, keyFunc func(i interface{}) string) {
	for _, i := range interfaces {
		AddTypeWithKey(i, keyFunc)
	}
}

// AddType add type to registry center
func AddType(i interface{}) string {
	return AddTypeWithKey(i, nil)
}

// AddTypeWithKey add type to registry center with key generator func
func AddTypeWithKey(i interface{}, keyFunc func(i interface{}) string) string {
	var key string
	tpe := reflect.TypeOf(i)
	if keyFunc == nil {
		switch tpe.Kind() {
		// case reflect.Ptr:
		// 	key = reflect.ValueOf(i).Type().Name()
		default:
			key = tpe.String()
		}
	} else {
		key = keyFunc(i)
	}

	typeRegistry[key] = tpe

	return key
}

// CleanRegistry clean registered types
func CleanRegistry() {
	if len(typeRegistry) > 0 {
		typeRegistry = make(map[string]reflect.Type)
	}
}

// RegistryLen return count of registied type
func RegistryLen() int {
	return len(typeRegistry)
}

// Create create type by key
// If type is pointer, Make will create an object, point to it and return none null pointer
func Create(key string) interface{} {
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

// CreateSlice create slice type by key
func CreateSlice(key string) interface{} {
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

// GetLen return lengh of slice stored in interface
func GetLen(i interface{}) int {
	value := reflect.ValueOf(i)

	if value.Kind() == reflect.Invalid {
		return -1
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return value.Len()
	default:
		return -1
	}
}
