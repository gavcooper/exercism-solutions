package collatzconjecture

import "errors"

var err = errors.New("Error")

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, err 
    }
	count := 0
	for (n > 0) {
        if n == 1 {
            break 
        }
        if (n % 2 == 0) {
            n = n/2
        } else {
            n = 3*n + 1
        }
    	count = count + 1
    }
	return count, nil
}
