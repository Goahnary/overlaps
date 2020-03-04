package main

import (
	"fmt"
	"reflect"
)

type event struct {
	start int
	end int
	groups []string
}

func main(){
	//Create Events
	campaigns := []event{
		event{
			start: 0,
			end: 4,
			groups: []string{"apple sale"},
		},
		event{
			start: 2,
			end: 5,
			groups: []string{"banna sale"},
		},
		event{
			start: 3,
			end: 10,
			groups: []string{"car sale"},
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

	overlaps := findOverlaps(campaigns)

	fmt.Printf("\n-------------\n| Overlaps  |\n-------------\n\n")

	for _, c := range overlaps {
		fmt.Printf("Campaign %v\nStart:%v\nEnd:%v\n\n", c.groups, c.start, c.end)
	}
}

func findOverlaps(campaigns []event) []event {

	//return array
	overlaps := make([]event, 0)

	fmt.Printf("-------------\n| Campaigns |\n-------------\n\n")

	//if(campaigns == nil){ //base case
	//	return overlaps
	//} else {
		//compare events
		for _, campA := range campaigns {
			for _, campB := range campaigns {
				if campB.start < campA.end && !reflect.DeepEqual(campA.groups, campB.groups) && campA.start <= campB.start{ //There's an overlap!
					//one encompasses another
					if campA.end > campB.end { //campA encompasses campB
						//append names and make it unique
						groups := appendUnique(campA.groups, campB.groups)
						overlaps = append( overlaps, event{campB.start, campB.end, groups})
					} else {
						//append names and make it unique
						groups := appendUnique(campA.groups, campB.groups)
						overlaps = append( overlaps, event{campB.start, campA.end, groups } )
					}
				}
			}
			fmt.Printf("Campaign %v\nStart:%v\nEnd:%v\n\n", campA.groups, campA.start, campA.end)
		}
		//return append(overlaps, findOverlaps(overlaps)...)
		return overlaps
	//}
}

func appendUnique(a []string, b []string) []string {

	check := make(map[string]int)
	d := append(a, b...)
	res := make([]string,0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res,letter)
	}

	return res
}
