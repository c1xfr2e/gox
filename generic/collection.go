package generic

import (
	"fmt"
	"reflect"
)

func In(elem interface{}, collection interface{}) (bool, error) {
	val := reflect.ValueOf(collection)
	switch val.Kind() {
	case reflect.Slice:
		return isElemInSlice(elem, collection, val)
	case reflect.Map:
		return isElemInMap(elem, collection, val)
	default:
		return false, fmt.Errorf("%T is not a collection type", collection)
	}

	return false, nil
}

func isElemInSlice(elem interface{}, obj interface{}, sliceVal reflect.Value) (bool, error) {
	elemtypeOfSlice := reflect.TypeOf(obj).Elem()
	elemtype := reflect.TypeOf(elem)
	if elemtype != elemtypeOfSlice {
		return false, fmt.Errorf("different element type for %v %s and %s[]",
			elem, elemtype.Kind(), elemtypeOfSlice.Kind())
	}

	for i := 0; i < sliceVal.Len(); i++ {
		val := sliceVal.Index(i).Interface()
		if elem == val {
			return true, nil
		}
	}
	return false, nil
}

func isElemInMap(elem interface{}, obj interface{}, mapVal reflect.Value) (bool, error) {
	keytypeOfMap := reflect.TypeOf(obj).Key()
	elemtype := reflect.TypeOf(elem)
	if elemtype != keytypeOfMap {
		return false, fmt.Errorf("different element type for %v %s and map[%s]",
			elem, elemtype.Kind(), keytypeOfMap.Kind())
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(elem))
	return elemVal.IsValid(), nil
}

func ForEach(constainer interface{}, f func(interface{}) error) error {
	return nil
}

func MergeMap(dst interface{}, src interface{}) error {
	return nil
}

func Replace() error {
	return nil
}

func IsCollectionType(v interface{}) {

}
