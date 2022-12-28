package muitl

func SafeEquals(input ,realPass string,maxSize int) bool {
	lenInput := len(input)
	lenRealPass := len(realPass)

	if lenInput > maxSize {
		return false
	}

	match := true
	for i := 0;i < maxSize;i++ {
		if i < lenInput && i < lenRealPass {
			if input[i] != realPass[i] {
				match = false
			}
		} else if  i < lenInput || i < lenRealPass {
			match = false
		} else {
			// pass
		}
	}
	return match
}