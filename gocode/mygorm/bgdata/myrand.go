package bgdata

import (
	"baotian0506.com/39_config/gocode/mygorm/fixdata"
	"fmt"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

var globalInt int64 = time.Now().UnixNano()

type MyRand struct {
}

func (s *MyRand) RandInt(min, max int64) int64 {
	rand.Seed(atomic.AddInt64(&globalInt, 1))
	return min + rand.Int63n(max-min)
}

// GetCn returns as a string, 一个随机长度为size的中文名称
// 返回值没有空格
func (s *MyRand) GetCn(size int64) string {
	a := make([]rune, size)
	for i := range a {
		a[i] = rune(s.RandInt(CN_MIN_UNICODE, CN_MAX_UNICODE))
	}
	return string(a)
}

// GetCnName 获取一个中文名
func (s *MyRand) GetCnName() string {
	size := int64(len(fixdata.FrequentlyUsedChineseList))
	nameSize := s.RandInt(2, 10)
	start := s.RandInt(0, size-nameSize)

	return strings.Join(fixdata.FrequentlyUsedChineseList[start:start+nameSize], "")
}

// GetGender returns as a string, 获取随机性别（中文名称，非代号）的方法
func (s *MyRand) GetGender() string {
	v := s.RandInt(0, 3)
	switch v {
	case 0:
		return "男"
	case 1:
		return "女"
	case 2:
		return "男生"
	case 3:
		return "女生"
	default:
		return "未知"
	}
}

// GetIp returns as a string, 随机获取一个ip
// 示例： 192.168.33.456
func (s *MyRand) GetIp() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		s.RandInt(100, 300),
		s.RandInt(301, 600),
		s.RandInt(0, 100),
		s.RandInt(601, 999),
	)
}

// GetCountry returns as as *fixdata.Country, 随机获取一个国家信息
func (s *MyRand) GetCountry() *fixdata.Country {
	rand.Seed(time.Now().UnixNano())
	return &fixdata.CountryJson[rand.Intn(len(fixdata.CountryJson))]
}

// GetSameSuffix 获取后缀一致的汉字字符串
func (s *MyRand) GetSameSuffix() string {
	return s.GetCnName() + "后缀一致"
}

// GetSamePrefix 获取前缀一致的汉字字符串
func (s *MyRand) GetSamePrefix() string {
	return "前缀一致" + s.GetCnName()
}

// GetCnChar 获取一个中文字符
func (s *MyRand) GetCnChar() string {
	size := int64(len(fixdata.FrequentlyUsedChineseList))

	start := s.RandInt(0, size-1)
	return strings.Join(fixdata.FrequentlyUsedChineseList[start:start+1], "")
}
