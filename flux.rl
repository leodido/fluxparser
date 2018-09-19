package pflux

%%{
machine flux;

alphtype uint8;

action mark {
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
vardecl = identifier __ '=' __;

# Option statement. # (todo) > complete
optdecl = 'option' __ vardecl;

# Return statement. # (todo) > complete
retdecl = 'return' __;

# Block declaration. # (todo) > complete
blkdecl = [^{}] | '{' __ @{ fcall closingbrace; };

closingbrace := blkdecl* __ '}' @{ fret; };

# Statement. # (todo) > complete
statement = (vardecl | optdecl | retdecl | blkdecl);

main := statement (__ statement __)*;

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
    m.stack = make([]int)
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