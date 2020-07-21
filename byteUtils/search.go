package byteUtils

func SearchByte(b byte, bytes []byte) int {
	for index, v := range bytes {
		if v == b {
			return index
		}
	}
	return -1
}

func SearchNextByte(b byte, bytes []byte, start int) int {
	if index := SearchByte(b, bytes[start:]);index != -1 {
		return index + start
	}else{
		return -1
	}
}

func SplitByByte(b byte, bytes []byte) [][]byte  {
	var bList [][]byte
	shrink := bytes
	for  {
		index := SearchByte(b, shrink)
		if index==-1{
			break
		}else{
			bList = append(bList, shrink[:index])
			shrink = shrink[index+1:]
		}
	}
	if len(shrink)>0 {
		bList = append(bList, shrink)
	}
	return bList
}