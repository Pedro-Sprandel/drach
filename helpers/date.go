package helpers

import "time"

func CurrentMonth() int {
	return int(time.Now().Month())
}

func CurrentYear() int {
	return int(time.Now().Year())
}

func MonthName(m int) string {
	names := map[int]string{
		1:  "Jan",
		2:  "Fev",
		3:  "Mar",
		4:  "Apr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	return names[m]
}
