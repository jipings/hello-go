// IntSet 是一个包含非负整数的集合
// 零值代表空集合
package intset
import (
	"bytes"
	"fmt"
)

type Inset struct {
	words []int64
}

// Has 方法的返回值表示是否存在非负数 x
func (s *Inset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 添加非负数x到集合中
func (s *Inset) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words,0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith 将会对s和t做并集并将结果存在s中
func (s *IntSet) UnionWith(t *Inset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Strings 方法以字符串 "{1,2,3}"的形式返回集中
func (s *Inset) Strings() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i + j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.Strings()
}