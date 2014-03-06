package trie

import (
	"testing"
)

func TestTrieAdd(t *testing.T) {
	trie := CreateTrie()

	i := trie.AddKey("foo")

	if i != 3 {
		t.Errorf("Expected 3, got: %d", i)
	}
}

func TestTrieKeys(t *testing.T) {
	trie := CreateTrie()
	expected := []string{"bar", "foo"}

	for _, key := range expected {
		trie.AddKey(key)
	}

	kl := len(trie.Keys())
	if kl != 2 {
		t.Errorf("Expected 2 keys, got %d, keys were: %v", kl, trie.Keys())
	}

	for i, key := range trie.Keys() {
		if key != expected[i] {
			t.Errorf("Expected %#v, got %#v", expected[i], key)
		}
	}
}

func TestKeysWithPrefix(t *testing.T) {
	trie := CreateTrie()
	expected := []string{"foosball", "football", "foreboding", "forementioned", "foretold", "foreverandeverandeverandever", "forbidden"}
	defer func() {
		r := recover()
		if r != nil {
			t.Error(r)
		}
	}()

	trie.AddKey("bar")
	for _, key := range expected {
		trie.AddKey(key)
	}

	tests := []struct {
		pre      string
		expected []string
		length   int
	}{
		{"fo", expected, len(expected)},
		{"foosbal", []string{"foosball"}, 1},
		{"abc", []string{}, 0},
	}

	for _, test := range tests {
		actual := trie.KeysWithPrefix(test.pre)
		if len(actual) != test.length {
			t.Errorf("Expected len(actual) to == %d for pre %s", test.length, test.pre)
		}

		for i, key := range actual {
			if key != test.expected[i] {
				t.Errorf("Expected %v got: %v", test.expected[i], key)
			}
		}
	}

	trie.KeysWithPrefix("fsfsdfasdf")
}

func BenchmarkTieKeys(b *testing.B) {
	trie := CreateTrie()
	keys := []string{"bar", "foo", "baz", "bur", "zum", "burzum", "bark", "barcelona", "football", "foosball", "footlocker"}

	for _, key := range keys {
		trie.AddKey(key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Keys()
	}
}

func BenchmarkKeysWithPrefix(b *testing.B) {
	trie := CreateTrie()
	expected := []string{
		"foosball",
		"football",
		"foreboding",
		"forementioned",
		"foretold",
		"foreverandeverandeverandever",
		"forbidden",
		"forsupercalifragilisticexpyaladocious",
		"forsupercalifragilisticexpyaladocious",
		"forsupercalifragilisticexpyaladocious/fors",
		"forsupercalifragilisticexpyvlsdocious/fors",
		"fofsupercrlifralilisticexpyaladocgous",
		"foo/bar/baz/ber/her/mer/fur/a.out",
		"foo/baz/baz/ber/her/mer/fur/a.out",
		"foo/baz/bur/ber/her/mer/fur/a.out",
		"foo/baz/bur/sher/her/mer/fur/a.out",
		"foo/curr/bur/sher/her/mer/fur/a.out",
		"foo/lurr/bur/sher/her/mer/fur/a.out",
		"foo/turr/bur/sher/her/mer/fur/a.out",
		"foors",
	}

	trie.AddKey("bar")
	for _, key := range expected {
		trie.AddKey(key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = trie.KeysWithPrefix("fo")
	}
}
