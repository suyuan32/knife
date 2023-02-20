// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package format

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// ChineseStandard contains date from year to second
	ChineseStandard = "2006年1月2日 3时4分5秒"

	// ChineseYearToDay contains date from year to month
	ChineseYearToDay = "2006年1月2日"

	// ChineseHourToSecond contains date from hour to second
	ChineseHourToSecond = "3时4分5秒"

	// ChineseYearToWeek contains date from year to weekday
	ChineseYearToWeek = "2006年1月2日 Mon"

	// ChineseYearToSecondWithWeek contains date from year to second with weekday
	ChineseYearToSecondWithWeek = "2006年1月2日 Mon 3时4分5秒"
)

type ChineseTime struct {
	Time time.Time
	Mode uint8
}

// ChineseFormat format date to Chinese string
// The layout include ChineseStandard, ChineseYearToWeek and so on.
// The mode 0 means the weekday use "星期" as prefix, else use "周" as prefix.
// e.g.
//
//	targetDate := time.Now()
//	chineseDate := ChineseTime{Time: &targetTime, mode: 1}
func (c *ChineseTime) ChineseFormat(layout string) string {
	switch layout {
	case ChineseStandard, ChineseYearToDay, ChineseHourToSecond:
		return c.Time.Format(layout)
	case ChineseYearToWeek, ChineseYearToSecondWithWeek:
		return convertWeekdayFromEnglishToChinese(c.Time.Format(layout), c.Mode)
	default:
		return c.Time.Format(ChineseStandard)
	}
}

func convertWeekdayFromEnglishToChinese(dateString string, mode uint8) string {
	dateSplit := strings.Split(dateString, " ")
	if mode == 0 {
		switch dateSplit[1] {
		case "Mon":
			dateSplit[1] = "星期一"
		case "Tue":
			dateSplit[1] = "星期二"
		case "Wed":
			dateSplit[1] = "星期三"
		case "Thu":
			dateSplit[1] = "星期四"
		case "Fri":
			dateSplit[1] = "星期五"
		case "Sat":
			dateSplit[1] = "星期六"
		case "Sun":
			dateSplit[1] = "星期日"
		}
	} else {
		switch dateSplit[1] {
		case "Mon":
			dateSplit[1] = "周一"
		case "Tue":
			dateSplit[1] = "周二"
		case "Wed":
			dateSplit[1] = "周三"
		case "Thu":
			dateSplit[1] = "周四"
		case "Fri":
			dateSplit[1] = "周五"
		case "Sat":
			dateSplit[1] = "周六"
		case "Sun":
			dateSplit[1] = "周日"
		}
	}
	return strings.Join(dateSplit, " ")
}

// Parse set the time from a date string such as 2006年1月2日 3时4分5秒
func (c *ChineseTime) Parse(dateString, layout string) (err error) {
	var year, month, day, hour, minute, second int
	dateRune := []rune(dateString)
	numberExtract := func(data []rune) (result []int, err error) {
		var tmpString strings.Builder
		for _, v := range dateRune {
			if v >= 48 && v <= 57 {
				tmpString.WriteRune(v)
			} else if tmpString.Len() > 0 {
				strConv, err := strconv.Atoi(tmpString.String())
				if err != nil {
					return nil, fmt.Errorf("failed to parse rune to int, err: %v", err)
				}
				result = append(result, strConv)
				tmpString.Reset()
			} else {
				tmpString.Reset()
			}
		}
		return result, err
	}

	numberArray, err := numberExtract(dateRune)
	if err != nil {
		return err
	}

	switch layout {
	case ChineseStandard:
		year = numberArray[0]
		month = numberArray[1]
		day = numberArray[2]
		hour = numberArray[3]
		minute = numberArray[4]
		second = numberArray[5]
	case ChineseYearToDay:
		year = numberArray[0]
		month = numberArray[1]
		day = numberArray[2]
	default:
		year = numberArray[0]
		month = numberArray[1]
		day = numberArray[2]
	}

	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	return err
}
