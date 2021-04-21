package incrementor

var values map[string]int = map[string]int{}

func Next(key string) int {
	values[key]++
	return values[key]
}
