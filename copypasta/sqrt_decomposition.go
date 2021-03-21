package copypasta

import (
	"math"
	"sort"
)

/* 一种技巧：组合两种算法从而降低复杂度 O(n^2) -> O(n√n)
参考 Competitive Programmer’s Handbook Ch.27

有 n 个对象，每个对象有一个「关于其他对象的统计量」ci（一个数、一个集合的元素个数，等等）
为方便起见，假设 ∑ci 的数量级和 n 一样，下面用 n 表示 ∑ci
当 ci > √n 时，这样的对象不超过 √n 个，暴力枚举这些对象之间的关系（或者，该对象与其他所有对象的关系），时间复杂度为 O(n) 或 O(n√n)。此乃算法一
当 ci ≤ √n 时，这样的对象有 O(n) 个，由于统计量 ci 很小，暴力枚举当前对象的统计量，时间复杂度为 O(n√n)。此乃算法二
这样，以 √n 为界，我们将所有对象划分成了两组，并用两个不同的算法处理
这两种算法是看待同一个问题的两种不同方式，通过恰当地组合这两个算法，复杂度由 O(n^2) 降至 O(n√n)
注意：**枚举时要做到不重不漏**

另一种题型是注意到 n 的整数分拆中，不同数字的个数至多有 O(√n) 种

好题 LCP16 https://leetcode-cn.com/problems/you-le-yuan-de-you-lan-ji-hua/
*/

/*
根号算法（分块）Sqrt Decomposition
https://oi-wiki.org/ds/decompose/
https://oi-wiki.org/ds/block-array/
浅谈基础根号算法——分块 https://www.luogu.com.cn/blog/deco/qian-tan-ji-chu-gen-hao-suan-fa-fen-kuai
TODO: 台湾的《根號算法》https://www.csie.ntu.edu.tw/~sprout/algo2018/ppt_pdf/root_methods.pdf

题目推荐 https://cp-algorithms.com/data_structures/sqrt_decomposition.html#toc-tgt-8
好题 https://codeforces.com/problemset/problem/91/E
todo 动态逆序对 https://www.luogu.com.cn/problem/P3157 https://www.luogu.com.cn/problem/UVA11990
https://cp-algorithms.com/sequences/rmq.html
*/
func sqrtDecompositionCollections() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	type block struct {
		l, r           int // [l,r]
		origin, sorted []int
		//lazyAdd int
	}
	var blocks []block
	sqrtInit := func(a []int) {
		n := len(a)
		blockSize := int(math.Sqrt(float64(n)))
		//blockSize := int(math.Sqrt(float64(n) * math.Log2(float64(n+1))))
		blockNum := (n-1)/blockSize + 1
		blocks = make([]block, blockNum)
		for i, v := range a {
			j := i / blockSize
			if i%blockSize == 0 {
				blocks[j] = block{l: i, origin: make([]int, 0, blockSize)}
			}
			blocks[j].origin = append(blocks[j].origin, v)
		}
		for i := range blocks {
			b := &blocks[i]
			b.r = b.l + len(b.origin) - 1
			b.sorted = append([]int(nil), b.origin...)
			sort.Ints(b.sorted)
		}
	}
	sqrtOp := func(l, r int, v int) { // [l,r], starts at 0
		for i := range blocks {
			b := &blocks[i]
			if b.r < l {
				continue
			}
			if b.l > r {
				break
			}
			if l <= b.l && b.r <= r {
				// do op on full block
			} else {
				// do op on part block
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl - b.l; j <= br-b.l; j++ {
					// do b.origin[j]...
				}
			}
		}
	}

	_ = []interface{}{sqrtInit, sqrtOp}
}