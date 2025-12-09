package day8

type UF struct {
	parent []int
	size   []int
	groups int
}

func NewUF(n int) *UF {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &UF{parent: p, size: s, groups: n}
}

func (u *UF) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UF) Union(a, b int) bool {
	pa, pb := u.Find(a), u.Find(b)
	if pa == pb {
		return false
	}
	if u.size[pa] < u.size[pb] {
		pa, pb = pb, pa
	}
	u.parent[pb] = pa
	u.size[pa] += u.size[pb]
	u.groups--
	return true
}
