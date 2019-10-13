package sql

import (
	"github.com/elliotcourant/melogale/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlan(t *testing.T) {
	db, cleanup := testutils.NewTestStore(t)
	defer cleanup()

	t.Run("create table", func(t *testing.T) {
		txn, err := db.Begin()
		assert.NoError(t, err, "could not create transaction")
		p := NewPlanner(txn)

		tree := testutils.MustParse(t, "CREATE TABLE users (user_id BIGINT PRIMARY KEY, email TEXT, password TEXT);")
		plan, err := p.Build(tree)
		assert.NoError(t, err, "could not build plan")
		assert.NotEmpty(t, plan)

		err = plan.Run()
		assert.NoError(t, err, "could not run plan")
	})

	t.Run("create table and insert", func(t *testing.T) {
		txn, err := db.Begin()
		assert.NoError(t, err, "could not create transaction")
		p := NewPlanner(txn)

		tree := testutils.MustParse(t, "CREATE TABLE users (user_id BIGINT PRIMARY KEY, email TEXT, password TEXT);")
		plan, err := p.Build(tree)
		assert.NoError(t, err, "could not build plan")
		assert.NotEmpty(t, plan)

		err = plan.Run()
		assert.NoError(t, err, "could not run plan")

		tree = testutils.MustParse(t, "INSERT INTO users (user_id, email, password) VALUES (1, 'email@email.com', 'password');")
		plan, err = p.Build(tree)
		assert.NoError(t, err, "could not build plan")
		assert.NotEmpty(t, plan)

		err = plan.Run()
		assert.NoError(t, err, "could not run plan")

		err = txn.Commit()
		assert.NoError(t, err, "could not commit")
	})
}
