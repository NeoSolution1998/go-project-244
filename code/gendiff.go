package code

func GenDiff(path1, path2 string) (map[string]interface{}, map[string]interface{}, error) {
	data1, err := parseFile(path1)
	if err != nil {
		return nil, nil, err
	}

	data2, err := parseFile(path2)
	if err != nil {
		return nil, nil, err
	}

	return data1, data2, nil
}
