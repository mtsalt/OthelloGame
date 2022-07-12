package othello

type point struct {
	x int
	y int
}

type pointList struct {
	pointList []point
}

func (p *pointList) Contain(target point) bool {

	for _, item := range p.pointList {
		if item == target {
			return true
		}
	}
	return false
}

func (p *pointList) Add(target point) {
	p.pointList = append(p.pointList, target)
}

func (p *pointList) Len() int {
	return len(p.pointList)
}
