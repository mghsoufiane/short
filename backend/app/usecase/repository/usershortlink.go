package repository

import "github.com/short-d/short/backend/app/entity"

// UserShortLink accesses User-ShortLink relationship from storage, such as database.
type UserShortLink interface {
	CreateRelation(user entity.User, shortLink entity.ShortLink) error
	FindAliasesByUser(user entity.User) ([]string, error)
	IsShortLinkRelated(alias string, user entity.User) (bool, error)
}