package sql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/stretchr/testify/require"
)

// Store provides all functions to execute DB query as well as their combinations within a txn
type Store struct {
	*Queries // Composition - Inheritance
	db       *sql.DB
}

// Create new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Executes a  function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %x", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

var txKey = struct{}{}

// TransferTxParams contains the input params of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is theresult of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another account
// It creates a transfer record, add account entries and update accounts' balance
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "Create Transfer")

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "Create Entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "Create Entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// Update accounts' balance
		fmt.Println(txName, "Get account 1 for update")

		fmt.Println(txName, "update account 1 balance")
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 2 for update")

		fmt.Println(txName, "update account 2 balance")
		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
