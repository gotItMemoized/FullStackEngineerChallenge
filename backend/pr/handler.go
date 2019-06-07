package pr

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"
)

type ReviewHandler struct {
	Data Data
}

func (rs *ReviewHandler) All(w http.ResponseWriter, r *http.Request) {

	reviews := rs.Data.getAllReviews()

	if reviews == nil {
		http.Error(w, "[]", http.StatusNotFound)
		return
	}

	writeToOutput(w, reviews)
}

func (rs *ReviewHandler) Get(w http.ResponseWriter, r *http.Request) {
	reviewID := chi.URLParam(r, "id")
	review := rs.Data.getReviewById(reviewID)

	if review == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeToOutput(w, review)
}

func (rs *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var review Review
	err := decoder.Decode(&review)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not decode user"))
		http.Error(w, "Error creating performance review", http.StatusBadRequest)
		return
	}

	if len(review.User.ID) == 0 {
		log.Print(errors.New("Invalid id for performance review"))
		http.Error(w, "Invalid user for performance review", http.StatusBadRequest)
		return
	}

	review.UserID = review.User.ID
	for ind, feedback := range review.Feedback {
		if len(feedback.Reviewer.ID) == 0 {
			log.Print(errors.New("Invalid feedback user id for performance review"))
			http.Error(w, "Invalid feedback user for performance review", http.StatusBadRequest)
			return
		}
		feedback.ReviewerID = feedback.Reviewer.ID
		review.Feedback[ind] = feedback
	}

	err = rs.Data.create(&review)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not save performance review"))
		http.Error(w, "Error creating performance review", http.StatusBadRequest)
		return
	}

	// in some cases it'd be nice to return the userid on create,
	// but in this case we'll skip that
	w.WriteHeader(http.StatusNoContent)
}

func (rs *ReviewHandler) Update(w http.ResponseWriter, r *http.Request) {
	idToUpdate := chi.URLParam(r, "id")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var updatedReview Review
	err := decoder.Decode(&updatedReview)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not decode review"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedReview.ID = idToUpdate

	existingReview := rs.Data.getReviewById(idToUpdate)
	if existingReview == nil {
		log.Print(errors.Wrap(err, "Could not update review"))
		http.Error(w, "Could not update review", http.StatusNotFound)
		return
	}

	err = rs.Data.updateReview(existingReview, &updatedReview)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not update review"))
		http.Error(w, "Could not update review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *ReviewHandler) GetPendingFeedbackForReviewer(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["id"].(string)
	if len(userID) == 0 {
		http.Error(w, "Could not confirm user", http.StatusUnauthorized)
		return
	}
	reviews := rs.Data.getAllFeedbackForReviewer(userID)

	if reviews == nil {
		http.Error(w, "[]", http.StatusNotFound)
		return
	}

	writeToOutput(w, reviews)
}

func (rs *ReviewHandler) GetFeedback(w http.ResponseWriter, r *http.Request) {
	feedbackID := chi.URLParam(r, "id")
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["id"].(string)
	if len(userID) == 0 {
		http.Error(w, "Could not confirm user", http.StatusUnauthorized)
		return
	}
	feedback := rs.Data.getFeedbackForReviewer(userID, feedbackID)

	if feedback == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeToOutput(w, feedback)
}

func (rs *ReviewHandler) GiveFeedback(w http.ResponseWriter, r *http.Request) {
	idToUpdate := chi.URLParam(r, "id")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var updatedReview FlatFeedback
	err := decoder.Decode(&updatedReview)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not decode review"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedReview.ID = idToUpdate

	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["id"].(string)
	if len(userID) == 0 {
		http.Error(w, "Could not confirm user", http.StatusUnauthorized)
		return
	}

	if !updatedReview.Message.Valid || len(updatedReview.Message.String) == 0 {
		log.Print(errors.New("Empty or invalid message"))
		http.Error(w, "Empty or invalid message", http.StatusBadRequest)
		return
	}

	existingReview := rs.Data.getFeedbackForReviewer(userID, idToUpdate)
	if existingReview == nil {
		log.Print(errors.Wrap(err, "Could not update feedback"))
		http.Error(w, "Could not update feedback", http.StatusNotFound)
		return
	}

	err = rs.Data.updateFeedback(&updatedReview)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not update review"))
		http.Error(w, "Could not update review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// this is technically inefficient, but allows for fast iterations and we can still get very fast responses locally
func writeToOutput(w http.ResponseWriter, object interface{}) {
	output, err := json.Marshal(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(output)
	if err != nil {
		log.Printf("error while writing output: %+v\n", err)
	}
}
