package pflux

const start int = 1
const first_final int = 4
const error int = 0

const en_closingbrace int = 2
const en_main int = 1

type machine struct {
	data       []byte
	stack      []int
	top        int
	cs         int
	p, pe, eof int
	pb         int
	curline    int
	err        error
}

func NewMachine() *machine {
	m := &machine{}

	return m
}

func (m *machine) Parse(input []byte) bool {
	m.data = input
	m.curline = 1
	m.p = 0
	m.pb = 0
	m.top = 0
	m.stack = make([]int)
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	{
		m.cs = start
		(m.top) = 0
	}

	{
		if (m.p) == (m.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch m.cs {
		case 1:
			goto st1
		case 4:
			goto st4
		case 5:
			goto st5
		case 0:
			goto st0
		case 2:
			goto st2
		case 3:
			goto st3
		case 6:
			goto st6
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof
		}
	_resume:
		switch m.cs {
		case 1:
			goto st_case_1
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 6:
			goto st_case_6
		}
		goto st_out
	st1:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch (m.data)[(m.p)] {
		case 123:
			goto st5
		case 125:
			goto st0
		}
		goto st4
	tr9:

		m.curline++

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr9
		case 123:
			goto st5
		case 125:
			goto st0
		}
		goto st4
	tr10:
		{
			(m.stack)[(m.top)] = 5
			(m.top)++
			goto st2
		}
		goto st5
	tr11:

		m.curline++

		{
			(m.stack)[(m.top)] = 5
			(m.top)++
			goto st2
		}
		goto st5
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr11
		case 32:
			goto tr10
		case 123:
			goto st5
		case 125:
			goto st0
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto tr10
		}
		goto st4
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	tr4:

		m.curline++

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr4
		case 123:
			goto st3
		case 125:
			goto tr6
		}
		goto st2
	tr7:
		{
			(m.stack)[(m.top)] = 3
			(m.top)++
			goto st2
		}
		goto st3
	tr8:

		m.curline++

		{
			(m.stack)[(m.top)] = 3
			(m.top)++
			goto st2
		}
		goto st3
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr8
		case 32:
			goto tr7
		case 123:
			goto st3
		case 125:
			goto tr6
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto tr7
		}
		goto st2
	tr6:
		{
			(m.top)--
			m.cs = (m.stack)[(m.top)]
			goto _again
		}
		goto st6
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof6
		}
	st_case_6:
		goto st0
	st_out:
	_test_eof1:
		m.cs = 1
		goto _test_eof
	_test_eof4:
		m.cs = 4
		goto _test_eof
	_test_eof5:
		m.cs = 5
		goto _test_eof
	_test_eof2:
		m.cs = 2
		goto _test_eof
	_test_eof3:
		m.cs = 3
		goto _test_eof
	_test_eof6:
		m.cs = 6
		goto _test_eof

	_test_eof:
		{
		}
	_out:
		{
		}
	}

	if m.cs < first_final {
		return false
	}

	return true
}
