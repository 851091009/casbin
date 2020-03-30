package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"strconv"
	"time"
)
// 统计接口前缀
var ApiAdminPrefix = "/v1/admin"

var CasbinTableName = "large_admin_casbin"

// md5 加密
func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source + "FtZj"))
	return hex.EncodeToString(md5h.Sum(nil))
}

// 验证手机号规则
func IsPhone(phone string) bool {
	reg := `^1[3-9][0-9]\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

// 日期转时间戳
func StringToUnix(timeSting string, timeLayout string) int64 {
	//timeLayout := "2006-01-02 15:04:05" // 转化所需模板
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, _ := time.ParseInLocation(timeLayout, timeSting, loc)
	return theTime.Unix()
}

// 时间戳转日期
func UnixToString(unixSting int64) string {
	timeLayout := "2006-01-02 15:04:05" // 转化所需模板
	return time.Unix(unixSting, 0).Format(timeLayout)
}

// float64 保留两位
func Float64ToDecimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// 判断slice中是否存在某个item
func IsExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if value == s.Index(i).Field(1).Interface() {
				fmt.Println("--red--")
				return true
			}
		}
	}
	return false
}

// ToString casts an interface to a string type.
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}


// ToStringE casts an interface to a string type.
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatInt(int64(s), 10), nil
	case uint64:
		return strconv.FormatInt(int64(s), 10), nil
	case uint32:
		return strconv.FormatInt(int64(s), 10), nil
	case uint16:
		return strconv.FormatInt(int64(s), 10), nil
	case uint8:
		return strconv.FormatInt(int64(s), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}


// ToInt64 casts an interface to an int64 type.
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToInt64E casts an interface to an int64 type.
func ToInt64E(i interface{}) (int64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}
