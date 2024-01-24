package reflection

import "reflect"

//func walk(x interface{}, fn func(input string)) {
//	fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
//}

//func walk(x interface{}, fn func(input string)) {
//
//	val := reflect.ValueOf(x)
//	field := val.Field(0)
//	fn(field.String())
//
//}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

//func walk(x interface{}, fn func(input string)) {
//
//	val := reflect.ValueOf(x)
//
//	for i := 0; i < val.NumField(); i++ {
//		field := val.Field(i)
//
//		if field.Kind() == reflect.String {
//			fn(field.String())
//		}
//
//		if field.Kind() == reflect.Struct {
//			walk(field.Interface(), fn)
//		}
//
//	}
//}

//func walk(x interface{}, fn func(input string)) {
//
//	val := reflect.ValueOf(x)
//
//	for i := 0; i < val.NumField(); i++ {
//		field := val.Field(i)
//
//		switch field.Kind() {
//		case reflect.String:
//			fn(field.String())
//		case reflect.Struct:
//			walk(field.Interface(), fn)
//
//		}
//	}
//}

//func walk(x interface{}, fn func(input string)) {
//
//	val := getValue(x)
//	if val.Kind() == reflect.Slice {
//
//		for i := 0; i < val.Len(); i++ {
//			walk(val.Index(i).Interface(), fn)
//		}
//		return
//	}
//
//	for i := 0; i < val.NumField(); i++ {
//		field := val.Field(i)
//
//		switch field.Kind() {
//		case reflect.String:
//			fn(field.String())
//		case reflect.Struct:
//			walk(field.Interface(), fn)
//
//		}
//	}
//}

//func walk(x interface{}, fn func(input string)) {
//
//	val := getValue(x)
//
//	switch val.Kind() {
//
//	case reflect.Struct:
//		for i := 0; i < val.NumField(); i++ {
//			walk(val.Field(i).Interface(), fn)
//		}
//	case reflect.Slice:
//		for i := 0; i < val.Len(); i++ {
//			walk(val.Index(i).Interface(), fn)
//		}
//	case reflect.String:
//		fn(val.String())
//
//	}
//
//}

func walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.String:
		fn(val.String())
	}
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
