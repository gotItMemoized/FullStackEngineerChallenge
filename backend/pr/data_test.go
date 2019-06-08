package pr

import (
	"reflect"
	"testing"

	"github.com/gotItMemoized/FullStackEngineerChallenge/backend/user"
)

func Test_feedbackChanges(t *testing.T) {
	type args struct {
		previous []Feedback
		current  []Feedback
	}
	tests := []struct {
		name    string
		args    args
		added   []string
		removed []string
	}{
		{name: "nil not explode check", args: args{previous: nil, current: nil}, added: nil, removed: nil},
		{name: "remove one", args: args{previous: []Feedback{Feedback{ID: "1", ReviewerID: "1", Reviewer: user.User{ID: "1"}}}, current: nil}, added: nil, removed: []string{"1"}},
		{
			name: "simple no change", args: args{
				previous: []Feedback{
					Feedback{ID: "1", ReviewerID: "1", Reviewer: user.User{ID: "1"}},
				}, current: []Feedback{
					Feedback{ID: "1", ReviewerID: "1", Reviewer: user.User{ID: "1"}},
				}}, added: nil, removed: nil,
		},
		{
			name: "simple add", args: args{
				previous: []Feedback{}, current: []Feedback{
					Feedback{ID: "1", ReviewerID: "1", Reviewer: user.User{ID: "1"}},
				}}, added: []string{"1"}, removed: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := feedbackChanges(tt.args.previous, tt.args.current)
			if !reflect.DeepEqual(got, tt.added) {
				t.Errorf("feedbackChanges() function added = %v, expected %v", got, tt.added)
			}
			if !reflect.DeepEqual(got1, tt.removed) {
				t.Errorf("feedbackChanges() function removed = %v, expected %v", got1, tt.removed)
			}
		})
	}
}
