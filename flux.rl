package fluxparser

import (
	"github.com/influxdata/flux/ast"
	// "strconv"
	// "bytes"
	// "github.com/davecgh/go-spew/spew"
	// "fmt"
)

%%{
machine flux;

include commonactions "commonactions.rl";

include integers "integers.rl";
include booleans "booleans.rl";
include durations "durations.rl";
include strings "strings.rl";

alphtype uint8;

newline = '\n' @{
	m.sol = m.p
    m.curline++;
};

comment = '//' (any - newline)* newline;

__ = (space | newline)*; # todo > comments ?

# Alpha-numberic characters or underscore.
alnum_ = alnum | '_';

# Alpha characters or underscore.
alpha_ = alpha | '_';

action ex_identifier {
	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}
}

# Identifier.
identifier = alpha_ >mark . alnum_* %ex_identifier;

# Equality operators.
equalityop = ('==' | '!=' | '=~' | '!~');

# # Logical operators.
relationalop = '<=' | '<' | '>=' | '>' | 'startswith' | 'STARTSWITH' | 'in' | 'IN' | 'not empty' | 'NOT EMPTY' | 'empty' | 'EMPTY';

# # Additivie operators.
additiveop = '+' | '-';

# # Multiplicative operators.
multiplicativeop = '*' | '/';

# # Unary operators.
unaryop = '-' | 'not';

# # Logical operators.
logicalop = 'or' | 'OR' | 'and' | 'AND';

pipeliteral = '<-';

literal = stringliteral | booleanliteral | pipeliteral | durationliteral; # (todo) > complete ...

primary = literal; # (todo) > complete ... pipexpr | array | literal | ...

# unaryexpr = __ unaryop __ primary | primary;

# multiplicative = unaryexpr (__ multiplicativeop __ unaryexpr)*;

# additive = multiplicative (__ additiveop __ multiplicative)*;

# relational = additive (__ relationalop __ additive)*;

# equality = relational (__ equalityop __ relational)*;

# logicalexpr = equality . (__ logicalop __ equality)*;

# Expression statement.
# expdecl = logicalexpr;

# fixme > we are in a rush


#todo(docmerlin) add ability to parse arguments
action ex_callexpression { 
	m.statements = append(m.statements, &ast.ExpressionStatement{
		Expression: &ast.CallExpression{
			Callee:    m.identifier,
			Arguments: nil,
			// BaseNode: base(m.text(), m.curline, m.col()),
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
}

callexpression = identifier >mark . '()' %ex_callexpression; #todo(docmerlin) add ability to parse arguments

expr = literal | callexpression;

expdecl = expr;

action ex_vardecl{
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
}

# Variable declaration. # (todo) > complete
vardecl = (identifier . __ . '=' . __ . expr) %eof(ex_vardecl);

# Option statement. # (todo) > complete
optdecl = 'option' . __ . vardecl;

# Return statement. # (todo) > complete
retdecl = 'return' . __;

blk_ini = '{' . __;

blk_end = __ . '}';

# Block declaration.
blkdecl = blk_ini @{ fcall blkbody; }; # fixme> in OR with? => [^{}] |

blkbody := (vardecl | optdecl | retdecl | blkdecl | expdecl)* . blk_end @{ fret; };

# action ex_statement {
# 	fmt.Println("ex_statement")
#	m.statements = append(m.statements, (ast.Statement)(nil))
# }

# Statement.
statement = (vardecl | optdecl | retdecl | blkdecl | expdecl);

action ex_program {
	//fmt.Println("ex_program")

	m.root = &ast.Program{
		Body:     m.statements,
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	// m.children = nil

	//fmt.Println(m.p)
	//spew.Dump(m.root)
}

# action endprogram {
# 	m.root.BaseNode.Loc.End = ast.Position{
# 		Line: m.curline,
# 		Column: ,
# 	}
# 	m.root.BaseNode.Loc.Source = m.text()
# }

program = (statement >mark (__ statement __)*) %ex_program;

main := program;

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
	sol 		 int // Position of the start of line
	err          error
	root 	     *ast.Program
	statements   []ast.Statement 	 
	identifier   *ast.Identifier
	expression   ast.Expression
	durations 	 []ast.Duration
	durationrank int
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
	m.durationrank = 11

    %% write init;
    %% write exec;

	if m.cs < first_final  {
		return nil
	}

	return m.root
}
