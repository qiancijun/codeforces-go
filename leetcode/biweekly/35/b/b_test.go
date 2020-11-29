// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,2,3,4,5]`, `[[1,3],[0,1]]`, 
			`19`,
		},
		{
			`[1,2,3,4,5,6]`, `[[0,1]]`, 
			`11`,
		},
		{
			`[1,2,3,4,5,10]`, `[[0,2],[1,3],[1,1]]`, 
			`47`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxSumRangeQuery, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-35/problems/maximum-sum-obtained-of-any-permutation/