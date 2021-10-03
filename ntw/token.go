package ntw

type (
	// interface is used because of the nature of "go" language
	// because inheritance is managed through composition
	iToken interface {
		kind() int
		value() string
		isEmpty() bool
		parse()
		setAnd()
		setOrdinal()
	}

	token struct {
		input   string
		pos     int
		ordinal bool
		val     string
	}

	tenToken struct {
		token
	}

	hundredToken struct {
		token
	}

	thousandToken struct {
		token
	}
)

func newTensToken(input string) *tenToken {
	t := &tenToken{
		token{
			input: input,
		},
	}
	t.parse()
	return t
}

func (t *tenToken) kind() int {
	return tens
}

// parse will parse last 2 digits of each 3 digit segment of a number
// last two digits are parsed in a group because this is most natural
// way to English language
func (t *tenToken) parse() {
	nr := t.input
	if len(nr) == 3 {
		nr = nr[1:]
	}
	if nr == "00" {
		return
	}

	if nr[0] == '0' {
		nr = nr[1:]
	}
	if len(nr) == 1 {
		m := carord(tokens.tenRight, tokens.tenRightOrdinal, t.ordinal)
		t.val = m[nr]
		return
	}

	m := carord(tokens.tenDouble, tokens.tenDoubleOrdinal, t.ordinal)
	if v, ok := m[nr]; ok {
		t.val = v
		return
	}

	if nr[1] == '0' {
		m = carord(tokens.tenLeft, tokens.tenLeftOrdinal, t.ordinal)
		t.val = m[string(nr[0])]
		return
	}

	m = carord(tokens.tenRight, tokens.tenRightOrdinal, t.ordinal)
	t.val = tokens.tenLeft[string(nr[0])] + "-" + m[string(nr[1])]
}

func newHundredsToken(input string) *hundredToken {
	t := &hundredToken{
		token{
			input: input,
		},
	}
	t.parse()
	return t
}

func (t *hundredToken) kind() int {
	return hundreds
}

// parse will parse hundred expression
// one hundred, two hundred...
func (t *hundredToken) parse() {
	if len(t.input) < 3 || t.input[0] == '0' {
		return
	}

	nr := string(t.input[0])
	if t.ordinal {
		t.val = tokens.tenRight[nr] + " " + tokens.other["100th"]
	} else {
		t.val = tokens.tenRight[nr] + " " + tokens.other["100"]
	}
}

func newThousandsToken(input string, pos int) *thousandToken {
	t := &thousandToken{
		token{
			input: input,
			pos:   pos,
		},
	}
	t.parse()
	return t
}

func (t *thousandToken) kind() int {
	return thousands
}

// parse will parse thousands expression:
// billion, million, thousand...
func (t *thousandToken) parse() {
	if t.input == "000" {
		t.val = ""
	} else if t.ordinal {
		t.val = tokens.thousandOrdinal[t.pos]
	} else {
		t.val = tokens.thousand[t.pos]
	}
}

func (t *token) value() string {
	return t.val
}

func (t *token) isEmpty() bool {
	return t.val == ""
}

func (t *token) setAnd() {
	t.val = tokens.other["and"] + " " + t.val
}

func (t *token) setOrdinal() {
	t.ordinal = true
}
