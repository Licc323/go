package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini 配置文件解析器

//MysqlConfig MySQL 配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Psssword string `ini:"password"`
}

type RedisConfig struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
	Password string `ini:"password"`
	Database int `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{})(err error)  {
	//0.参数的校验
	//0.1传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") //新创建一个错误
		return
	}
	//0.2传进来的data参数必须是结构体类型的指针(因为配置文件中各种键值对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct{
		err = errors.New("data param should be a struct pointer")
	}
	//1.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b)//将字节类型的文件内容转化成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Println("%#v\n",lineSlice)
	//2.一行一行的读数据
	var structName string //
	for idx, line := range lineSlice {
		//去掉每个字符串首尾的空格
		line = strings.TrimSpace(line)
		//如果是空行跳过
		if len(line) == 0 {
			continue
		}
		//2.1如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#"){
			continue
		}
		//2.2如果是[开头的就表示节(section)
		if strings.HasPrefix(line, "["){
			if line[0] == '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//把这一行首尾的[]去掉 取到中间的内容把首尾的空格去掉拿到内容
			 sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
				//根据字符串sectionName去掉data里面根据反射找到对应的结构体
				for i := 0; i < t.Elem().NumField(); i++ {
					field := t.Elem().Field(i)
					if sectionName == field.Tag.Get("ini") {
						//说明找到了对应的嵌套结构体，把字段名记下来
						structName = field.Name
						fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
					}
				}
			}
		}else {
			//2.3如果不是[开头就是=分割的键值对
			//1.以等号分割这一行， 等号左边是key 等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			//value := strings.TrimSpace(line[index+1:])
			//2.根据structName 去 data 里面把对应的嵌套结构体找出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type() //拿到嵌套结构体的类型信息

			//structObj := v.Elem().FieldByName(structName)
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s 字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			//3.遍历嵌套结构体的每一个字段，判断tag 是不是等于key
			var fileType reflect.StructField
			for i:=0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) //tag 信息是存储在累心信息中
				if filed.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			//4.如果key = tag， 给这个字段赋值
			//4.1根据fieldName 去取出这个字段
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			//4.2对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString("value")
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

				var valueInt int64
				valueInt, err := strconv.ParseInt("value", 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
				}
				fileObj.SetInt(valueInt)

			}
		}
	}
	return
}
func main()  {
	var cfg Config 
	//var x = new(int)
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
