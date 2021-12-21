package service

import (
	"testing"

	"github.com/tim5wang/selfman/common/util"
)

var (
	users = map[string]int{ // 乙, 甲, 丙, 丁
		"甲": 2,
		"乙": 1,
		"丙": 3,
		"丁": 4,
	}
	assertion = map[string]func(u map[string]int) int{
		"甲": func(u map[string]int) int {
			res := 0
			if u["甲"] < u["丙"] {
				res++
			}
			if u["甲"] > u["丁"] {
				res++
			}
			return res
		},
		"乙": func(u map[string]int) int {
			res := 0
			if u["乙"] == 1 {
				res++
			}
			if u["甲"] > u["丙"] {
				res++
			}
			return res
		},
		"丙": func(u map[string]int) int {
			res := 0
			if u["甲"] < u["丙"] {
				res++
			}
			if u["乙"] > u["丁"] {
				res++
			}
			return res
		},
		"丁": func(u map[string]int) int {
			res := 0
			if u["丁"] < 4 {
				res++
			}
			if u["丙"] < u["乙"] {
				res++
			}
			return res
		},
	}
	rankScore = map[int]int{
		1: 2,
		2: 1,
		3: 1,
		4: 0,
	}
	rank = []int{1, 2, 3, 4}
)

func Test_logic(t *testing.T) {
	util.PrintJSON(countScore(users))
	ccc := 0
	solve(rank, 0, 3, func(in []int) {
		ccc++
		users["甲"] = in[0]
		users["乙"] = in[1]
		users["丙"] = in[2]
		users["丁"] = in[3]
		//util.PrintJSON(users)
		count := countScore(users)
		//util.PrintJSON(count)
		if count == 4 {
			util.Print(users)
		}
	})
	util.Print(ccc)

}

func swap(in []int, i, j int) {
	temp := in[i]
	in[i] = in[j]
	in[j] = temp
}

func solve(in []int, k, m int, handle func(in []int)) {
	if k == m {
		for i := 0; i <= m; i++ {
			handle(in)
		}
	} else {
		for i := k; i <= m; i++ {
			swap(in, i, k)
			solve(in, k+1, m, handle)
			swap(in, i, k)
		}
	}
}

func countScore(us map[string]int) int {
	count := 0
	for u, aser := range assertion {
		score := aser(us)
		if rankScore[us[u]] == score {
			count++
		}
	}
	return count
}
