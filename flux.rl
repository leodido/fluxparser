package fluxparser

import (
	"fmt"
)

%%{
machine flux;

alphtype uint8;

action mark {
	fmt.Println("mark", m.p, m.pb, m.stack)
	m.pb = m.p
}

newline = '\n' @{
    m.curline++;
};

comment = '//' (any - newline)* newline;

__ = (space | newline)*; # todo > comments ?

# Alpha-numberic characters or underscore.
alnum_ = alnum | '_';

# Alpha characters or underscore.
alpha_ = alpha | '_';

# String literal.
stringliteral = [^"\\] | !newline | ('\\"');

doublestringchar = '"' . stringliteral* . '"';

# Identifier.
identifier = alpha_ . alnum_*;

# Variable declaration. # (todo) > complete
vardecl = identifier . __ . '=' . __;

# Option statement. # (todo) > complete
optdecl = 'option' __ vardecl;

# Return statement. # (todo) > complete
retdecl = 'return' __;

blkopen = '{' . __;

blkclose = __ . '}';

# Block declaration. # (todo) > complete
blkdecl = blkopen >mark @{ fcall closingbrace; }; # fixme> in OR with? => [^{}] |

closingbrace := (vardecl | optdecl | retdecl | blkdecl)* blkclose >mark @{ fret; };

# Statement. # (todo) > complete
statement = (vardecl | optdecl | retdecl | blkdecl);

# # Equality operators.
# equalityop = ('==' | '!=' | '=~' | '!~');

# # Logical operators.
# relationalop = '<=' | '<' | '>=' | '>' | 'startswith' | 'STARTSWITH' | 'in' | 'IN' | 'not empty' | 'NOT EMPTY' | 'empty' | 'EMPTY';

# # Additivie operators.
# additiveop = '+' | '-';

# # Multiplicative operators.
# multiplicativeop = '*' | '/';

# # Unary operators.
# unaryop = '-' | 'not';

# # Logical operators.
# logicalop = 'or' | 'OR' | 'and' | 'AND';

# array = '[' __ arrayelems? __ ']';

# primary = array | literal; # (todo) > complete ... pipexpr | array | literal | ...

# arrayrest = ',' __ primary;

# arrayelems = primary __ arrayrest*;

# literal = stringliteral; # (todo) > complete ... | booleanliteral | regexpliteral

# unaryexpr = __ unaryop __ primary | primary;

# multiplicative = unaryexpr (__ multiplicativeop __ unaryexpr)*;

# additive = multiplicative (__ additiveop __ multiplicative)*;

# relational = additive (__ relationalop __ additive)*;

# equality = relational (__ equalityop __ relational)*;

# logicalexpr = equality (__ logicalop __ equality)*;

main := statement (__ statement __)*;

}%%

%%{
prepush {
	m.stack = append(m.stack, 0)
}

postpop {
	m.stack = m.stack[:len(m.stack) - 1]
}
}%%

%% write data noprefix noerror;

type machine struct {
	data         []byte
    stack        []int
    top          int
	cs           int
	p, pe, eof   int
	pb           int
    curline      int
	err          error
}

func NewMachine() *machine {
	m := &machine{}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;
    %% variable stack m.stack;
    %% variable top m.top;

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

    %% write init;
    %% write exec;

	if m.cs < first_final  {
		return false
	}

	return true
}