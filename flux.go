
//line /tmp/flux.rl:1
package fluxparser

import (
	"github.com/influxdata/flux/ast"
)


//line /tmp/flux.rl:183



//line /tmp/flux.rl:193



//line /tmp/flux.go:19
const start int = 1
const first_final int = 86

const en_closingbrace int = 53
const en_main int = 1


//line /tmp/flux.rl:196

type machine struct {
	data         []byte
    stack        []int
    top          int
	cs           int
	p, pe, eof   int
	pb           int
    curline      int
	sol 		 int // Position of the start of line
	err          error
	root 	     *ast.Program
	statements   []ast.Statement 	 
	identifier   *ast.Identifier
	expression   ast.Expression
}

func NewMachine() *machine {
	m := &machine{}

	
//line /tmp/flux.rl:217
	
//line /tmp/flux.rl:218
	
//line /tmp/flux.rl:219
	
//line /tmp/flux.rl:220
	
//line /tmp/flux.rl:221
    
//line /tmp/flux.rl:222
    
//line /tmp/flux.rl:223

	return m
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

func (m *machine) col() int {
	return m.p - m.sol
}

func (m *machine) Parse(input []byte) *ast.Program {
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
	m.root = nil

    
//line /tmp/flux.go:88
	{
	 m.cs = start
	( m.top) = 0
	}

//line /tmp/flux.rl:249
    
//line /tmp/flux.go:96
	{
	if ( m.p) == ( m.pe) {
		goto _test_eof
	}
	goto _resume

_again:
	switch  m.cs {
	case 1:
		goto st1
	case 0:
		goto st0
	case 2:
		goto st2
	case 3:
		goto st3
	case 86:
		goto st86
	case 4:
		goto st4
	case 5:
		goto st5
	case 6:
		goto st6
	case 87:
		goto st87
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
	case 88:
		goto st88
	case 12:
		goto st12
	case 13:
		goto st13
	case 89:
		goto st89
	case 14:
		goto st14
	case 15:
		goto st15
	case 16:
		goto st16
	case 17:
		goto st17
	case 18:
		goto st18
	case 90:
		goto st90
	case 91:
		goto st91
	case 92:
		goto st92
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
	case 93:
		goto st93
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
	case 94:
		goto st94
	case 36:
		goto st36
	case 37:
		goto st37
	case 38:
		goto st38
	case 95:
		goto st95
	case 39:
		goto st39
	case 96:
		goto st96
	case 40:
		goto st40
	case 41:
		goto st41
	case 42:
		goto st42
	case 43:
		goto st43
	case 44:
		goto st44
	case 45:
		goto st45
	case 46:
		goto st46
	case 47:
		goto st47
	case 48:
		goto st48
	case 49:
		goto st49
	case 50:
		goto st50
	case 51:
		goto st51
	case 52:
		goto st52
	case 97:
		goto st97
	case 98:
		goto st98
	case 53:
		goto st53
	case 54:
		goto st54
	case 99:
		goto st99
	case 55:
		goto st55
	case 56:
		goto st56
	case 57:
		goto st57
	case 58:
		goto st58
	case 59:
		goto st59
	case 60:
		goto st60
	case 61:
		goto st61
	case 62:
		goto st62
	case 63:
		goto st63
	case 64:
		goto st64
	case 65:
		goto st65
	case 66:
		goto st66
	case 67:
		goto st67
	case 68:
		goto st68
	case 69:
		goto st69
	case 70:
		goto st70
	case 71:
		goto st71
	case 72:
		goto st72
	case 73:
		goto st73
	case 74:
		goto st74
	case 75:
		goto st75
	case 76:
		goto st76
	case 77:
		goto st77
	case 78:
		goto st78
	case 79:
		goto st79
	case 80:
		goto st80
	case 81:
		goto st81
	case 82:
		goto st82
	case 83:
		goto st83
	case 84:
		goto st84
	case 85:
		goto st85
	}

	if ( m.p)++; ( m.p) == ( m.pe) {
		goto _test_eof
	}
_resume:
	switch  m.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 86:
		goto st_case_86
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 87:
		goto st_case_87
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
	case 88:
		goto st_case_88
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 89:
		goto st_case_89
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
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
	case 93:
		goto st_case_93
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
	case 94:
		goto st_case_94
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 95:
		goto st_case_95
	case 39:
		goto st_case_39
	case 96:
		goto st_case_96
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 99:
		goto st_case_99
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	}
	goto st_out
	st1:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr0
		case 95:
			goto tr2
		case 111:
			goto tr3
		case 114:
			goto tr4
		case 123:
			goto tr5
		}
		switch {
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 65:
			goto tr2
		}
		goto st0
st_case_0:
	st0:
		 m.cs = 0
		goto _out
tr0:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st2
	st2:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof2
		}
	st_case_2:
