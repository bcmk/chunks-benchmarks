package chunks

func chunks(xs []string, chunkSize int) [][]string {
	divided := make([][]string, (len(xs)+chunkSize-1)/chunkSize)
	i := chunkSize
	j := 0
	for ; i < len(xs); i += chunkSize {
		divided[j] = xs[i-chunkSize : i]
		j++
	}
	if i-chunkSize < len(xs) {
		divided[j] = xs[i-chunkSize : len(xs)]
	}
	return divided
}

func chunks2(logs []string, chunkSize int) [][]string {
	var divided [][]string
	for i := 0; i < len(logs); i += chunkSize {
		end := i + chunkSize
		if end > len(logs) {
			end = len(logs)
		}
		divided = append(divided, logs[i:end])
	}
	return divided
}
