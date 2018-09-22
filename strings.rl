%%{
machine strings;

include commonactions "commonactions.rl";

action ex_doublestringchars {
	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}
}

fieldstringchar = [^\n\r\\"] | '\\' [\\"];

doublestringchars = fieldstringchar* >mark %ex_doublestringchars;

stringliteral = '"' doublestringchars '"';
}%%