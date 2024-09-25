// Code generated by volcago. DO NOT EDIT.
// generated version: v1.11.1
package infrastructures

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	model "github.com/topgate/gcim-temporary/back/app/internal/volcago"
)

//go:generate ../../../../../bin/mockgen -source $GOFILE -destination mocks/user_session_gen.go

// UserSessionRepository - Repository of UserSession
type UserSessionRepository interface {
	// Single
	Get(ctx context.Context, id string, opts ...GetOption) (*model.UserSession, error)
	GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*model.UserSession, error)
	Insert(ctx context.Context, subject *model.UserSession) (_ string, err error)
	Update(ctx context.Context, subject *model.UserSession) (err error)
	StrictUpdate(ctx context.Context, id string, param *UserSessionUpdateParam, opts ...firestore.Precondition) error
	Delete(ctx context.Context, subject *model.UserSession, opts ...DeleteOption) (err error)
	DeleteByID(ctx context.Context, id string, opts ...DeleteOption) (err error)
	// Multiple
	GetMulti(ctx context.Context, ids []string, opts ...GetOption) ([]*model.UserSession, error)
	InsertMulti(ctx context.Context, subjects []*model.UserSession) (_ []string, er error)
	UpdateMulti(ctx context.Context, subjects []*model.UserSession) (er error)
	DeleteMulti(ctx context.Context, subjects []*model.UserSession, opts ...DeleteOption) (er error)
	DeleteMultiByIDs(ctx context.Context, ids []string, opts ...DeleteOption) (er error)
	// Single(Transaction)
	GetWithTx(tx *firestore.Transaction, id string, opts ...GetOption) (*model.UserSession, error)
	GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*model.UserSession, error)
	InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.UserSession) (_ string, err error)
	UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.UserSession) (err error)
	StrictUpdateWithTx(tx *firestore.Transaction, id string, param *UserSessionUpdateParam, opts ...firestore.Precondition) error
	DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.UserSession, opts ...DeleteOption) (err error)
	DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...DeleteOption) (err error)
	// Multiple(Transaction)
	GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...GetOption) ([]*model.UserSession, error)
	InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.UserSession) (_ []string, er error)
	UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.UserSession) (er error)
	DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.UserSession, opts ...DeleteOption) (er error)
	DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...DeleteOption) (er error)
	// Search
	Search(ctx context.Context, param *UserSessionSearchParam, q *firestore.Query) ([]*model.UserSession, error)
	SearchWithTx(tx *firestore.Transaction, param *UserSessionSearchParam, q *firestore.Query) ([]*model.UserSession, error)
	SearchByParam(ctx context.Context, param *UserSessionSearchParam) ([]*model.UserSession, *PagingResult, error)
	SearchByParamWithTx(tx *firestore.Transaction, param *UserSessionSearchParam) ([]*model.UserSession, *PagingResult, error)
	// misc
	GetCollection() *firestore.CollectionRef
	GetCollectionName() string
	GetDocRef(id string) *firestore.DocumentRef
	RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error)
}

// UserSessionRepositoryMiddleware - middleware of UserSessionRepository
type UserSessionRepositoryMiddleware interface {
	BeforeInsert(ctx context.Context, subject *model.UserSession) (bool, error)
	BeforeUpdate(ctx context.Context, old, subject *model.UserSession) (bool, error)
	BeforeDelete(ctx context.Context, subject *model.UserSession, opts ...DeleteOption) (bool, error)
	BeforeDeleteByID(ctx context.Context, ids []string, opts ...DeleteOption) (bool, error)
}

type userSessionRepository struct {
	collectionName   string
	firestoreClient  *firestore.Client
	middleware       []UserSessionRepositoryMiddleware
	uniqueRepository *uniqueRepository
}

// NewUserSessionRepository - constructor
func NewUserSessionRepository(firestoreClient *firestore.Client, middleware ...UserSessionRepositoryMiddleware) UserSessionRepository {
	return &userSessionRepository{
		collectionName:   "user_session",
		firestoreClient:  firestoreClient,
		middleware:       middleware,
		uniqueRepository: newUniqueRepository(firestoreClient, "user_session"),
	}
}

