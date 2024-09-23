package dal

import (
	"survey/graph/model"

	"gorm.io/gorm"
)

type SurveyRepository interface {
	CreateSurvey(survey *model.Survey) (*model.Survey, error)
	GetSurveys() ([]*model.Survey, error)
	GetSurveyByID(id string) (*model.Survey, error)
	GetQuestionsBySurveyID(surveyID string) ([]*model.Question, error)
	GetOptionsByQuestionID(questionID string) ([]*model.Option, error)

	UpdateSurvey(survey *model.Survey) (*model.Survey, error)
	DeleteQuestionsBySurveyID(surveyID string) error
	DeleteOptionsByQuestionID(questionID string) error
}

type Surveyrepo struct {
	db *gorm.DB
}

func NewSurveyRepository(db *gorm.DB) SurveyRepository {
	return &Surveyrepo{db: db}
}

func (s *Surveyrepo) CreateSurvey(survey *model.Survey) (*model.Survey, error) {
	if err := s.db.Create(survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

func (s *Surveyrepo) GetSurveys() ([]*model.Survey, error) {
	var surveys []*model.Survey

	// Implement soft delete
	if err := s.db.Find(&surveys).Error; err != nil {
		return nil, err
	}

	return surveys, nil
}

func (s *Surveyrepo) GetSurveyByID(id string) (*model.Survey, error) {
	var survey model.Survey

	// Implement soft delete
	if err := s.db.Where("id = ?", id).First(&survey).Error; err != nil {
		return nil, err
	}

	return &survey, nil
}

func (s *Surveyrepo) CreateQuestion(question *model.Question) (*model.Question, error) {
	if err := s.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (s *Surveyrepo) GetQuestionsBySurveyID(surveyID string) ([]*model.Question, error) {
	var questions []*model.Question
	// Implement soft delete
	if err := s.db.Where("survey_id = ?", surveyID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *Surveyrepo) GetOptionsByQuestionID(questionID string) ([]*model.Option, error) {
	var options []*model.Option

	// Implement soft delete
	if err := s.db.Where("question_id = ?", questionID).Find(&options).Error; err != nil {
		return nil, err
	}

	return options, nil
}

func (s *Surveyrepo) UpdateSurvey(survey *model.Survey) (*model.Survey, error) {
	if err := s.db.Save(survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

func (s *Surveyrepo) DeleteQuestionsBySurveyID(surveyID string) error {
	// Implement soft delete
	if err := s.db.Where("survey_id = ?", surveyID).Delete(&model.Question{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *Surveyrepo) DeleteOptionsByQuestionID(questionID string) error {
	// Implement soft delete
	if err := s.db.Where("question_id = ?", questionID).Delete(&model.Option{}).Error; err != nil {
		return err
	}

	return nil
}
