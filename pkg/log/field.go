package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type field struct {
	Skip        func() zap.Field
	Binary      func(key string, val []byte) zap.Field
	Bool        func(key string, val bool) zap.Field
	Boolp       func(key string, val *bool) zap.Field
	ByteString  func(key string, val []byte) zap.Field
	Complex128  func(key string, val complex128) zap.Field
	Complex128p func(key string, val *complex128) zap.Field
	Complex64   func(key string, val complex64) zap.Field
	Complex64p  func(key string, val *complex64) zap.Field
	Error       func(err error) zap.Field
	Float64     func(key string, val float64) zap.Field
	Float64p    func(key string, val *float64) zap.Field
	Float32     func(key string, val float32) zap.Field
	Float32p    func(key string, val *float32) zap.Field
	Int         func(key string, val int) zap.Field
	Intp        func(key string, val *int) zap.Field
	Int64       func(key string, val int64) zap.Field
	Int64p      func(key string, val *int64) zap.Field
	Int32       func(key string, val int32) zap.Field
	Int32p      func(key string, val *int32) zap.Field
	Int16       func(key string, val int16) zap.Field
	Int16p      func(key string, val *int16) zap.Field
	Int8        func(key string, val int8) zap.Field
	Int8p       func(key string, val *int8) zap.Field
	String      func(key string, val string) zap.Field
	Stringp     func(key string, val *string) zap.Field
	Uint        func(key string, val uint) zap.Field
	Uintp       func(key string, val *uint) zap.Field
	Uint64      func(key string, val uint64) zap.Field
	Uint64p     func(key string, val *uint64) zap.Field
	Uint32      func(key string, val uint32) zap.Field
	Uint32p     func(key string, val *uint32) zap.Field
	Uint16      func(key string, val uint16) zap.Field
	Uint16p     func(key string, val *uint16) zap.Field
	Uint8       func(key string, val uint8) zap.Field
	Uint8p      func(key string, val *uint8) zap.Field
	Uintptr     func(key string, val uintptr) zap.Field
	Uintptrp    func(key string, val *uintptr) zap.Field
	Reflect     func(key string, val interface{}) zap.Field
	Namespace   func(key string) zap.Field
	Stringer    func(key string, val fmt.Stringer) zap.Field
	Time        func(key string, val time.Time) zap.Field
	Timep       func(key string, val *time.Time) zap.Field
	Stack       func(key string) zap.Field
	StackSkip   func(key string, skip int) zap.Field
	Duration    func(key string, val time.Duration) zap.Field
	Durationp   func(key string, val *time.Duration) zap.Field
	Any         func(key string, val interface{}) zap.Field
}

var Field = createField()

func createField() *field {
	return &field{
		Skip:        zap.Skip,
		Binary:      zap.Binary,
		Bool:        zap.Bool,
		Boolp:       zap.Boolp,
		ByteString:  zap.ByteString,
		Complex128:  zap.Complex128,
		Complex128p: zap.Complex128p,
		Complex64:   zap.Complex64,
		Complex64p:  zap.Complex64p,
		Error:       zap.Error,
		Float64:     zap.Float64,
		Float64p:    zap.Float64p,
		Float32:     zap.Float32,
		Float32p:    zap.Float32p,
		Int:         zap.Int,
		Intp:        zap.Intp,
		Int64:       zap.Int64,
		Int64p:      zap.Int64p,
		Int32:       zap.Int32,
		Int32p:      zap.Int32p,
		Int16:       zap.Int16,
		Int16p:      zap.Int16p,
		Int8:        zap.Int8,
		Int8p:       zap.Int8p,
		String:      zap.String,
		Stringp:     zap.Stringp,
		Uint:        zap.Uint,
		Uintp:       zap.Uintp,
		Uint64:      zap.Uint64,
		Uint64p:     zap.Uint64p,
		Uint32:      zap.Uint32,
		Uint32p:     zap.Uint32p,
		Uint16:      zap.Uint16,
		Uint16p:     zap.Uint16p,
		Uint8:       zap.Uint8,
		Uint8p:      zap.Uint8p,
		Uintptr:     zap.Uintptr,
		Uintptrp:    zap.Uintptrp,
		Reflect:     zap.Reflect,
		Namespace:   zap.Namespace,
		Stringer:    zap.Stringer,
		Time:        zap.Time,
		Timep:       zap.Timep,
		Stack:       zap.Stack,
		StackSkip:   zap.StackSkip,
		Duration:    zap.Duration,
		Durationp:   zap.Durationp,
		Any:         zap.Any,
	}
}
