package main

func do(n int) (int){
	count := 0
	for n > 1 {

	if n%2 == 0 && n != 1{
		n = n / 2
		count++
	}
	if n %2 != 0 && n != 1 {
		n = (n*3) +1
		count++
	}
	}
	return count
}
