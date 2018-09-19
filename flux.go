package fluxparser

import (
	"fmt"
)



const start int = 1
const first_final int = 20

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
		case 20:
			goto st20
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
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 16:
			goto st16
		case 17:
			goto st17
		case 25:
			goto st25
		case 18:
			goto st18
		case 19:
			goto st19
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
		case 20:
			goto st_case_20
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
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 25:
			goto st_case_25
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		}
		goto st_out
	st1:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch (m.data)[(m.p)] {
		case 95:
			goto st2
		case 111:
			goto st4
		case 114:
			goto st11
		case 123:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st2
			}
		case (m.data)[(m.p)] >= 65:
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	tr34:

		fmt.Println("mark", m.p, m.pb, m.stack)
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
			goto st20
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
			goto st20
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto st3
		}
		goto st0
	tr29:

		m.curline++

		goto st20
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr29
		case 32:
			goto st20
		case 95:
			goto st2
		case 111:
			goto st4
		case 114:
			goto st11
		case 123:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st20
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	tr35:

		fmt.Println("mark", m.p, m.pb, m.stack)
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
			goto st20
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
			goto st20
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
			goto st20
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
			goto st20
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
			goto st20
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
			goto tr14
		case 32:
			goto st10
		case 61:
			goto st20
		case 95:
			goto st2
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
					goto st2
				}
			case (m.data)[(m.p)] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	tr14:

		m.curline++

		goto st10
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr14
		case 32:
			goto st10
		case 61:
			goto st20
		case 95:
			goto st2
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st10
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	tr36:

		fmt.Println("mark", m.p, m.pb, m.stack)
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
			goto st20
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
			goto st20
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
			goto st20
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
			goto st20
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
			goto st20
		case 95:
			goto st2
		case 110:
			goto st21
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
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr31
		case 32:
			goto st22
		case 61:
			goto st20
		case 95:
			goto st2
		case 111:
			goto st4
		case 114:
			goto st11
		case 123:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] < 48:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st22
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
	tr31:

		m.curline++

		goto st22
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr31
		case 32:
			goto st22
		case 61:
			goto st20
		case 95:
			goto st2
		case 111:
			goto st4
		case 114:
			goto st11
		case 123:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto st22
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	tr37:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p

		goto st23
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof23
		}
	st_case_23:
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
	tr32:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 24
				(m.top)++
				goto st16
			}
		}
		goto st24
	tr33:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p


		m.curline++

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 24
				(m.top)++
				goto st16
			}
		}
		goto st24
	tr38:
		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 24
				(m.top)++
				goto st16
			}
		}
		goto st24
	tr39:

		m.curline++

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 24
				(m.top)++
				goto st16
			}
		}
		goto st24
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr39
		case 32:
			goto tr38
		case 95:
			goto st2
		case 111:
			goto st4
		case 114:
			goto st11
		case 123:
			goto st23
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				goto tr38
			}
		case (m.data)[(m.p)] > 90:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr21
		case 32:
			goto st17
		case 123:
			goto st18
		case 125:
			goto tr23
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto st17
		}
		goto st0
	tr21:

		m.curline++

		goto st17
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr21
		case 32:
			goto st17
		case 125:
			goto tr23
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto st17
		}
		goto st0
	tr23:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p

		{
			(m.top)--
			m.cs = (m.stack)[(m.top)]
			goto _again
		}
		goto st25
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof25
		}
	st_case_25:
		goto st0
	tr26:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p

		goto st18
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr25
		case 32:
			goto tr24
		case 123:
			goto tr26
		case 125:
			goto tr23
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto tr24
		}
		goto st0
	tr24:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 19
				(m.top)++
				goto st16
			}
		}
		goto st19
	tr25:

		fmt.Println("mark", m.p, m.pb, m.stack)
		m.pb = m.p


		m.curline++

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 19
				(m.top)++
				goto st16
			}
		}
		goto st19
	tr27:
		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 19
				(m.top)++
				goto st16
			}
		}
		goto st19
	tr28:

		m.curline++

		{
			m.stack = append(m.stack, 0)
			{
				(m.stack)[(m.top)] = 19
				(m.top)++
				goto st16
			}
		}
		goto st19
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch (m.data)[(m.p)] {
		case 10:
			goto tr28
		case 32:
			goto tr27
		case 123:
			goto st18
		case 125:
			goto tr23
		}
		if 9 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
			goto tr27
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
	_test_eof20:
		m.cs = 20
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
	_test_eof16:
		m.cs = 16
		goto _test_eof
	_test_eof17:
		m.cs = 17
		goto _test_eof
	_test_eof25:
		m.cs = 25
		goto _test_eof
	_test_eof18:
		m.cs = 18
		goto _test_eof
	_test_eof19:
		m.cs = 19
		goto _test_eof

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 23:

				fmt.Println("mark", m.p, m.pb, m.stack)
				m.pb = m.p

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
