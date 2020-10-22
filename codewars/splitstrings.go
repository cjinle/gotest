package codewars


func SplitStrings(str string) []string {
	out := []string{}
	length := len(str)
	if length == 0 {
		return out
	} 
	for i := 0; i < length; i += 2 {
		if i + 2 > length {
			out = append(out, str[i:]+"_")
		} else {
			out = append(out, str[i:i+2])
		}
	}
	return out
}

func SplitStrings2(str string) []string {
	out := []string{}
	length := len(str)
	if length % 2 == 1 {
		str += "_"
	}

	for i := 0; i < length; i += 2 {
		out = append(out, str[i:i+2])
	}
	return out
}