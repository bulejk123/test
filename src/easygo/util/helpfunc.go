package util

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const TIMEFORMAT = "2006-01-02 15:04:05" //! 时间格式化


func HF_I64stoas(i64s ...int64) []string {
	res := make([]string, len(i64s))
	for i, i64 := range i64s {
		res[i] = strconv.FormatInt(i64, 10)
	}
	return res
}
func HF_F64toa(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func RemoveFromSlice(dest []int64, val int64) []int64 {
	if dest == nil {
		return nil
	}
	length := len(dest)
	for i, v := range dest {
		if v == val {
			if i == length-1 {
				dest = dest[0:i]
			} else {
				dest = append(dest[0:i], dest[i+1:]...)
			}
		}
	}
	return dest
}

//GetTimeLastDayStartAndEnd 返回给定时间的昨天的开始和结束时间
func GetTimeLastDayStartAndEnd(t time.Time) (time.Time, time.Time) {
	y, m, d := t.Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	end := start.Add(-24 * time.Hour)
	return end, start
}


func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func In64(slice []int64, elem int64) bool {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == elem {
			return true
		}
	}
	return false
}


// 浮点数除法取整
func HF_DecimalDivide(dividend, divisor float64, n uint8) float64 {
	if divisor == 0 {
		return dividend
	}
	f := fmt.Sprintf("%%0.%df", n)
	// fmt.Println(f)
	v, _ := strconv.ParseFloat(fmt.Sprintf(f, dividend/divisor), 64)
	return v
}




//! int取最小
func HF_MinInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

//! int取最大
func HF_MaxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func HF_MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func HF_Abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

//! 数字 转 []byte

func HF_Itobytes(i int) []byte {
	return []byte(strconv.Itoa(i))
}
func HF_I64tobytes(i int64) []byte {
	return HF_Atobytes(HF_I64toa(i))
}

//! []byte 转 数字
func HF_Bytestoi(b []byte) int {
	return HF_Atoi(HF_Bytestoa(b))
}
func HF_Bytestoi64(b []byte) int64 {
	return HF_Atoi64(HF_Bytestoa(b))
}

//! 字符串转数字
func HF_Atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func HF_Itoa(i int) string {
	return strconv.Itoa(i)
}

func HF_Atoi64(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

//! 数字转string
func HF_I64toa(i int64) string {
	return strconv.FormatInt(i, 10)
}

//! string 转 []byte
func HF_Atobytes(s string) []byte {
	return []byte(s)
}

//! []byte 转 string
func HF_Bytestoa(b []byte) string {
	return string(b[:])
}

//! 结构转json串
func HF_JtoA(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

//! 结构转json串
func HF_JtoB(v interface{}) []byte {
	s, _ := json.Marshal(v)
	return s
}

//! 字符串转float32
func HF_Atof(s string) float32 {
	num, _ := strconv.ParseFloat(s, 32)
	return float32(num)
}

//! 字符串转float64
func HF_Atof64(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}

//! 得到一个随机数
func HF_GetRandom(num int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(1000))).Intn(num)
}

