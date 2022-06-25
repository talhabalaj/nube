package nube

import "io"

const CHUNK_SIZE = 1024

func ReadFull(rd io.Reader) []byte {
	result := []byte{}
	totalSize := 0

	for {		
		buffer := make([]byte, CHUNK_SIZE)
		n, _ :=	rd.Read(buffer)

		if n < CHUNK_SIZE {
			break
		}

		totalSize += n

		result = append(result, buffer...)
	}

	return result[:totalSize]
}
