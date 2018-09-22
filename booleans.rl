%%{
machine booleans;

action ex_booleanliteral {
    var v bool
    if string(m.text()) == "true" {
        v = true
    }
    m.expression = &ast.BooleanLiteral{
		Value: v,
	}
}

booleanliteral = ('true' | 'false') >mark %ex_booleanliteral;

}%%