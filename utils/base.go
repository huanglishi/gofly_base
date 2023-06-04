package utils

import (
	"fmt"
	"gofly/app/model"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取ip函数
func GetIp(c *gin.Context) string {
	reqIP := c.Request.Header.Get("X-Forwarded-For")
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// 判断元素是否存在数组中
func IsContain(items []interface{}, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 多维数组合并
func ArrayMerge(data []interface{}) []interface{} {
	var rule_ids_arr []interface{}
	for _, mainv := range data {
		ids_arr := strings.Split(mainv.(string), `,`)
		for _, intv := range ids_arr {
			rule_ids_arr = append(rule_ids_arr, intv)
		}
	}
	return rule_ids_arr
}

// 多维数组合并
func ArraymoreMerge(data []interface{}) []interface{} {
	var rule_ids_arr []interface{}
	for _, mainv := range data {
		ids_arr := strings.Split(mainv.(string), `,`)
		for _, intv := range ids_arr {
			rule_ids_arr = append(rule_ids_arr, intv)
		}
	}
	return rule_ids_arr
}

// 合并数组-两个数组合并为一个数组
func MergeArr(a, b []interface{}) []interface{} {
	var arr []interface{}
	for _, i := range a {
		arr = append(arr, i)
	}
	for _, j := range b {
		arr = append(arr, j)
	}
	return arr
}

// JSONMarshalToString JSON编码为字符串
func JSONMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

// 获取子菜单包含的父级ID-返回全部ID
func GetRulesID(tablename string, field string, menus interface{}) interface{} {
	menus_rang := menus.([]interface{})
	var fnemuid []interface{}
	for _, v := range menus_rang {
		fid := getParentID(tablename, field, v)
		if fid != nil {
			fnemuid = MergeArr(fnemuid, fid)
		}
	}
	r_nemu := MergeArr(menus_rang, fnemuid)
	uni_fnemuid := UniqueArr(r_nemu) //去重
	return uni_fnemuid
}

// 获取所有父级ID
func getParentID(tablename string, field string, id interface{}) []interface{} {
	var pids []interface{}
	pid, _ := model.DB().Table(tablename).Where("id", id).Value(field)
	if pid != nil {
		a_pid := pid.(int64)
		var zr_pid int64 = 0
		if a_pid != zr_pid {
			pids = append(pids, a_pid)
			getParentID(tablename, field, pid)
		}
	}
	return pids
}

// 去重
func UniqueArr(m []interface{}) []interface{} {
	d := make([]interface{}, 0)
	tempMap := make(map[int]bool, len(m))
	for _, v := range m { // 以值作为键名
		keyv := GetInterfaceToInt(v)
		if tempMap[keyv] == false {
			tempMap[keyv] = true
			d = append(d, v)
		}
	}
	return d
}

// interface{}转int
func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 获取随机数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
