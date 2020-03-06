package main

import (
	"fmt"
)

type event struct {
	start int
	end int
	groups map[string]bool
}

func main(){
	//Create Events
	campaigns := []event{
		event{
			start: 0,
			end: 4,
			groups: map[string]bool{
				"apple sale": true,
			},
		},
		event{
			start: 2,
			end: 5,
			groups: map[string]bool{			
				"banna sale": true,			
			},
		},
		event{
			start: 3,
			end: 10,
			groups: map[string]bool{
				"car sale": true,
			},
		},
		event{
			start: 11,
			end: 20,
			groups: map[string]bool{
				"pickle festival": true,
			},
		},
		/*
		event{
			start: 1,
			end: 7,
			groups: []string{"dog adoption"},
		},
		event{
			start: 8,
			end: 11,
			groups: []string{"elderly community day"},
		},
		event{
			start: 0,
			end: 4,
			groups: []string{"food function"},
		},
		*/
	}

	overlaps := recursiveOverlaps(campaigns, make([]event, 0))

	fmt.Printf("\n-------------\n| Overlaps  |\n-------------\n\n")

	for _, c := range overlaps {
		fmt.Printf("Campaign %v\nStart:%v\nEnd:%v\n\n", c.groups, c.start, c.end)
	}
}

func recursiveOverlaps(events []event, overlaps []event) []event {
	//Remove event, that we shall call "comparisonEvent", and compare to the rest of the events to find overlaps.
	//Throw that event away and repeat this process.
	//base case: events array is of length 2. return comparisonEvent

	// fmt.Printf("Events slice before pop: %v\n", events)
	//pop comparisonEvent (first item in array)
	comparisonEvent, events := events[len(events)-1], events[:len(events)-1]
	// fmt.Printf("Events slice after pop: %v\n", events)
	fmt.Printf("comparisonEvent: %v\n", comparisonEvent )	
	

	if len(events) == 0 {//base case
		return overlaps;
	}

	//Find overlaps
	for _, eventItem := range events {

		if 	overlapExists(comparisonEvent, eventItem) {

			if comparisonEvent.end > eventItem.end { //comparisonEvent encompasses eventItem

				//append names and make it unique
				groups := mergeKeys(comparisonEvent.groups, eventItem.groups)

				//add to overlaps
				overlaps = append( overlaps, event{eventItem.start, eventItem.end, groups} )

			} else {

				//append names and make it unique
				groups := mergeKeys(comparisonEvent.groups, eventItem.groups)

				//add to overlaps
				overlaps = append( overlaps, event{eventItem.start, comparisonEvent.end, groups } )

			}
		}
	}


	// literalEvent := []event{ event{len(events),0,map[string]bool{"Error: Nil Case Reached. Length of Events in start of object.":true}} }

	//call function again if basecase not met.
	return recursiveOverlaps(events, overlaps)
}

func overlapExists(a event, b event) bool {
	if a.end > b.start {
		return true
	}
	return false
}

// Given two maps, recursively merge right into left, NEVER replacing any key that already exists in left
func mergeKeys(left, right map[string]bool) map[string]bool {
    for key, rightVal := range right {
        left[key] = rightVal
    }
    return left
}
