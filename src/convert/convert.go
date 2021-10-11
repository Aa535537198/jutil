package jconvert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 转换成string
func ToStr(src interface{}) string {

	switch src.(type) {
	case nil:
		return ""
	case int, int8, int16, int64, int32:
		return fmt.Sprintf("%d", src)
	case float32, float64:
		return fmt.Sprintf("%f", src)
	default:
		bs, _ := json.Marshal(src)
		return string(bs)
	}
}

// 转换成int
func ToInt(src interface{}) int {
	switch src.(type) {
	case int:
		return src.(int)
	case int8:
		return (int)(src.(int8))
	case int16:
		return (int)(src.(int16))
	case int32:
		return (int)(src.(int32))
	case int64:
		return (int)(src.(int64))
	case string:
		i, _ := strconv.Atoi(src.(string))
		return i
	case []int8:
		var buffer bytes.Buffer
		for _, i8 := range src.([]int8) {
			buffer.WriteString(ToStr(i8))
		}
		return ToInt(buffer.String())
	case []int16:
		var buffer bytes.Buffer
		for _, i16 := range src.([]int8) {
			buffer.WriteString(ToStr(i16))
		}
		return ToInt(buffer.String())
	case []int32:
		var buffer bytes.Buffer
		for _, i32 := range src.([]int32) {
			buffer.WriteString(ToStr(i32))
		}
		return ToInt(buffer.String())
	case []int64:
		var buffer bytes.Buffer
		for _, i64 := range src.([]int64) {
			buffer.WriteString(ToStr(i64))
		}
		return ToInt(buffer.String())
	default:
		panic("toInt is not s expect type")
	}
}

const (
	timeLayout = "2006-01-02 15:04:05"
)

var suffixTime = make(map[int]string)

func init() {
	suffixTime[10] = " 00:00:00"
	suffixTime[13] = ":00:00"
	suffixTime[16] = ":00"
	suffixTime[19] = ""
}

//  转化为日期格式
//  @param 2017-12-23 13:23:59
//  @param 2017/12/23 13:23:59
//  example: convert.ToDateStr("2017/12/23 12-56-36", "/", "-")
func StrToTime(date string, params ...string) time.Time {
	l := len(date) // 日期字符串的长度
	switch l {
	case 10, 13, 16, 19:
		date += suffixTime[l]
		space := strings.Count(date, " ")
		if space == 0 {
			panic("input time str is error, please check")
		}

		switch len(params) {
		case 0:
			date = strings.Replace(date, "/", "-", 2)
		case 2:
			ss := strings.Split(date, " ")
			s := ss[1]
			s = strings.Replace(s, params[1], ":", 2)
			date = ss[0] + " " + s
			fallthrough
		case 1:
			date = strings.Replace(date, params[0], "-", 2)
		default:
			panic("input param is too long")
		}
	default:
		panic("input time str is error, please check")
	}
	location, _ := time.LoadLocation("Local") // 获得当前时区
	theTime, _ := time.ParseInLocation(timeLayout, date, location)
	return theTime
}

// @example toStr := convert.TimeToStr(&str, "/", "/")
// 2017/12/23 12/56/36
func TimeToStr(t *time.Time, param ...string) string {
	timeStr := t.Format(timeLayout)
	switch len(param) {
	case 2:
		timeStr = strings.ReplaceAll(timeStr, ":", param[1])
		fallthrough
	case 1:
		timeStr = strings.Replace(timeStr, "-", param[0], 2)
	case 0:
	default:
		panic("input param is too long")
	}
	return timeStr
}

// 时间转化成时间戳
func TimeToStamp(t *time.Time) int64 {
	millisecond := t.UnixNano() / 1e6
	return millisecond
}

// 时间戳转换成字符串
func StampToStr(stamp int64, param ...string) string {
	timeStr := time.Unix(stamp/1000, 0).Format(timeLayout)
	switch len(param) {
	case 2:
		timeStr = strings.ReplaceAll(timeStr, ":", param[1])
		fallthrough
	case 1:
		timeStr = strings.Replace(timeStr, "-", param[0], 2)
	case 0:
	default:
		panic("input param is too long")
	}
	return timeStr
}

// 时间戳转换成时间
func StampToTime(stamp int64) time.Time {
	timeStr := StampToStr(stamp)
	return StrToTime(timeStr)
}
