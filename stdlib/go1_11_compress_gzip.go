// Code generated by 'goexports compress/gzip'. DO NOT EDIT.

// +build go1.11,!go1.12

package stdlib

import (
	"compress/gzip"
	"reflect"
)

func init() {
	Symbols["compress/gzip"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BestCompression":    reflect.ValueOf(gzip.BestCompression),
		"BestSpeed":          reflect.ValueOf(gzip.BestSpeed),
		"DefaultCompression": reflect.ValueOf(gzip.DefaultCompression),
		"ErrChecksum":        reflect.ValueOf(&gzip.ErrChecksum).Elem(),
		"ErrHeader":          reflect.ValueOf(&gzip.ErrHeader).Elem(),
		"HuffmanOnly":        reflect.ValueOf(gzip.HuffmanOnly),
		"NewReader":          reflect.ValueOf(gzip.NewReader),
		"NewWriter":          reflect.ValueOf(gzip.NewWriter),
		"NewWriterLevel":     reflect.ValueOf(gzip.NewWriterLevel),
		"NoCompression":      reflect.ValueOf(gzip.NoCompression),

		// type definitions
		"Header": reflect.ValueOf((*gzip.Header)(nil)),
		"Reader": reflect.ValueOf((*gzip.Reader)(nil)),
		"Writer": reflect.ValueOf((*gzip.Writer)(nil)),
	}
}
