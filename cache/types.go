package cache

type (
	// Lists represents a sequences of string
	Lists []string
	// Dict represent some sort of dictionary datapwd
	Dict map[string]string
)

// AsSlice represent dictionary as a slice where each element regardless of whether it
// was map key or map value turns to be slice value
func (d Dict) AsSlice() []string {
	num := len(d) * 2
	res := make([]string, num, num)

	for key, value := range d {
		res = append(res, key)
		res = append(res, value)
	}

	return res
}
