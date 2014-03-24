package util

import (
	"fmt"
	"github.com/Shopify/go-lua"
	"io"
)

func DumpStack(l *lua.State, w io.Writer, prefix string) (n int, err error) {
	n, err = fmt.Fprintf(w, "%s: top=%d, 'real (pseudo): val':\n", prefix, lua.Top(l))
	if err != nil {
		return
	}
	var m int
	for i, pseudo := 1, 0-lua.Top(l); i <= lua.Top(l); {
		m, err = fmt.Fprintf(w, "\t %d (%d): %#v\n", i, pseudo, lua.ToValue(l, i))
		n += m
		if err != nil {
			return
		}

		i++
		pseudo++
	}
	return
}
