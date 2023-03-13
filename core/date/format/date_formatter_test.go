package format

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateFormatter(t *testing.T) {
	targetTime := time.Date(2023, 11, 12, 11, 11, 11, 11, time.Local)

	assert.Equal(t, "2023-11-12 11:11:11", targetTime.Format(DashStandard))
	assert.Equal(t, "2023-11-12", targetTime.Format(DashYearToDay))
	assert.Equal(t, "2023-11-12 Sun", targetTime.Format(DashYearToWeek))
	assert.Equal(t, "2023-11-12 Sun 11:11:11", targetTime.Format(DashYearToSecondWithWeek))
	assert.Equal(t, "12/11/2023 11:11:11", targetTime.Format(SlashStandard))
	assert.Equal(t, "12/11/2023 Sun", targetTime.Format(SlashYearToWeek))
	assert.Equal(t, "12/11/2023", targetTime.Format(SlashYearToDay))
	assert.Equal(t, "12/11/2023 Sun 11:11:11", targetTime.Format(SlashYearToSecondWithWeek))
	assert.Equal(t, "11:11:11", targetTime.Format(HourToSecond))
}
