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

func Concat(l []string, delimiter string) string {
	if len(l) == 0 {
		return ""
	}

	concatTail := Concat(l[1:], delimiter)

	if concatTail == "" {
		return l[0]
	}

	return l[0] + delimiter + Concat(l[1:], delimiter)
}

func Height[T any](n *Node[T]) int {
	if n == nil {
		return 0
	}

	return 1 + max(Height(n.l), Height(n.r))
}

type Nat struct {
	*Nat
}

func Pred(n *Nat) Opt[*Nat] {
	if n == nil {
		return Opt[*Nat]{}
	}

	return OptOfOk(n.Nat)
}

func Add(m, n *Nat) *Nat {
	if m == nil {
		return n
	}

	return &Nat{
		Nat: Add(m.Nat, n),
	}
}
