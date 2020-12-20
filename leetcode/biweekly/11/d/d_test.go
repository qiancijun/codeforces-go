// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[1,2,3,4,5,6,7,8,9]`, `5`, 
			`6`,
		},
		{
			`[5,6,7,8,9,1,2,3,4]`, `8`, 
			`1`,
		},
		{
			`[1,2,2,1,2,2,1,2,2]`, `2`, 
			`5`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, maximizeSweetness, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-11/problems/divide-chocolate/