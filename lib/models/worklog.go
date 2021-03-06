package models

import (
	"encoding/json"
	"fmt"
)

// Worklog struct
type Worklog struct {
	IssueKey         string `json:"issueKey"`
	Issue            Issue  `json:"issue"`
	TimeSpentSeconds int64  `json:"timeSpentSeconds"`
	StartDate        string `json:"startDate"`
	StartTime        string `json:"startTime"`
	Description      string `json:"description"`
	AuthorAccountID  string `json:"authorAccountId"`
}

// Issue struct
type Issue struct {
	Key string `json:"key"`
}

// ToJSONData encode to json
func (w Worklog) ToJSONData() ([]byte, error) {
	dat, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func (w Worklog) String() string {
	key := w.IssueKey
	if key == "" {
		key = w.Issue.Key
	}
	return fmt.Sprintf("%s: %s", key, w.Description)
}

// NewWorklog func
func NewWorklog(date, details, jiraUser string, timeSpent int64) Worklog {
	issue := "TIQ-684" // TODO: map json via env?
	return Worklog{
		IssueKey:         issue,
		TimeSpentSeconds: timeSpent,
		StartDate:        date,
		StartTime:        "13:00:00",
		Description:      details,
		AuthorAccountID:  jiraUser,
	}
}

// Worklogs struct
type Worklogs struct {
	Results []Worklog `json:"results"`
}

// Logged timespent in seconds
func (logs *Worklogs) Logged() (logged int64) {
	for _, worklog := range logs.Results {
		logged += worklog.TimeSpentSeconds
	}
	return logged
}

// NewWorklogs func
func NewWorklogs(data []byte) (*Worklogs, error) {
	logs := &Worklogs{}

	if err := json.Unmarshal(data, logs); err != nil {
		return nil, err
	}

	return logs, nil
}