func (repo *userSessionRepository) setMeta(subject *model.UserSession, isInsert bool) {
	now := time.Now()

	if isInsert {
		subject.CreatedAt = now
	}
	subject.UpdatedAt = now
	subject.Version++
}

func (repo *userSessionRepository) setMetaWithStrictUpdate(param *UserSessionUpdateParam) {
	param.UpdatedAt = firestore.ServerTimestamp
	param.Version = firestore.Increment(1)
}

func (repo *userSessionRepository) beforeInsert(ctx context.Context, subject *model.UserSession) error {
	if subject.Version != 0 {
		return xerrors.Errorf("insert data must be Version == 0 %+v: %w", subject, ErrVersionConflict)
	}
	if subject.DeletedAt != nil {
		return xerrors.Errorf("insert data must be DeletedAt == nil: %+v", subject)
	}
	repo.setMeta(subject, true)
	repo.uniqueRepository.setMiddleware(ctx)
	err := repo.uniqueRepository.CheckUnique(ctx, nil, subject)
	if err != nil {
		return xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeInsert(ctx, subject)
		if err != nil {
			return xerrors.Errorf("beforeInsert.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return nil
}

func (repo *userSessionRepository) beforeUpdate(ctx context.Context, old, subject *model.UserSession) error {
	if ctx.Value(transactionInProgressKey{}) != nil && old == nil {
		var err error
		doc := repo.GetDocRef(subject.ID)
		old, err = repo.get(context.Background(), doc)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return ErrNotFound
			}
			return xerrors.Errorf("error in Get method: %w", err)
		}
	}
	if old.Version > subject.Version {
		return xerrors.Errorf(
			"The data in the database is newer: (db version: %d, target version: %d) %+v: %w",
			old.Version, subject.Version, subject, ErrVersionConflict,
		)
	}
	if subject.DeletedAt != nil {
		return xerrors.Errorf("update data must be DeletedAt == nil: %+v", subject)
	}
	repo.setMeta(subject, false)
	repo.uniqueRepository.setMiddleware(ctx)
	err := repo.uniqueRepository.CheckUnique(ctx, old, subject)
	if err != nil {
		return xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeUpdate(ctx, old, subject)
		if err != nil {
			return xerrors.Errorf("beforeUpdate.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return nil
}

func (repo *userSessionRepository) beforeDelete(ctx context.Context, subject *model.UserSession, opts ...DeleteOption) error {
	repo.setMeta(subject, false)
	repo.uniqueRepository.setMiddleware(ctx)
	err := repo.uniqueRepository.DeleteUnique(ctx, subject)
	if err != nil {
		return xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeDelete(ctx, subject, opts...)
		if err != nil {
			return xerrors.Errorf("beforeDelete.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return nil
}

// GetCollection - *firestore.CollectionRef getter
func (repo *userSessionRepository) GetCollection() *firestore.CollectionRef {
	return repo.firestoreClient.Collection(repo.collectionName)
}

// GetCollectionName - CollectionName getter
func (repo *userSessionRepository) GetCollectionName() string {
	return repo.collectionName
}

// GetDocRef - *firestore.DocumentRef getter
func (repo *userSessionRepository) GetDocRef(id string) *firestore.DocumentRef {
	return repo.GetCollection().Doc(id)
}

// RunInTransaction - (*firestore.Client).RunTransaction getter
func (repo *userSessionRepository) RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error) {
	return repo.firestoreClient.RunTransaction
}

// UserSessionSearchParam - params for search
type UserSessionSearchParam struct {
	ID        *QueryChainer
	UserID    *QueryChainer
	ExpiresAt *QueryChainer
	CreatedAt *QueryChainer
	CreatedBy *QueryChainer
	UpdatedAt *QueryChainer
	UpdatedBy *QueryChainer
	DeletedAt *QueryChainer
	DeletedBy *QueryChainer
	Version   *QueryChainer

	IncludeSoftDeleted bool
	CursorKey          string
	CursorLimit        int
}

// UserSessionUpdateParam - params for strict updates
type UserSessionUpdateParam struct {
	UserID    interface{}
	ExpiresAt interface{}
	CreatedAt interface{}
	CreatedBy interface{}
	UpdatedAt interface{}
	UpdatedBy interface{}
	DeletedAt interface{}
	DeletedBy interface{}
	Version   interface{}
}

// Search - search documents
// The third argument is firestore.Query, basically you can pass nil
func (repo *userSessionRepository) Search(ctx context.Context, param *UserSessionSearchParam, q *firestore.Query) ([]*model.UserSession, error) {
	return repo.search(ctx, param, q)
}

// SearchByParam - search documents by search param
func (repo *userSessionRepository) SearchByParam(ctx context.Context, param *UserSessionSearchParam) ([]*model.UserSession, *PagingResult, error) {
	return repo.searchByParam(ctx, param)
}

// Get - get `UserSession` by `UserSession.ID`
func (repo *userSessionRepository) Get(ctx context.Context, id string, opts ...GetOption) (*model.UserSession, error) {
	doc := repo.GetDocRef(id)
	return repo.get(ctx, doc, opts...)
}

// GetWithDoc - get `UserSession` by *firestore.DocumentRef
func (repo *userSessionRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*model.UserSession, error) {
	return repo.get(ctx, doc, opts...)
}

// Insert - insert of `UserSession`
func (repo *userSessionRepository) Insert(ctx context.Context, subject *model.UserSession) (_ string, err error) {
	if err := repo.beforeInsert(ctx, subject); err != nil {
		return "", xerrors.Errorf("before insert error: %w", err)
	}

	return repo.insert(ctx, subject)
}

// Update - update of `UserSession`
func (repo *userSessionRepository) Update(ctx context.Context, subject *model.UserSession) (err error) {
	doc := repo.GetDocRef(subject.ID)

	old, err := repo.get(ctx, doc)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return ErrNotFound
		}
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if err := repo.beforeUpdate(ctx, old, subject); err != nil {
		return xerrors.Errorf("before update error: %w", err)
	}

	return repo.update(ctx, subject)
}

// StrictUpdate - strict update of `UserSession`
func (repo *userSessionRepository) StrictUpdate(ctx context.Context, id string, param *UserSessionUpdateParam, opts ...firestore.Precondition) error {
	return repo.strictUpdate(ctx, id, param, opts...)
}

// Delete - delete of `UserSession`
func (repo *userSessionRepository) Delete(ctx context.Context, subject *model.UserSession, opts ...DeleteOption) (err error) {
	if err := repo.beforeDelete(ctx, subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(ctx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

	return repo.deleteByID(ctx, subject.ID)
}

// DeleteByID - delete `UserSession` by `UserSession.ID`
func (repo *userSessionRepository) DeleteByID(ctx context.Context, id string, opts ...DeleteOption) (err error) {
	subject, err := repo.Get(ctx, id)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if err := repo.beforeDelete(ctx, subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(ctx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

	return repo.Delete(ctx, subject, opts...)
}

// GetMulti - get `UserSession` in bulk by array of `UserSession.ID`
func (repo *userSessionRepository) GetMulti(ctx context.Context, ids []string, opts ...GetOption) ([]*model.UserSession, error) {
	return repo.getMulti(ctx, ids, opts...)
}

// InsertMulti - bulk insert of `UserSession`
func (repo *userSessionRepository) InsertMulti(ctx context.Context, subjects []*model.UserSession) (_ []string, er error) {

	ids := make([]string, 0, len(subjects))
	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		var ref *firestore.DocumentRef
		if subject.ID == "" {
			ref = collect.NewDoc()
			subject.ID = ref.ID
		} else {
			ref = collect.Doc(subject.ID)
			if s, err := ref.Get(ctx); err == nil {
				return nil, xerrors.Errorf("already exists [%v]: %#v", subject.ID, s)
			}
		}

		if err := repo.beforeInsert(ctx, subject); err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subject.ID, err)
		}

		batch.Set(ref, subject)
		ids = append(ids, ref.ID)
		i++
		if (i%500) == 0 && len(subjects) != i {
			batches = append(batches, batch)
			batch = repo.firestoreClient.Batch()
		}
	}
	batches = append(batches, batch)

	for _, b := range batches {
		if _, err := b.Commit(ctx); err != nil {
			return nil, xerrors.Errorf("error in Commit method: %w", err)
		}
	}

	return ids, nil
}

// UpdateMulti - bulk update of `UserSession`
func (repo *userSessionRepository) UpdateMulti(ctx context.Context, subjects []*model.UserSession) (er error) {

	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		ref := collect.Doc(subject.ID)
		snapShot, err := ref.Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.ID, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.ID, err)
		}

		old := new(model.UserSession)
		if err = snapShot.DataTo(&old); err != nil {
			return xerrors.Errorf("error in DataTo method: %w", err)
		}

		if err := repo.beforeUpdate(ctx, old, subject); err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subject.ID, err)
		}

		batch.Set(ref, subject)
		i++
		if (i%500) == 0 && len(subjects) != i {
			batches = append(batches, batch)
			batch = repo.firestoreClient.Batch()
		}
	}
	batches = append(batches, batch)

	for _, b := range batches {
		if _, err := b.Commit(ctx); err != nil {
			return xerrors.Errorf("error in Commit method: %w", err)
		}
	}

	return nil
}

// DeleteMulti - bulk delete of `UserSession`
func (repo *userSessionRepository) DeleteMulti(ctx context.Context, subjects []*model.UserSession, opts ...DeleteOption) (er error) {

	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		ref := collect.Doc(subject.ID)
		if _, err := ref.Get(ctx); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.ID, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.ID, err)
		}

		if err := repo.beforeDelete(ctx, subject, opts...); err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.ID, err)
		}

		if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
			t := time.Now()
			subject.DeletedAt = &t
			batch.Set(ref, subject)
		} else {
			batch.Delete(ref)
		}

		i++
		if (i%500) == 0 && len(subjects) != i {
			batches = append(batches, batch)
			batch = repo.firestoreClient.Batch()
		}
	}
	batches = append(batches, batch)

	for _, b := range batches {
		if _, err := b.Commit(ctx); err != nil {
			return xerrors.Errorf("error in Commit method: %w", err)
		}
	}

	return nil
}