//line /tmp/flux.go:556
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr7
		case 92:
			goto tr8
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto tr6
tr6:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st3
	st3:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof3
		}
	st_case_3:
//line /tmp/flux.go:581
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr10
		case 92:
			goto st30
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto st3
tr7:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st86
tr10:
//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st86
	st86:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof86
		}
	st_case_86:
//line /tmp/flux.go:620
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr13
		case 32:
			goto st4
		case 34:
			goto st5
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st4
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr13:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st4
tr207:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st4
tr208:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st4
tr211:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

	goto st4
tr212:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st4
tr209:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st4
tr210:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st4
	st4:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line /tmp/flux.go:797
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr13
		case 32:
			goto st4
		case 34:
			goto st5
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st4
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr183:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st5
tr190:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

	goto st5
tr202:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st5
	st5:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof5
		}
	st_case_5:
//line /tmp/flux.go:890
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr20
		case 92:
			goto tr21
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto tr19
tr19:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st6
	st6:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof6
		}
	st_case_6:
//line /tmp/flux.go:915
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr23
		case 92:
			goto st29
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto st6
tr20:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st87
tr23:
//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st87
tr180:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st87
tr181:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st87
tr182:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st87
tr188:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

	goto st87
tr189:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st87
tr200:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st87
tr201:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st87
	st87:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof87
		}
	st_case_87:
//line /tmp/flux.go:1096
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr180
		case 32:
			goto st87
		case 34:
			goto st5
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st87
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr15:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st7
tr184:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st7
tr191:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st7
tr203:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st7
	st7:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof7
		}
	st_case_7:
//line /tmp/flux.go:1211
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
tr31:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st8
tr25:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st8
tr26:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st8
	st8:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof8
		}
	st_case_8:
//line /tmp/flux.go:1279
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr31
		case 32:
			goto st8
		case 61:
			goto st9
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st8
		}
		goto st0
tr33:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st9
tr29:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st9
	st9:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof9
		}
	st_case_9:
//line /tmp/flux.go:1314
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr33
		case 32:
			goto st9
		case 34:
			goto st10
		case 95:
			goto tr35
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st9
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr35
			}
		default:
			goto tr35
		}
		goto st0
	st10:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr37
		case 92:
			goto tr38
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto tr36
tr36:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st11
	st11:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof11
		}
	st_case_11:
//line /tmp/flux.go:1367
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr40
		case 92:
			goto st26
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto st11
tr37:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st88
tr40:
//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st88
	st88:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof88
		}
	st_case_88:
//line /tmp/flux.go:1406
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr182
		case 32:
			goto tr181
		case 34:
			goto tr183
		case 95:
			goto tr184
		case 111:
			goto tr185
		case 114:
			goto tr186
		case 123:
			goto tr187
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr181
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr184
			}
		default:
			goto tr184
		}
		goto st0
tr16:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st12
tr185:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st12
tr192:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st12
tr204:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st12
	st12:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof12
		}
	st_case_12:
//line /tmp/flux.go:1521
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 112:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
tr27:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st13
	st13:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof13
		}
	st_case_13:
//line /tmp/flux.go:1568
		if ( m.data)[( m.p)] == 41 {
			goto st89
		}
		goto st0
	st89:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr189
		case 32:
			goto tr188
		case 34:
			goto tr190
		case 95:
			goto tr191
		case 111:
			goto tr192
		case 114:
			goto tr193
		case 123:
			goto tr194
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr188
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr191
			}
		default:
			goto tr191
		}
		goto st0
tr17:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st14
tr186:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st14
tr193:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st14
tr205:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st14
	st14:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof14
		}
	st_case_14:
