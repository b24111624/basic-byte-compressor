package compressor

type Store interface {
	Read([]byte) (int, error)
}
