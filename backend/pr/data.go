package pr

type Data interface {
	getAllReviews() []Review
	getReviewById(reviewId string) *Review
	getAllFeedbackForReviewer(userID string) []FlatFeedback
	updateFeedback(updatedReview *FlatFeedback) error
	getFeedbackForReviewer(userID, idToUpdate string) *FlatFeedback
	create(review *Review) error
	updateReview(previous, review *Review) error
	Start()
	Stop()
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// determines the differences in the two feedback arrays
// returns an array of user ids that were added and an array of user ids that were removed
func feedbackChanges(previous, current []Feedback) ([]string, []string) {
	var reviewerToAdd []string
	var reviewerToRemove []string
	// We have to determine what needs to be added or deleted
	feedbackMap := make(map[string]bool, max(len(previous), len(current)))
	// first set everything in the map
	for _, prevFB := range previous {
		feedbackMap[prevFB.ReviewerID] = false
	}
	// if it's not in the map, we need to add it else mark it as true
	for _, newFB := range current {
		_, ok := feedbackMap[newFB.Reviewer.ID]
		if !ok {
			reviewerToAdd = append(reviewerToAdd, newFB.Reviewer.ID)
		} else {
			feedbackMap[newFB.Reviewer.ID] = true
		}
	}

	// anything still false in the map will need to be removed
	for id, wasFound := range feedbackMap {
		if !wasFound {
			reviewerToRemove = append(reviewerToRemove, id)
		}
	}

	return reviewerToAdd, reviewerToRemove
}
