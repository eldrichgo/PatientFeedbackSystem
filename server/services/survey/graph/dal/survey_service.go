package dal

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"survey/graph/model"
)

type SurveyService struct {
	repo SurveyRepository
}

func NewSurveyService(repo SurveyRepository) *SurveyService {
	return &SurveyService{repo: repo}
}

func (s *SurveyService) CreateSurvey(ctx context.Context, input *model.SurveyInput) (*model.Survey, error) {
	// Validate the survey
	if input.Name == "" {
		return nil, errors.New("survey name cannot be blank")
	}

	if len(input.Questions) == 0 {
		return nil, errors.New("survey must have at least one question")
	}

	survey := &model.Survey{
		Name:        input.Name,
		Description: input.Description,
	}

	// Add questions and their options to the survey
	for _, question := range input.Questions {
		survey.Questions = append(survey.Questions, &model.Question{
			QuestionText: strings.TrimSpace(question.QuestionText),
			Options: func(o []*model.OptionInput) []*model.Option {
				var options []*model.Option

				for _, option := range o {
					options = append(options, &model.Option{
						OptionText: option.OptionText,
					})
				}

				return options
			}(question.Options),
		})
	}

	return s.repo.CreateSurvey(survey)
}

func (s *SurveyService) GetSurveys() ([]*model.Survey, error) {
	return s.repo.GetSurveys()
}

func (s *SurveyService) GetSurveyByID(surveyID string) (*model.Survey, error) {
	if surveyID == "" {
		return nil, errors.New("survey ID cannot be blank")
	}

	return s.repo.GetSurveyByID(surveyID)
}

func (s *SurveyService) GetQuestionsBySurveyID(surveyID string) ([]*model.Question, error) {
	if surveyID == "" {
		return nil, errors.New("survey ID cannot be blank")
	}

	return s.repo.GetQuestionsBySurveyID(surveyID)
}

func (s *SurveyService) GetOptionsByQuestionID(questionID string) ([]*model.Option, error) {
	if questionID == "" {
		return nil, errors.New("question ID cannot be blank")
	}

	return s.repo.GetOptionsByQuestionID(questionID)
}

func (s *SurveyService) UpdateSurvey(ctx context.Context, surveyID string, input *model.SurveyInput) (*model.Survey, error) {
	if surveyID == "" {
		return nil, errors.New("survey ID cannot be blank")
	}

	if input.Name == "" {
		return nil, errors.New("survey name cannot be blank")
	}

	if len(input.Questions) == 0 {
		return nil, errors.New("survey must have at least one question")
	}

	survey, err := s.repo.GetSurveyByID(surveyID)
	if err != nil {
		return nil, err
	}

	survey.Name = input.Name
	survey.Description = input.Description

	// Delete all questions
	// Make sure its soft delete
	if err := s.repo.DeleteQuestionsBySurveyID(surveyID); err != nil {
		return nil, err
	}

	// Delete all options
	// Make sure its soft delete
	if err := s.repo.DeleteOptionsByQuestionID(surveyID); err != nil {
		return nil, err
	}

	// Clear the questions slice in the survey object
	survey.Questions = nil

	// Add new questions and their options to the survey
	for _, question := range input.Questions {
		survey.Questions = append(survey.Questions, &model.Question{
			QuestionText: strings.TrimSpace(question.QuestionText),
			Options: func(o []*model.OptionInput) []*model.Option {
				var options []*model.Option

				for _, option := range o {
					options = append(options, &model.Option{
						OptionText: option.OptionText,
					})
				}

				return options
			}(question.Options),
		})
		// debug and show me question and its options in console
		fmt.Println("Question: ", question.QuestionText)
		for _, option := range question.Options {
			fmt.Println("Option: ", option.OptionText)
		}
	}

	return s.repo.UpdateSurvey(survey)
}
