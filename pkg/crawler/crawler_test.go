package crawler

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_getBodyString(t *testing.T) {
	type args struct {
		b *html.Node
	}

	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	b, _ := html.Parse(strings.NewReader(s))

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{b: b},
			want: "Links:FooBarBaz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBodyString(tt.args.b); got != tt.want {
				t.Errorf("getBodyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
