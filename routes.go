package golwf

type Routes []Route

func (rs *Routes) Append(r Route) {
	*rs = append(*rs, r)
}

func (rs Routes) Len() int {
	return len(rs)
}

func (rs Routes) Less(i, j int) bool {
	r1 := rs[i]
	r2 := rs[j]

	return r1.path < r2.path || (r1.path == r2.path && methodOrder[r1.method] < methodOrder[r2.method])
}

func (rs Routes) Swap(i, j int) {
	temp := rs[i]
	rs[i] = rs[j]
	rs[j] = temp
}
