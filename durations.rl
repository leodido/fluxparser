%%{
machine durations;

include commonactions "commonactions.rl";

action ranky {
    fmt.Println(string(m.text()), "y")
    m.durationrank = 10
}

action rankmo {
    fmt.Println(string(m.text()), "mo")
    m.durationrank = 9
}

action rankw {
    fmt.Println(string(m.text()), "w")
    m.durationrank = 8
}

action rankd {
    fmt.Println(string(m.text()), "d")
    m.durationrank = 7
}

action rankh {
    fmt.Println(string(m.text()), "h")
    m.durationrank = 6
}

action rankm {
    fmt.Println(string(m.text()), "m")
    m.durationrank = 5
}

action ranks {
    fmt.Println(string(m.text()), "s")
    m.durationrank = 4
}

action rankms {
    fmt.Println(string(m.text()), "ms")
    m.durationrank = 3
}

action rankus {
    fmt.Println(string(m.text()), "us")
    m.durationrank = 2
}

action rankns {
    fmt.Println(string(m.text()), "ns")
    m.durationrank = 1
}

durationliteral = 
    start: (
        ('1'..'9' . digit*) >mark -> units
    ),
    units: (
        ('y' when { m.durationrank > 10 } %ranky) -> cont |
        ('y' when { m.durationrank > 10 } %ranky) -> final |

        ('mo' when { m.durationrank > 9 } @(emme, 2) %rankmo) -> cont |
        ('mo' when { m.durationrank > 9 } %(emme, 2) %rankmo) -> final | 

        ('w' when { m.durationrank > 8 } %rankw) -> cont | 
        ('w' when { m.durationrank > 8 } %rankw) -> final |

        ('d' when { m.durationrank > 7 } %rankd) -> cont |
        ('d' when { m.durationrank > 7 } %rankd) -> final |

        ('h' when { m.durationrank > 6 } %rankh) -> cont |
        ('h' when { m.durationrank > 6 } %rankh) -> final |

        ('m' when { m.durationrank > 5 } @(emme, 1) %rankm) -> cont |
        ('m' when { m.durationrank > 5 } %(emme, 1) %rankm) -> final |

        ('s' when { m.durationrank > 4 } %ranks) -> cont |
        ('s' when { m.durationrank > 4 } %ranks) -> final |

        ('ms' when { m.durationrank > 3 } @(emme, 2) %rankms) -> cont |
        ('ms' when { m.durationrank > 3 } %(emme, 2) %rankms) -> final |

        ((0xCE . 0xBC . 's' | 'us') when { m.durationrank > 2 } %rankus) -> cont |  
        ((0xCE . 0xBC . 's' | 'us') when { m.durationrank > 2 } %rankus) -> final |
        
        ('ns' when { m.durationrank > 1 } %rankns) -> final
    ),
    cont: (
        ('1'..'9' . digit*) >mark -> units
    );

# 1y3mo2w1d4h1m30s5ms2µs70ns
# 500ms

}%%