//line /tmp/flux.go:1692
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 101:
			goto st15
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st15:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 116:
			goto st16
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st16:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 117:
			goto st17
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st17:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 114:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st18:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 110:
			goto st90
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st90:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr196
		case 32:
			goto tr195
		case 34:
			goto st5
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr195
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto tr15
				}
			case ( m.data)[( m.p)] >= 65:
				goto tr15
			}
		default:
			goto st7
		}
		goto st0
tr198:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st91
tr195:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st91
tr196:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st91
	st91:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof91
		}
	st_case_91:
//line /tmp/flux.go:1953
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr198
		case 32:
			goto st91
		case 34:
			goto st5
		case 61:
			goto st9
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st91
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr5:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 92; ( m.top)++; goto st53 }} 
	goto st92
tr18:
//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 92; ( m.top)++; goto st53 }} 
	goto st92
tr187:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 92; ( m.top)++; goto st53 }} 
	goto st92
tr194:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 92; ( m.top)++; goto st53 }} 
	goto st92
tr199:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 92; ( m.top)++; goto st53 }} 
	goto st92
tr206:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 92; ( m.top)++; goto st53 }} 
	goto st92
	st92:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof92
		}
	st_case_92:
//line /tmp/flux.go:2089
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr199
		case 32:
			goto tr18
		case 34:
			goto st5
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr18
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
	st19:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 116:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st20:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 105:
			goto st21
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st21:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 111:
			goto st22
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st22:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto st7
		case 110:
			goto st23
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st7
				}
			case ( m.data)[( m.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st23:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr54
		case 32:
			goto tr53
		case 40:
			goto tr27
		case 61:
			goto tr29
		case 95:
			goto tr15
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr53
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto tr15
				}
			case ( m.data)[( m.p)] >= 65:
				goto tr15
			}
		default:
			goto st7
		}
		goto st0
tr56:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st24
tr53:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st24
tr54:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st24
	st24:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof24
		}
	st_case_24:
//line /tmp/flux.go:2339
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr56
		case 32:
			goto st24
		case 61:
			goto st9
		case 95:
			goto tr57
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st24
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr57
			}
		default:
			goto tr57
		}
		goto st0
tr57:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st25
	st25:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof25
		}
	st_case_25:
//line /tmp/flux.go:2375
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr26
		case 32:
			goto tr25
		case 61:
			goto tr29
		case 95:
			goto st25
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr25
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st25
				}
			case ( m.data)[( m.p)] >= 65:
				goto st25
			}
		default:
			goto st25
		}
		goto st0
tr38:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st26
	st26:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof26
		}
	st_case_26:
//line /tmp/flux.go:2416
		switch ( m.data)[( m.p)] {
		case 34:
			goto st11
		case 92:
			goto st11
		}
		goto st0
tr35:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st27
	st27:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof27
		}
	st_case_27:
//line /tmp/flux.go:2436
		switch ( m.data)[( m.p)] {
		case 40:
			goto tr59
		case 95:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st27
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto st27
			}
		default:
			goto st27
		}
		goto st0
tr59:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st28
	st28:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof28
		}
	st_case_28:
//line /tmp/flux.go:2470
		if ( m.data)[( m.p)] == 41 {
			goto st93
		}
		goto st0
	st93:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr201
		case 32:
			goto tr200
		case 34:
			goto tr202
		case 95:
			goto tr203
		case 111:
			goto tr204
		case 114:
			goto tr205
		case 123:
			goto tr206
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr200
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr203
			}
		default:
			goto tr203
		}
		goto st0
tr21:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st29
	st29:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof29
		}
	st_case_29:
//line /tmp/flux.go:2521
		switch ( m.data)[( m.p)] {
		case 34:
			goto st6
		case 92:
			goto st6
		}
		goto st0
tr8:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st30
	st30:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof30
		}
	st_case_30:
//line /tmp/flux.go:2541
		switch ( m.data)[( m.p)] {
		case 34:
			goto st3
		case 92:
			goto st3
		}
		goto st0
tr2:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st31
	st31:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof31
		}
	st_case_31:
//line /tmp/flux.go:2561
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
tr68:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st32
tr62:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st32
tr63:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st32
	st32:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line /tmp/flux.go:2629
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr68
		case 32:
			goto st32
		case 61:
			goto st33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st32
		}
		goto st0
