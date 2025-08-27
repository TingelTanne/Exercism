package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    itterationCounter := 0

    if n < 1{
        return itterationCounter, errors.New("invalid Value for n")
    }
    for n > 1{
        if n % 2 == 0{
            n /= 2
        }else{
            n = n*3+1
        }
        itterationCounter++
    }
    return itterationCounter, nil
}
