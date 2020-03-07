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
		/*
		event{
			start: 11,
			end: 20,
			groups: map[string]bool{
				"pickle festival": true,
			},
		},
		event{
			start: 1,
			end: 7,
			groups: map[string]bool{
				"dog adoption": true,
			},
		},
		event{
			start: 8,
			end: 11,
			groups: map[string]bool{
				"elderly community day": true,
			},
		},
		event{
			start: 0,
			end: 4,
			groups: map[string]bool{
				"food function": true,
			},
		},
		*/
	}

	fmt.Printf("\n-------------\n| Events    |\n-------------\n\n")

	for _, c := range campaigns {
		fmt.Printf("Name: ")
		for name, _ := range c.groups {
			fmt.Printf("%v, ", name)
		}
		fmt.Printf("\nStart:%v\nEnd:%v\n\n", c.start, c.end)
	}

	overlaps := recursiveOverlaps(campaigns, make([]event, 0))

	fmt.Printf("\n-------------\n| Overlaps  |\n-------------\n\n")

	for _, c := range overlaps {
		fmt.Printf("Events: ")
		for name, _ := range c.groups {
			fmt.Printf("%v, ", name)
		}
		fmt.Printf("\nStart:%v\nEnd:%v\n\n", c.start, c.end)
	}
}

func recursiveOverlaps(events []event, overlaps []event) []event {

	//pop comparisonEvent (first item in array)
	comparisonEvent, events := events[len(events)-1], events[:len(events)-1]
	

	if len(events) == 0 {//check base case
		return overlaps;
	}

	//Find overlaps
	for _, eventItem := range events {

		overlaping, overlapCase := overlapExists(comparisonEvent, eventItem)

		if 	overlaping {

			groups := mergeKeys(comparisonEvent.groups, eventItem.groups)

			switch overlapCase {
			
				case 1:
					overlaps = append( overlaps, event{eventItem.start, eventItem.end, groups} )

				case 2:
					overlaps = append( overlaps, event{comparisonEvent.start, comparisonEvent.end, groups} )

				case 3:
					overlaps = append( overlaps, event{eventItem.start, comparisonEvent.end, groups} )

				case 4:
					overlaps = append( overlaps, event{comparisonEvent.start, eventItem.end, groups} )
			}
		}
	}

	return recursiveOverlaps(events, overlaps)
}

func overlapExists(a event, b event) (bool, int) {
	
	if between(a.start, a.end, b.start) && between(a.start, a.end, b.end) {
	// [----------]
	//    [-----]
		return true, 1
	}

	if between(b.start, b.end, a.start) && between(b.start, b.end, a.end) {
	//    [-----]
	// [----------]
		return true, 2
	}

	if between(a.start, a.end, b.start) {
	// [----------]
	//    [--------------
		return true, 3
	}

	if between(a.start, a.end, b.end) {
	//            [----------]
	//    --------------]
		return true, 4
	}

	return false, 0
}

func between(a, b, c int) bool {
	//is c between a and b?
	if c > a && c < b {
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
