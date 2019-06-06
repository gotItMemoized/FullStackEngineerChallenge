package pr

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ReviewService struct {
	db *sqlx.DB
}

func (rs *ReviewService) All(w http.ResponseWriter, r *http.Request) {
	// TODO: get all
	// maybe add filters
	reviews := rs.getAllReviews()

	if reviews == nil {
		http.Error(w, "[]", http.StatusNotFound)
		return
	}

	writeToOutput(w, reviews)
}

func (rs *ReviewService) Get(w http.ResponseWriter, r *http.Request) {
	reviewID := chi.URLParam(r, "id")
	review := rs.getReviewById(reviewID)

	if review == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	writeToOutput(w, review)
}

func (rs *ReviewService) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: create
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

	// TODO: check for existing active performance reviews for this user

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

	err = rs.create(&review)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not save performance review"))
		http.Error(w, "Error creating performance review", http.StatusBadRequest)
		return
	}

	// in some cases it'd be nice to return the userid on create,
	// but in this case we'll skip that
	w.WriteHeader(http.StatusNoContent)
}

func (rs *ReviewService) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: updates
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

	existingReview := rs.getReviewById(idToUpdate)
	if existingReview == nil {
		log.Print(errors.Wrap(err, "Could not update review"))
		http.Error(w, "Could not update review", http.StatusNotFound)
		return
	}

	if existingReview.User.ID != updatedReview.User.ID {
		log.Print(errors.New("Trieid to change the target reviewee"))
		http.Error(w, "Could not update review", http.StatusBadRequest)
		return
	}

	err = rs.updateReview(existingReview, &updatedReview)
	if err != nil {
		log.Print(errors.Wrap(err, "Could not update review"))
		http.Error(w, "Could not update review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *ReviewService) Done(w http.ResponseWriter, r *http.Request) {
	// TODO: sets reviews as done
	// doneReviewId := chi.URLParam(r, "id")

	// parsedReviewId, err := strconv.ParseInt(doneReviewId, 10, 64)
	// if err != nil {
	// 	log.Printf("Couldn't parse int for marking as done\n%+v\n", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// _, claims, _ := jwtauth.FromContext(r.Context())
	// userID := claims["id"].(string)
	// if len(userID) == 0 {
	// 	http.Error(w, "Could not confirm user", http.StatusUnauthorized)
	// 	return
	// }
	// parsedUserId, err := strconv.ParseInt(userID, 10, 64)
	// if err != nil {
	// 	log.Printf("Couldn't parse int for deleting\n%+v\n", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// if parsedUserId == parsedDeleteId {
	// 	log.Printf("User tried to delete themselves\n%+v\n", err)
	// 	http.Error(w, "Cannot delete yourself", http.StatusUnauthorized)
	// }

	// err = u.delete(parsedDeleteId)
	// if err != nil {
	// 	log.Printf("Couldn't delete user\n%+v\n", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
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
