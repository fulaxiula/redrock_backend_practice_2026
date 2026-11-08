package visibility

const adultAge = 18

func stage(age int) string {
	if age >= adultAge {
		return "adult"
	}
	return "minor"
}