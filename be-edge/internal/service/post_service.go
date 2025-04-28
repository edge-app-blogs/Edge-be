package service

import (
	"be-edge/internal/domain"
	"be-edge/internal/repository"
	"errors"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *domain.Post) error {
	post = applyTitleValidation(post)
	return s.repo.Create(post)
}

func (s *PostService) GetPosts() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) UpdatePost(post *domain.Post) error {
	post = applyTitleValidation(post)
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.Delete(id)
}

func (s *PostService) GetPostByID(id int) (*domain.Post, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, errors.New("post not found")
	}

	return post, nil
}
