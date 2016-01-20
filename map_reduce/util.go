package map_reduce

import "sort"

type LessFunc func(string, string) bool

type sortableStrings struct {
	strings []string
	less    LessFunc
}

func (strings sortableStrings) Len() int {
	return len(strings.strings)
}

func (strings sortableStrings) Less(i, j int) bool {
	return strings.less(strings.strings[i], strings.strings[j])
}

func (strings sortableStrings) Swap(i, j int) {
	strings.strings[i], strings.strings[j] = strings.strings[j], strings.strings[i]
}

func sortStrings(strings []string, less LessFunc) {
	sort.Sort(sortableStrings{strings: strings, less: less})
}

func DoReduce(
	inputs []string,
	partFunc func(string) string,
	lessFunc func(string, string) bool,
	reduceFunc func([]string, string) []string,
) []string {
	// partition
	parts := make(map[string][]string)
	for _, input := range inputs {
		part := partFunc(input)
		parts[part] = append(parts[part], input)
	}

	// sort partitions
	sortedKeys := make([]string, len(parts))
	index := 0
	for key := range parts {
		sortedKeys[index] = key
		index = index + 1
	}
	sortStrings(sortedKeys, lessFunc)

	// reduce partitions
	output := []string{}
	for _, key := range sortedKeys {
		strings := parts[key]
		partOutput := []string{}
		for _, s := range strings {
			partOutput = reduceFunc(partOutput, s)
		}
		output = append(output, partOutput...)
	}
	return output
}

func DoMap(inputs []string, f func(string) (string, bool)) []string {
	outputs := []string{}
	for _, input := range inputs {
		if output, ok := f(input); ok {
			outputs = append(outputs, output)
		}
	}
	return outputs
}
