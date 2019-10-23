package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		var res int = 1
		for i := 1; i <= nb; i++ {
			res *= i
		}
		return res
	}
}
