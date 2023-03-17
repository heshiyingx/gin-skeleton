package initmodule

import (
	"github.com/spf13/viper"
	"reflect"
	"strconv"
)

const DEFAULT_KEY = "default"

func ConfigToModel(file string, m interface{}) error {
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")
	//v.AddConfigPath(wd)
	//reflectT(m)

	v.SetConfigName("config")
	v.SetConfigFile(file)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(m)
	if err != nil {
		return err
	}
	fillDefault(m)
	return nil
}

func fillDefault(m interface{}) {
	tm := reflect.TypeOf(m)
	if tm.Kind() != reflect.Pointer {
		return
	}
	vmps := reflect.ValueOf(m).Elem()
	if vmps.Kind() != reflect.Struct {
		return
	}
	fileStructDefaultValue(vmps)

}
func fileStructDefaultValue(vStruct reflect.Value) {
	fieldNum := vStruct.NumField()
	tempTm := vStruct.Type()
	for i := 0; i < fieldNum; i++ {
		field := vStruct.Field(i)
		switch field.Kind() {
		case reflect.Bool:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseBool(v)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(vt))
		case reflect.Int:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(int(vt)))
		case reflect.Int8:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(int8(vt)))
		case reflect.Int16:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseUint(v, 10, 16)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(int16(vt)))
		case reflect.Int32:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(int32(vt)))
		case reflect.Int64:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(vt))
		case reflect.Uint:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(uint(vt)))
		case reflect.Uint8:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 8)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(uint8(vt)))
		case reflect.Uint16:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 16)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(uint16(vt)))
		case reflect.Uint32:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(uint32(vt)))
		case reflect.Uint64:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(uint64(vt)))
		case reflect.Float32:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseFloat(v, 32)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(float32(vt)))
		case reflect.Float64:

			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			vt, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}
			field.Set(reflect.ValueOf(float64(vt)))
		case reflect.Array:
			fileValueSlice(field)
		case reflect.Pointer:
			fileValuePointer(field)
		case reflect.Slice:
			fileValueSlice(field)
		case reflect.Struct:
			fileStructDefaultValue(field)
		case reflect.String:
			v := tempTm.Field(i).Tag.Get(DEFAULT_KEY)
			if !field.IsZero() || v == "" {
				continue
			}
			field.Set(reflect.ValueOf(v))
		}
	}
}
func fileValuePointer(v reflect.Value) {
	elem := v.Elem()
	if v.IsZero() || elem.Type().Kind() != reflect.Struct {
		return
	}
	fileStructDefaultValue(elem)
}
func fileValueSlice(v reflect.Value) {
	for j := 0; j < v.Len(); j++ {
		indexObj := v.Index(j)
		switch indexObj.Kind() {
		case reflect.Struct:
			fileStructDefaultValue(indexObj)
		case reflect.Pointer:
			fileValuePointer(indexObj)
		}
	}
}
