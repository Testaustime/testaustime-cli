package datahelper

import "testing"

func TestRawTimeToHumanReadable(t *testing.T) {
	testTime := func(time int, expectedOutput string) {
		timeInHumanReadableFormat := rawTimeToHumanReadable(float32(time))
		if timeInHumanReadableFormat != expectedOutput {
			t.Fatalf(
				"Failed to get time in human readable format: %d, expected %s got %s",
				time,
				expectedOutput,
				timeInHumanReadableFormat)
		}
	}

	testTime(0, "0m")
	testTime(50, "50m")
	testTime(120, "2h, 0m")
	testTime(1500, "1d, 1h, 0m")
}
