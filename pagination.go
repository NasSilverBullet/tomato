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
	FilterCurrent(i interface{}) (interface{}, error)
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
	if current < 1 {
		return nil, fmt.Errorf("can't use current: %d", current)
	}

	if per < 1 {
		return nil, fmt.Errorf("can't use per page: %d", per)
	}

	if count < 0 {
		return nil, fmt.Errorf("can't use records count: %d", count)
	}

	return &pagination{
		Count:    count,
		Per:      per,
		Current:  calcCurrent(current, per, count),
		First:    calcFirst(),
		Last:     calcLast(count, per),
		Previous: calcPrevious(current, per, count),
		Next:     calcNext(current, per, count),
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

	if v.Len() != p.Count {
		return nil, fmt.Errorf("length of Slice doesn't match length of pagination")
	}

	currentFirstIndex := p.Per * (p.Current - 1)

	currentLastIndex := p.Per * p.Current
	if currentLastIndex > p.Count {
		currentLastIndex = p.Count
	}

	if currentFirstIndex > currentLastIndex {
		return nil, fmt.Errorf("pagination's fields is broken")
	}

	if v.Len() < currentLastIndex {
		return nil, fmt.Errorf("Can't filter %v", i)
	}

	return v.Slice(currentFirstIndex, currentLastIndex).Interface(), nil
}

func calcFirst() int {
	return 1
}

func calcLast(count, per int) int {
	return (count + (per - 1)) / per
}

func calcCurrent(current, per, count int) int {
	if (current-1)*per >= count {
		return calcLast(count, per)
	}
	return current
}

func calcNext(current, per, count int) int {
	current = calcCurrent(current, per, count)
	if last := calcLast(count, per); current >= last {
		return last
	}
	return current + 1
}

func calcPrevious(current, per, count int) int {
	current = calcCurrent(current, per, count)
	if first := calcFirst(); current <= first {
		return first
	}
	return current - 1
}
