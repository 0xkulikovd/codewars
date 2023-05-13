// https://www.codewars.com/kata/52742f58faf5485cae000b9a

package kata

import "fmt"

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	time := [5]int64{0, 0, 0, 0, seconds}
	dur := [4]int64{365, 24, 60, 60}
	strings := [5]string{"year", "day", "hour", "minute", "second"}

	// count time
	for i := 3; i > -1; i-- {
		if time[i+1] < dur[i] {
			break
		}
		time[i] = time[i+1] / dur[i]
		time[i+1] = time[i+1] % (time[i] * dur[i])
	}

	// check non-zero values count
	count := 0
	for i := 0; i < 5; i++ {
		if time[i] != 0 {
			count++
		}
	}

	// format string
	output := ""
	for i := 0; i < 5; i++ {
		if time[i] == 0 {
			continue
		}
		name := ""
		end := ""
		if time[i] > 1 {
			name = strings[i] + "s"
		} else {
			name = strings[i]
		}

		if count > 2 {
			end = ", "
		}
		if count == 2 {
			end = " and "
		}
		count--

		output += fmt.Sprintf("%d %s%s", time[i], name, end)
	}

	return output
}