tr70:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st33
tr66:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st33
	st33:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof33
		}
	st_case_33:
//line /tmp/flux.go:2664
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr70
		case 32:
			goto st33
		case 34:
			goto st34
		case 95:
			goto tr72
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st33
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr72
			}
		default:
			goto tr72
		}
		goto st0
	st34:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr74
		case 92:
			goto tr75
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto tr73
tr73:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st35
	st35:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof35
		}
	st_case_35:
//line /tmp/flux.go:2717
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr77
		case 92:
			goto st36
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto st35
tr74:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st94
tr77:
//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st94
	st94:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof94
		}
	st_case_94:
//line /tmp/flux.go:2756
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr208
		case 32:
			goto tr207
		case 34:
			goto tr183
		case 95:
			goto tr184
		case 111:
			goto tr185
		case 114:
			goto tr186
		case 123:
			goto tr187
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr207
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr184
			}
		default:
			goto tr184
		}
		goto st0
tr75:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st36
	st36:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof36
		}
	st_case_36:
//line /tmp/flux.go:2798
		switch ( m.data)[( m.p)] {
		case 34:
			goto st35
		case 92:
			goto st35
		}
		goto st0
tr72:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st37
	st37:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof37
		}
	st_case_37:
//line /tmp/flux.go:2818
		switch ( m.data)[( m.p)] {
		case 40:
			goto tr79
		case 95:
			goto st37
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st37
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st0
tr79:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st38
	st38:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof38
		}
	st_case_38:
//line /tmp/flux.go:2852
		if ( m.data)[( m.p)] == 41 {
			goto st95
		}
		goto st0
	st95:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr210
		case 32:
			goto tr209
		case 34:
			goto tr202
		case 95:
			goto tr203
		case 111:
			goto tr204
		case 114:
			goto tr205
		case 123:
			goto tr206
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr209
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr203
			}
		default:
			goto tr203
		}
		goto st0
tr64:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st39
	st39:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof39
		}
	st_case_39:
//line /tmp/flux.go:2905
		if ( m.data)[( m.p)] == 41 {
			goto st96
		}
		goto st0
	st96:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr212
		case 32:
			goto tr211
		case 34:
			goto tr190
		case 95:
			goto tr191
		case 111:
			goto tr192
		case 114:
			goto tr193
		case 123:
			goto tr194
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr211
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr191
			}
		default:
			goto tr191
		}
		goto st0
tr3:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st40
	st40:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof40
		}
	st_case_40:
//line /tmp/flux.go:2956
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 112:
			goto st41
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st41:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 116:
			goto st42
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st42:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 105:
			goto st43
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st43:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 111:
			goto st44
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st44:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 110:
			goto st45
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st45:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr89
		case 32:
			goto tr88
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto tr2
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr88
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto tr2
				}
			case ( m.data)[( m.p)] >= 65:
				goto tr2
			}
		default:
			goto st31
		}
		goto st0
tr91:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st46
tr88:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st46
tr89:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st46
	st46:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof46
		}
	st_case_46:
//line /tmp/flux.go:3209
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr91
		case 32:
			goto st46
		case 61:
			goto st33
		case 95:
			goto tr92
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st46
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr92
			}
		default:
			goto tr92
		}
		goto st0
tr92:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st47
	st47:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof47
		}
	st_case_47:
//line /tmp/flux.go:3245
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 61:
			goto tr66
		case 95:
			goto st47
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st47
				}
			case ( m.data)[( m.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
tr4:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st48
	st48:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof48
		}
	st_case_48:
//line /tmp/flux.go:3286
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 101:
			goto st49
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st49:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 116:
			goto st50
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st50:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 117:
			goto st51
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st51:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 114:
			goto st52
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st52:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr63
		case 32:
			goto tr62
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto st31
		case 110:
			goto st97
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr62
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st31
				}
			case ( m.data)[( m.p)] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st97:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr214
		case 32:
			goto tr213
		case 34:
			goto st5
		case 40:
			goto tr64
		case 61:
			goto tr66
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr213
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto tr15
				}
			case ( m.data)[( m.p)] >= 65:
				goto tr15
			}
		default:
			goto st31
		}
		goto st0
tr216:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st98
tr213:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st98
tr214:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st98
	st98:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof98
		}
	st_case_98:
