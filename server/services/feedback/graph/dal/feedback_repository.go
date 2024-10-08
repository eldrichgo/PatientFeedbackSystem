package dal

import (
	"feedback/graph/model"

	"gorm.io/gorm"
)

type FeedbackRepository interface {
	CreateFeedback(feedback *model.Feedback) (*model.Feedback, error)
	GetFeedbacksBySurveyID(surveyID string) ([]*model.Feedback, error)
	GetAnswersByFeedbackID(feedbackID string) ([]*model.FeedbackAnswer, error)
}

type Feedbackrepo struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &Feedbackrepo{db: db}
}

func (f *Feedbackrepo) CreateFeedback(feedback *model.Feedback) (*model.Feedback, error) {
	if err := f.db.Create(feedback).Error; err != nil {
		return nil, err
	}

	return feedback, nil
}

func (f *Feedbackrepo) GetFeedbacksBySurveyID(surveyID string) ([]*model.Feedback, error) {
	var feedbacks []*model.Feedback

	if err := f.db.Where("survey_id = ?", surveyID).Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (f *Feedbackrepo) GetAnswersByFeedbackID(feedbackID string) ([]*model.FeedbackAnswer, error) {
	var answers []*model.FeedbackAnswer

	if err := f.db.Where("feedback_id = ?", feedbackID).Find(&answers).Error; err != nil {
		return nil, err
	}

	return answers, nil
}
