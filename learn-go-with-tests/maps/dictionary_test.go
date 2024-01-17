package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"
	assertString(t, got, want)

}

//func TestSearch1(t *testing.T) {
//	dictionary := Dictionary{"test": "this is just a test"}
//	got := dictionary.Search("test")
//	want := "this is just a test"
//	assertString(t, got, want)
//
//}

func TestSearch2(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertString(t, err.Error(), want)
	})

}
func assertString(tb testing.TB, got, want string) {
	tb.Helper()
	if got != want {
		tb.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
