package main

func main() {
	s := []string{"A", "B", "C"}

	f := s[1:2]

	_ = cap(f)
	return
}