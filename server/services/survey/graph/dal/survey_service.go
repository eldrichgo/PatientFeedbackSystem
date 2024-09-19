package dal

import (
	"context"
	"errors"
	"strings"
	"survey/graph/model"
)

type SurveyService struct {
	repo SurveyRepository
}

func NewSurveyService(repo SurveyRepository) *SurveyService {
	return &SurveyService{repo: repo}
}

func ProcessQuestions(questions []string) []string {
	processedQuestions := []string{}

	for _, question := range questions {
		if question != "" {
			if strings.TrimSpace(question) != "" {
				processedQuestions = append(processedQuestions, question)
			}
		}
	}

	return processedQuestions
}

func (s *SurveyService) CreateSurvey(ctx context.Context, input *model.CreateSurveyInput) (*model.Survey, error) {
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

	for _, question := range input.Questions {
		survey.Questions = append(survey.Questions, &model.Question{
			QuestionText: strings.Trim(question.QuestionText, " "),
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
