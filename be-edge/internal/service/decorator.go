package service

import "be-edge/internal/domain"

func applyTitleValidation(post *domain.Post) *domain.Post {
	if len(post.Title) < 5 {
		post.Title = "Título muy corto"
	}
	return post
}
