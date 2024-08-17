package iteration


func Iterator(character string) string {
	var repeated string
	for i:= 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}