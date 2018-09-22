%%{
machine durations;

include commonactions "commonactions.rl";

action ranky {
    m.durations = append(m.durations, getDuration(m.text(), "y"))
    m.durationrank = 10
}

action rankmo {
    m.durations = append(m.durations, getDuration(m.text(), "mo"))
    m.durationrank = 9
}

action rankw {
    m.durations = append(m.durations, getDuration(m.text(), "w"))
    m.durationrank = 8
}

action rankd {
    m.durations = append(m.durations, getDuration(m.text(), "d"))    
    m.durationrank = 7
}

action rankh {
    m.durations = append(m.durations, getDuration(m.text(), "h"))    
    m.durationrank = 6
}

action rankm {
    m.durations = append(m.durations, getDuration(m.text(), "m"))    
    m.durationrank = 5
}

action ranks {
    m.durations = append(m.durations, getDuration(m.text(), "s"))    
    m.durationrank = 4
}

action rankms {
    m.durations = append(m.durations, getDuration(m.text(), "ms"))    
    m.durationrank = 3
}

action rankus {
    m.durations = append(m.durations, getDuration(m.text(), "us"))    
    m.durationrank = 2
}

action rankus2 {
    m.durations = append(m.durations, getDuration(m.text(), "μs"))    
    m.durationrank = 2
}

action rankns {
    m.durations = append(m.durations, getDuration(m.text(), "ns"))    
    m.durationrank = 1
}

action ex_durationsliteral {
    // reset duration rank
    m.durationrank = 11
    //
    m.expression = &ast.DurationLiteral{
        Values: m.durations,
    }
    // (todo) > reset m.durations
    fmt.Println(m.durations)
}

non0digit = '1'..'9';

duration = 
    start: (
        (non0digit . digit*) >mark -> units
    ),
    units: (
        ('y' when { m.durationrank > 10 } %ranky) -> again |
        ('y' when { m.durationrank > 10 } %ranky) -> final |

        ('mo' when { m.durationrank > 9 } @(mp, 2) %rankmo) -> again |
        ('mo' when { m.durationrank > 9 } %(mp, 2) %rankmo) -> final |

        ('w' when { m.durationrank > 8 } %rankw) -> again |
        ('w' when { m.durationrank > 8 } %rankw) -> final |

        ('d' when { m.durationrank > 7 } %rankd) -> again |
        ('d' when { m.durationrank > 7 } %rankd) -> final |

        ('h' when { m.durationrank > 6 } %rankh) -> again |
        ('h' when { m.durationrank > 6 } %rankh) -> final |

        ('m' when { m.durationrank > 5 } @(mp, 1) %rankm) -> again |
        ('m' when { m.durationrank > 5 } %(mp, 1) %rankm) -> final |

        ('s' when { m.durationrank > 4 } %ranks) -> again |
        ('s' when { m.durationrank > 4 } %ranks) -> final |

        ('ms' when { m.durationrank > 3 } @(mp, 2) %rankms) -> again |
        ('ms' when { m.durationrank > 3 } %(mp, 2) %rankms) -> final |

        ('us' when { m.durationrank > 2 } %rankus) -> again |  
        ('us' when { m.durationrank > 2 } %rankus) -> final |

        ((0xCE . 0xBC . 's') when { m.durationrank > 2 } %rankus2) -> again |  
        ((0xCE . 0xBC . 's') when { m.durationrank > 2 } %rankus2) -> final |
        
        ('ns' when { m.durationrank > 1 } %rankns) -> final
    ),
    again: (
        (non0digit . digit*) >mark -> units
    );

# (fixme)
# action ex_integerliteral {
#     // (todo) > handle errors
#     vi, _ := strconv.Atoi(string(m.text()))
#     m.expression = &ast.IntegerLiteral{
#         Value: int64(vi),
#     }
#     fmt.Println("ex_int", m.expression)
# }
# 
# action ex_floatliteral {
#     // (todo) > handle errors
#     vf, _ := strconv.ParseFloat(string(m.text()), 64)
#     m.expression = &ast.FloatLiteral{
#         Value: vf,
#     }
#     fmt.Println("ex_float", m.expression)
# }
# floatliteral = (('0' | non0digit . digit*) . '.' digit+) >mark %eof(ex_floatliteral);
# integerliteral = ('0' | non0digit . digit*) >mark %eof(ex_integerliteral);

durationliteral = duration %eof(ex_durationsliteral);
}%%
