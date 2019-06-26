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

package stmts_test

import (
	. "github.com/pingcap/check"
	"github.com/Dong-Chan/alloydb"
	"github.com/Dong-Chan/alloydb/stmt/stmts"
)

func (s *testStmtSuite) TestAlterTable(c *C) {
	testSQL := "drop table if exists t; create table t (c1 int); alter table t add column c2 int;"

	stmtList, err := alloydb.Compile(testSQL)
	c.Assert(err, IsNil)

	stmtLen := len(stmtList)
	c.Assert(stmtLen, Greater, 0)

	testStmt, ok := stmtList[stmtLen-1].(*stmts.AlterTableStmt)
	c.Assert(ok, IsTrue)

	c.Assert(testStmt.IsDDL(), IsTrue)
	c.Assert(len(testStmt.OriginText()), Greater, 0)

	mf := newMockFormatter()
	testStmt.Explain(nil, mf)
	c.Assert(mf.Len(), Greater, 0)

	mustExec(c, s.testDB, testSQL)
}
