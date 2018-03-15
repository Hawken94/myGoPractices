package goExamples

func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
func deferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
func deferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