// DeleteMultiByIDs - delete `UserSession` in bulk by array of `UserSession.ID`
func (repo *userSessionRepository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...DeleteOption) (er error) {
	subjects := make([]*model.UserSession, len(ids))

	opt := GetOption{}
	if len(opts) > 0 {
		opt.IncludeSoftDeleted = opts[0].Mode == DeleteModeHard
	}
	for i, id := range ids {
		subject, err := repo.Get(ctx, id, opt)
		if err != nil {
			return xerrors.Errorf("error in Get method: %w", err)
		}
		subjects[i] = subject
	}

	return repo.DeleteMulti(ctx, subjects, opts...)
}

// SearchWithTx - search documents in transaction
func (repo *userSessionRepository) SearchWithTx(tx *firestore.Transaction, param *UserSessionSearchParam, q *firestore.Query) ([]*model.UserSession, error) {
	return repo.search(tx, param, q)
}

// SearchByParamWithTx - search documents by search param in transaction
func (repo *userSessionRepository) SearchByParamWithTx(tx *firestore.Transaction, param *UserSessionSearchParam) ([]*model.UserSession, *PagingResult, error) {
	return repo.searchByParam(tx, param)
}

// GetWithTx - get `UserSession` by `UserSession.ID` in transaction
func (repo *userSessionRepository) GetWithTx(tx *firestore.Transaction, id string, opts ...GetOption) (*model.UserSession, error) {
	doc := repo.GetDocRef(id)
	return repo.get(tx, doc, opts...)
}

