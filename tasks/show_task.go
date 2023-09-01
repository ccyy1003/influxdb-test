package tasks

type ShowTask struct {
	Task
}

var show_list = map[string]string{
	"SHOW DATABASES": "SHOW DATABASES",

	"SHOW SERIES":             `SHOW SERIES FROM yottadb_partition_replicas_num_lzl where cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 5 OFFSET 5`,
	"SHOW SERIES ON DATABASE": `SHOW SERIES ON mydb FROM yottadb_partition_replicas_num_lzl where cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 5 OFFSET 5`,
	"SHOW SERIES ORDER BY":    `SHOW SERIES  FROM yottadb_partition_replicas_num_lzl where cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 5 OFFSET 5 ORDER BY ASC`,

	"SHOW MEASUREMENTS":             `SHOW MEASUREMENTS WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 3 OFFSET 0`,
	"SHOW MEASUREMENTS ON DATABASE": `SHOW MEASUREMENTS ON mydb WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 3 OFFSET 0`,
	"SHOW MEASUREMENTS ORDER BY":    `SHOW MEASUREMENTS WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 3 OFFSET 0 ORDER BY ASC`,

	"SHOW TAG KEYS": `SHOW TAG KEYS WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 5 OFFSET 1`,
	"SHOW TAG KEYS [ SLIMIT_clause] [SOFFSET_clause] [ORDER BY]": "SHOW TAG KEYS WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 5 OFFSET 1 ORDER BY ASC SLIMIT 2 SOFFSET 2 ",
	"SHOW TAG KEYS ON DATABASE":                                  `SHOW TAG KEYS ON mydb WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 5 OFFSET 1`,

	"SHOW TAG VALUES": `SHOW TAG VALUES WITH KEY IN (account_id, cluster_display_name, cluster_name)  LIMIT 8 OFFSET 1`,
	"SHOW TAG VALUES WITH [ SLIMIT_clause] [SOFFSET_clause] [ORDER BY]": "SHOW TAG VALUES WITH KEY IN (account_id, cluster_display_name, cluster_name)   LIMIT 8 OFFSET 1 ORDER BY ASC SLIMIT 2 SOFFSET 1",
	"SHOW TAG VALUES ON DATABASE":                                       `SHOW TAG VALUES ON mydb WITH KEY IN (account_id, cluster_display_name, cluster_name)  LIMIT 8 OFFSET 1`,

	"SHOW FIELD KEYS": `SHOW FIELD KEYS FROM yottadb_partition_replicas_num`,
	"SHOW FIELD KEYS WITH [LIMIT_clause] [OFFSET_clause] [ SLIMIT_clause] [SOFFSET_clause] [ORDER BY]": "SHOW FIELD KEYS FROM yottadb_partition_replicas_num LIMIT 4 OFFSET 4 ORDER BY ASC SLIMIT 1 SOFFSET 1",
	"SHOW FIELD KEYS ON DATABASE": `SHOW FIELD KEYS ON mydb FROM yottadb_partition_replicas_num`,

	"SHOW SERIES CARDINALITY": `SHOW SERIES CARDINALITY from yottadb_partition_replicas_num group by partition_id LIMIT 1 OFFSET 1`,
	"SHOW SERIES CARDINALITY WITH [ SLIMIT_clause] [SOFFSET_clause]": "SHOW SERIES CARDINALITY from yottadb_partition_replicas_num group by partition_id LIMIT 1 OFFSET 1 SLIMIT 1 SOFFSET 1",
	"SHOW SERIES CARDINALITY ON DATABASE":                            `SHOW SERIES CARDINALITY on mydb from yottadb_partition_replicas_num group by partition_id LIMIT 1 OFFSET 1`,
}

func (s *ShowTask) Prepare() {
	s.TaskName = "show task"
	s.CmdList = show_list
}

func (s *ShowTask) GetTaskName() string {
	return "show"
}
