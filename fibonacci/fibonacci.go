package fibonacci

var F1 int32
var F2 int32

func Fibo(k int) (Fibo []int32) {
	F1, F2 = 0, 1
	Fibo = make([]int32, 1)
	Fibo = append(Fibo, F2)
	//fibo = append(fibo, f1, f2)
	if k > 2 {
		for i := 2; i < k; i++ {
			F1, F2 = F2, F1+F2
			Fibo = append(Fibo, F2)
		}
		return
	} else {
		return
	}
}
