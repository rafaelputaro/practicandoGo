package main


func main() {
	path := string("trips_2021.csv")
	p, _ := Parser(path)
	if p != nil {
		Print(p)
	}
}
