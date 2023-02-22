package idcard

import (
	"regexp"
	"time"

	"github.com/suyuan32/knife/core/regex/matcher"
)

var (
	// Power is the weighting factor.
	Power = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	// ProvinceCode stores the codes which represent Chinese provinces.
	ProvinceCode = map[string]string{
		"11": "北京",
		"12": "天津",
		"13": "河北",
		"14": "山西",
		"15": "内蒙古",
		"21": "辽宁",
		"22": "吉林",
		"23": "黑龙江",
		"31": "上海",
		"32": "江苏",
		"33": "浙江",
		"34": "安徽",
		"35": "福建",
		"36": "江西",
		"37": "山东",
		"41": "河南",
		"42": "湖北",
		"43": "湖南",
		"44": "广东",
		"45": "广西",
		"46": "海南",
		"50": "重庆",
		"51": "四川",
		"52": "贵州",
		"53": "云南",
		"54": "西藏",
		"61": "陕西",
		"62": "甘肃",
		"63": "青海",
		"64": "宁夏",
		"65": "新疆",
		"71": "台湾",
		"81": "香港",
		"82": "澳门",
		"83": "台湾",
		"91": "国外",
	}

	// TaiwanFirstCode stores the first codes'  values.
	TaiwanFirstCode = map[byte]int{
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
		'G': 16,
		'H': 17,
		'J': 18,
		'K': 19,
		'L': 20,
		'M': 21,
		'N': 22,
		'P': 23,
		'Q': 24,
		'R': 25,
		'S': 26,
		'T': 27,
		'U': 28,
		'V': 29,
		'X': 30,
		'Y': 31,
		'W': 32,
		'Z': 33,
		'I': 34,
		'O': 35,
	}
)

// ChineseID stores the ID data and provides some functions.
type ChineseID struct {
	Id       string
	Province string
	Sex      string
	BirthDay time.Time
}

// IsValidCard judges whether the ID is valid.
func (c *ChineseID) IsValidCard() bool {
	if c.Id == "" {
		return false
	}

	switch len(c.Id) {
	case 18:
		if _, ok := ProvinceCode[c.Id[:2]]; !ok {
			return false
		}

		parse, err := time.Parse("20060102", c.Id[6:14])
		if err != nil {
			return false
		}

		c.BirthDay = parse

		if val, ok := ProvinceCode[c.Id[:2]]; ok {
			c.Province = val
		}

		if (int(c.Id[16]-'0') % 2) == 0 {
			c.Sex = "女"
		} else {
			c.Sex = "男"
		}

		if !matcher.NewMatcher(matcher.DigitPure).MatchString(c.Id[:17]) {
			return false
		} else {
			if c.Id[17] != c.GetCheckCode() || c.GetCheckCode() == ' ' {
				return false
			}
		}
		return true
	case 10:
		if ok, err := regexp.MatchString("^[a-zA-Z][0-9]{9}$", c.Id); err == nil && ok {
			c.Province = "台湾"
			return c.IsValidTWCard()
		} else if ok, err := regexp.MatchString("^[157][0-9]{6}\\(?[0-9A-Z]\\)?$", c.Id); err == nil && ok {
			c.Province = "澳门"
			return ok
		} else if ok, err := regexp.MatchString("^[A-Z][0-9]*\\(?[0-9A-Z]\\)?$", c.Id); err == nil && ok {
			c.Province = "香港"
			// Todo: more accuracy validate
			return ok
		} else {
			return false
		}
	default:
		return false
	}
}

// GetCheckCode returns the check code to validate ID
func (c *ChineseID) GetCheckCode() byte {
	if len(c.Id) == 18 || len(c.Id) == 10 {
		if len(c.Id) == 18 {
			valSum := 0
			for i := 0; i < 17; i++ {
				valSum += int(c.Id[i]-'0') * Power[i]
			}

			switch valSum % 11 {
			case 10:
				return '2'
			case 9:
				return '3'
			case 8:
				return '4'
			case 7:
				return '5'
			case 6:
				return '6'
			case 5:
				return '7'
			case 4:
				return '8'
			case 3:
				return '9'
			case 2:
				return 'X'
			case 1:
				return '0'
			case 0:
				return '1'
			}
		}
	}

	return ' '
}

// IsValidTWCard judges whether it is a valid Taiwan Card.
func (c *ChineseID) IsValidTWCard() bool {
	if val, ok := TaiwanFirstCode[c.Id[0]]; ok {
		sumData := val/10 + (val%10)*9
		for i := 8; i >= 1; i-- {
			sumData += int(c.Id[(9-i)]-'0') * i
		}

		var target int
		if sumData%10 == 0 {
			target = 0
		} else {
			target = 10 - sumData%10
		}

		return target == int(c.Id[9]-'0')
	} else {
		return false
	}
}
