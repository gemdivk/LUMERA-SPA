package usecase

import (
	"fmt"
	"time"
)

func (s *SalonInteractor) GetProcedureSlots(procedureID string, date string) ([]string, error) {
	override, err := s.repo.GetScheduleOverride(procedureID, date)
	if err != nil {
		return nil, err
	}
	var startTime, endTime string
	if override != nil {
		startTime = override.StartTime
		endTime = override.EndTime
	} else {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}
		weekday := parsedDate.Weekday()
		sched, err := s.repo.GetWeeklySchedule(procedureID, int32(weekday))
		if err != nil {
			return nil, err
		}
		if sched == nil {
			return []string{}, nil
		}
		startTime = sched.StartTime
		endTime = sched.EndTime
	}

	procs, err := s.repo.GetAllProcedures()
	if err != nil {
		return nil, err
	}
	var duration int32
	for _, p := range procs {
		if p.ID == procedureID {
			duration = p.Duration
			break
		}
	}
	if duration == 0 {
		return nil, fmt.Errorf("procedure not found")
	}

	layout := "15:04"
	start, _ := time.Parse(layout, startTime)
	end, _ := time.Parse(layout, endTime)
	var slots []string
	for start.Add(time.Duration(duration)*time.Minute).Before(end) || start.Add(time.Duration(duration)*time.Minute).Equal(end) {
		slots = append(slots, start.Format("15:04"))
		start = start.Add(time.Duration(duration+15) * time.Minute)
	}
	return slots, nil
}
