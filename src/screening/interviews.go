package screening

import "time"

type Interviews struct {
	interviews []Interview
}

func NewInterviews() Interviews {
	return Interviews{}
}

// AddNextInterview 次の面接を追加する
func (i *Interviews) AddNextInterview(interviewDate time.Time) *Interviews {
	nextInterview := NewInterview(interviewDate, i.getNextInterviewNumber())
	i.interviews = append(i.interviews, nextInterview)
	return i
}

func (i Interviews) getNextInterviewNumber() int {
	return len(i.interviews) + 1
}
