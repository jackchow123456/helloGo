package split

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	want := []string{"我", "你"}
// 	got := Split("我爱你", "爱")
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("want:%v, got:%v\n", want, got)
// 	}
// }

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple": test{input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"en":     test{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"en_2":   test{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"china":  test{input: "小明小不小", sep: "小", want: []string{"", "明", "不", ""}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%v, want:%v, got:%v\n", name, tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的数
	for i := 0; i < b.N; i++ {
		Split("小明小不小", "小")
	}
}
