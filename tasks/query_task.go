package tasks

type QueryTask struct {
	Task
}

var query_list = map[string]string{
	"SELECT *":              "select * from car",
	"WHERE":                 "select * from car where speed > 110",
	"GROUP BY":              `select * from car group by "city"`,
	"ORDER BY TIME":         `select * from car group by city order by time ASC`,
	"INTO":                  `SELECT "water_level" INTO "h2o_feet_copy_1" FROM "h2o_feet" WHERE "location" = 'coyote_creek'`,
	"LIMIT":                 "select * from car limit 3 ",
	"SLIMIT":                "select * from car GROUP BY * SLIMIT 1",
	"OFFSET":                "select * from car LIMIT 3 OFFSET 3",
	"SOFFSET":               "select * from car GROUP BY * SLIMIT 1 SOFFSET 1",
	"SQL ORDER BY":          `select * from car where speed > 100 group by "city" limit 3 OFFSET 3 SLIMIT 1 SOFFSET 1 SQL ORDER BY "speed"`,
	"Time Zone":             `SELECT "water_level" FROM "h2o_feet" WHERE "location" = 'santa_monica'  tz('America/Chicago')`,
	"COMPLEX EXPRESSION":    "select (max(speed) + max(temp) - min(speed)) * count(speed) / 20 % 5 as f1, floor(min(speed)) & ceil(max(temp)) | round(min(speed)) ^ round(max(temp)) as f2, sqrt(pow(abs(max(speed)*min(speed)), 2)) as f3 from car group by time(3s), type fill(none)",
	"AGGREGATE CALCULATING": "select count(speed), mean(speed), count(temp), percentile(temp, 50) from car group by time(3s), type fill(none) limit 1 offset 1 slimit 2 soffset 1",
	"MIX SELECTOR WITH FIELD, GROUP BY TIME INTERVALS": "select max(speed)*6, speed*10 from car where time <= 2000000000ns group by time(1s)",
	"MIX SELECTOR WITH FIELD, GROUP BY TAG":            "select max(speed), pow(speed, temp/100.0) from car group by type",
	"ONLY FIELDS,GROUP BY TAG":                         "select *, temp, sqrt(pow(abs(max(speed)), 3)) FROM car group by city",
}

func (q *QueryTask) Prepare() {
	q.TaskName = "query task"
	q.CmdList = query_list
}

func (q *QueryTask) GetTaskName() string {
	return "query"
}
