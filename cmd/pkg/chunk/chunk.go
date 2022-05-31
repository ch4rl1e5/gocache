package chunk

type Chunk struct {
	value []byte
	chunk *Chunk
}

func AppendChunk(value []byte, chunk *Chunk) {
	if chunk.value == nil {
		chunk.value = value
		return
	}

	if chunk.chunk == nil {
		chunk.chunk = &Chunk{}
	}
	AppendChunk(value, chunk.chunk)
}

func Join(s ...[]byte) []byte {
	n := 0
	for _, v := range s {
		n += len(v)
	}

	b, i := make([]byte, n), 0
	for _, v := range s {
		i += copy(b[i:], v)
	}
	return b
}

func (c Chunk) Bytes() []byte {
	if c.chunk == nil {
		return c.value
	}

	return Join(c.value, c.chunk.Bytes())
}

func (c Chunk) String() string {
	return string(c.Bytes())
}
