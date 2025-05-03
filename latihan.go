package main
import "fmt"
const NMAX int = 100

func main() {
	var A [NMAX]int
	var left, right int
	fmt.Scan(&left, &right)
	result := selfDividingNumbers(left, right)

	// Menyalin hasil ke array A dan menampilkannya
	for i := 0; i < len(result); i++ {
		A[i] = result[i]
		fmt.Print(A[i], " ")
	}
}
func selfDividingNumbers(left, right int) []int {
	var B []int
	for i := left; i <= right; i++ {
		n := i
		isValid := true
		for n > 0 {
			d := n % 10
			if d == 0 || i%d != 0 {
				isValid = false
				break
			}
			n /= 10
		}
		if isValid {
			B = append(B, i)
		}
	}
	return B
}
