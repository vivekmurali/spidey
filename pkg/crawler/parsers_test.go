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

func Test_absoluteURL(t *testing.T) {
	type args struct {
		u string
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{u: "https://github.com", v: "/vivekmurali"},
			want:    "https://github.com/vivekmurali",
			wantErr: false,
		},
		{
			name:    "2",
			args:    args{u: "https://github.com", v: "/vivekmurali/tempbin/issues/20"},
			want:    "https://github.com/vivekmurali/tempbin/issues/20",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := absoluteURL(tt.args.u, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("absoluteURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("absoluteURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
