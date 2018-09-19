package fluxparser

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/influxdata/flux/ast"
)



const start int = 1
const first_final int = 36

const en_closingbrace int = 16
const en_main int = 1


type machine struct {
	data       []byte
	stack      []int
	top        int
	cs         int
	p, pe, eof int
	pb         int
	curline    int
	sol        int // Position of the start of line
	err        error
	root       *ast.Program
	statements []ast.Statement
}

func NewMachine() *machine {
	m := &machine{}








	return m
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

func (m *machine) col() int {
	return m.p - m.sol
}

func (m *machine) Parse(input []byte) bool {
	m.data = input
	m.curline = 1
	m.sol = 0
	m.p = 0
	m.pb = 0
	m.top = 0
	m.stack = nil
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
		case 0:
			goto st0
		case 2:
			goto st2
		case 3:
			goto st3
		case 36:
			goto st36
		case 4:
			goto st4
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 16:
			goto st16
		case 17:
			goto st17
		case 40:
			goto st40
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		}

		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof
		}
	_resume:
		switch m.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 36:
			goto st_case_36
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 40:
			goto st_case_40
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		}
		goto st_out
	st1:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch (m.data)[(m.p)] {
		case 95:
			goto tr0
		case 111:
			goto tr2
		case 114:
			goto tr3
		case 123:
			goto tr4
		}
		switch {
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr0
			}
		case (m.data)[(m.p)] >= 65:
			goto tr0
		}
		goto st0
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	tr0:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st2
	tr56:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})


		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	tr6:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st3
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto st3
		}
		goto st0
	tr54:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		goto st36
	tr55:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++


		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		goto st36
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr55
		case 32:
			goto tr54
		case 95:
			goto tr56
		case 111:
			goto tr57
		case 114:
			goto tr58
		case 123:
			goto tr59
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto tr54
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
	tr2:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st4
	tr57:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})


		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 112:
			goto st5
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 116:
			goto st6
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 105:
			goto st7
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st7:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 111:
			goto st8
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 110:
			goto st9
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr15
		case 32:
			goto st10
		case 61:
			goto st36
		case 95:
			goto tr0
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st10
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr0
				}
			case (m.data)[(m.p)] >= 65:
				goto tr0
			}
		default:
			goto st2
		}
		goto st0
	tr15:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st10
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr15
		case 32:
			goto st10
		case 61:
			goto st36
		case 95:
			goto tr0
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st10
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
	tr3:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st11
	tr58:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})


		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st11
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 101:
			goto st12
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st12:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 116:
			goto st13
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 117:
			goto st14
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 114:
			goto st15
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st15:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr6
		case 32:
			goto st3
		case 61:
			goto st36
		case 95:
			goto st2
		case 110:
			goto st37
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st3
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr61
		case 32:
			goto st38
		case 61:
			goto st36
		case 95:
			goto tr0
		case 111:
			goto tr2
		case 114:
			goto tr3
		case 123:
			goto tr62
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st38
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr0
				}
			case (m.data)[(m.p)] >= 65:
				goto tr0
			}
		default:
			goto st2
		}
		goto st0
	tr61:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st38
	st38:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr61
		case 32:
			goto st38
		case 61:
			goto st36
		case 95:
			goto tr0
		case 111:
			goto tr2
		case 114:
			goto tr3
		case 123:
			goto tr62
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st38
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
	tr4:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 39
				(m.top)++
				goto st16
			}
		}
		goto st39
	tr62:
		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 39
				(m.top)++
				goto st16
			}
		}
		goto st39
	tr59:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 39
				(m.top)++
				goto st16
			}
		}
		goto st39
	tr63:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 39
				(m.top)++
				goto st16
			}
		}
		goto st39
	st39:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 95:
			goto tr0
		case 111:
			goto tr2
		case 114:
			goto tr3
		case 123:
			goto tr62
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto tr62
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr22
		case 32:
			goto st17
		case 95:
			goto tr23
		case 111:
			goto tr24
		case 114:
			goto tr25
		case 123:
			goto tr26
		case 125:
			goto tr27
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st17
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr23
			}
		default:
			goto tr23
		}
		goto st0
	tr22:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st17
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr22
		case 32:
			goto st17
		case 125:
			goto tr27
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto st17
		}
		goto st0
	tr27:
		{
			(m.top)--
			m.cs = (m.stack)[(m.top)]
			{
				m.stack = m.stack[:len(m.stack)-1]
			}
			goto _again
		}
		goto st40
	tr38:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		{
			(m.top)--
			m.cs = (m.stack)[(m.top)]
			{
				m.stack = m.stack[:len(m.stack)-1]
			}
			goto _again
		}
		goto st40
	st40:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof40
		}
	st_case_40:
		goto st0
	tr23:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st18
	tr34:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})


		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st18
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	tr29:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st19
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto st19
		}
		goto st0
	tr32:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		goto st20
	tr33:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++


		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		goto st20
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr33
		case 32:
			goto tr32
		case 95:
			goto tr34
		case 111:
			goto tr35
		case 114:
			goto tr36
		case 123:
			goto tr37
		case 125:
			goto tr38
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto tr32
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr34
			}
		default:
			goto tr34
		}
		goto st0
	tr24:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st21
	tr35:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})


		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st21
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 112:
			goto st22
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 116:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 105:
			goto st24
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 111:
			goto st25
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 110:
			goto st26
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr45
		case 32:
			goto st27
		case 61:
			goto st20
		case 95:
			goto tr23
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st27
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr23
				}
			case (m.data)[(m.p)] >= 65:
				goto tr23
			}
		default:
			goto st18
		}
		goto st0
	tr45:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st27
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr45
		case 32:
			goto st27
		case 61:
			goto st20
		case 95:
			goto tr23
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st27
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr23
			}
		default:
			goto tr23
		}
		goto st0
	tr25:

		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st28
	tr36:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})


		fmt.Println("mark", m.pb, m.p, m.stack)
		m.pb = m.p

		goto st28
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 101:
			goto st29
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 116:
			goto st30
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 117:
			goto st31
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 114:
			goto st32
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st19
		case 61:
			goto st20
		case 95:
			goto st18
		case 110:
			goto st33
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st19
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto st18
				}
			case (m.data)[(m.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr52
		case 32:
			goto st34
		case 61:
			goto st20
		case 95:
			goto tr23
		case 111:
			goto tr24
		case 114:
			goto tr25
		case 123:
			goto tr26
		case 125:
			goto tr27
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st34
			}
		case (m.data)[(m.p)] > 57:
			switch {
			case (m.data)[(m.p)] > 90:
				if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
					goto tr23
				}
			case (m.data)[(m.p)] >= 65:
				goto tr23
			}
		default:
			goto st18
		}
		goto st0
	tr52:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		goto st34
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr52
		case 32:
			goto st34
		case 61:
			goto st20
		case 95:
			goto tr23
		case 111:
			goto tr24
		case 114:
			goto tr25
		case 123:
			goto tr26
		case 125:
			goto tr27
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st34
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr23
			}
		default:
			goto tr23
		}
		goto st0
	tr26:
		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 35
				(m.top)++
				goto st16
			}
		}
		goto st35
	tr37:

		m.statements = append(m.statement, ast.VariableDeclaration{
			Declarations: []*ast.VariableDeclarator{
				{
					ID:   nil,
					Init: nil,
				},
			},
			BaseNode: base(m.text(), m.curline, m.col()),
		})

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 35
				(m.top)++
				goto st16
			}
		}
		goto st35
	tr53:

		fmt.Println("INSIDEN")
		m.sol = m.p
		m.curline++

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 35
				(m.top)++
				goto st16
			}
		}
		goto st35
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr53
		case 32:
			goto tr26
		case 95:
			goto tr23
		case 111:
			goto tr24
		case 114:
			goto tr25
		case 123:
			goto tr26
		case 125:
			goto tr27
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto tr26
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto tr23
			}
		default:
			goto tr23
		}
		goto st0
	st_out:
	_test_eof1:
		m.cs = 1
		goto _test_eof
	_test_eof2:
		m.cs = 2
		goto _test_eof
	_test_eof3:
		m.cs = 3
		goto _test_eof
	_test_eof36:
		m.cs = 36
		goto _test_eof
	_test_eof4:
		m.cs = 4
		goto _test_eof
	_test_eof5:
		m.cs = 5
		goto _test_eof
	_test_eof6:
		m.cs = 6
		goto _test_eof
	_test_eof7:
		m.cs = 7
		goto _test_eof
	_test_eof8:
		m.cs = 8
		goto _test_eof
	_test_eof9:
		m.cs = 9
		goto _test_eof
	_test_eof10:
		m.cs = 10
		goto _test_eof
	_test_eof11:
		m.cs = 11
		goto _test_eof
	_test_eof12:
		m.cs = 12
		goto _test_eof
	_test_eof13:
		m.cs = 13
		goto _test_eof
	_test_eof14:
		m.cs = 14
		goto _test_eof
	_test_eof15:
		m.cs = 15
		goto _test_eof
	_test_eof37:
		m.cs = 37
		goto _test_eof
	_test_eof38:
		m.cs = 38
		goto _test_eof
	_test_eof39:
		m.cs = 39
		goto _test_eof
	_test_eof16:
		m.cs = 16
		goto _test_eof
	_test_eof17:
		m.cs = 17
		goto _test_eof
	_test_eof40:
		m.cs = 40
		goto _test_eof
	_test_eof18:
		m.cs = 18
		goto _test_eof
	_test_eof19:
		m.cs = 19
		goto _test_eof
	_test_eof20:
		m.cs = 20
		goto _test_eof
	_test_eof21:
		m.cs = 21
		goto _test_eof
	_test_eof22:
		m.cs = 22
		goto _test_eof
	_test_eof23:
		m.cs = 23
		goto _test_eof
	_test_eof24:
		m.cs = 24
		goto _test_eof
	_test_eof25:
		m.cs = 25
		goto _test_eof
	_test_eof26:
		m.cs = 26
		goto _test_eof
	_test_eof27:
		m.cs = 27
		goto _test_eof
	_test_eof28:
		m.cs = 28
		goto _test_eof
	_test_eof29:
		m.cs = 29
		goto _test_eof
	_test_eof30:
		m.cs = 30
		goto _test_eof
	_test_eof31:
		m.cs = 31
		goto _test_eof
	_test_eof32:
		m.cs = 32
		goto _test_eof
	_test_eof33:
		m.cs = 33
		goto _test_eof
	_test_eof34:
		m.cs = 34
		goto _test_eof
	_test_eof35:
		m.cs = 35
		goto _test_eof

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 37, 38, 39:

				fmt.Println("ex_program")

				m.root = &ast.Program{
					Body:     m.statements,
					BaseNode: base(m.text(), m.curline, m.col()),
				}

				// m.children = nil

				fmt.Println(m.p)
				spew.Dump(m.root)

			case 36:

				m.statements = append(m.statement, ast.VariableDeclaration{
					Declarations: []*ast.VariableDeclarator{
						{
							ID:   nil,
							Init: nil,
						},
					},
					BaseNode: base(m.text(), m.curline, m.col()),
				})


				fmt.Println("ex_program")

				m.root = &ast.Program{
					Body:     m.statements,
					BaseNode: base(m.text(), m.curline, m.col()),
				}

				// m.children = nil

				fmt.Println(m.p)
				spew.Dump(m.root)

			}
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