// GetWithDocWithTx - get `UserSession` by *firestore.DocumentRef in transaction
func (repo *userSessionRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*model.UserSession, error) {
	return repo.get(tx, doc, opts...)
}

// InsertWithTx - insert of `UserSession` in transaction
func (repo *userSessionRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.UserSession) (_ string, err error) {
	if err := repo.beforeInsert(context.WithValue(ctx, transactionInProgressKey{}, tx), subject); err != nil {
		return "", xerrors.Errorf("before insert error: %w", err)
	}

	return repo.insert(tx, subject)
}

// UpdateWithTx - update of `UserSession` in transaction
func (repo *userSessionRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.UserSession) (err error) {
	if err := repo.beforeUpdate(context.WithValue(ctx, transactionInProgressKey{}, tx), nil, subject); err != nil {
		return xerrors.Errorf("before update error: %w", err)
	}

	return repo.update(tx, subject)
}

// StrictUpdateWithTx - strict update of `UserSession` in transaction
func (repo *userSessionRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *UserSessionUpdateParam, opts ...firestore.Precondition) error {
	return repo.strictUpdate(tx, id, param, opts...)
}

// DeleteWithTx - delete of `UserSession` in transaction
func (repo *userSessionRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.UserSession, opts ...DeleteOption) (err error) {
	if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(tx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

	return repo.deleteByID(tx, subject.ID)
}

// DeleteByIDWithTx - delete `UserSession` by `UserSession.ID` in transaction
func (repo *userSessionRepository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...DeleteOption) (err error) {
	subject, err := repo.Get(context.Background(), id)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(tx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

	if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	return repo.deleteByID(tx, id)
}

// GetMultiWithTx - get `UserSession` in bulk by array of `UserSession.ID` in transaction
func (repo *userSessionRepository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...GetOption) ([]*model.UserSession, error) {
	return repo.getMulti(tx, ids, opts...)
}

// InsertMultiWithTx - bulk insert of `UserSession` in transaction
func (repo *userSessionRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.UserSession) (_ []string, er error) {

	for i := range subjects {
		if _, err := tx.Get(new(firestore.DocumentRef)); err == nil {
			return nil, xerrors.Errorf("already exists(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	ids := make([]string, len(subjects))

	for i := range subjects {
		if err := repo.beforeInsert(ctx, subjects[i]); err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subjects[i].ID, err)
		}

		id, err := repo.insert(tx, subjects[i])
		if err != nil {
			return nil, xerrors.Errorf("error in insert method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
		ids[i] = id
	}

	return ids, nil
}

// UpdateMultiWithTx - bulk update of `UserSession` in transaction
func (repo *userSessionRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.UserSession) (er error) {
	ctx = context.WithValue(ctx, transactionInProgressKey{}, tx)

	for i := range subjects {
		if err := repo.beforeUpdate(ctx, nil, subjects[i]); err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	for i := range subjects {
		if err := repo.update(tx, subjects[i]); err != nil {
			return xerrors.Errorf("error in update method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	return nil
}

// DeleteMultiWithTx - bulk delete of `UserSession` in transaction
func (repo *userSessionRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.UserSession, opts ...DeleteOption) (er error) {

	t := time.Now()
	var isHardDeleteMode bool
	if len(opts) > 0 {
		isHardDeleteMode = opts[0].Mode == DeleteModeHard
	}
	opt := GetOption{
		IncludeSoftDeleted: isHardDeleteMode,
	}
	for i := range subjects {
		dr := repo.GetDocRef(subjects[i].ID)
		if _, err := repo.get(context.Background(), dr, opt); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, subjects[i].ID)
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, subjects[i].ID, err)
		}

		if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subjects[i], opts...); err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subjects[i].ID, err)
		}

		if !isHardDeleteMode {
			subjects[i].DeletedAt = &t
			if err := repo.update(tx, subjects[i]); err != nil {
				return xerrors.Errorf("error in update method(%d) [%v]: %w", i, subjects[i].ID, err)
			}
		}
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		return nil
	}

	for i := range subjects {
		if err := repo.deleteByID(tx, subjects[i].ID); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	return nil
}

// DeleteMultiByIDWithTx - delete `UserSession` in bulk by array of `UserSession.ID` in transaction
func (repo *userSessionRepository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...DeleteOption) (er error) {

	t := time.Now()
	for i := range ids {
		dr := repo.GetDocRef(ids[i])
		subject, err := repo.get(context.Background(), dr)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, ids[i])
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, ids[i], err)
		}

		if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subject, opts...); err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.ID, err)
		}

		if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
			subject.DeletedAt = &t
			if err := repo.update(tx, subject); err != nil {
				return xerrors.Errorf("error in update method(%d) [%v]: %w", i, ids[i], err)
			}
		}
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		return nil
	}

	for i := range ids {
		if err := repo.deleteByID(tx, ids[i]); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, ids[i], err)
		}
	}

	return nil
}

