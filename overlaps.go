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
	//Create Requests array
	campaigns := make([]event, 3)

	//Populate array
	campaigns[0] = event{0,  4, []string{"apple sale"}}
	campaigns[1] = event{2,  5, []string{"banna sale"}}
	campaigns[2] = event{3, 10, []string{"car sale"}}
	//campaigns[3] = event{1,  7, []string{"dog adoption"}}
	//campaigns[4] = event{8, 11, []string{"elderly community day"}}
	//campaigns[5] = event{0,  4, []string{"food function"}}

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
