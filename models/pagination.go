package models

type Paginator struct {
	Page      int
	PerPage   int
	TotalPage int64
	TotalRows int64

	NavPage int
}

func NewDefaultPaginator(perPage int) Paginator {
	return Paginator{Page: 1, PerPage: perPage, TotalPage: 1, NavPage: 7}
}

func (p *Paginator) SetFromTotalRows(totalRows int64) {
	p.TotalPage = (totalRows + int64(p.PerPage) - int64(1)) / int64(p.PerPage)
	p.TotalRows = totalRows
}

func (p *Paginator) GetOffset() int64 {
	return (int64(p.Page) - 1) * int64(p.PerPage)
}

func (p *Paginator) GetRangeData() (int64, int64) {
	start := p.GetOffset()
	end := p.GetOffset() + int64(p.PerPage)
	if end > p.TotalRows {
		end = p.TotalRows
	}
	return start, end
}

func (p *Paginator) GetNavPagination(currentPage int) []int64 {
	var maxPagesToShow = p.NavPage
	var chuckPage = (p.TotalPage + int64(maxPagesToShow) - int64(1)) / int64(maxPagesToShow)
	if chuckPage == 0 {
		chuckPage = 1
	}

	var m = map[int64][]int64{}
	var rangePageIndex = -1
	for i := int64(1); i <= chuckPage; i++ {
		st := ((i - 1) * int64(maxPagesToShow))

		end := st + int64(maxPagesToShow)
		if st <= 0 {
			st = 1
			end = int64(maxPagesToShow)
		}
		if i > 1 {
			st += 1
		}
		if end > p.TotalPage {
			end = p.TotalPage
		}

		m[i] = make([]int64, 0)
		for j := st; j <= end; j++ {
			m[i] = append(m[i], int64(j))
			if int64(currentPage) == j {
				rangePageIndex = int(i)
			}
		}
		if rangePageIndex != -1 {
			break
		}
	}

	return m[int64(rangePageIndex)]
}
