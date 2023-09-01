package tasks

import (
	"fmt"
	"influxdb-test/common"
	"strings"

	"github.com/influxdata/influxdb/client/v2"
)

type Pair struct {
	First  string
	Second string
}
type MgDbTask struct {
	Task
	Mycmd []Pair
}

// var mgdb_list = map[string]string{
// 	// syntax : sentence
// 	"CREATE DATABASE WITH SHARD DURATION": "CREATE DATABASE test_db WITH DURATION 3d REPLICATION 3 SHARD DURATION 3d",
// 	"CREATE DATABASE WITH PARTITIONS":     "CREATE DATABASE test_db WITH DURATION 3d REPLICATION 3 PARTITIONS 16",
// 	"CREATE DATABASE WITH CTSDB_OPTION":   "CREATE DATABASE test_db WITH DURATION 3d REPLICATION 3 PARTITIONS 16 ctsdb_option '{\"route_tag\": {\"measurements\": { \"m1\": [ \"t1\", \"t2\", \"t3\"], \"m2\": [ \"t4\", \"t5\"]} } }'",
// 	"DROP MEASUREMENT":                    "DROP MEASUREMENT car",
// 	"DELETE":                              "delete from car where city = 'city_0'",
// 	"DROP SERIES":                         `DROP SERIES FROM "h2o_feet" WHERE "location" = 'santa_monica'`,
// 	"CREATE RETENTION POLICY":             `CREATE RETENTION POLICY "one_day_only" ON "test_db" DURATION 1d REPLICATION 1`,
// 	"ALTER RETENTION POLICY":              `ALTER RETENTION POLICY "one_day_only" ON "test_db" DURATION 3w SHARD DURATION 2h DEFAULT`,
// 	"DROP RETENTION POLICY":               `DROP RETENTION POLICY "one_day_only" ON "test_db"`,
// 	"DROP DATABASE ":                      "DROP DATABASE test_db",
// }

func (m *MgDbTask) Prepare() {
	m.TaskName = "mgdb task"
	//m.CmdList = mgdb_list
	m.Mycmd = append(m.Mycmd, Pair{"CREATE DATABASE WITH SHARD DURATION", "CREATE DATABASE test_db WITH DURATION 3d REPLICATION 3 SHARD DURATION 3d NAME autogen"})
	m.Mycmd = append(m.Mycmd, Pair{"CREATE DATABASE WITH PARTITIONS", "CREATE DATABASE test_db WITH DURATION 3d REPLICATION 3 PARTITIONS 16"})
	m.Mycmd = append(m.Mycmd, Pair{"CREATE DATABASE WITH CTSDB_OPTION", "CREATE DATABASE test_db WITH DURATION 3d REPLICATION 3 PARTITIONS 16 ctsdb_option '{\"route_tag\": {\"measurements\": { \"m1\": [ \"t1\", \"t2\", \"t3\"], \"m2\": [ \"t4\", \"t5\"]} } }'"})
	m.Mycmd = append(m.Mycmd, Pair{"CREATE RETENTION POLICY", `CREATE RETENTION POLICY "one_day_only" ON "test_db" DURATION 1d REPLICATION 1`})
	m.Mycmd = append(m.Mycmd, Pair{"ALTER RETENTION POLICY", `ALTER RETENTION POLICY "one_day_only" ON "test_db" DURATION 3w SHARD DURATION 2h DEFAULT`})
	m.Mycmd = append(m.Mycmd, Pair{"DELETE", "delete from car where city = 'city_0'"})
	m.Mycmd = append(m.Mycmd, Pair{"DROP SERIES", `DROP SERIES FROM "h2o_feet" WHERE "location" = 'santa_monica'`})
	m.Mycmd = append(m.Mycmd, Pair{"DROP MEASUREMENT", "DROP MEASUREMENT car"})
	m.Mycmd = append(m.Mycmd, Pair{"DROP RETENTION POLICY", `DROP RETENTION POLICY "one_day_only" ON "test_db"`})
	m.Mycmd = append(m.Mycmd, Pair{"DROP DATABASE ", "DROP DATABASE test_db"})
}

func (m *MgDbTask) Start() common.TestRes {
	var clnt = common.HttpClnt.Client

	var tr common.TestRes
	tr.TaskName = m.TaskName
	tr.TotalCnt = len(m.Mycmd)

	var q client.Query

	tr.PassCnt = 0
	for _, p := range m.Mycmd {
		q = client.NewQuery(p.Second, m.DbName(), "ns")

		if response, err := clnt.Query(q); err == nil && response.Error() == nil {
			tr.SupportedSyntax = append(tr.SupportedSyntax, fmt.Sprintf("[ %s ] %s ", p.First, p.Second))
			tr.PassCnt++
		} else {
			// record errinfo
			format_cmd := strings.ReplaceAll(p.Second, "\\", "")
			info := fmt.Sprintf("[ %s ] %s :", p.First, format_cmd)
			if err != nil {
				info += " " + err.Error() + " "
			}
			if response.Error() != nil {
				info += " " + response.Error().Error() + " "
			}

			tr.ErrInfos = append(tr.ErrInfos, info)
		}

	}

	return tr
}
func (m *MgDbTask) GetTaskName() string {
	return "manage db"
}
