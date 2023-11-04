package baseApi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	N   = 1000005
	mod = 998244353
)

var (
	__in  *bufio.Scanner
	__out *bufio.Writer
)

func read() int {
	__in.Scan()
	return atoi(__in.Text())
}
func readByte() []byte {
	__in.Scan()
	return __in.Bytes()
}
func setMaxSize() {
	buf := make([]byte, 0, N)
	__in.Buffer(buf, int(1e18))
}
func print(x ...interface{}) {
	fmt.Fprint(__out, x...)
}
func println(x ...interface{}) {
	fmt.Fprintln(__out, x...)
}
func max(a ...int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	ret := a[0]
	for i := 1; i < n; i++ {
		if a[i] > ret {
			ret = a[i]
		}
	}
	return ret
}
func min(a ...int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	ret := a[0]
	for i := 1; i < n; i++ {
		if a[i] < ret {
			ret = a[i]
		}
	}
	return ret
}
func __sum(a ...int) int {
	ret := 0
	for _, val := range a {
		ret += val
	}
	return ret
}
func __gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return __gcd(y, x%y)
}
func unique(a []int) []int {
	n := len(a)
	cnt := 0
	for i := 0; i < n; i++ {
		j := i
		for j+1 < n && a[j+1] == a[i] {
			j++
		}
		i = j
		a[cnt] = a[i]
		cnt++
	}
	return a[:cnt]
}
func lower_bound(a []int, x int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := (l + r) >> 1
		if a[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if a[l] >= x {
		return l
	}
	return -1
}
func atoi(s string) int {
	ret := 0
	for _, val := range s {
		ret = ret*10 + int(val-'0')
	}
	return ret
}
func parseSlice(s string) []string {
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, " ", "")
	tmp := strings.Split(s, ",")
	return tmp
}
func parseTree(s string) interface{} {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tmp := parseSlice(s)
	n := len(tmp)
	f := make(map[int]*TreeNode)
	for i := 0; i < n; i++ {
		if tmp[i] == "null" {
			continue
		}
		f[i] = &TreeNode{Val: atoi(tmp[i])}
	}
	for i := 0; i < n; i++ {
		if i*2+1 < n {
			f[i].Left = f[i*2+1]
		}
		if i*2+2 < n {
			f[i].Right = f[i*2+2]
		}
	}
	return f[0]
}
func parseTwoDimensionalSlice(s string) [][]int {
	tmp := parseSlice(s)
	n := len(tmp)
	m := 3
	ret := make([][]int, 0)
	for i := 0; i < n; i += m {
		a := make([]int, 0)
		for j := i; j <= i+m-1; j++ {
			a = append(a, atoi(tmp[j]))
		}
		ret = append(ret, a)
	}
	return ret
}

func solve() {
	for __in.Scan() {
		s := __in.Text()
		n := len(s)
		if s[n-1] != ';' {
			println(s)
			continue
		}
		name := strings.Split(s, " ")[3]
		//println("name = ", name)
		print(s[:n-1])
		ret := ""
		for i, val := range name {
			if i > 0 && 'A' <= val && val <= 'Z' {
				ret += "_" + strings.ToLower(string(val))
				continue
			}
			ret += strings.ToLower(string(val))
		}
		println(" [json_name = \"" + ret + "\", (gogoproto.jsontag) = \"" + ret + "\"];")
	}
}
func main() {
	inputFile, err := os.Open("my_data.in")
	if err != nil {
		inputFile = os.Stdin
	}
	__in = bufio.NewScanner(inputFile)
	__out = bufio.NewWriter(os.Stdout)
	defer __out.Flush()
	//setMaxSize()
	//__in.Split(bufio.ScanWords)
	tt := 1
	//tt = read()
	for tt > 0 {
		tt--
		solve()
	}
}