//line /tmp/flux.go:3547
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr216
		case 32:
			goto st98
		case 34:
			goto st5
		case 61:
			goto st33
		case 95:
			goto tr15
		case 111:
			goto tr16
		case 114:
			goto tr17
		case 123:
			goto tr18
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st98
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr108:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st53
tr111:
//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st53
	st53:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof53
		}
	st_case_53:
//line /tmp/flux.go:3605
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr100
		case 32:
			goto st54
		case 34:
			goto st55
		case 95:
			goto tr102
		case 111:
			goto tr103
		case 114:
			goto tr104
		case 123:
			goto tr105
		case 125:
			goto tr106
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st54
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto st0
tr100:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st54
tr130:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st54
tr131:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st54
tr140:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

	goto st54
tr141:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st54
tr171:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st54
tr172:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st54
	st54:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof54
		}
	st_case_54:
//line /tmp/flux.go:3784
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr100
		case 32:
			goto st54
		case 125:
			goto tr106
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st54
		}
		goto st0
tr106:
//line /tmp/flux.rl:147
 {( m.top)--;  m.cs = ( m.stack)[( m.top)];{
	m.stack = m.stack[:len(m.stack) - 1]
}
goto _again } 
	goto st99
tr137:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:147
 {( m.top)--;  m.cs = ( m.stack)[( m.top)];{
	m.stack = m.stack[:len(m.stack) - 1]
}
goto _again } 
	goto st99
tr147:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:147
 {( m.top)--;  m.cs = ( m.stack)[( m.top)];{
	m.stack = m.stack[:len(m.stack) - 1]
}
goto _again } 
	goto st99
tr178:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:147
 {( m.top)--;  m.cs = ( m.stack)[( m.top)];{
	m.stack = m.stack[:len(m.stack) - 1]
}
goto _again } 
	goto st99
	st99:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof99
		}
	st_case_99:
//line /tmp/flux.go:3882
		goto st0
tr132:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st55
tr142:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

	goto st55
tr173:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

	goto st55
	st55:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof55
		}
	st_case_55:
//line /tmp/flux.go:3947
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr108
		case 92:
			goto tr109
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto tr107
tr107:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st56
	st56:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof56
		}
	st_case_56:
//line /tmp/flux.go:3972
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr111
		case 92:
			goto st57
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto st56
tr109:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st57
	st57:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof57
		}
	st_case_57:
//line /tmp/flux.go:3997
		switch ( m.data)[( m.p)] {
		case 34:
			goto st56
		case 92:
			goto st56
		}
		goto st0
tr102:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st58
tr133:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st58
tr143:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st58
tr174:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st58
	st58:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof58
		}
	st_case_58:
//line /tmp/flux.go:4090
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
tr119:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st59
tr113:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st59
tr114:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st59
	st59:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof59
		}
	st_case_59:
//line /tmp/flux.go:4158
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr119
		case 32:
			goto st59
		case 61:
			goto st60
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st59
		}
		goto st0
tr121:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st60
tr117:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st60
	st60:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof60
		}
	st_case_60:
//line /tmp/flux.go:4193
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr121
		case 32:
			goto st60
		case 34:
			goto st61
		case 95:
			goto tr123
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st60
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr123
			}
		default:
			goto tr123
		}
		goto st0
	st61:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr125
		case 92:
			goto tr126
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto tr124
tr124:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st62
	st62:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof62
		}
	st_case_62:
//line /tmp/flux.go:4246
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 34:
			goto tr128
		case 92:
			goto st82
		}
		if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
			goto st0
		}
		goto st62
tr125:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st63
tr128:
//line /tmp/flux.rl:33

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

	goto st63
	st63:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof63
		}
	st_case_63:
//line /tmp/flux.go:4285
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr131
		case 32:
			goto tr130
		case 34:
			goto tr132
		case 95:
			goto tr133
		case 111:
			goto tr134
		case 114:
			goto tr135
		case 123:
			goto tr136
		case 125:
			goto tr137
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr130
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr133
			}
		default:
			goto tr133
		}
		goto st0
tr103:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st64
tr134:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st64
tr144:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st64
tr175:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st64
	st64:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof64
		}
	st_case_64:
//line /tmp/flux.go:4402
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 112:
			goto st75
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
tr115:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st65
	st65:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof65
		}
	st_case_65:
