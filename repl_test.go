package main

// import "testing"


// func TestCleanInput(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{
// 		{"  Hello, World!  ", []string{"hello,", "world!"}},
// 		{"  Go is great!  ", []string{"go", "is", "great!"}},
// 		{"  Test 123  ", []string{"test", "123"}},
// 		{"  Multiple   spaces  ", []string{"multiple", "spaces"}},
// 		{"  Special @# characters!  ", []string{"special", "@#", "characters!"}},
// 		{"  ", []string{""}}}


// 		for _, c := range cases {
// 			actual := cleanInput(c.input)

// 			for i := range actual {
// 				if actual[i] != c.expected[i] {
// 					t.Errorf("cleanInput(%q) = %v; want %v", c.input, actual, c.expected)
// 				}
// 			}
// 		}
// 	}