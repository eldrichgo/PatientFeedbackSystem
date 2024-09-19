package dal

import (
	"context"
	"errors"
	"feedback/graph/model"
)

type FeedbackService struct {
	repo FeedbackRepository
}

func NewFeedbackService(repo FeedbackRepository) *FeedbackService {
	return &FeedbackService{repo: repo}
}

func (f *FeedbackService) CreateFeedback(ctx context.Context, input *model.SubmitFeedbackInput) (*model.Feedback, error) {
	// Validate the feedback
	if input.SurveyID == "" {
		return nil, errors.New("survey ID cannot be blank")
	}

	if input.UserID == "" {
		return nil, errors.New("user ID cannot be blank")
	}

	if len(input.Answers) == 0 {
		return nil, errors.New("feedback must have at least one answer")
	}

	feedback := &model.Feedback{
		SurveyID: input.SurveyID,
		UserID:   input.UserID,
	}

	for _, answer := range input.Answers {
		feedback.Answers = append(feedback.Answers, &model.FeedbackAnswer{
			QuestionID: answer.QuestionID,
			Answer:     answer.Answer,
		})
	}

	if _, err := f.repo.CreateFeedback(feedback); err != nil {
		return nil, err
	}

	return feedback, nil
}
