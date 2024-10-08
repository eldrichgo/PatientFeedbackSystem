package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"feedback/graph/dal"
	"feedback/graph/model"
)

// Answers is the resolver for the answers field.
func (r *feedbackResolver) Answers(ctx context.Context, obj *model.Feedback) ([]*model.FeedbackAnswer, error) {
	svc := dal.NewFeedbackService(dal.NewFeedbackRepository(r.Db))
	answers, err := svc.GetAnswersByFeedbackID(obj.ID)
	if err != nil {
		return nil, err
	}

	return answers, nil
}

// SubmitFeedback is the resolver for the submitFeedback field.
func (r *mutationResolver) SubmitFeedback(ctx context.Context, input model.SubmitFeedbackInput) (*model.Feedback, error) {
	svc := dal.NewFeedbackService(dal.NewFeedbackRepository(r.Db))
	feedback, err := svc.CreateFeedback(ctx, &input)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

// Feedback is the resolver for the feedback field.
func (r *queryResolver) Feedback(ctx context.Context, surveyID string) ([]*model.Feedback, error) {
	svc := dal.NewFeedbackService(dal.NewFeedbackRepository(r.Db))
	feedback, err := svc.GetFeedback(surveyID)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

// Feedback returns FeedbackResolver implementation.
func (r *Resolver) Feedback() FeedbackResolver { return &feedbackResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type feedbackResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
