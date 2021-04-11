package utils

// import (
// 	"errors"
// 	"fmt"
// 	"beego-learn/conf"
// 	. "reflect"
// 	"strconv"
// 	"strings"
// )
//
// const ConfigTagName = "conf"
//
// func MapToConstruct(configMap map[string] interface{}, bean interface{}) {
// 	rv, ok := bean.(Value)
// 	if !ok {
// 		rv = Indirect(ValueOf(bean))
// 	}
// 	fieldType := rv.Type()
// 	fieldCount := fieldType.NumField()
//
// 	// 遍历结构体字段（有标签就读标签，没有就默认以字段变量名的匹配规则）
// 	// 支持 conf 标签，没有标签
// 	for i := 0; i < fieldCount; i++ {
// 		theField := rv.Field(i)
// 		structField := fieldType.Field(i)
// 		fieldName := structField.Name
//
// 		// ”-“ 忽略标签项支持
// 		fieldTag := structField.Tag.Get(ConfigTagName)
// 		if fieldTag == "-" {
// 			continue
// 		}
//
// 		// 没有定义标签才考虑结构体字段名（要考虑的情况比较多，一下没法搞好）
// 		match, matchValue := mapNameMatch(fieldTag, fieldName, configMap)
// 		if !match {
// 			continue
// 		}
//
// 		// 确认解析，还要根据字段是否是结构体类型决定是否递归解析
// 		// 1、字段是结构体类型
// 		// 2、【配置Map】匹配值是Map类型
// 		matchValueMap, isMap := matchValue.(map[string]interface{})
// 		if theField.Kind() == Struct && isMap {
// 			MapToConstruct(matchValueMap, theField)
// 		}
// 		matchValueString, _ := matchValue.(string)
//
// 		// 读取配置文件参数，为空则报错
// 		if matchValueString == "" {
// 			panic(fmt.Sprintf("配置文件 %s 参数 %s 为空", getFieldTypeName(theField), fieldTag))
// 		}
//
// 		// 根据字段类型来解析参数
// 		err := parseParam(matchValueString, theField)
// 		if err != nil {
// 			panic(fmt.Sprintf("配置文件 %s 参数 %s 解析失败：%v", getFieldTypeName(theField), fieldTag, err))
// 		}
// 	}
// }
//
// func mapNameMatch(tag string, name string, configMap map[string]interface{}) (bool, interface{}) {
// 	// lowerFieldName := utils.CamelCase(fieldName)
// 	// snakeFileName := utils.SnakeString(fieldName)
// 	//
// 	// var matchValue interface{}
// 	// var lowerMatch, snakeMatch bool
// 	// matchValue, lowerMatch = configMap[lowerFieldName]
// 	// matchValue, snakeMatch = configMap[snakeFileName]
// 	// if !lowerMatch && !snakeMatch {
// 	// 	continue
// 	// }
// 	return true, configMap[""]
// }
//
// func getFieldTypeName(fv Value) string {
// 	result := fv.Type().Name()
// 	if result == "" {
// 		result = fv.Kind().String()
// 	}
// 	return result
// }
//
// func parseParam(param string, fv Value) error {
// 	switch fv.Kind() {
// 	case String:
// 		fv.SetString(param)
// 		return nil
// 	case Int, Int8, Int16, Int32, Int64:
// 		intVal, err := strconv.Atoi(param)
// 		if err != nil {
// 			return err
// 		}
// 		fv.SetInt(int64(intVal))
// 		return nil
// 	case Array:
// 		// 检查参数长度，与字段定义的长度是否匹配
// 		arrLen := fv.Len()
// 		values := strings.Split(param, ",")
// 		if len(values) != arrLen {
// 			return fmt.Errorf("列表长度不匹配，期望为[%d]，实际为[%d]", arrLen, len(values))
// 		}
// 		// 根据数组类型来赋值
// 		switch fv.Type().Elem().Kind() {
// 		case String:
// 			for i := 0; i < arrLen; i++ {
// 				fv.Index(i).SetString(values[i])
// 			}
// 			return nil
// 		case Int, Int8, Int16, Int32, Int64:
// 			// 解析参数为数字切片
// 			intValues, err := StrSlice(values).ParseIntSlice()
// 			if err != nil {
// 				return err
// 			}
// 			for i := 0; i < arrLen; i++ {
// 				fv.Index(i).SetInt(int64(intValues[i]))
// 			}
// 			return nil
// 		}
// 	case Slice:
// 		values := strings.Split(param, ",")
// 		// 根据数组类型来赋值
// 		switch fv.Type().Elem().Kind() {
// 		case String:
// 			fv.Set(ValueOf(values))
// 			return nil
// 		case Int:
// 			// 解析参数为数字切片
// 			intValues, err := StrSlice(values).ParseIntSlice()
// 			if err != nil {
// 				return err
// 			}
// 			fv.Set(ValueOf(intValues))
// 			return nil
// 		}
// 	}
// 	return errors.New("暂不支持此类型的参数")
// }
//
// func call(configMap map[string] interface{}) {
// 	// 第二个参数一定得是结构体类型（不清楚现在的设计是不是这样）
// 	MapToConstruct(configMap, conf.Center)
// }
