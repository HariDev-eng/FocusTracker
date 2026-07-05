package service

import (
	"sort"
	"time"

	"focustracker/internal/models"
)

type StreakService struct{}

func NewStreakService() *StreakService {
	return &StreakService{}
}

func (s *StreakService) Compute(completions []models.Completion) (current int, longest int) {
	if len(completions) == 0 {
		return 0, 0
	}

	dateSet := make(map[string]bool, len(completions))
	for _, c := range completions {
		dateSet[c.Date.Format("2006-01-02")] = true
	}

	d := time.Now()
	if !dateSet[d.Format("2006-01-02")] {
		d = d.AddDate(0, 0, -1)
	}
	for dateSet[d.Format("2006-01-02")] {
		current++
		d = d.AddDate(0, 0, -1)
	}

	// Copy before sorting — this slice came from the caller (the repo's
	// return value), and mutating their ordering as a side effect of
	// computing a streak would be a nasty surprise if they use it after.
	sorted := make([]models.Completion, len(completions))
	copy(sorted, completions)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Date.Before(sorted[j].Date)
	})

	run := 0
	var prev time.Time
	for i, c := range sorted {
		if i == 0 || c.Date.Sub(prev).Hours() != 24 {
			run = 1
		} else {
			run++
		}
		if run > longest {
			longest = run
		}
		prev = c.Date
	}

	return current, longest
}
