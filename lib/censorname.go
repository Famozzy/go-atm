package lib

func CensorName(name string) string {
	var censoredName string
	for i := 0; i < len(name); i++ {
		if i == 0 || i == len(name)-1 {
			censoredName += string(name[i])
		} else {
			censoredName += "*"
		}
	}
	return censoredName
}
