package gofn

// Partition splits slice items into 2 lists.
func Partition[T any, S ~[]T](s S, partitionFunc func(T, int) bool) (S, S) {
	partition0, partitionRemaining := S{}, S{}
	for i, v := range s {
		if partitionFunc(v, i) {
			partition0 = append(partition0, v)
		} else {
			partitionRemaining = append(partitionRemaining, v)
		}
	}
	return partition0, partitionRemaining
}

// PartitionN splits slice items into N lists.
// partitionFunc should return index of the partition to put the corresponding item into.
func PartitionN[T any, S ~[]T](s S, numPartitions uint, partitionFunc func(T, int) int) []S {
	if numPartitions <= 0 {
		return []S{}
	}
	partitions := make([]S, numPartitions)
	for i := range partitions {
		partitions[i] = S{}
	}
	lastIndex := int(numPartitions) - 1
	for i, v := range s {
		pIndex := partitionFunc(v, i)
		if pIndex < 0 || pIndex > lastIndex {
			pIndex = lastIndex
		}
		partitions[pIndex] = append(partitions[pIndex], v)
	}
	return partitions
}
