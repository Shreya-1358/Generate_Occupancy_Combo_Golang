package model

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOccupancy(t *testing.T) {
	t.Run("TestGenerateOccupancy_Success", func(t *testing.T) {
		occupReq := OccupancyRequest{
			BaseAdult:     2,
			BaseChild:     0,
			MaxAdult:      3,
			MaxChild:      0,
			MaxGuest:      5,
			ChildAgeRange: []int{5, 10, 15},
		}

		occpResp := OccupancyResponse{
			Occupancy: []occupancy{{OccupancyCounter: 1, OccupancyDetails: []occupancyDetails{{Adult: 1, Child: 0}}},
				{OccupancyCounter: 2, OccupancyDetails: []occupancyDetails{{Adult: 1, Child: 1, ChildRange: "5"}, {Adult: 1, Child: 1, ChildRange: "10"}, {Adult: 1, Child: 1, ChildRange: "15"}, {Adult: 2, Child: 0}}},
				{OccupancyCounter: 3, OccupancyDetails: []occupancyDetails{{Adult: 3, Child: 0}}},
				{OccupancyCounter: 4}, {OccupancyCounter: 5},
			},
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockoccup := NewMockOccupancyInterface(cntrl)
		occupancyImpl := NewOccupancyImpl()

		comb1 := [][]int{{5}, {10}, {15}}

		mockoccup.EXPECT().GenerateCombinations(occupReq.ChildAgeRange, 1).Return(comb1).AnyTimes()

		actual_occup := occupancyImpl.GenerateOccupancy(&occupReq)
		assert.ObjectsAreEqualValues(occpResp, actual_occup)
	})
}

func TestGenerateOccupancy1(t *testing.T) {
	t.Run("TestGenerateOccupancy_Success", func(t *testing.T) {
		occupReq := OccupancyRequest{
			BaseAdult:     2,
			BaseChild:     2,
			MaxAdult:      3,
			MaxChild:      3,
			MaxGuest:      4,
			ChildAgeRange: []int{5, 10, 15},
		}

		occpResp := OccupancyResponse{
			Occupancy: []occupancy{{OccupancyCounter: 1, OccupancyDetails: []occupancyDetails{{Adult: 1, Child: 0}}},
				{OccupancyCounter: 2, OccupancyDetails: []occupancyDetails{{Adult: 1, Child: 1, ChildRange: "5"}, {Adult: 1, Child: 1, ChildRange: "10"}, {Adult: 1, Child: 1, ChildRange: "15"}, {Adult: 2, Child: 0}}},
				{OccupancyCounter: 3, OccupancyDetails: []occupancyDetails{{Adult: 1, Child: 2, ChildRange: "5#5"}, {Adult: 1, Child: 2, ChildRange: "5#10"}, {Adult: 1, Child: 2, ChildRange: "5#15"}, {Adult: 1, Child: 2, ChildRange: "10#10"}, {Adult: 1, Child: 2, ChildRange: "10#15"}, {Adult: 1, Child: 2, ChildRange: "15#15"}, {Adult: 2, Child: 1, ChildRange: "5"}, {Adult: 2, Child: 1, ChildRange: "10"}, {Adult: 2, Child: 1, ChildRange: "15"}, {Adult: 3, Child: 0}}},
				{OccupancyCounter: 4, OccupancyDetails: []occupancyDetails{{Adult: 1, Child: 3, ChildRange: "5#5#5"}, {Adult: 1, Child: 3, ChildRange: "5#5#10"}, {Adult: 1, Child: 3, ChildRange: "5#5#15"}, {Adult: 1, Child: 3, ChildRange: "5#10#10"}, {Adult: 1, Child: 3, ChildRange: "5#10#15"}, {Adult: 1, Child: 3, ChildRange: "5#15#15"}, {Adult: 1, Child: 3, ChildRange: "10#10#10"}, {Adult: 1, Child: 3, ChildRange: "10#10#15"}, {Adult: 1, Child: 3, ChildRange: "10#15#15"}, {Adult: 1, Child: 3, ChildRange: "15#15#15"}, {Adult: 2, Child: 2, ChildRange: "5#5"}, {Adult: 2, Child: 2, ChildRange: "5#10"}, {Adult: 1, Child: 2, ChildRange: "5#15"}, {Adult: 2, Child: 2, ChildRange: "10#10"}, {Adult: 2, Child: 2, ChildRange: "10#15"}, {Adult: 2, Child: 2, ChildRange: "15#15"}, {Adult: 3, Child: 1, ChildRange: "5"}, {Adult: 3, Child: 1, ChildRange: "10"}, {Adult: 3, Child: 1, ChildRange: "15"}}},
			},
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockoccup := NewMockOccupancyInterface(cntrl)
		occupancyImpl := NewOccupancyImpl()

		comb1 := [][]int{{5}, {10}, {15}}
		comb2 := [][]int{{5, 5}, {5, 10}, {5, 15}, {10, 10}, {10, 15}, {15, 15}}
		comb3 := [][]int{{5, 5, 5}, {5, 5, 10}, {5, 5, 15}, {5, 10, 10}, {5, 10, 15}, {5, 15, 15}, {10, 10, 10}, {10, 10, 15}, {10, 15, 15}, {15, 15, 15}}

		var occup occupancyDetails
		if occup.Child == 1 {
			mockoccup.EXPECT().GenerateCombinations(occupReq.ChildAgeRange, 1).Return(comb1).AnyTimes()
		} else if occup.Child == 2 {
			mockoccup.EXPECT().GenerateCombinations(occupReq.ChildAgeRange, 2).Return(comb2).AnyTimes()
		} else if occup.Child == 3 {
			mockoccup.EXPECT().GenerateCombinations(occupReq.ChildAgeRange, 3).Return(comb3).AnyTimes()
		}

		actual_occup := occupancyImpl.GenerateOccupancy(&occupReq)
		assert.ObjectsAreEqualValues(occpResp, actual_occup)
	})
}

func TestGenerateCombinations(t *testing.T) {
	t.Run("TestGenerateCombinations_Success", func(t *testing.T) {
		occupReq := OccupancyRequest{ChildAgeRange: []int{5, 10, 15}}
		occupResp := [][]int{{5}, {10}, {15}}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockoccup := NewMockOccupancyInterface(cntrl)
		occupancyImpl := NewOccupancyImpl()

		mockoccup.EXPECT().Combinations(occupReq, 0, 2, 1, []int{}, [][]int{{}}).AnyTimes()
		actual_occup := occupancyImpl.GenerateCombinations(occupReq.ChildAgeRange, 1)
		assert.Equal(t, occupResp, actual_occup)
	})
}

var table = []struct {
	occupReq OccupancyRequest
}{
	{occupReq: OccupancyRequest{
		BaseAdult:     10,
		BaseChild:     10,
		MaxAdult:      13,
		MaxChild:      13,
		MaxGuest:      25,
		ChildAgeRange: []int{5, 10, 15},
	},
	},
	{occupReq: OccupancyRequest{
		BaseAdult:     20,
		BaseChild:     12,
		MaxAdult:      30,
		MaxChild:      18,
		MaxGuest:      40,
		ChildAgeRange: []int{5, 10, 15},
	},
	},
}

func BenchmarkGenerateOccupancy(b *testing.B) {

	occup := NewOccupancyImpl()
	b.Run(fmt.Sprintf("BenchmarkGenerateOccupancy_Size_%d", table[0].occupReq), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			occup.GenerateOccupancy(&table[0].occupReq)
		}
	})

	b.Run(fmt.Sprintf("BenchmarkGenerateOccupancy_Size_%d", table[1].occupReq), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			occup.GenerateOccupancy(&table[1].occupReq)
		}
	})
}