func (repo *userSessionRepository) get(v interface{}, doc *firestore.DocumentRef, opts ...GetOption) (*model.UserSession, error) {
	var (
		snapShot *firestore.DocumentSnapshot
		err      error
	)

	switch x := v.(type) {
	case *firestore.Transaction:
		snapShot, err = x.Get(doc)
	case context.Context:
		snapShot, err = doc.Get(x)
	default:
		return nil, xerrors.Errorf("invalid type: %v", x)
	}

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, ErrNotFound
		}
		return nil, xerrors.Errorf("error in Get method: %w", err)
	}

	subject := new(model.UserSession)
	if err := snapShot.DataTo(&subject); err != nil {
		return nil, xerrors.Errorf("error in DataTo method: %w", err)
	}

	if len(opts) == 0 || !opts[0].IncludeSoftDeleted {
		if subject.DeletedAt != nil {
			return nil, ErrAlreadyDeleted
		}
	}
	subject.ID = snapShot.Ref.ID

	return subject, nil
}

func (repo *userSessionRepository) getMulti(v interface{}, ids []string, opts ...GetOption) ([]*model.UserSession, error) {
	var (
		snapShots []*firestore.DocumentSnapshot
		err       error
		collect   = repo.GetCollection()
		drs       = make([]*firestore.DocumentRef, len(ids))
	)

	for i, id := range ids {
		ref := collect.Doc(id)
		drs[i] = ref
	}

	switch x := v.(type) {
	case *firestore.Transaction:
		snapShots, err = x.GetAll(drs)
	case context.Context:
		snapShots, err = repo.firestoreClient.GetAll(x, drs)
	default:
		return nil, xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return nil, xerrors.Errorf("error in GetAll method: %w", err)
	}

	subjects := make([]*model.UserSession, 0, len(ids))
	mErr := NewMultiErrors()
	for i, snapShot := range snapShots {
		if !snapShot.Exists() {
			mErr = append(mErr, NewMultiError(i, ErrNotFound))
			continue
		}

		subject := new(model.UserSession)
		if err = snapShot.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		if len(opts) == 0 || !opts[0].IncludeSoftDeleted {
			if subject.DeletedAt != nil {
				mErr = append(mErr, NewMultiError(i, ErrLogicallyDeletedData))
				continue
			}
		}
		subject.ID = snapShot.Ref.ID
		subjects = append(subjects, subject)
	}

	if len(mErr) == 0 {
		return subjects, nil
	}

	return subjects, mErr
}

