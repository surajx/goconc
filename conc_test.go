package conc

import (
	"testing";
//	"fmt";
)

func TestFor(t *testing.T) {
	numbers := make(chan Box);
	go func() {
		for i:=0; i<20; i++ {
			numbers <- i;
		}
		close(numbers);
	}();
	
	var vals [20]int;
	wait := For(numbers, 3, func(i Box) {vals[i.(int)] = 1});
	wait();
	
	total := 0;
	for i:=0; i<20; i++ {
		total += vals[i];
	}
	
	if total != 20 {
		t.Fail();
	}
}

func TestMap(t *testing.T) {
	incr := func(a Box) Box {
		return a.(int)+1;
	};
	numbers := make(chan Box);
	go func() {
		for i:=0; i<20; i++ {
			numbers <- i;
		}
		close(numbers);
	}();
	incrNumbers := Map(incr, numbers);
	for i:=0; i<20; i++ {
		j := <- incrNumbers;
		if i+1 != j.(int) {
			t.Fail();
		}
	}
}

func TestFold(t *testing.T) {
	sum := func(a Box, b Box) Box {
		return a.(int)+b.(int);
	};
	numbers := make(chan Box);
	go func() {
		for i:=0; i<10; i++ {
			numbers <- i;
		}
		close(numbers);
	}();
	totalSum := Fold(sum, numbers);

	if totalSum.(int) != 45 {
		t.Fail();
	}
}

func TestFilter(t *testing.T) {
	numbers := make(chan Box);
	go func() {
		for i:=0; i<10; i++ {
			numbers <- i;
		}
		close(numbers);
	}();
	
	results := Filter(func(i Box) bool { return i.(int)%2==0 }, numbers);
	for i := 0; i < 10; i+=2 {
		if i != (<-results).(int) {
			t.Fail();
		}
	}
}
