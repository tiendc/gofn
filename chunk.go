package gofn

func Chunk[T any, S ~[]T](s S, chunkSize int) []S {
	if chunkSize <= 0 {
		return []S{}
	}

	chunks := make([]S, 0, len(s)/chunkSize+1)
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

func ChunkByPieces[T any, S ~[]T](s S, chunkCount int) []S {
	if chunkCount <= 0 {
		return []S{}
	}
	chunkSize := len(s) / chunkCount
	if chunkSize*chunkCount < len(s) {
		chunkSize++
	}

	return Chunk(s, chunkSize)
}
