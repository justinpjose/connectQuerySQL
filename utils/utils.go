package utils

import "time"

func GetFilenameDate() string {
	const layout = "2006-01-02-1504"
	t := time.Now()
	return t.Format(layout)
}