func (repo *userSessionRepository) insert(v interface{}, subject *model.UserSession) (string, error) {
	var (
		dr  = repo.GetDocRef(subject.ID)
		err error
	)

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Create(dr, subject)
	case context.Context:
		_, err = dr.Create(x, subject)
	default:
		return "", xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return "", xerrors.Errorf("error in Create method: err=%+v: %w", err, ErrAlreadyExists)
		}
		return "", xerrors.Errorf("error in Create method: %w", err)
	}

	subject.ID = dr.ID

	return dr.ID, nil
}

func (repo *userSessionRepository) update(v interface{}, subject *model.UserSession) error {
	var (
		dr  = repo.GetDocRef(subject.ID)
		err error
	)

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Set(dr, subject)
	case context.Context:
		_, err = dr.Set(x, subject)
	default:
		return xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return xerrors.Errorf("error in Set method: %w", err)
	}

	return nil
}

func (repo *userSessionRepository) strictUpdate(v interface{}, id string, param *UserSessionUpdateParam, opts ...firestore.Precondition) error {
	var (
		dr  = repo.GetDocRef(id)
		err error
	)

	repo.setMetaWithStrictUpdate(param)

	updates := updater(model.UserSession{}, param)

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Update(dr, updates, opts...)
	case context.Context:
		_, err = dr.Update(x, updates, opts...)
	default:
		return xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return xerrors.Errorf("error in Update method: %w", err)
	}

	return nil
}

