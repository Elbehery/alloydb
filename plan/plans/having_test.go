// 
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package plans

import (
	. "github.com/pingcap/check"
	"github.com/Dong-Chan/alloydb/expression/expressions"
	"github.com/Dong-Chan/alloydb/model"
	"github.com/Dong-Chan/alloydb/parser/opcode"
)

type testHavingPlan struct{}

var _ = Suite(&testHavingPlan{})

var havingTestData = []*testRowData{
	&testRowData{1, []interface{}{10, "10"}},
	&testRowData{2, []interface{}{10, "20"}},
	&testRowData{3, []interface{}{10, "30"}},
	&testRowData{4, []interface{}{40, "40"}},
	&testRowData{6, []interface{}{60, "60"}},
}

func (t *testHavingPlan) TestHaving(c *C) {
	tblPlan := &testTablePlan{groupByTestData, []string{"id", "name"}}
	havingPlan := &HavingPlan{
		Src: tblPlan,
		Expr: &expressions.BinaryOperation{
			Op: opcode.GE,
			L: &expressions.Ident{
				CIStr: model.NewCIStr("id"),
			},
			R: &expressions.Value{
				Val: 20,
			},
		},
	}

	// having's behavior just like where
	cnt := 0
	havingPlan.Do(nil, func(id interface{}, data []interface{}) (bool, error) {
		cnt++
		return true, nil
	})
	c.Assert(cnt, Equals, 2)
}
