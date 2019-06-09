package pr

import (
	"log"
	"strconv"

	"github.com/gotItMemoized/FullStackEngineerChallenge/backend/user"
	"github.com/pkg/errors"
)

type MapData struct {
	reviewCount   int
	reviews       map[string]Review
	feedbackCount int

	// Accidental bad design. This means we're tightly coupled to this service
	// for the sqlReviewService we had foreign key constraints that
	// we need to be careful of here
	UserData *user.MapData
}

func (rs *MapData) getAllReviews() []Review {
	results := make([]Review, len(rs.reviews))
	ind := 0
	for _, review := range rs.reviews {
		results[ind] = review
		ind += 1
	}

	return results
}

func (rs *MapData) getReviewById(reviewId string) *Review {
	var result Review
	review, ok := rs.reviews[reviewId]

	if !ok {
		log.Printf("Could not get the reviews\n")
		return nil
	}
	result = review

	return &result
}

func (rs *MapData) getAllFeedbackForReviewer(userID string) []FlatFeedback {
	var results []FlatFeedback

	for _, review := range rs.reviews {
		if review.Feedback == nil {
			continue
		}
		var feedback Feedback
		for _, fb := range review.Feedback {
			if fb.ReviewerID != userID {
				continue
			}
			feedback = fb
			break
		}
		if len(feedback.ID) == 0 {
			continue
		}
		result := FlatFeedback{
			ID:         feedback.ID,
			ReviewID:   review.ID,
			ReviewerID: feedback.ReviewID,
			UserID:     review.User.ID,
			Name:       review.User.Name,
			Username:   review.User.Username,
			Message:    feedback.Message,
		}
		results = append(results, result)
	}

	if results == nil {
		results = make([]FlatFeedback, 0)
	}

	return results
}

func (rs *MapData) updateFeedback(updatedReview *FlatFeedback) error {

	found := false
	for key, review := range rs.reviews {
		if review.Feedback == nil {
			continue
		}
		updated := false
		for i, fb := range review.Feedback {
			if fb.ID == updatedReview.ID {
				fb.Message.String = updatedReview.Message.String
				fb.Message.Valid = len(updatedReview.Message.String) != 0
				review.Feedback[i] = fb
				updated = true
				break
			}
		}
		if updated {
			rs.reviews[key] = review
			found = true
			break
		}
	}

	if !found {
		log.Printf("Error updating feedback: not found\n")
		return errors.New("Not found")
	}

	return nil
}

func (rs *MapData) getFeedbackForReviewer(userID, idToUpdate string) *FlatFeedback {
	var results FlatFeedback

	for _, review := range rs.reviews {
		if review.Feedback == nil {
			continue
		}
		var feedback Feedback
		for _, fb := range review.Feedback {
			if fb.ReviewerID != idToUpdate {
				continue
			}
			feedback = fb
			break
		}
		if len(feedback.ID) == 0 {
			continue
		}
		result := FlatFeedback{
			ID:         feedback.ID,
			ReviewID:   review.ID,
			ReviewerID: feedback.ReviewID,
			UserID:     review.User.ID,
			Name:       review.User.Name,
			Username:   review.User.Username,
			Message:    feedback.Message,
		}
		results = result
		break
	}

	return &results
}

func (rs *MapData) create(review *Review) error {
	rs.reviewCount += 1
	review.ID = strconv.Itoa(rs.reviewCount)
	if review.Feedback != nil && len(review.Feedback) > 0 {
		for i, fb := range review.Feedback {
			rs.feedbackCount += 1
			fb.ID = strconv.Itoa(rs.feedbackCount)
			fb.Reviewer = *rs.UserData.GetUserById(fb.Reviewer.ID)
			review.Feedback[i] = fb
		}
	}
	review.User = *rs.UserData.GetUserById(review.User.ID)
	rs.reviews[review.ID] = *review
	return nil
}

func (rs *MapData) updateReview(previous, review *Review) error {

	previous.IsActive = review.IsActive
	feedbackArray := previous.Feedback

	addFeedback, removeFeedback := feedbackChanges(previous.Feedback, review.Feedback)

	// disallow removing if inactive
	if !review.IsActive && len(removeFeedback) != 0 {
		return errors.New("Cannot remove reviewers on completed performance review")
	}

	// remove anything we identified to remove
	// this one's a bit slow
	for _, val := range removeFeedback {
		ind := -1
		for removeIndex, feedback := range feedbackArray {
			if feedback.Reviewer.ID == val {
				ind = removeIndex
				break
			}
		}
		if ind >= 0 {
			feedbackArray = remove(feedbackArray, ind)
		}
	}

	// add anything that was missing
	for _, val := range addFeedback {
		rs.feedbackCount += 1
		feedbackId := strconv.Itoa(rs.feedbackCount)
		feedbackArray = append(feedbackArray, Feedback{
			ID:         feedbackId,
			ReviewID:   previous.ID,
			ReviewerID: val,
			Reviewer:   *rs.UserData.GetUserById(val),
		})
	}
	previous.Feedback = feedbackArray
	rs.reviews[review.ID] = *previous
	return nil
}

// 'swap' from the last spot and then return everything but the last spot
func remove(s []Feedback, i int) []Feedback {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (rs *MapData) Start() {
	rs.reviews = make(map[string]Review)
}

func (rs *MapData) Stop() {
}
