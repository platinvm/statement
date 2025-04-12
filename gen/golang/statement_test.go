package statement

import (
	"image/color"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestParsePrimitives(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`true`, true},
		{`false`, false},
		{`42`, int64(42)},
		{`3.14`, float64(3.14)},
		{`0b1010`, int64(10)},
		{`0x1f`, int64(31)},
		{`"hello"`, "hello"},
		{`192.168.1.1`, net.ParseIP("192.168.1.1")},
		{`00:0a:95:9d:68:16`, mustMAC("00:0a:95:9d:68:16")},
		{`#ffcc00`, color.RGBA{255, 204, 0, 255}},
		{`"1.2.3"`, "1.2.3"}, // VERSION
		{`1h30m`, mustDuration("1h30m")},
		{`2023-12-25T12:00:00Z`, mustTime("2023-12-25T12:00:00Z")},
		{`42%`, float64(42)},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := Parse(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Parse(%q) = %#v, want %#v", test.input, got, test.expected)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	got := Parse(`[1, "two", false]`)
	want := []any{int64(1), "two", false}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseObject(t *testing.T) {
	got := Parse(`{a: 1, "b": "text", flag: true}`)
	want := map[string]any{
		"a":    int64(1),
		"b":    "text",
		"flag": true,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func mustMAC(s string) net.HardwareAddr {
	h, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}
	return h
}

func mustDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func mustTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}
