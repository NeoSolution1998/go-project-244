package code

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"

	"github.com/samber/lo"
)

type DiffNode struct {
	Key    string
	Status string
	OldVal any
	NewVal any
}

const (
	StatusAdded     = "added"
	StatusRemoved   = "removed"
	StatusUpdated   = "updated"
	StatusUnchanged = "unchanged"
)

func GenDiff(path1, path2 string) (string, error) {
	data1, err := parseFile(path1)
	if err != nil {
		return "", err
	}

	data2, err := parseFile(path2)
	if err != nil {
		return "", err
	}
	node := buildDiff(data1, data2)

	res := formatStylish(node)
	return res, nil
}

func formatStylish(node []DiffNode) string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString("\n")
	for _, value := range node {
		line := ""

		switch value.Status {
		case StatusRemoved:
			line = fmt.Sprintf("  - %s: %v\n", value.Key, value.OldVal)
		case StatusAdded:
			line = fmt.Sprintf("  + %s: %v\n", value.Key, value.NewVal)
		case StatusUnchanged:
			line = fmt.Sprintf("    %s: %v\n", value.Key, value.OldVal)
		case StatusUpdated:
			line = fmt.Sprintf("  - %s: %v\n  + %s: %v\n",
				value.Key, value.OldVal,
				value.Key, value.NewVal,
			)
		}
		builder.WriteString(line)
	}

	builder.WriteString("}")
	result := builder.String()
	return result
}

func buildDiff(data1, data2 map[string]interface{}) []DiffNode {

	k1 := slices.Sorted(maps.Keys(data1))
	k2 := slices.Sorted(maps.Keys(data2))

	keys := lo.Union(k1, k2)
	slices.Sort(keys)

	node := make([]DiffNode, 0, len(keys))

	for _, key := range keys {
		val1, ok1 := data1[key]
		val2, ok2 := data2[key]

		switch {
		case ok1 && !ok2:
			node = append(node, DiffNode{
				Key:    key,
				Status: StatusRemoved,
				OldVal: val1,
			})
		case !ok1 && ok2:
			node = append(node, DiffNode{
				Key:    key,
				Status: StatusAdded,
				NewVal: val2,
			})

		case ok1 && ok2 && reflect.DeepEqual(val1, val2):
			node = append(node, DiffNode{
				Key:    key,
				Status: StatusUnchanged,
				OldVal: val1,
			})

		case ok1 && ok2 && !reflect.DeepEqual(val1, val2):
			node = append(node, DiffNode{
				Key:    key,
				Status: StatusUpdated,
				OldVal: val1,
				NewVal: val2,
			})
		}
	}
	return node
}
