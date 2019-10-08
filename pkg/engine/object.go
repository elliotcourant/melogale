package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
)

type ObjectIdPlan struct {
	objectType base.ObjectType
	name       string
	prefix     []byte
}

func (o *ObjectIdPlan) Explain() Explanation {
	return Explanation{
		Level:  0,
		Action: ID,
		Name:   "new object",
		Desc:   fmt.Sprintf("new object id for %s - %s", o.objectType, o.name),
		Key:    o.prefix,
	}
}

func (o *ObjectIdPlan) Execute(ctx ExecuteContext) error {
	id, err := ctx.Txn().NewObjectId(o.prefix)
	if err != nil {
		return fmt.Errorf("failed to generate object Id for %s - %s: %v", o.objectType, o.name, err)
	}
	ctx.SetObjectId(o.objectType, o.name, id)
	return nil
}

func (p *plannerBase) NewObjectIdPlan(objectType base.ObjectType, name string) PlanNode {
	return &ObjectIdPlan{
		objectType: objectType,
		name:       name,
		prefix:     base.NewObjectIdPrefix(objectType, name),
	}
}
