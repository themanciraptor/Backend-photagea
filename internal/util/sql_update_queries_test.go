package util_test

import (
	"testing"

	"github.com/themanciraptor/Backend-photagea/internal/util"

	"github.com/stretchr/testify/assert"
)

type UpdateQueryTest struct {
	b *util.UpdateQueryBuilder
}

func setup() *UpdateQueryTest {
	return &UpdateQueryTest{b: util.InitUpdateQueryBuilder("test")}
}

func Test_Add__ReturnErrorIfValueisEmpty(t *testing.T) {
	ut := setup()
	ut.b.Add("cool", "").AddFilter("cool", "news")
	assert.Equal(t, len(ut.b.GetErrors()), 1)
}

func Test_Add__ReturnErrorIfFieldisEmpty(t *testing.T) {
	ut := setup()
	ut.b.Add("", "beans").AddFilter("cool", "news")
	assert.Equal(t, 1, len(ut.b.GetErrors()))
}

func Test_BuildUpdateStatement__CorrectlyAssemblesAndReturns(t *testing.T) {
	ut := setup()
	ut.b.Add("cool", "beans").AddFilter("cool", "news")
	q, err := ut.b.BuildUpdateQuery()
	assert.NoError(t, err, "No error should be returned")
	assert.Equal(t, len(ut.b.GetErrors()), 0)
	assert.Equal(t, "UPDATE test SET cool=? WHERE cool=?", q)
}

func Test_BuildUpdateStatement__CorrectlyAssemblesMultipleStatements(t *testing.T) {
	ut := setup()
	ut.b.Add("me good", "at talking").Add("but you good", "at smilin'").Add("take", "my bacons").AddFilter("cool", "news")
	q, err := ut.b.BuildUpdateQuery()
	assert.NoError(t, err, "No error should be returned")
	assert.Equal(t, len(ut.b.GetErrors()), 0)
	assert.Equal(t, "UPDATE test SET me good=?,but you good=?,take=? WHERE cool=?", q)
}
