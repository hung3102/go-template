package repositories

import (
	"context"
	"fmt"
	"gcim/example/internal/domain/model"
	"gcim/example/internal/domain/repositories"
	"log"

	"cloud.google.com/go/firestore"
)

type TaskRepository struct {
	client *firestore.Client
}

func NewTaskRepository(c *firestore.Client) repositories.TaskRepository {
	return &TaskRepository{client: c}
}

// Delete implements repositories.TaskRepository.
func (r *TaskRepository) Delete(ctx context.Context, subject *model.Task, opts ...repositories.DeleteOption) (err error) {
	panic("unimplemented")
}

// DeleteByID implements repositories.TaskRepository.
func (r *TaskRepository) DeleteByID(ctx context.Context, id string, opts ...repositories.DeleteOption) (err error) {
	panic("unimplemented")
}

// DeleteByIDWithTx implements repositories.TaskRepository.
func (r *TaskRepository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...repositories.DeleteOption) (err error) {
	panic("unimplemented")
}

// DeleteMulti implements repositories.TaskRepository.
func (r *TaskRepository) DeleteMulti(ctx context.Context, subjects []*model.Task, opts ...repositories.DeleteOption) (er error) {
	panic("unimplemented")
}

// DeleteMultiByIDs implements repositories.TaskRepository.
func (r *TaskRepository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...repositories.DeleteOption) (er error) {
	panic("unimplemented")
}

// DeleteMultiByIDsWithTx implements repositories.TaskRepository.
func (r *TaskRepository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...repositories.DeleteOption) (er error) {
	panic("unimplemented")
}

// DeleteMultiWithTx implements repositories.TaskRepository.
func (r *TaskRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.Task, opts ...repositories.DeleteOption) (er error) {
	panic("unimplemented")
}

// DeleteWithTx implements repositories.TaskRepository.
func (r *TaskRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.Task, opts ...repositories.DeleteOption) (err error) {
	panic("unimplemented")
}

// Get implements repositories.TaskRepository.
func (r *TaskRepository) Get(ctx context.Context, id string, opts ...repositories.GetOption) (*model.Task, error) {
	panic("unimplemented")
}

// GetByDesc implements repositories.TaskRepository.
func (r *TaskRepository) GetByDesc(ctx context.Context, description string) (*model.Task, error) {
	panic("unimplemented")
}

// GetByDescWithTx implements repositories.TaskRepository.
func (r *TaskRepository) GetByDescWithTx(tx *firestore.Transaction, description string) (*model.Task, error) {
	panic("unimplemented")
}

// GetCollection implements repositories.TaskRepository.
func (r *TaskRepository) GetCollection() *firestore.CollectionRef {
	panic("unimplemented")
}

// GetCollectionName implements repositories.TaskRepository.
func (r *TaskRepository) GetCollectionName() string {
	panic("unimplemented")
}

// GetDocRef implements repositories.TaskRepository.
func (r *TaskRepository) GetDocRef(id string) *firestore.DocumentRef {
	panic("unimplemented")
}

// GetMulti implements repositories.TaskRepository.
func (r *TaskRepository) GetMulti(ctx context.Context, ids []string, opts ...repositories.GetOption) ([]*model.Task, error) {
	panic("unimplemented")
}

// GetMultiWithTx implements repositories.TaskRepository.
func (r *TaskRepository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...repositories.GetOption) ([]*model.Task, error) {
	panic("unimplemented")
}

// GetWithDoc implements repositories.TaskRepository.
func (r *TaskRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...repositories.GetOption) (*model.Task, error) {
	panic("unimplemented")
}

// GetWithDocWithTx implements repositories.TaskRepository.
func (r *TaskRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...repositories.GetOption) (*model.Task, error) {
	panic("unimplemented")
}

// GetWithTx implements repositories.TaskRepository.
func (r *TaskRepository) GetWithTx(tx *firestore.Transaction, id string, opts ...repositories.GetOption) (*model.Task, error) {
	panic("unimplemented")
}

// Insert implements repositories.TaskRepository.
func (r *TaskRepository) Insert(ctx context.Context, subject *model.Task) (_ string, err error) {

	projectID := "test-project"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	rep := repositories.NewTaskRepository(client)
	_, err = rep.Insert(ctx, subject)

	if err != nil {
		return "", fmt.Errorf("failed to StoreTask: %w", err)
	}

	return "", nil
}

// InsertMulti implements repositories.TaskRepository.
func (r *TaskRepository) InsertMulti(ctx context.Context, subjects []*model.Task) (_ []string, er error) {
	panic("unimplemented")
}

// InsertMultiWithTx implements repositories.TaskRepository.
func (r *TaskRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.Task) (_ []string, er error) {
	panic("unimplemented")
}

// InsertWithTx implements repositories.TaskRepository.
func (r *TaskRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.Task) (_ string, err error) {
	panic("unimplemented")
}

// RunInTransaction implements repositories.TaskRepository.
func (r *TaskRepository) RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error) {
	panic("unimplemented")
}

// Search implements repositories.TaskRepository.
func (r *TaskRepository) Search(ctx context.Context, param *repositories.TaskSearchParam, q *firestore.Query) ([]*model.Task, error) {
	panic("unimplemented")
}

// SearchByParam implements repositories.TaskRepository.
func (r *TaskRepository) SearchByParam(ctx context.Context, param *repositories.TaskSearchParam) ([]*model.Task, *repositories.PagingResult, error) {
	panic("unimplemented")
}

// SearchByParamWithTx implements repositories.TaskRepository.
func (r *TaskRepository) SearchByParamWithTx(tx *firestore.Transaction, param *repositories.TaskSearchParam) ([]*model.Task, *repositories.PagingResult, error) {
	panic("unimplemented")
}

// SearchWithTx implements repositories.TaskRepository.
func (r *TaskRepository) SearchWithTx(tx *firestore.Transaction, param *repositories.TaskSearchParam, q *firestore.Query) ([]*model.Task, error) {
	panic("unimplemented")
}

// StrictUpdate implements repositories.TaskRepository.
func (r *TaskRepository) StrictUpdate(ctx context.Context, id string, param *repositories.TaskUpdateParam, opts ...firestore.Precondition) error {
	panic("unimplemented")
}

// StrictUpdateWithTx implements repositories.TaskRepository.
func (r *TaskRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *repositories.TaskUpdateParam, opts ...firestore.Precondition) error {
	panic("unimplemented")
}

// Update implements repositories.TaskRepository.
func (r *TaskRepository) Update(ctx context.Context, subject *model.Task) (err error) {
	panic("unimplemented")
}

// UpdateMulti implements repositories.TaskRepository.
func (r *TaskRepository) UpdateMulti(ctx context.Context, subjects []*model.Task) (er error) {
	panic("unimplemented")
}

// UpdateMultiWithTx implements repositories.TaskRepository.
func (r *TaskRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.Task) (er error) {
	panic("unimplemented")
}

// UpdateWithTx implements repositories.TaskRepository.
func (r *TaskRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.Task) (err error) {
	panic("unimplemented")
}

func (r *TaskRepository) Store(ctx context.Context, m *model.Task) error {
	// projectID := "test-project"

	// client, err := firestore.NewClient(ctx, projectID)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	rep := repositories.NewTaskRepository(r.client)
	_, err := rep.Insert(ctx, m)

	if err != nil {
		return fmt.Errorf("failed to StoreTask: %w", err)
	}

	return nil
}
