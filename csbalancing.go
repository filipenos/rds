package csbalancing

import (
	"sort"
)

// Entity ...
type Entity struct {
	ID    int
	Score int
}

// CustomerSuccessBalancing ...
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {

	var (
		css   = AvailableCss(customerSuccess, customerSuccessAway)
		calls = make(Calls)
	)

	for _, customer := range customers {
		id := css.Lookup(customer.Score)
		if id == 0 {
			continue
		}
		calls.Add(id)
	}

	return calls.Get()
}

//CustomerSuccess represent the available customer success
type CustomerSuccess struct {
	availables map[int]int
	scoreToId  map[int]int
	scores     []int
}

//Lookup find id for score
func (css CustomerSuccess) Lookup(customerScore int) int {
	pos := sort.SearchInts(css.scores, customerScore)
	if pos >= len(css.scores) {
		return 0
	}
	score := css.scores[pos]
	id := css.scoreToId[score]
	return id
}

//AvailableCss create CustomerSuccess struct with available css
func AvailableCss(customerSuccess []Entity, customerSuccessAway []int) CustomerSuccess {
	css := CustomerSuccess{
		availables: make(map[int]int),
		scoreToId:  make(map[int]int),
		scores:     make([]int, 0, len(customerSuccess)),
	}

	for _, customer := range customerSuccess {
		css.availables[customer.ID] = customer.Score
	}
	for _, customer := range customerSuccessAway {
		delete(css.availables, customer)
	}

	for id, score := range css.availables {
		css.scores = append(css.scores, score)
		css.scoreToId[score] = id
	}
	sort.Ints(css.scores)

	return css
}

//Calls represent the calls of CustomerSuccess
type Calls map[int]int

//Add add a call to manager
func (c Calls) Add(id int) {
	c[id]++
}

//Get retrieve the manager with more calls
func (c Calls) Get() int {
	var id, big int
	for k, v := range c {
		if v > big {
			id, big = k, v
		} else if v == big {
			id = 0
			break
		}
	}
	return id
}
