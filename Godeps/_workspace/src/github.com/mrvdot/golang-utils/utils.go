// Package utils is a collection of helpful utilities for common actions within GoLang development
package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"unicode"
)

// type ApiResponse is a generic API response struct
type ApiResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Result  interface{}            `json:"result"`
	Data    map[string]interface{} `json:"data"` // Generic extra data to be sent along in response
}

// GenerateSlug converts a string into a lowercase dasherized slug
//
// For example: GenerateSlug("My cool object") returns "my-cool-object"
func GenerateSlug(str string) (slug string) {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
		return -1
	}, strings.ToLower(strings.TrimSpace(str)))
}

// InChain returns a boolean if a string is already in a slice of strings
//
// [TODO] Extend this to work for all standard types
func InChain(needle string, haystack []string) bool {
	if haystack == nil {
		return false
	}
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}

// Similar to "extend" in JS, only updates fields that are specified and not empty in newData
//
// Both newData and mainObj must be pointers to struct objects
func Update(mainObj interface{}, newData interface{}) bool {
	newDataVal, mainObjVal := reflect.ValueOf(newData).Elem(), reflect.ValueOf(mainObj).Elem()
	fieldCount := newDataVal.NumField()
	changed := false
	for i := 0; i < fieldCount; i++ {
		newField := newDataVal.Field(i)
		// They passed in a value for this field, update our DB user
		if newField.IsValid() && !IsEmpty(newField) {
			dbField := mainObjVal.Field(i)
			dbField.Set(newField)
			changed = true
		}
	}
	return changed
}

// IsEmpty checks to see if a field has a set value
//
// Goes beyond usual reflect.IsZero check to handle numbers, strings, and slices
// For structs, iterates over all accessible properties and returns true only if all nested fields
// are also empty.
func IsEmpty(val reflect.Value) bool {
	valType := val.Kind()
	switch valType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.String:
		return val.String() == ""
	case reflect.Interface, reflect.Slice, reflect.Ptr, reflect.Map, reflect.Chan, reflect.Func:
		// Check for empty slices and props
		if val.IsNil() {
			return true
		} else if valType == reflect.Slice || valType == reflect.Map {
			return val.Len() == 0
		}
	case reflect.Struct:
		fieldCount := val.NumField()
		for i := 0; i < fieldCount; i++ {
			field := val.Field(i)
			if field.IsValid() && !IsEmpty(field) {
				return false
			}
		}
		return true
	default:
		return false
	}
	return false
}

// func CorsHandler provides an *extremely* broad Cors handler for development
// Not suitable for production use, as origin, method, and headers should all be
// more extensively restricted for a production environment
func CorsHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		jsonp := false
		if cbName := req.FormValue("callback"); cbName != "" && req.Method == "GET" {
			rw.Header().Set("Content-Type", "text/javascript")
			rw.Write([]byte(fmt.Sprintf("%v(", cbName)))
			jsonp = true
		} else {
			rw.Header().Set("Content-Type", "application/json")
			rw.Header().Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
			rw.Header().Add("Access-Control-Allow-Methods", req.Header.Get("Access-Control-Request-Method"))
			rw.Header().Add("Access-Control-Allow-Headers", req.Header.Get("Access-Control-Request-Headers"))
			rw.Header().Add("Access-Control-Allow-Credentials", "true")

			// If we're getting an OPTIONS request, just send response
			if req.Method == "OPTIONS" {
				rw.WriteHeader(http.StatusOK)
				return
			}
		}
		handler.ServeHTTP(rw, req)
		if jsonp {
			rw.Write([]byte(");"))
		}
	})
}
