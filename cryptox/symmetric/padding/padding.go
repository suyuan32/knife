package padding

// PaddingType is the padding method such as PKCS7.
type PaddingType uint8

const (
	// ISO97971 add padding at the end of byte slice with zero bytes which are separated by 0x80.
	ISO97971 PaddingType = 1 + iota
	// No add no padding.
	No
	// PKCS5 is similar to PKCS7, but its block size is 8.
	PKCS5
	// PKCS7 padding is a generalization of PKCS5 padding (also known as standard padding). PKCS7 padding works by
	// appending N bytes with the value of chr(N) , where N is the number of bytes required to make the final block of
	// data the same size as the block size.
	PKCS7
	// Zero add padding at the end of byte slice with byte 0.
	Zero
)

// Padding pads data with method provided such as Zero Padding.
func Padding(data []byte, method PaddingType, blockSize int) ([]byte, error) {
	switch method {
	case Zero:
		return PaddingZero(data, blockSize), nil
	case PKCS5:
		return PaddingPKCS5(data), nil
	case PKCS7:
		return PaddingPKCS7(data, blockSize), nil
	case ISO97971:
		return PaddingISO97971(data, blockSize), nil
	case No:
		return data, nil
	}
	return data, nil
}

// DePadding depads data with method provided such as Zero Padding.
func DePadding(data []byte, method PaddingType, blockSize int) ([]byte, error) {
	switch method {
	case Zero:
		return DePaddingZero(data), nil
	case PKCS5:
		return DePaddingPKCS5(data), nil
	case PKCS7:
		return DePaddingPKCS7(data), nil
	case ISO97971:
		return DePaddingISO97971(data), nil
	case No:
		return data, nil
	}
	return data, nil
}
