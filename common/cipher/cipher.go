package cipher

// The default key
var _key = []byte("vtun@2022")

// SetKey sets the key
func SetKey(key string) {
	getSha := sha512.Sum512([]byte(key))
	_key = getSha[:]
}

// XOR encrypts the data
func XOR(src []byte) []byte {
	// Dynamic key size
	if len(src) > len(_key) {
		_klen := len(src) / 64
		for i := 0; i < _klen; i++ {
			getSha := sha512.Sum512(_key)
			_key = append(_key, getSha[:]...)
		}
	}

	for i := 0; i < len(src); i++ {
		src[i] ^= _key[i]
	}
	
	return src
}
