package pr

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (rs *ReviewService) getAllReviews() []Review {
	var results []Review
	rows, err := rs.db.Queryx(`
		select r.*,
					 u.id as "user.id",
					 u.name as "user.name"
			from reviews r
			join users u on u.id = r.userid
	`)

	if err != nil {
		log.Printf("Could not get all the reviews\n%+v\n", err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var review Review
		err = rows.StructScan(&review)
		if err != nil {
			log.Print(err)
		}
		results = append(results, review)
	}

	if results == nil {
		results = make([]Review, 0)
	}

	return results
}

func (rs *ReviewService) getReviewById(reviewId string) *Review {
	var result Review
	pid, err := strconv.ParseInt(reviewId, 10, 64)
	if err != nil {
		log.Printf("Couldn't parse int for getting\n%+v\n", err)
		return nil
	}
	rows, err := rs.db.Queryx(`
		select r.*,
					 u.id as "user.id",
					 u.name as "user.name",
					 u.username as "user.username"
			from reviews r
			join users u on u.id = r.userid
			where r.id = $1
	`, pid)

	if err != nil {
		log.Printf("Could not get all the reviews\n%+v\n", err)
		return nil
	}

	defer rows.Close()
	var review Review
	for rows.Next() {
		err = rows.StructScan(&review)
		if err != nil {
			log.Print(err)
			continue
		}
		result = review
	}
	if len(result.ID) != 0 {
		result.Feedback = []Feedback{}
		rows2, err := rs.db.Queryx(`
			select fb.*,
					u.id as "reviewer.id",
					u.name as "reviewer.name",
					u.username as "reviewer.username"
				from reviews_feedback fb
				join users u on u.id = fb.reviewerid
				where reviewid = $1
		`, pid)

		if err != nil {
			log.Printf("Could not get all the feedback\n%+v\n", err)
			return nil
		}

		for rows2.Next() {
			var feedback Feedback
			err = rows2.StructScan(&feedback)
			if err != nil {
				log.Print(err)
				continue
			}
			result.Feedback = append(result.Feedback, feedback)
		}
	}

	if err != nil {
		log.Printf("Could not get the user\n%+v\n", err)
		return nil
	}

	return &result
}

func (rs *ReviewService) getAllFeedbackForReviewer(userID string) []FlatFeedback {
	var results []FlatFeedback
	rows, err := rs.db.Queryx(`
		select fb.id as "id",
					 r.id as "reviewid",
					 fb.reviewerid as "reviewerid",
					 r.userid as "userid",
					 u.username as "username",
					 u.name as "name",
					 fb.*
			from reviews r
			join users u on u.id = r.userid
			join reviews_feedback fb on fb.reviewid = r.id
			where fb.reviewerid = $1 and r.isactive = true
	`, userID)

	if err != nil {
		log.Printf("Could not get all the reviews\n%+v\n", err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var review FlatFeedback
		err = rows.StructScan(&review)
		if err != nil {
			log.Print(err)
		}
		results = append(results, review)
	}

	if results == nil {
		results = make([]FlatFeedback, 0)
	}

	return results
}

func (rs *ReviewService) updateFeedback(updatedReview *FlatFeedback) error {
	_, err := rs.db.Exec(`
		update reviews_feedback
			set message=$1
			where id=$2
	`, updatedReview.Message.String, updatedReview.ID)
	if err != nil {
		log.Printf("Error updating feedback \n%+v\n", err)
		return err
	}
	return nil
}

func (rs *ReviewService) getFeedbackForReviewer(userID, idToUpdate string) *FlatFeedback {
	var results FlatFeedback
	rows, err := rs.db.Queryx(`
		select fb.id as "id",
					 r.id as "reviewid",
					 fb.reviewerid as "reviewerid",
					 r.userid as "userid",
					 u.username as "username",
					 u.name as "name",
					 fb.*
			from reviews r
			join users u on u.id = r.userid
			join reviews_feedback fb on fb.reviewid = r.id
			where fb.reviewerid = $1 and r.isactive = true and fb.id = $2
	`, userID, idToUpdate)

	if err != nil {
		log.Printf("Could not get all the reviews\n%+v\n", err)
		return nil
	}

	var review FlatFeedback
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&review)
		if err != nil {
			log.Print(err)
		}
		results = review
	}

	return &results
}

func (rs *ReviewService) create(review *Review) error {
	m := map[string]interface{}{"userid": review.UserID, "isActive": review.IsActive}
	sql := `
		with ids as (
		insert into reviews
			(userid, isActive)
		values
			(:userid, :isActive)
		returning id
		)
	`
	if review.Feedback != nil {
		for ind, feedback := range review.Feedback {
			sql += fmt.Sprintf(`
			insert into reviews_feedback(reviewid, reviewerid)
			select id, :reviewer%v from ids`, ind)
			m[fmt.Sprintf("reviewer%v", ind)] = feedback.ReviewerID
		}
	}
	_, err := rs.db.NamedExec(sql, m)

	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}

func (rs *ReviewService) updateReview(previous, review *Review) error {
	// this one will be gross
	m := map[string]interface{}{"id": review.ID, "isactive": review.IsActive}
	tx, err := rs.db.Beginx()
	if err != nil {
		return err
	}
	sql := `update reviews
		set isactive=:isactive
		where id=:id`

	_, err = tx.NamedExec(sql, m)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	var reviewerToAdd []string
	var reviewerToRemove []string
	// We have to determine what needs to be added or deleted
	feedbackMap := make(map[string]bool, max(len(previous.Feedback), len(review.Feedback)))
	// first set everything in the map
	for _, prevFB := range previous.Feedback {
		feedbackMap[prevFB.ReviewerID] = false
	}
	// if it's not in the map, we need to add it else mark it as true
	for _, newFB := range review.Feedback {
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

	// disallow removing if inactive
	if !review.IsActive && len(reviewerToRemove) != 0 {
		_ = tx.Rollback()
		return errors.New("Cannot remove reviewers on completed performance review")
	}

	// add anything that was missing
	for _, val := range reviewerToAdd {
		_, err := tx.Exec(`
		insert into reviews_feedback 
            (reviewid, reviewerid) 
		values 
			($1, $2)
		`, review.ID, val)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	// remove anything we identified to remove
	for _, val := range reviewerToRemove {
		_, err := tx.Exec(`
			delete from reviews_feedback 
      where reviewid = $1 and reviewerid = $2
		`, review.ID, val)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	return nil
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (rs *ReviewService) Start(db *sqlx.DB) {
	rs.db = db
}

func (rs *ReviewService) Stop() {
	if rs.db != nil {
		rs.db.Close()
	}
}
