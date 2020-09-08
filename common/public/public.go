package public

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	随机数
*/

// 固定位数数字(6位验证码)
func GetCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
func GetCurrentTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// mongodb时间加8小时得到正常时间
func TimeUtcToCst(t time.Time) time.Time {
	return t.Add(time.Hour * time.Duration(8))
}

// 自定义范围内随机整数
func GetRangeRand(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// 在orgStr中以subStr为查找位置获取长度为cutlen的字符串
// ("Colour the interception dddd","interception",4)获取dddd,如果带空格需要设置长度为4+1=5
func GetSubString(parentStr, subStr string, cutlen int) string {
	var (
		index, subStrLen, orgLen int
		part                     string
	)
	index = strings.Index(parentStr, subStr)
	subStrLen = len(subStr)
	orgLen = index + subStrLen
	cutlen = orgLen + cutlen
	part = parentStr[orgLen:cutlen]
	part = strings.Replace(part, " ", "", -1)
	return part
}

// 移除父串中的子串
func RemoveSubString(parentStr, subStr string) string {
	var (
		index, subStrLen, orgLen int
	)
	index = strings.Index(parentStr, subStr)
	subStrLen = len(subStr)
	orgLen = index + subStrLen
	var part string
	if index > 0 {
		part = parentStr[:index] + parentStr[orgLen:]
	} else {
		part = parentStr[orgLen:]
	}
	return part
}

// 移除所有空格和制表符
func RemoveTrimAndTabs(oldStr string) string {
	var (
		reg    *regexp.Regexp
		newStr string
	)
	reg = regexp.MustCompile("\\s+")
	newStr = reg.ReplaceAllString(oldStr, "")
	return newStr
}

// 移除所有空格和换行符
// 例如 " 去 除空 \n格与换行 \n后 "
func RemoveTrimAndLineTabs(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

// 串拼接
func StringBuilder(s ...string) string {
	var str strings.Builder
	for _, v := range s {
		str.WriteString(v)
	}
	return str.String()
}

// 切片拼接
func StringBuilderBySlice(s []string) string {
	var str strings.Builder
	for i := range s {
		str.WriteString(s[i])
	}
	return str.String()
}

// str在串中出现的次数
func SubStringCount(s string, sep string) int {
	return strings.Count(s, sep)
}

// 串中是否包含子串
func IsContainSubString(s, sub string) bool {
	return strings.Contains(s, sub)
}

// 替换字符串中制定的部分
// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n=-1会替换所有old子串。
func StringReplace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// json to struct
func JsonToStruct(js string, data interface{}) (interface{}, error) {
	var err error
	if err = json.Unmarshal([]byte(js), data); err != nil {
		return nil, err
	}
	return data, nil
}

// struct to json
func StructTojson(data interface{}) (string, error) {
	value, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(value), err
}

// string slice 去重
func StringSliceRepeated(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// object slice 去重
func SliceRepeated(arr []interface{}) (newArr []interface{}) {
	newArr = make([]interface{}, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//func GetCurrentTimeString() string {
//	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
//}

//公共处理错误
func CommonERROR(err error) {
	if err != nil {
		fmt.Println("unusual")
	}
}

//time.now()转字符串
func TimeNowToStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//string转时间戳
func StrToTimeInt64(str string) int64 {
	local, _ := time.LoadLocation("Local")
	bTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, local)
	return bTime.Unix()
}

//time 时间戳
func TimeStamp() int64 {
	return time.Now().Unix() //当前时间
}

func DealErr() { //defer就是把匿名函数压入到defer栈中，等到执行完毕后或者发生异常后调用匿名函数
	err := recover() //recover是内置函数，可以捕获到异常
	if err != nil {  //说明有错误
		fmt.Println("defer函数内部错误捕捉=============", err)
		//当然这里可以把错误的详细位置发送给开发人员
		//send email to admin
	}
}

//保留两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//获取上个月月份
func FindLastMonth() string {
	currentYear, currentMonth, _ := time.Now().Date()
	firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, time.Now().Location())
	S_day := firstOfMonth.Format("2006-01")
	return S_day
}

//获取上个月第一天和最后一天
func QueryLastMonthFirstAndLast() (string, string) {
	currentYear, currentMonth, _ := time.Now().Date()
	firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, time.Now().Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	S_day := firstOfMonth.Format("2006-01-02")
	E_day := lastOfMonth.Format("2006-01-02")
	return S_day, E_day
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//获取上个月日期  2019-09
func QueryLastMonth() string {
	currentYear, currentMonth, _ := time.Now().Date()
	firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, time.Now().Location())
	S_day := firstOfMonth.Format("2006-01")
	return S_day
}

//float to  str
func FloatToStr(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

//#string到float64
func StringToFloat64(str string) float64 {
	float1, _ := strconv.ParseFloat(str, 64)
	return float1
}

func StringToInt(str string) int {
	strToint, _ := strconv.Atoi(str)
	return strToint
}

func ToJsonStr(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func StringToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

/*
函数功能:字符串时间转时间戳
*/
func DateToTimeStamp(s string) int64 {
	local, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", s, local)
	return tm.Unix()
}

/*
函数功能:时间戳转时间字符串
*/
func TimeStampToDate(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02")
}

/*
函数功能:返回float64的较小值
*/
func ReturnFloatMin(f1, f2 float64) float64 {
	f := f1
	if f1 > f2 {
		f = f2
	}
	return f
}
