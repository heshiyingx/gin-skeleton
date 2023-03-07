package log

import (
	"go.uber.org/zap"
)

//type FieldType uint8
//
//const (
//	// FieldTypeUnknownType is the default field type. Attempting to add it to an encoder will panic.
//	FieldTypeUnknownType FieldType = iota
//	// FieldTypeArrayMarshalerType indicates that the field carries an ArrayMarshaler.
//	FieldTypeArrayMarshalerType
//	// FieldTypeObjectMarshalerType indicates that the field carries an ObjectMarshaler.
//	FieldTypeObjectMarshalerType
//	// FieldTypeBinaryType indicates that the field carries an opaque binary blob.
//	FieldTypeBinaryType
//	// FieldTypeBoolType indicates that the field carries a bool.
//	FieldTypeBoolType
//	// FieldTypeByteStringType indicates that the field carries UTF-8 encoded bytes.
//	FieldTypeByteStringType
//	// FieldTypeComplex128Type indicates that the field carries a complex128.
//	FieldTypeComplex128Type
//	// FieldTypeComplex64Type indicates that the field carries a complex128.
//	FieldTypeComplex64Type
//	// FieldTypeDurationType indicates that the field carries a time.Duration.
//	FieldTypeDurationType
//	// FieldTypeFloat64Type indicates that the field carries a float64.
//	FieldTypeFloat64Type
//	// FieldTypeFloat32Type indicates that the field carries a float32.
//	FieldTypeFloat32Type
//	// FieldTypeInt64Type indicates that the field carries an int64.
//	FieldTypeInt64Type
//	// FieldTypeInt32Type indicates that the field carries an int32.
//	FieldTypeInt32Type
//	// FieldTypeInt16Type indicates that the field carries an int16.
//	FieldTypeInt16Type
//	// FieldTypeInt8Type indicates that the field carries an int8.
//	FieldTypeInt8Type
//	// FieldTypeStringType indicates that the field carries a string.
//	FieldTypeStringType
//	// FieldTypeFieldTypeTimeType indicates that the field carries a time.Time that is
//	// representable by a UnixNano() stored as an int64.
//	FieldTypeFieldTypeTimeType
//	// FieldTypeTimeFullType indicates that the field carries a time.Time stored as-is.
//	FieldTypeTimeFullType
//	// FieldTypeUint64Type indicates that the field carries a uint64.
//	FieldTypeUint64Type
//	// FieldTypeUint32Type indicates that the field carries a uint32.
//	FieldTypeUint32Type
//	// FieldTypeUint16Type indicates that the field carries a uint16.
//	FieldTypeUint16Type
//	// FieldTypeUint8Type indicates that the field carries a uint8.
//	FieldTypeUint8Type
//	// FieldTypeUintptrType indicates that the field carries a uintptr.
//	FieldTypeUintptrType
//	// FieldTypeReflectType indicates that the field carries an interface{}, which should
//	// be serialized using reflection.
//	FieldTypeReflectType
//	// FieldTypeNamespaceType signals the beginning of an isolated namespace. All
//	// subsequent fields should be added to the new namespace.
//	FieldTypeNamespaceType
//	// FieldTypeStringerType indicates that the field carries a fmt.Stringer.
//	FieldTypeStringerType
//	// FieldTypeErrorType indicates that the field carries an error.
//	FieldTypeErrorType
//	// FieldTypeSkipType indicates that the field is a no-op.
//	FieldTypeSkipType
//
//	// FieldTypeInlineMarshalerType indicates that the field carries an ObjectMarshaler
//	// that should be inlined.
//	FieldTypeInlineMarshalerType
//)

//	type Field struct {
//		Key       string
//		Type      FieldType
//		Integer   int64
//		String    string
//		Interface interface{}
//	}
type Log interface {
	Debug(template string, args ...interface{})
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
}

// impl ...
type impl struct {
	log *zap.Logger
}

func NewLog() Log {
	config := zap.NewProductionConfig()
	logger, _ := config.Build()
	return &impl{log: logger}
}
func (i *impl) Debug(template string, args ...interface{}) {
	i.log.Sugar().Infof(template, args...)
}

func (i *impl) Info(msg string, fields ...zap.Field) {
	i.log.Info(msg, fields...)
}

func (i *impl) Warn(msg string, fields ...zap.Field) {
	i.log.Warn(msg, fields...)
}

func (i *impl) Error(msg string, fields ...zap.Field) {
	stackFiled := zap.StackSkip("stack", 1)
	fields = append(fields, stackFiled)
	i.log.Error(msg, fields...)
}

func (i *impl) Panic(msg string, fields ...zap.Field) {
	stackFiled := zap.StackSkip("stack", 1)
	fields = append(fields, stackFiled)
	i.log.Panic(msg, fields...)
}

//func fieldsToZapFields(fields []Field) []zap.Field {
//	f := make([]zap.Field, 0, len(fields))
//	for _, field := range fields {
//		f = append(f, zap.Field{
//			Key:       field.Key,
//			Type:      zapcore.FieldType(field.Type),
//			Integer:   field.Integer,
//			String:    field.String,
//			Interface: field.Interface,
//		})
//	}
//	return f
//}

var _ Log = &impl{}
