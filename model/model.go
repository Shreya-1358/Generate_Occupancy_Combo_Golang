package model

import (
	"strconv"
)

type OccupancyRequest struct {
	BaseAdult     int   `json:"base_adult"`
	BaseChild     int   `json:"base_child"`
	MaxAdult      int   `json:"max_adult"`
	MaxChild      int   `json:"max_child"`
	MaxGuest      int   `json:"max_guest"`
	ChildAgeRange []int `json:"child_age_range"`
}

type OccupancyResponse struct {
	Occupancy []occupancy `json:"occupancy"`
}
type occupancy struct {
	OccupancyCounter int                `json:"occupancy_counter"`
	OccupancyDetails []occupancyDetails `json:"occupancy_details,omitempty"`
}

type occupancyDetails struct {
	Adult      int    `json:"adult"`
	Child      int    `json:"child"`
	ChildRange string `json:"childRange,omitempty"`
}

type OccupancyInterface interface {
	GenerateOccupancy(occupancyRequest *OccupancyRequest) OccupancyResponse
	GenerateCombinations(nums []int, child_no int) [][]int
	Combinations(nums []int, start, end, child_no int, combination []int, res *[][]int)
}

type OccupancyImpl struct {
}

func NewOccupancyImpl() OccupancyImpl {
	return OccupancyImpl{}
}

func (o *OccupancyImpl) Combinations(child_age_range []int, start, end, child_no int, combination []int, res *[][]int) {
	if len(combination) == child_no {
		comb := make([]int, len(combination))
		copy(comb, combination)
		*res = append(*res, comb)
		return
	}
	for i := start; i <= end; i++ {
		combination = append(combination, child_age_range[i])
		o.Combinations(child_age_range, i, end, child_no, combination, res)
		combination = combination[:len(combination)-1]
	}
}

func (o *OccupancyImpl) GenerateCombinations(child_age_range []int, child_no int) [][]int {
	var res [][]int
	var combination []int
	o.Combinations(child_age_range, 0, len(child_age_range)-1, child_no, combination, &res)
	return res
}

func (o *OccupancyImpl) GenerateOccupancy(occupancyRequest *OccupancyRequest) OccupancyResponse {

	occup_map := make(map[string]int)
	for i := 1; i <= occupancyRequest.MaxGuest; i++ {
		for j := 1; j <= occupancyRequest.MaxAdult; j++ {
			if occupancyRequest.MaxChild == 0 && occupancyRequest.BaseAdult > 1 {
				for k := 0; k <= occupancyRequest.BaseAdult-j; k++ {
					if j+k <= i {
						occup_map[strconv.Itoa(j)+"a"+strconv.Itoa(k)+"c"]++
					}
				}
			} else if j <= i {
				for k := 0; k <= occupancyRequest.MaxChild; k++ {
					if j+k <= i {
						occup_map[strconv.Itoa(j)+"a"+strconv.Itoa(k)+"c"]++
					}
				}
			}
		}
	}
	var occupancyResponse OccupancyResponse
	for i := 1; i <= occupancyRequest.MaxGuest; i++ {
		var occup occupancy
		for j := range occup_map {
			adult := int(j[0] - '0')
			child := int(j[2] - '0')

			if adult+child == i {
				if child == 0 {
					occup.OccupancyDetails = append(occup.OccupancyDetails, occupancyDetails{Adult: adult, Child: child, ChildRange: ""})
				} else {
					combs := o.GenerateCombinations(occupancyRequest.ChildAgeRange, child)
					for _, comb := range combs {
						ch := ""
						for _, c := range comb {
							ch += strconv.Itoa(c) + "#"
						}
						ch = ch[:len(ch)-1]
						occup.OccupancyDetails = append(occup.OccupancyDetails, occupancyDetails{Adult: adult, Child: child, ChildRange: ch})
					}
				}
			}
		}
		occupancyResponse.Occupancy = append(occupancyResponse.Occupancy, occupancy{OccupancyCounter: i, OccupancyDetails: occup.OccupancyDetails})
	}
	return occupancyResponse
}
