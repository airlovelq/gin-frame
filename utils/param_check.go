package utils

import (
	//	"fmt"
	//	"reflect"
	"regexp"
)

func Check_Num(s ...interface{}) bool {
	for _, atom := range s {
		// fmt.Print(reflect.TypeOf(atom))
		// fmt.Print(atom)
		_, ok := atom.(float64)
		// fmt.Print(ok)
		if !ok {
			return false
		}
		// _, ok = atom.(int16)
		// fmt.Print(ok)
	}
	return true
}

func CheckNumIfExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			continue
		}
		_, ok := atom.(float64)
		if !ok {
			return false
		}
	}
	return true
}

func Check_String(s ...interface{}) bool {
	for _, atom := range s {
		_, ok := atom.(string)
		if !ok {
			return false
		}
	}
	return true
}

func CheckStringIfExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			continue
		}
		_, ok := atom.(string)
		if !ok {
			return false
		}
	}
	return true
}

func Check_Bytes(s ...interface{}) bool {
	for _, atom := range s {
		_, ok := atom.([]byte)
		if !ok {
			return false
		}
	}
	return true
}

func Check_UUID4(s ...interface{}) bool {
	for _, atom := range s {
		if !Check_String(atom) {
			return false
		}
		is, err := regexp.MatchString("^[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}$", atom.(string))
		if (err != nil) || (!is) {
			return false
		}
	}
	return true
}

func Check_Email(s ...interface{}) bool {
	for _, atom := range s {
		if !Check_String(atom) {
			return false
		}
		str := atom.(string)
		is, err := regexp.MatchString(`^[\.A-Za-z0-9]+@[\.a-zA-Z0-9_-]+$`, str)
		if (err != nil) || (!is) {
			return false
		}
	}
	return true
}
