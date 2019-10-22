// ContainerUtils.go
package arr

import (
	"reflect"
)

//容器是否包含某个元素。适用：array slice map (https://blog.csdn.net/qq_28612967/article/details/97392727)
func ContainElem(arr interface{}, elem interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrValue.Len(); i++ {
			if arrValue.Index(i).Interface() == elem {
				return true
			}
		}
	case reflect.Map:
		if arrValue.MapIndex(reflect.ValueOf(elem)).IsValid() {
			return true
		}
		// case reflect.Chan://
		// case reflect.String:用strings.index()代替
	}

	return false
}
