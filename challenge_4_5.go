package taofp

func longer(a, b string) string {
	if len(a) > len(b) {
		return a
	}

	return b
}

func LongestString(l []string) Opt[string] {
	if len(l) == 0 {
		return Opt[string]{}
	}

	tailLongest := LongestString(l[1:])

	if !tailLongest.ok {
		return OptOfOk(l[0])
	}

	return OptOfOk(longer(l[0], tailLongest.v))
}
