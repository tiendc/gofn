package gofn

func Chunk[T any](s []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return [][]T{}
	}

	chunks := make([][]T, 0, len(s)/chunkSize+1)
	for {
		if len(s) == 0 {
			break
		}
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		chunks = append(chunks, s[0:chunkSize])
		s = s[chunkSize:]
	}
	return chunks
}

func ChunkByPieces[T any](s []T, chunkCount int) [][]T {
	if chunkCount <= 0 {
		return [][]T{}
	}
	chunkSize := len(s) / chunkCount
	if chunkSize*chunkCount < len(s) {
		chunkSize++
	}

	return Chunk(s, chunkSize)
}
