package frontmatter

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func ExampleParse() {
	doc := `+++
key1 = first key
key2 = 2ndKey 
+++
# My Document
...
`
	b := &bytes.Buffer{}
	fm, _ := Parse("+++", strings.NewReader(doc), b)
	fmt.Printf("%#v\n", fm)
	fmt.Println(b)
	// Output: map[string]string{"key1":"first key", "key2":"2ndKey"}
	// # My Document
	// ...
}

func Test_splitLine(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name    string
		args    args
		wantKey string
		wantVal string
	}{
		{
			"simple",
			args{"key1=val1"},
			"key1",
			"val1",
		},
		{
			"spaces",
			args{"key1  =  val1"},
			"key1",
			"val1",
		},
		{
			"just key",
			args{"key1"},
			"key1",
			"",
		},
		{
			"empty line",
			args{" "},
			"",
			"",
		},
		{
			"text",
			args{"key1  =  val1 and some text"},
			"key1",
			"val1 and some text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotVal := splitLine(tt.args.l)
			if gotKey != tt.wantKey {
				t.Errorf("splitLine() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotVal != tt.wantVal {
				t.Errorf("splitLine() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}
