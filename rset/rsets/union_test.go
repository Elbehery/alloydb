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

package rsets

import (
	. "github.com/pingcap/check"
	"github.com/Dong-Chan/alloydb/plan/plans"
)

var _ = Suite(&testUnionRsetSuite{})

type testUnionRsetSuite struct {
	r *UnionRset
}

func (s *testUnionRsetSuite) SetUpSuite(c *C) {
	s.r = &UnionRset{}
}

func (s *testUnionRsetSuite) TestShowRsetPlan(c *C) {
	p, err := s.r.Plan(nil)
	c.Assert(err, IsNil)

	_, ok := p.(*plans.UnionPlan)
	c.Assert(ok, IsTrue)
}
