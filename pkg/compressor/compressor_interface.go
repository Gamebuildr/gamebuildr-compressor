package compressor

// Compression is the type of compression to use
// when moving data to a source destination
type Compression interface {
	Encode(source string, target string) error
}
