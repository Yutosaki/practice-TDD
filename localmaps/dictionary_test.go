package localmaps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known words", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		asserString(t, got, want)
	})
	t.Run("unknown words", func(t *testing.T){
		_, err := dictionary.Search("unknown")

		asserError(t, err, ErrNoFound)
	})
}

func asserString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want: %q, but got %q", want, got)
	}
}

func asserError(t *testing.T, got, want error){
	t.Helper()

	if got != want {
		t.Errorf("want: %q, but got %q", want, got)
	}
}

