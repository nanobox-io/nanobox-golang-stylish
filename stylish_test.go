package stylish

import "testing"
// import "fmt"

func TestNest(t *testing.T) {
	if len(ProcessStart("hi")) != len(Nest(2, ProcessStart("hi"))) {
		t.Errorf("the lenghts dont match after nesting progress bars")
	}


	// offset the message by each newline lenght in this type of thing
	// warnings add 2 new lines
	if len(Warning("this is a warning")) != len(Nest(1, Warning("this is a warning")))-4 {
		t.Errorf("the lenghts dont match after nesting warnings")
		
	}

	// offset the message by each newline lenght in this type of thing
	// errors add 3 new lines
	if len(Error("ERROR","this is a Error")) != len(Nest(1, Error("ERROR","this is a Error")))-6 {
		t.Errorf("the lenghts dont match after nesting errors")
	}

}
