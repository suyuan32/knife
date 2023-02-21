package format

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChineseTime_ChineseFormat(t *testing.T) {
	targetTime := time.Date(2023, 11, 12, 11, 11, 11, 11, time.Local)
	targetChineseTime := ChineseTime{Time: targetTime, Mode: 0}

	assert.Equal(t, "2023年11月12日 11时11分11秒", targetChineseTime.ChineseFormat("2006"))
	assert.Equal(t, "2023年11月12日 11时11分11秒", targetChineseTime.ChineseFormat(ChineseStandard))
	assert.Equal(t, "2023年11月12日", targetChineseTime.ChineseFormat(ChineseYearToDay))
	assert.Equal(t, "11时11分11秒", targetChineseTime.ChineseFormat(ChineseHourToSecond))
	assert.Equal(t, "2023年11月12日 星期日", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTime.Time = targetChineseTime.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月13日 星期一", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTime.Time = targetChineseTime.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月14日 星期二", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTime.Time = targetChineseTime.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月15日 星期三", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTime.Time = targetChineseTime.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月16日 星期四", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTime.Time = targetChineseTime.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月17日 星期五", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTime.Time = targetChineseTime.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月18日 星期六", targetChineseTime.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1 := ChineseTime{Time: targetTime, Mode: 1}
	assert.Equal(t, "2023年11月12日 周日", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1.Time = targetChineseTimeMode1.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月13日 周一", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1.Time = targetChineseTimeMode1.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月14日 周二", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1.Time = targetChineseTimeMode1.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月15日 周三", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1.Time = targetChineseTimeMode1.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月16日 周四", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1.Time = targetChineseTimeMode1.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月17日 周五", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
	targetChineseTimeMode1.Time = targetChineseTimeMode1.Time.AddDate(0, 0, 1)
	assert.Equal(t, "2023年11月18日 周六", targetChineseTimeMode1.ChineseFormat(ChineseYearToWeek))
}

func TestChineseTime_Parse(t *testing.T) {
	wantedTime := time.Date(2023, 11, 12, 11, 11, 11, 0, time.Local)
	targetTime := &ChineseTime{}
	err := targetTime.Parse("2023年11月12日 11时11分11秒", ChineseStandard)
	assert.Nil(t, err)
	assert.Equal(t, targetTime.Time.String(), wantedTime.String())

	wantedTime = time.Date(2023, 11, 12, 0, 0, 0, 0, time.Local)
	err = targetTime.Parse("2023年11月12日", ChineseYearToDay)
	assert.Nil(t, err)
	assert.Equal(t, targetTime.Time.String(), wantedTime.String())

	wantedTime = time.Date(2023, 11, 12, 0, 0, 0, 0, time.Local)
	err = targetTime.Parse("2023年11月12日", "2006")
	assert.Nil(t, err)
	assert.Equal(t, targetTime.Time.String(), wantedTime.String())

	err = targetTime.Parse("", "2006")
	assert.NotNil(t, err)
}
