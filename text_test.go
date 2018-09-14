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
			name: "empty",
			args: args{},
			want: "",
		},
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
			want: "it-s-a-special-character",
		},
		{
			name: "multiple special characters",
			args: args{":;^&%$#@|[]{}' That's a *lot* of \"special\" characters!"},
			want: "that-s-a-lot-of-special-characters",
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
			name: "single quotes",
			args: args{"We're in here, and they're out there."},
			want: "we-re-in-here-and-they-re-out-there",
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
			name: "slashes",
			args: args{"this/that/the other"},
			want: "this-that-the-other",
		},
		{
			name: "pipes",
			args: args{"this|that|the other"},
			want: "this-that-the-other",
		},
		{
			name: "fun",
			args: args{"I'm so excited, and I just can't hide it. I'm about to lose control, and I think I like it! -- The Pointer Sisters (1982) "},
			want: "i-m-so-excited-and-i-just-can-t-hide-it-i-m-about-to-lose-control-and-i-think-i-like-it-the-pointer-sisters-1982",
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
