package utils


import (
  "reflect"
	"strconv"
)

func StrToInt[T interface{}](str string, valueNum *T, size int) error {
  num, err := strconv.ParseUint(str, 10, size)
  if err != nil {
    return err
  }
  value := reflect.ValueOf(valueNum).Elem()
  switch value.Kind() {
  case reflect.Uint8:
    value.SetUint(uint64(num))
  case reflect.Uint64:
    value.SetUint(num)
  default:
    return err
  }
  return nil
}
