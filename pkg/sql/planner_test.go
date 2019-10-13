package sql

import (
	"github.com/elliotcourant/melogale/pkg/engine"
	"github.com/elliotcourant/melogale/pkg/testutils"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func DoQuery(t *testing.T, txn engine.Transaction, query string) {
	start := time.Now()
	defer timber.Tracef("`%s` took %s", query, time.Since(start))
	p := NewPlanner(txn)
	tree := testutils.MustParse(t, query)
	plan, err := p.Build(tree)
	assert.NoError(t, err, "could not build plan")
	assert.NotEmpty(t, plan)
	err = plan.Run()
	assert.NoError(t, err, "could not run plan")
}

func TestPlan(t *testing.T) {
	db, cleanup := testutils.NewTestStore(t)
	defer cleanup()

	t.Run("create table", func(t *testing.T) {
		txn, err := db.Begin()
		assert.NoError(t, err, "could not create transaction")

		DoQuery(t, txn, "CREATE TABLE users (user_id BIGINT PRIMARY KEY, email TEXT, password TEXT);")
	})

	t.Run("create table and insert", func(t *testing.T) {
		txn, err := db.Begin()
		assert.NoError(t, err, "could not create transaction")

		DoQuery(t, txn, "CREATE TABLE users (user_id BIGINT PRIMARY KEY, email TEXT, password TEXT);")
		DoQuery(t, txn, "INSERT INTO users (user_id, email, password) VALUES (1, 'email@email.com', 'password');")
		DoQuery(t, txn, "INSERT INTO users (user_id, email, password) VALUES (2, 'email@email.com', 'password');")
		DoQuery(t, txn, "INSERT INTO users (user_id, email, password) VALUES (3, 'email@email.com', 'password');")
		DoQuery(t, txn, "INSERT INTO users (user_id, email, password) VALUES (4, 'email@email.com', 'password');")

		err = txn.Commit()
		assert.NoError(t, err, "could not commit")
	})
}
