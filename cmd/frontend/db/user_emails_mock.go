package db

import (
	"context"
)

type MockUserEmails struct {
	GetPrimaryEmail    func(ctx context.Context, id int32) (email string, verified bool, err error)
	Get                func(userID int32, email string) (emailCanonicalCase string, verified bool, err error)
	SetVerified        func(ctx context.Context, userID int32, email string, verified bool) error
	GetVerifiedEmails  func(ctx context.Context, emails ...string) ([]*UserEmail, error)
	ListByUser         func(id int32) ([]*UserEmail, error)
	ListVerifiedByUser func(id int32) ([]*UserEmail, error)
}
