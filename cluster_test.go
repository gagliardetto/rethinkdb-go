// +build cluster

package gorethink

import (
	"fmt"

	test "gopkg.in/check.v1"
)

func (s *RethinkSuite) TestClusterConnect(c *test.C) {
	cluster, err := ConnectCluster(url, url2, url3)
	c.Assert(err, test.IsNil)

	row, err := Expr("Hello World").Run(cluster)
	c.Assert(err, test.IsNil)

	var response string
	err = row.One(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World")
}

func (s *RethinkSuite) TestClusterMultipleQueries(c *test.C) {
	cluster, err := ConnectCluster(url, url2, url3)
	c.Assert(err, test.IsNil)

	for i := 0; i < 10; i++ {
		row, err := Expr(fmt.Sprintf("Hello World", i)).Run(cluster)
		c.Assert(err, test.IsNil)

		var response string
		err = row.One(&response)
		c.Assert(err, test.IsNil)
		c.Assert(response, test.Equals, fmt.Sprintf("Hello World", i))
	}
}

func (s *RethinkSuite) TestClusterConnectError(c *test.C) {
	var err error
	_, err = ConnectCluster("nonexistanturl")
	c.Assert(err, test.NotNil)
}

func (s *RethinkSuite) TestClusterConnectDatabase(c *test.C) {
	session, err := ConnectClusterWithOpts(ConnectOpts{
		Database: "test2",
	}, url, url2, url3)
	c.Assert(err, test.IsNil)

	_, err = Table("test2").Run(session)
	c.Assert(err, test.NotNil)
	c.Assert(err.Error(), test.Equals, "gorethink: Database `test2` does not exist. in: \nr.Table(\"test2\")")
}