func (repo *userSessionRepository) deleteByID(v interface{}, id string) error {
	dr := repo.GetDocRef(id)
	var err error

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Delete(dr, firestore.Exists)
	case context.Context:
		_, err = dr.Delete(x, firestore.Exists)
	default:
		return xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return xerrors.Errorf("error in Delete method: %w", err)
	}

	return nil
}

func (repo *userSessionRepository) runQuery(v interface{}, query firestore.Query) ([]*model.UserSession, error) {
	var iter *firestore.DocumentIterator

	switch x := v.(type) {
	case *firestore.Transaction:
		iter = x.Documents(query)
	case context.Context:
		iter = query.Documents(x)
	default:
		return nil, xerrors.Errorf("invalid type: %v", v)
	}

	defer iter.Stop()

	subjects := make([]*model.UserSession, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, xerrors.Errorf("error in Next method: %w", err)
		}

		subject := new(model.UserSession)

		if err = doc.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		subject.ID = doc.Ref.ID
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

// BUG(54m): there may be potential bugs
func (repo *userSessionRepository) searchByParam(v interface{}, param *UserSessionSearchParam) ([]*model.UserSession, *PagingResult, error) {
	query := func() firestore.Query {
		return repo.GetCollection().Query
	}()
	if param.ID != nil {
		for _, chain := range param.ID.QueryGroup {
			var value interface{}
			switch val := chain.Value.(type) {
			case string:
				value = repo.GetDocRef(val)
			case []string:
				docRefs := make([]*firestore.DocumentRef, len(val))
				for i := range val {
					docRefs[i] = repo.GetDocRef(val[i])
				}
				value = docRefs
			default:
				return nil, nil, xerrors.Errorf("document id can only be of type `string` and `[]string`. value: %#v", chain.Value)
			}
			query = query.Where(firestore.DocumentID, chain.Operator, value)
		}
		if direction := param.ID.OrderByDirection; direction > 0 {
			query = query.OrderBy(firestore.DocumentID, direction)
			query = param.ID.BuildCursorQuery(query)
		}
	}
	if param.UserID != nil {
		for _, chain := range param.UserID.QueryGroup {
			query = query.Where("user_id", chain.Operator, chain.Value)
		}
		if direction := param.UserID.OrderByDirection; direction > 0 {
			query = query.OrderBy("user_id", direction)
			query = param.UserID.BuildCursorQuery(query)
		}
	}
	if param.ExpiresAt != nil {
		for _, chain := range param.ExpiresAt.QueryGroup {
			query = query.Where("expires_at", chain.Operator, chain.Value)
		}
		if direction := param.ExpiresAt.OrderByDirection; direction > 0 {
			query = query.OrderBy("expires_at", direction)
			query = param.ExpiresAt.BuildCursorQuery(query)
		}
	}
	if param.CreatedAt != nil {
		for _, chain := range param.CreatedAt.QueryGroup {
			query = query.Where("created_at", chain.Operator, chain.Value)
		}
		if direction := param.CreatedAt.OrderByDirection; direction > 0 {
			query = query.OrderBy("created_at", direction)
			query = param.CreatedAt.BuildCursorQuery(query)
		}
	}
	if param.CreatedBy != nil {
		for _, chain := range param.CreatedBy.QueryGroup {
			query = query.Where("created_by", chain.Operator, chain.Value)
		}
		if direction := param.CreatedBy.OrderByDirection; direction > 0 {
			query = query.OrderBy("created_by", direction)
			query = param.CreatedBy.BuildCursorQuery(query)
		}
	}
	if param.UpdatedAt != nil {
		for _, chain := range param.UpdatedAt.QueryGroup {
			query = query.Where("updated_at", chain.Operator, chain.Value)
		}
		if direction := param.UpdatedAt.OrderByDirection; direction > 0 {
			query = query.OrderBy("updated_at", direction)
			query = param.UpdatedAt.BuildCursorQuery(query)
		}
	}
	if param.UpdatedBy != nil {
		for _, chain := range param.UpdatedBy.QueryGroup {
			query = query.Where("updated_by", chain.Operator, chain.Value)
		}
		if direction := param.UpdatedBy.OrderByDirection; direction > 0 {
			query = query.OrderBy("updated_by", direction)
			query = param.UpdatedBy.BuildCursorQuery(query)
		}
	}
	if param.DeletedAt != nil {
		for _, chain := range param.DeletedAt.QueryGroup {
			query = query.Where("deleted_at", chain.Operator, chain.Value)
		}
		if direction := param.DeletedAt.OrderByDirection; direction > 0 {
			query = query.OrderBy("deleted_at", direction)
			query = param.DeletedAt.BuildCursorQuery(query)
		}
	}
	if param.DeletedBy != nil {
		for _, chain := range param.DeletedBy.QueryGroup {
			query = query.Where("deleted_by", chain.Operator, chain.Value)
		}
		if direction := param.DeletedBy.OrderByDirection; direction > 0 {
			query = query.OrderBy("deleted_by", direction)
			query = param.DeletedBy.BuildCursorQuery(query)
		}
	}
	if param.Version != nil {
		for _, chain := range param.Version.QueryGroup {
			query = query.Where("version", chain.Operator, chain.Value)
		}
		if direction := param.Version.OrderByDirection; direction > 0 {
			query = query.OrderBy("version", direction)
			query = param.Version.BuildCursorQuery(query)
		}
	}
	if !param.IncludeSoftDeleted {
		query = query.Where("deleted_at", OpTypeEqual, nil)
	}

	limit := param.CursorLimit + 1

	if param.CursorKey != "" {
		var (
			ds  *firestore.DocumentSnapshot
			err error
		)
		switch x := v.(type) {
		case *firestore.Transaction:
			ds, err = x.Get(repo.GetDocRef(param.CursorKey))
		case context.Context:
			ds, err = repo.GetDocRef(param.CursorKey).Get(x)
		default:
			return nil, nil, xerrors.Errorf("invalid x type: %v", v)
		}
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return nil, nil, ErrNotFound
			}
			return nil, nil, xerrors.Errorf("error in Get method: %w", err)
		}
		query = query.StartAt(ds)
	}

	if limit > 1 {
		query = query.Limit(limit)
	}

	subjects, err := repo.runQuery(v, query)
	if err != nil {
		return nil, nil, xerrors.Errorf("error in runQuery method: %w", err)
	}

	pagingResult := &PagingResult{
		Length: len(subjects),
	}
	if limit > 1 && limit == pagingResult.Length {
		next := pagingResult.Length - 1
		pagingResult.NextCursorKey = subjects[next].ID
		subjects = subjects[:next]
		pagingResult.Length--
	}

	return subjects, pagingResult, nil
}

func (repo *userSessionRepository) search(v interface{}, param *UserSessionSearchParam, q *firestore.Query) ([]*model.UserSession, error) {
	if (param == nil && q == nil) || (param != nil && q != nil) {
		return nil, xerrors.New("either one should be nil")
	}

	query := func() firestore.Query {
		if q != nil {
			return *q
		}
		return repo.GetCollection().Query
	}()

	if q == nil {
		subjects, _, err := repo.searchByParam(v, param)
		if err != nil {
			return nil, xerrors.Errorf("error in searchByParam method: %w", err)
		}

		return subjects, nil
	}

	return repo.runQuery(v, query)
}
