package validator

import "testing"

func TestIsPrintableString_Valid(t *testing.T) {
    valids := []string{
        "RAN-Node1",
        "My RAN (east)",
        "ABCxyz 0123',-./:+=?",
    }
    for _, s := range valids {
        if !IsPrintableString(s) {
            t.Fatalf("expected valid printable string: %q", s)
        }
    }
}

func TestIsPrintableString_Invalid(t *testing.T) {
    invalids := []string{
        "Bad\nName",
        "Null\x00Byte",
        "Unicode✓",
        "Emoji😀",
    }
    for _, s := range invalids {
        if IsPrintableString(s) {
            t.Fatalf("expected invalid printable string: %q", s)
        }
    }
}
