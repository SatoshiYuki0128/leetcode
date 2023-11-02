package easy

func CanPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)

	if n == 0 {
		return true
	}

	if size < 3 {
		switch size {
		case 0:
			return false
		case 1:
			if flowerbed[0] == 0 {
				n--
			}
		case 2:
			if flowerbed[0] == 0 && flowerbed[1] == 0 {
				n--
			}
		}

		return n == 0
	}

	for i := range flowerbed {
		if flowerbed[i] == 1 {
			continue
		}

		if i == 0 {
			if flowerbed[i+1] != 1 {
				flowerbed[i] = 1
				n--
			}
		} else if i == size-1 {
			if flowerbed[i-1] != 1 {
				flowerbed[i] = 1
				n--
			}
		} else {
			if flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
				flowerbed[i] = 1
				n--
			}
		}

		//fmt.Println(flowerbed)

		if n == 0 {
			return true
		}
	}

	return false
}
