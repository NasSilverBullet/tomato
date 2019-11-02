package tomato

import (
	"fmt"
	"reflect"
)

type Pagination interface {
	GetCount() int
	GetPer() int
	GetCurrent() int
	GetFirst() int
	GetLast() int
	GetPrevious() int
	GetNext() int
}

type pagination struct {
	Count    int
	Per      int
	Current  int
	First    int
	Last     int
	Previous int
	Next     int
}

func New(current, per, count int) (Pagination, error) {
	if current < 0 {
		return nil, fmt.Errorf("can't use current: %d", current)
	}

	if per < 1 {
		return nil, fmt.Errorf("can't use per page: %d", per)
	}

	if count < 0 {
		return nil, fmt.Errorf("can't use records count: %d", count)
	}

	last := calcLast(count, per)

	current = calcCurrent(current, per, last, count)

	first := calcFirst(count)
	previous := calcPrevious(current, first)
	next := calcNext(current, last)

	return &pagination{
		Count:    count,
		Per:      per,
		Current:  current,
		First:    first,
		Last:     last,
		Previous: previous,
		Next:     next,
	}, nil
}

func (p *pagination) GetCurrent() int {
	return p.Current
}

func (p *pagination) GetPer() int {
	return p.Per
}

func (p *pagination) GetFirst() int {
	return p.First
}

func (p *pagination) GetLast() int {
	return p.Last
}

func (p *pagination) GetPrevious() int {
	return p.Previous
}

func (p *pagination) GetNext() int {
	return p.Next
}

func (p *pagination) GetCount() int {
	return p.Count
}

func (p *pagination) FilterCurrent(i interface{}) (interface{}, error) {
	if reflect.TypeOf(i).Kind() != reflect.Slice {
		return nil, fmt.Errorf("%v type is not supported, please give %v", reflect.TypeOf(i).Kind(), reflect.Slice)
	}

	v := reflect.ValueOf(i)

	currentFirstRecIdx := p.Per * (p.Current - 1)

	currentLastRecIdx := p.Per * p.Current
	if currentLastRecIdx > p.Count {
		currentLastRecIdx = p.Count
	}

	if currentFirstRecIdx > currentLastRecIdx {
		return nil, fmt.Errorf("pagination's fields is broken")
	}

	if v.Len() < currentLastRecIdx {
		return nil, fmt.Errorf("Can't filter %v", i)
	}

	return v.Slice(currentFirstRecIdx, currentLastRecIdx).Interface(), nil
}

func calcFirst(count int) int {
	if count > 0 {
		return 1
	}
	return 0
}

func calcLast(count, per int) int {
	return (count + (per - 1)) / per
}

func calcCurrent(current, per, last, count int) int {
	if current == 0 {
		return 1
	}

	if (current-1)*per >= count {
		return last
	}

	return current
}

func calcNext(current, last int) int {
	if current < last {
		return current + 1
	}
	return 0
}

func calcPrevious(current, first int) int {
	if current > first {
		return current - 1
	}
	return 0
}
