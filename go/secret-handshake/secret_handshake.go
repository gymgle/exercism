package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake convert a decimal number to the appropriate sequence of events for a secret handshake
func Handshake(code uint) []string {
	var h []string
	var i uint

	for i = 0; i < 4; i++ {
		if 1 == (code>>i)&1 {
			h = append(h, actions[i])
		}
	}

	if code > 15 && 1 == (code>>1)&1 {
		return reverse(h)
	}

	return h
}

// reverse a []string
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
