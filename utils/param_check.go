package utils

import (
	//	"fmt"
	//	"reflect"
	"regexp"

	"gopkg.in/guregu/null.v3"
)

func CheckIntIfExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			continue
		}
		_, ok := atom.(int)
		if ok {
			continue
		}
		_, ok = atom.(null.Int)
		if !ok {
			return false
		}
	}
	return true
}

func CheckIntMustExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			return false
		}
		_, ok := atom.(int)
		if ok {
			continue
		}
		nullInt, ok := atom.(null.Int)
		if !ok {
			return false
		}
		if !nullInt.Valid {
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
		if ok {
			continue
		}
		_, ok = atom.(null.String)
		if !ok {
			return false
		}
	}
	return true
}

func CheckStringMustExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			return false
		}
		_, ok := atom.(string)
		if ok {
			continue
		}
		nullString, ok := atom.(null.String)
		if !ok {
			return false
		}
		if !nullString.Valid {
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

func CheckUUID4MustExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			return false
		}
		var str string
		str, ok := atom.(string)
		if !ok {
			nullString, ok := atom.(null.String)
			if !ok {
				return false
			}
			if !nullString.Valid {
				return false
			}
			str = nullString.String
		}
		is, err := regexp.MatchString("^[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}$", str)
		if (err != nil) || (!is) {
			return false
		}
	}
	return true
}

func CheckEmailMustExist(s ...interface{}) bool {
	for _, atom := range s {
		if atom == nil {
			return false
		}
		var str string
		str, ok := atom.(string)
		if !ok {
			nullString, ok := atom.(null.String)
			if !ok {
				return false
			}
			if !nullString.Valid {
				return false
			}
			str = nullString.String
		}
		is, err := regexp.MatchString(`^[\.A-Za-z0-9]+@[\.a-zA-Z0-9_-]+$`, str)
		if (err != nil) || (!is) {
			return false
		}
	}
	return true
}
