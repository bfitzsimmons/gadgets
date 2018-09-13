package gadgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	t.Parallel()

	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{"name"},
			want: "name",
		},
		{
			name: "two words",
			args: args{"one two"},
			want: "one-two",
		},
		{
			name: "three words",
			args: args{"one two three"},
			want: "one-two-three",
		},
		{
			name: "basic special characters",
			args: args{"it's a special character"},
			want: "its-a-special-character",
		},
		{
			name: "underscores",
			args: args{"we_need_hyphens"},
			want: "we-need-hyphens",
		},
		{
			name: "punctuation",
			args: args{"This is a sentence."},
			want: "this-is-a-sentence",
		},
		{
			name: "more punctuation",
			args: args{"This is a sentence?"},
			want: "this-is-a-sentence",
		},
		{
			name: "special character prefix and suffix",
			args: args{"¿Are you kidding?"},
			want: "are-you-kidding",
		},
		{
			name: "begin and end with spaces",
			args: args{" name "},
			want: "name",
		},
		{
			name: "crazy",
			args: args{"I'm so excited, and I just can't hide it. I'm about to lose control, and I think I like it! -- The Pointer Sisters (1982) "},
			want: "im-so-excited-and-i-just-cant-hide-it-im-about-to-lose-control-and-i-think-i-like-it-the-pointer-sisters-1982",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Slugify(tt.args.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