//line /tmp/flux.go:4449
		if ( m.data)[( m.p)] == 41 {
			goto st66
		}
		goto st0
	st66:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr141
		case 32:
			goto tr140
		case 34:
			goto tr142
		case 95:
			goto tr143
		case 111:
			goto tr144
		case 114:
			goto tr145
		case 123:
			goto tr146
		case 125:
			goto tr147
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr140
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr143
			}
		default:
			goto tr143
		}
		goto st0
tr104:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st67
tr135:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st67
tr145:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st67
tr176:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st67
	st67:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof67
		}
	st_case_67:
//line /tmp/flux.go:4575
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 101:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st68:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 116:
			goto st69
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st69:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 117:
			goto st70
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st70:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 114:
			goto st71
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st71:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 110:
			goto st72
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st72:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr154
		case 32:
			goto tr153
		case 34:
			goto st55
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto tr102
		case 111:
			goto tr103
		case 114:
			goto tr104
		case 123:
			goto tr105
		case 125:
			goto tr106
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr153
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto tr102
				}
			case ( m.data)[( m.p)] >= 65:
				goto tr102
			}
		default:
			goto st58
		}
		goto st0
tr156:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st73
tr153:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st73
tr154:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st73
	st73:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof73
		}
	st_case_73:
//line /tmp/flux.go:4838
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr156
		case 32:
			goto st73
		case 34:
			goto st55
		case 61:
			goto st60
		case 95:
			goto tr102
		case 111:
			goto tr103
		case 114:
			goto tr104
		case 123:
			goto tr105
		case 125:
			goto tr106
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st73
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto st0
tr105:
//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 74; ( m.top)++; goto st53 }} 
	goto st74
tr136:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 74; ( m.top)++; goto st53 }} 
	goto st74
tr146:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 74; ( m.top)++; goto st53 }} 
	goto st74
tr157:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 74; ( m.top)++; goto st53 }} 
	goto st74
tr177:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:145
 {
	m.stack = append(m.stack, 0)
{( m.stack)[( m.top)] = 74; ( m.top)++; goto st53 }} 
	goto st74
	st74:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof74
		}
	st_case_74:
//line /tmp/flux.go:4965
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr157
		case 32:
			goto tr105
		case 34:
			goto st55
		case 95:
			goto tr102
		case 111:
			goto tr103
		case 114:
			goto tr104
		case 123:
			goto tr105
		case 125:
			goto tr106
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr105
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto st0
	st75:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 116:
			goto st76
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st76:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 105:
			goto st77
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st77:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 111:
			goto st78
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st78:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto st58
		case 110:
			goto st79
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st58
				}
			case ( m.data)[( m.p)] >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st79:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr163
		case 32:
			goto tr162
		case 40:
			goto tr115
		case 61:
			goto tr117
		case 95:
			goto tr102
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr162
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto tr102
				}
			case ( m.data)[( m.p)] >= 65:
				goto tr102
			}
		default:
			goto st58
		}
		goto st0
tr165:
//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st80
tr162:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st80
tr163:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

//line /tmp/flux.rl:17

	//fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

	goto st80
	st80:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof80
		}
	st_case_80:
//line /tmp/flux.go:5217
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr165
		case 32:
			goto st80
		case 61:
			goto st60
		case 95:
			goto tr166
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st80
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr166
			}
		default:
			goto tr166
		}
		goto st0
tr166:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st81
	st81:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof81
		}
	st_case_81:
//line /tmp/flux.go:5253
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr114
		case 32:
			goto tr113
		case 61:
			goto tr117
		case 95:
			goto st81
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr113
			}
		case ( m.data)[( m.p)] > 57:
			switch {
			case ( m.data)[( m.p)] > 90:
				if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
					goto st81
				}
			case ( m.data)[( m.p)] >= 65:
				goto st81
			}
		default:
			goto st81
		}
		goto st0
tr126:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st82
	st82:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof82
		}
	st_case_82:
//line /tmp/flux.go:5294
		switch ( m.data)[( m.p)] {
		case 34:
			goto st62
		case 92:
			goto st62
		}
		goto st0
tr123:
//line /tmp/flux.rl:12

	//fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

	goto st83
	st83:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof83
		}
	st_case_83:
