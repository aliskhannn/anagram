package anagram

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name       string
		dictionary []string
		want       map[string][]string
	}{
		{
			name:       "simple anagrams",
			dictionary: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:       "case insensitivity",
			dictionary: []string{"Пятак", "пятка", "ТЯПКА"},
			want: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:       "duplicates",
			dictionary: []string{"пятак", "пятак", "пятка"},
			want: map[string][]string{
				"пятак": {"пятак", "пятка"},
			},
		},
		{
			name:       "no anagrams",
			dictionary: []string{"стол", "дом", "камень"},
			want:       map[string][]string{},
		},
		{
			name:       "latin anagrams",
			dictionary: []string{"listen", "silent", "enlist", "inlets", "google", "gooogle"},
			want: map[string][]string{
				"enlist": {"enlist", "inlets", "listen", "silent"},
			},
		},
		{
			name:       "empty dictionary",
			dictionary: []string{},
			want:       map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAnagrams(tt.dictionary)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abc", "abc"},
		{"cba", "abc"},
		{"пятак", "акптя"},
		{"тяпка", "акптя"},
		{"", ""},
		{"a", "a"},
		{"абвгд", "абвгд"},
		{"гдба", "абгд"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := sortString(tt.input)
			if got != tt.want {
				t.Errorf("sortString(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