//! 得到一个随机字符串（例如：邀请码）
func HF_GetRandomString(max int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < max; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//! 得到一个随机数字符串
func HF_GetRandomNumberString(max int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < max; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//! 克隆对象 dst为指针
func HF_DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

//! 是否合法
func HF_IsLicitName(name []byte) bool {
	for i := 0; i < len(name); i++ {
		if name[i] == '\r' {
			return false
		} else if name[i] == '\'' {
			return false
		} else if name[i] == '\n' {
			return false
		} else if name[i] == ' ' {
			return false
		} else if name[i] == '	' {
			return false
		}
	}

	return true
}

//! 得到ip
func HF_GetHttpIP(req *http.Request) string {
	ip := req.Header.Get("Remote_addr")
	if ip == "" {
		ip = req.RemoteAddr
	}
	return strings.Split(ip, ":")[0]
}

/*
	往切片中插入一个元素，需提供：切片，插入坐标，插入元素
	这个函数不会检查类型是否一致，使用前确保切片和插入元素类型一致
*/
func HF_InsertOne(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface()
}

/*
	根据字段名字设置字段的值
	struc_dst	结构体地址
	column   	字段名
	set      	设置的目标值
*/
func HF_Setter(struc_dst interface{}, column string, set interface{}) error {
	struc := reflect.ValueOf(struc_dst).Elem()
	if !(struc.Kind() == reflect.Struct) {
		return errors.New("the args indexed 0 is not a struct")
	}
	field := struc.FieldByName(column)
	_set := reflect.ValueOf(set)
	if !field.IsValid() {
		return errors.New(fmt.Sprintf("Setter->the column named:%s not exist", column))
	}
	if !field.CanSet() {
		return errors.New("Setter->can ont set")
	}
	if _set.Kind() == reflect.Invalid {
		field.Set(reflect.Zero(field.Type()))
		return nil
	}
	if !(field.Kind() == _set.Kind()) {
		return errors.New("Setter->type mismatch")
	}
	field.Set(_set)
	return nil
}



/*
	判断文件或文件夹是否存在
*/
func HF_PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

/*
	元素去重
*/
func HF_RemoveRep(slc []byte) []byte {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepByMap(slc)
	}
}

// 通过两重循环过滤重复元素
func removeRepByLoop(slc []byte) []byte {
	result := make([]byte, 0) // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []byte) []byte {
	result := make([]byte, 0)
	tempMap := map[byte]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// slice int 转 byte
func HF_IntsToBytes(slc []int) []byte {
	byteCard := make([]byte, 0)
	for _, value := range slc {
		byteCard = append(byteCard, byte(value))
	}
	return byteCard
}

// slice byte 转 []int
func HF_BytesToInts(slc []byte) []int {
	intCard := make([]int, 0)
	for _, value := range slc {
		intCard = append(intCard, int(value))
	}
	return intCard
}

// int64 转 []byte 2进制流
func HF_Int64ToBytes(num int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(num))
	return buf
}

// []byte 转 int64
func HF_BytesToInt64(slc []byte) int64 {
	return int64(binary.BigEndian.Uint64(slc))
}

// int 转 []byte
func HF_IntToBytes(num int) []byte {
	return []byte(strconv.Itoa(num))
}



// 通过方法名调用对象的方法
func HF_CallMethod(structdst interface{} /*结构体指针*/, callMethodname string /*调用的方法名字*/, args ...interface{} /*方法所需要的参数*/) ([]reflect.Value, error) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	if v := reflect.ValueOf(structdst).MethodByName(callMethodname); v.String() == "<invalid Value>" {
		//syslog.Logger().Error("Call method named: [" + callMethodname + "] #FAILED#!")
		return nil, errors.New("The method invoked does not exist")
	} else {
		//syslog.Logger().Error("Call method named: [" + callMethodname + "] #SUCCEED#!")
		return v.Call(inputs), nil
	}
}



// 得到执行路径
func HF_GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return p, nil
}

// 得到进程名字
func HF_GetCourseName() string {
	courseName := filepath.Base(os.Args[0])
	if runtime.GOOS == "windows" {
		courseName = strings.Replace(courseName, ".exe", "", -1)
	}
	return courseName
}

// 读取yaml文件
func HF_ReadYaml(path string, name string, v interface{}) error {
	cf := filepath.Join(path, name)
	data, err := ioutil.ReadFile(cf)
	if err != nil {
		return fmt.Errorf("HF_ReadYaml:%s", err)
	}
	if err := yaml.Unmarshal(data, v); err != nil {
		return fmt.Errorf("HF_ReadYaml:%s", err)
	}
	//fmt.Println(fmt.Sprintf("ReadYaml succeed : %+v ",v))
	return nil
}

// 合并二进制流
func HF_BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}



// 得到今天剩余的秒数
func HF_GetTodayRemainSecond() int64 {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)
	return todayLastTime.Unix() - time.Now().Local().Unix()
}

func RAN(min, max int) int {
	return INT_N(max-min+1) + min
}

func INT_N(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}






// 根据身份证号码获取年龄
func HF_GetAgeFromIdcard(idcard string) int {
	// 身份证解密
	//idcard, _ = HF_DecodeStr(HF_Atobytes(idcard), UserEncodeKey)
	if len(idcard) <= 0 || len(idcard) > 18 {
		return -1
	}
	// 获取一下当前时间
	timeObj := time.Now()
	year := timeObj.Year()
	month := int(timeObj.Month())
	day := timeObj.Day()
	// 合法的身份证号 获取一下出身年月日
	birthYear, err1 := strconv.Atoi(idcard[6:10])
	birthMonth, err2 := strconv.Atoi(idcard[10:12])
	birthDay, err3 := strconv.Atoi(idcard[12:14])
	if err1 != nil || err2 != nil || err3 != nil {
		return -1
	} else {
		if year > birthYear {
			if month > birthMonth || (month == birthMonth && day >= birthDay) {
				return year - birthYear
			} else {
				return year - birthYear - 1
			}
		} else if year == birthYear {
			return 0
		} else {
			return -1
		}
	}
}

func HF_SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
