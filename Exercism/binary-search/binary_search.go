// Package implement binary search
package binarysearch

// SearchInts func try to find key in ordered list
func SearchInts(list []int, key int) int {
	listLen := len(list)
	if listLen == 0 {
		return -1
	}
	if listLen == 1 {
		return 0
	}
	leftBorder := 0
	rightBorder := listLen - 1

	if list[leftBorder] > key || list[rightBorder] < key {
		return -1
	}

	if list[rightBorder] == key {
		return rightBorder
	}

	if list[leftBorder] == key {
		return leftBorder
	}

	midelement := listLen / 2
	if listLen%2 == 0 {
		midelement -= 1
	}
	count := 0

	for {

		if list[midelement] == key {
			return midelement
		}
		oldmidelement := midelement
		count += 1
		if count > listLen {
			return -1
		}

		if list[midelement] < key {
			midelement += 1
			if midelement >= listLen {
				return -1
			}
			leftBorder = midelement
			if leftBorder > rightBorder {
				return -1
			}
			sliceLen := len(list[leftBorder:rightBorder])

			midelement += sliceLen / 2
			if midelement < listLen-1 && sliceLen%2 == 1 {
				midelement += 1
			}

			continue
		}

		if list[midelement] > key {
			midelement -= 1
			if midelement < 0 {
				return -1
			}
			rightBorder = midelement
			if leftBorder > rightBorder {
				return -1
			}
			sliceLen := len(list[leftBorder:rightBorder])

			midelement -= sliceLen / 2
			if midelement > 0 && sliceLen%2 == 1 {
				midelement -= 1
			}
			continue
		}
		if oldmidelement == midelement {
			return -1
		}
	}
}
