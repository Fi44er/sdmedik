package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func SortByNumericRange(items []string) []string {
	type itemRange struct {
		min   float64
		max   float64
		value string
	}

	ranges := lo.Map(items, func(item string, _ int) itemRange {
		numbers := lo.Map(
			regexp.MustCompile(`(\d+\.?\d*)`).FindAllString(item, -1),
			func(s string, _ int) float64 {
				num, _ := strconv.ParseFloat(s, 64)
				return num
			},
		)

		switch len(numbers) {
		case 0:
			return itemRange{0, 0, item}
		case 1:
			return itemRange{numbers[0], numbers[0], item}
		default:
			return itemRange{
				min:   lo.Min(numbers),
				max:   lo.Max(numbers),
				value: item,
			}
		}
	})

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].min != ranges[j].min {
			return ranges[i].min < ranges[j].min
		}
		return ranges[i].max < ranges[j].max
	})

	return lo.Map(ranges, func(r itemRange, _ int) string {
		return r.value
	})
}

func InterfaceToStrSlice(input []interface{}) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

func StringSliceToInterfaceSlice(strs []string) []interface{} {
	result := make([]interface{}, len(strs))
	for i, v := range strs {
		result[i] = v
	}
	return result
}
