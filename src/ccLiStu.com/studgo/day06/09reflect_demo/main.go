package main

import (
	"fmt"
	"reflect"
)

type Cat struct {

}
func reflectType(x interface{})  {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())

}

func reflectValue(x interface{})  {
	v := reflect.ValueOf(x)
	k := v.Kind() //值的类型种类
	switch k {
	case reflect.Int64:
		//v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is float32, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		//v.Float()从反射中获取整型的原始值，然后通过flo32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		//V.Float()从反射中获取整个原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))

	}
}

func reflectSetvalue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}
func reflectSetvalue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}
func main()  {
	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int64 = 100
	reflectType(b) //type : int64

	var c = Cat{}
	reflectType(c)
	//valueOf
	reflectValue(a)
	//设置值
	//reflectSetvalue1(&b)
	reflectSetvalue2(&b)
	fmt.Println(b)

}