package typeregistry

import (
	"reflect"
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

func TestCleanRegistry(t *testing.T) {
	CleanRegistry()
	AddType(1)

	assert.Equal(t, 1, RegistryLen())

	CleanRegistry()

	assert.Equal(t, 0, RegistryLen())
}

func TestMake(t *testing.T) {
	CleanRegistry()
	t.Run("test make basic type", func(t *testing.T) {
		AddType(1)
		AddType(3.14)

		in := Make("int")
		fl := Make("float64")

		assert.Equal(t, reflect.Int, reflect.TypeOf(in).Kind())
		assert.Equal(t, 0, in.(int))
		assert.Equal(t, reflect.Float64, reflect.TypeOf(fl).Kind())
		assert.Equal(t, 0.0, fl.(float64))
	})

	t.Run("test make struct type", func(t *testing.T) {
		AddType(Student{})
		AddType(new(Student))

		student := Make("typeregistry.Student")
		studentPtr := Make("*typeregistry.Student")

		assert.Equal(t, reflect.Struct, reflect.TypeOf(student).Kind())
		assert.Equal(t, 0, student.(Student).Age)
		assert.Equal(t, reflect.Ptr, reflect.TypeOf(studentPtr).Kind())
		assert.Equal(t, 0, studentPtr.(*Student).Age)
	})

	t.Run("test make slice", func(t *testing.T) {
		AddType([]int{})
		AddType([]Student{})
		AddType([]*Student{})

		intSlice := Make("[]int")
		structSlice := Make("[]typeregistry.Student")
		structPtrSlice := Make("[]*typeregistry.Student")

		assert.Equal(t, reflect.Slice, reflect.TypeOf(intSlice).Kind())
		assert.Equal(t, reflect.Slice, reflect.TypeOf(structSlice).Kind())
		assert.Equal(t, reflect.Slice, reflect.TypeOf(structPtrSlice).Kind())

		assert.Equal(t, 0, len(intSlice.([]int)))
		assert.Equal(t, 0, len(structSlice.([]Student)))
		assert.Equal(t, 0, len(structPtrSlice.([]*Student)))

	})
}

func TestMakeSlice(t *testing.T) {
	CleanRegistry()
	t.Run("test make basic slice type", func(t *testing.T) {
		AddType(1)

		in := MakeSlice("int")

		assert.Equal(t, reflect.Slice, reflect.TypeOf(in).Kind())
		assert.Equal(t, 0, len(in.([]int)))
	})

	t.Run("test make struct slice type", func(t *testing.T) {
		AddType(Student{})
		AddType(new(Student))

		students := MakeSlice("typeregistry.Student")
		studentsPtr := MakeSlice("*typeregistry.Student")

		assert.Equal(t, reflect.Slice, reflect.TypeOf(students).Kind())
		assert.Equal(t, reflect.Slice, reflect.TypeOf(studentsPtr).Kind())

		_, studentsOk := students.([]Student)
		_, studentsPtrOk := studentsPtr.([]*Student)
		assert.Equal(t, true, studentsOk)
		assert.Equal(t, true, studentsPtrOk)
	})
}
