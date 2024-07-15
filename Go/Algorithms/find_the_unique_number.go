package kata

func FindUniq(arr []float32) float32 {
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] == arr[i+1] && arr[i] != arr[i+2] {
			return arr[i+2]
		} else if arr[i] == arr[i+2] && arr[i] != arr[i+1] {
			return arr[i+1]
		} else if arr[i+1] == arr[i+2] && arr[i] != arr[i+1] {
			return arr[i]
		}
	}
	return arr[0]
}