//line /tmp/flux.go:5314
		switch ( m.data)[( m.p)] {
		case 40:
			goto tr168
		case 95:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st83
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto st83
			}
		default:
			goto st83
		}
		goto st0
tr168:
//line /tmp/flux.rl:48

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	goto st84
	st84:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof84
		}
	st_case_84:
//line /tmp/flux.go:5348
		if ( m.data)[( m.p)] == 41 {
			goto st85
		}
		goto st0
	st85:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof85
		}
	st_case_85:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr172
		case 32:
			goto tr171
		case 34:
			goto tr173
		case 95:
			goto tr174
		case 111:
			goto tr175
		case 114:
			goto tr176
		case 123:
			goto tr177
		case 125:
			goto tr178
		}
		switch {
		case ( m.data)[( m.p)] < 65:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr171
			}
		case ( m.data)[( m.p)] > 90:
			if 97 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 122 {
				goto tr174
			}
		default:
			goto tr174
		}
		goto st0
	st_out:
	_test_eof1:  m.cs = 1; goto _test_eof
	_test_eof2:  m.cs = 2; goto _test_eof
	_test_eof3:  m.cs = 3; goto _test_eof
	_test_eof86:  m.cs = 86; goto _test_eof
	_test_eof4:  m.cs = 4; goto _test_eof
	_test_eof5:  m.cs = 5; goto _test_eof
	_test_eof6:  m.cs = 6; goto _test_eof
	_test_eof87:  m.cs = 87; goto _test_eof
	_test_eof7:  m.cs = 7; goto _test_eof
	_test_eof8:  m.cs = 8; goto _test_eof
	_test_eof9:  m.cs = 9; goto _test_eof
	_test_eof10:  m.cs = 10; goto _test_eof
	_test_eof11:  m.cs = 11; goto _test_eof
	_test_eof88:  m.cs = 88; goto _test_eof
	_test_eof12:  m.cs = 12; goto _test_eof
	_test_eof13:  m.cs = 13; goto _test_eof
	_test_eof89:  m.cs = 89; goto _test_eof
	_test_eof14:  m.cs = 14; goto _test_eof
	_test_eof15:  m.cs = 15; goto _test_eof
	_test_eof16:  m.cs = 16; goto _test_eof
	_test_eof17:  m.cs = 17; goto _test_eof
	_test_eof18:  m.cs = 18; goto _test_eof
	_test_eof90:  m.cs = 90; goto _test_eof
	_test_eof91:  m.cs = 91; goto _test_eof
	_test_eof92:  m.cs = 92; goto _test_eof
	_test_eof19:  m.cs = 19; goto _test_eof
	_test_eof20:  m.cs = 20; goto _test_eof
	_test_eof21:  m.cs = 21; goto _test_eof
	_test_eof22:  m.cs = 22; goto _test_eof
	_test_eof23:  m.cs = 23; goto _test_eof
	_test_eof24:  m.cs = 24; goto _test_eof
	_test_eof25:  m.cs = 25; goto _test_eof
	_test_eof26:  m.cs = 26; goto _test_eof
	_test_eof27:  m.cs = 27; goto _test_eof
	_test_eof28:  m.cs = 28; goto _test_eof
	_test_eof93:  m.cs = 93; goto _test_eof
	_test_eof29:  m.cs = 29; goto _test_eof
	_test_eof30:  m.cs = 30; goto _test_eof
	_test_eof31:  m.cs = 31; goto _test_eof
	_test_eof32:  m.cs = 32; goto _test_eof
	_test_eof33:  m.cs = 33; goto _test_eof
	_test_eof34:  m.cs = 34; goto _test_eof
	_test_eof35:  m.cs = 35; goto _test_eof
	_test_eof94:  m.cs = 94; goto _test_eof
	_test_eof36:  m.cs = 36; goto _test_eof
	_test_eof37:  m.cs = 37; goto _test_eof
	_test_eof38:  m.cs = 38; goto _test_eof
	_test_eof95:  m.cs = 95; goto _test_eof
	_test_eof39:  m.cs = 39; goto _test_eof
	_test_eof96:  m.cs = 96; goto _test_eof
	_test_eof40:  m.cs = 40; goto _test_eof
	_test_eof41:  m.cs = 41; goto _test_eof
	_test_eof42:  m.cs = 42; goto _test_eof
	_test_eof43:  m.cs = 43; goto _test_eof
	_test_eof44:  m.cs = 44; goto _test_eof
	_test_eof45:  m.cs = 45; goto _test_eof
	_test_eof46:  m.cs = 46; goto _test_eof
	_test_eof47:  m.cs = 47; goto _test_eof
	_test_eof48:  m.cs = 48; goto _test_eof
	_test_eof49:  m.cs = 49; goto _test_eof
	_test_eof50:  m.cs = 50; goto _test_eof
	_test_eof51:  m.cs = 51; goto _test_eof
	_test_eof52:  m.cs = 52; goto _test_eof
	_test_eof97:  m.cs = 97; goto _test_eof
	_test_eof98:  m.cs = 98; goto _test_eof
	_test_eof53:  m.cs = 53; goto _test_eof
	_test_eof54:  m.cs = 54; goto _test_eof
	_test_eof99:  m.cs = 99; goto _test_eof
	_test_eof55:  m.cs = 55; goto _test_eof
	_test_eof56:  m.cs = 56; goto _test_eof
	_test_eof57:  m.cs = 57; goto _test_eof
	_test_eof58:  m.cs = 58; goto _test_eof
	_test_eof59:  m.cs = 59; goto _test_eof
	_test_eof60:  m.cs = 60; goto _test_eof
	_test_eof61:  m.cs = 61; goto _test_eof
	_test_eof62:  m.cs = 62; goto _test_eof
	_test_eof63:  m.cs = 63; goto _test_eof
	_test_eof64:  m.cs = 64; goto _test_eof
	_test_eof65:  m.cs = 65; goto _test_eof
	_test_eof66:  m.cs = 66; goto _test_eof
	_test_eof67:  m.cs = 67; goto _test_eof
	_test_eof68:  m.cs = 68; goto _test_eof
	_test_eof69:  m.cs = 69; goto _test_eof
	_test_eof70:  m.cs = 70; goto _test_eof
	_test_eof71:  m.cs = 71; goto _test_eof
	_test_eof72:  m.cs = 72; goto _test_eof
	_test_eof73:  m.cs = 73; goto _test_eof
	_test_eof74:  m.cs = 74; goto _test_eof
	_test_eof75:  m.cs = 75; goto _test_eof
	_test_eof76:  m.cs = 76; goto _test_eof
	_test_eof77:  m.cs = 77; goto _test_eof
	_test_eof78:  m.cs = 78; goto _test_eof
	_test_eof79:  m.cs = 79; goto _test_eof
	_test_eof80:  m.cs = 80; goto _test_eof
	_test_eof81:  m.cs = 81; goto _test_eof
	_test_eof82:  m.cs = 82; goto _test_eof
	_test_eof83:  m.cs = 83; goto _test_eof
	_test_eof84:  m.cs = 84; goto _test_eof
	_test_eof85:  m.cs = 85; goto _test_eof

	_test_eof: {}
	if ( m.p) == ( m.eof) {
		switch  m.cs {
		case 86, 87, 90, 91, 92, 97, 98:
//line /tmp/flux.rl:157

	//fmt.Println("ex_program")

	m.root = &ast.Program{
		Body:     m.statements,
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	// m.children = nil

	//fmt.Println(m.p)
	//spew.Dump(m.root)

		case 89, 96:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:157

	//fmt.Println("ex_program")

	m.root = &ast.Program{
		Body:     m.statements,
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	// m.children = nil

	//fmt.Println(m.p)
	//spew.Dump(m.root)

		case 88, 94:
//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:157

	//fmt.Println("ex_program")

	m.root = &ast.Program{
		Body:     m.statements,
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	// m.children = nil

	//fmt.Println(m.p)
	//spew.Dump(m.root)

		case 93, 95:
//line /tmp/flux.rl:99
 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil

//line /tmp/flux.rl:117

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

//line /tmp/flux.rl:157

	//fmt.Println("ex_program")

	m.root = &ast.Program{
		Body:     m.statements,
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	// m.children = nil

	//fmt.Println(m.p)
	//spew.Dump(m.root)

//line /tmp/flux.go:5605
		}
	}

	_out: {}
	}

//line /tmp/flux.rl:250

	if m.cs < first_final  {
		return nil
	}

	return m.root
}
