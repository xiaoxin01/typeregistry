package typeregistry

import (
	"reflect"
	"strings"
	"testing"

	"gotest.tools/assert"
)

type Student struct {
	Age  int
	Name string
}

func TestAddTypes(t *testing.T) {
	CleanRegistry()
	t.Run("test add multiple basic types", func(t *testing.T) {
		types := []interface{}{}
		types = append(types,
			1,
			3.14,
			new(int),
		)
		assert.Equal(t, 0, RegistryLen())

		AddTypes(types)

		assert.Equal(t, 3, len(types))
		assert.Equal(t, 3, RegistryLen())
	})
}

func TestAddType(t *testing.T) {
	CleanRegistry()
	t.Run("test add basic types", func(t *testing.T) {
		intKey := AddType(1)
		floatKey := AddType(3.14)
		intPtrKey := AddType(new(int))

		assert.Equal(t, "int", intKey)
		assert.Equal(t, "float64", floatKey)
		assert.Equal(t, "*int", intPtrKey)
	})

	t.Run("test add struct types", func(t *testing.T) {
		structKey := AddType(Student{})
		structPtrKey := AddType(new(Student))

		assert.Equal(t, "typeregistry.Student", structKey)
		assert.Equal(t, "*typeregistry.Student", structPtrKey)
	})

	t.Run("test add basic slice types", func(t *testing.T) {
		sliceKey := AddType([]int{})
		slicePtrKey := AddType([]*int{})

		assert.Equal(t, "[]int", sliceKey)
		assert.Equal(t, "[]*int", slicePtrKey)
	})

	t.Run("test add struct slice types", func(t *testing.T) {
		sliceKey := AddType([]Student{})
		slicePtrKey := AddType([]*Student{})

		assert.Equal(t, "[]typeregistry.Student", sliceKey)
		assert.Equal(t, "[]*typeregistry.Student", slicePtrKey)
	})
}

func TestAddTypeWithKey(t *testing.T) {
	CleanRegistry()
	t.Run("test use struct lowercase name for key", func(t *testing.T) {
		var i interface{} = new(Student)
		name := AddTypeWithKey(i, func(i interface{}) string {
			tpe := reflect.TypeOf(i).Elem()
			return strings.ToLower(tpe.Name())
		})

		student, ok := Create("student").(*Student)

		assert.Equal(t, "student", name)
		assert.Equal(t, true, ok)
		assert.Equal(t, 0, student.Age)
	})
}

func TestAddTypesWithKey(t *testing.T) {
	CleanRegistry()
	t.Run("test use struct lowercase name for key", func(t *testing.T) {
		var is []interface{} = []interface{}{new(Student)}
		AddTypesWithKey(is, func(i interface{}) string {
			tpe := reflect.TypeOf(i).Elem()
			return strings.ToLower(tpe.Name())
		})

		student, ok := Create("student").(*Student)

		assert.Equal(t, true, ok)
		assert.Equal(t, 0, student.Age)
	})
}

func TestCleanRegistry(t *testing.T) {
	CleanRegistry()
	AddType(1)

	assert.Equal(t, 1, RegistryLen())

	CleanRegistry()

	assert.Equal(t, 0, RegistryLen())
}

func TestCreate(t *testing.T) {
	CleanRegistry()
	t.Run("test make basic type", func(t *testing.T) {
		AddType(1)
		AddType(3.14)

		in := Create("int")
		fl := Create("float64")

		assert.Equal(t, reflect.Int, reflect.TypeOf(in).Kind())
		assert.Equal(t, 0, in.(int))
		assert.Equal(t, reflect.Float64, reflect.TypeOf(fl).Kind())
		assert.Equal(t, 0.0, fl.(float64))
	})

	t.Run("test make struct type", func(t *testing.T) {
		AddType(Student{})
		AddType(new(Student))

		student := Create("typeregistry.Student")
		studentPtr := Create("*typeregistry.Student")

		assert.Equal(t, reflect.Struct, reflect.TypeOf(student).Kind())
		assert.Equal(t, 0, student.(Student).Age)
		assert.Equal(t, reflect.Ptr, reflect.TypeOf(studentPtr).Kind())
		assert.Equal(t, 0, studentPtr.(*Student).Age)
	})

	t.Run("test make slice", func(t *testing.T) {
		AddType([]int{})
		AddType([]Student{})
		AddType([]*Student{})

		intSlice := Create("[]int")
		structSlice := Create("[]typeregistry.Student")
		structPtrSlice := Create("[]*typeregistry.Student")

		assert.Equal(t, reflect.Slice, reflect.TypeOf(intSlice).Kind())
		assert.Equal(t, reflect.Slice, reflect.TypeOf(structSlice).Kind())
		assert.Equal(t, reflect.Slice, reflect.TypeOf(structPtrSlice).Kind())

		assert.Equal(t, 0, len(intSlice.([]int)))
		assert.Equal(t, 0, len(structSlice.([]Student)))
		assert.Equal(t, 0, len(structPtrSlice.([]*Student)))

	})
}

func TestCreateSlice(t *testing.T) {
	CleanRegistry()
	t.Run("test make basic slice type", func(t *testing.T) {
		AddType(1)

		in := CreateSlice("int")

		assert.Equal(t, reflect.Slice, reflect.TypeOf(in).Kind())
		assert.Equal(t, 0, len(in.([]int)))
	})

	t.Run("test make struct slice type", func(t *testing.T) {
		AddType(Student{})
		AddType(new(Student))

		students := CreateSlice("typeregistry.Student")
		studentsPtr := CreateSlice("*typeregistry.Student")

		assert.Equal(t, reflect.Slice, reflect.TypeOf(students).Kind())
		assert.Equal(t, reflect.Slice, reflect.TypeOf(studentsPtr).Kind())

		_, studentsOk := students.([]Student)
		_, studentsPtrOk := studentsPtr.([]*Student)
		assert.Equal(t, true, studentsOk)
		assert.Equal(t, true, studentsPtrOk)
	})
}

func TestGetLen(t *testing.T) {
	t.Run("get len of slice", func(t *testing.T) {
		intSlice := []int{1, 2, 3}
		len := GetLen(intSlice)

		assert.Equal(t, 3, len)
	})

	t.Run("get len of invalid", func(t *testing.T) {
		nilLen := GetLen(nil)
		intLen := GetLen(1)

		assert.Equal(t, -1, nilLen)
		assert.Equal(t, -1, intLen)
	})
}
