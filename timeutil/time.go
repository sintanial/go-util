package timeutil

import "time"

const MysqlDateFormat = "2006-01-02"
const MysqlDateTimeFormat = "2006-01-02 15:04:05"

const Day time.Duration = 24 * time.Hour
const Week time.Duration = 7 * Day
const Month time.Duration = 30 * Day
