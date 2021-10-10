package clone

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// 浅克隆
// Elem()就是获取反射数据的实际类型和实际值
// 此方法只能实现非指针遍历的clone，若为指针变量，则为同一元素
func Clone(src interface{}) interface{} {
	if src == nil {
		panic("clone:(src nil)")
	}

	sourceTyp := reflect.TypeOf(src)

	// 指针类型
	if sourceTyp.Kind() == reflect.Ptr {
		// 获得源实际类型
		sourceTyp = sourceTyp.Elem()
		// 创建对象
		dst := reflect.New(sourceTyp).Elem()
		// 源对象
		data := reflect.ValueOf(src).Elem()
		// 将dst的值设置为data的值
		dst.Set(data)

		// address
		dst = dst.Addr()

		return dst.Interface()
	} else {

		// 创建对象
		dst := reflect.New(sourceTyp).Elem()
		data := reflect.ValueOf(src)
		dst.Set(data)
		return dst.Interface()
	}
}

//  基于json的形式深度clone
func DeepClone(src interface{}) interface{} {
	if src == nil {
		panic("clone:(src nil)")
	}
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		dst := reflect.New(typ.Elem()).Elem()
		// 导出json
		b, _ := json.Marshal(src)
		_ = json.Unmarshal(b, dst.Addr().Interface()) //json序列化
		return dst.Addr().Interface()
	} else {
		dst := reflect.New(typ).Elem()                //创建对象
		b, _ := json.Marshal(src)                     //导出json
		_ = json.Unmarshal(b, dst.Addr().Interface()) //json序列化
		return dst.Interface()                        //返回值
	}
}

// 仅限结构体的复制，指针字段则指向同一对象
func CopyProperties(src, dst interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}

// 基于Json的方式，深度复制属性
// clone.DeepCopyProperties(&a, &b)
func DeepCopyProperties(src, dst interface{}) {
	bs, _ := json.Marshal(src)
	_ = json.Unmarshal(bs, dst)
}
