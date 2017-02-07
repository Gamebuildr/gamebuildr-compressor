package examples

import "github.com/Gamebuildr/gamebuildr-compressor/pkg/compressor"

// ZipCompressionExample shows how to use compressor to store
// files in a zip archive
func ZipCompressionExample(source string, target string) {
	zipCompression := compressor.Zip{}
	zipCompression.Encode(source, target)
}
