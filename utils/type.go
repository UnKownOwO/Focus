package utils

func SliceIntToUint32(data []int) []uint32 {
	result := make([]uint32, 0, len(data))
	for _, v := range data {
		result = append(result, uint32(v))
	}
	return result
}

func MapIntFloat32ToUint32Float32(data map[int]float32) map[uint32]float32 {
	result := make(map[uint32]float32)
	for k, v := range data {
		result[uint32(k)] = float32(v)
	}
	return result
}

func MapIntIntToUint32Uint32(data map[int]int) map[uint32]uint32 {
	result := make(map[uint32]uint32)
	for k, v := range data {
		result[uint32(k)] = uint32(v)
	}
	return result
}
