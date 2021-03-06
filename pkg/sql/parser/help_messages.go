// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1140
	`ALTER`: {
		//line sql.y: 1141
		Category: hGroup,
		//line sql.y: 1142
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1157
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1158
		Category: hDDL,
		//line sql.y: 1159
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... ALTER PRIMARY KEY USING INDEX <name>
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER TABLE ... UNSPLIT AT <selectclause>
  ALTER TABLE ... UNSPLIT ALL
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1197
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1211
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1212
		Category: hDDL,
		//line sql.y: 1213
		Text: `
ALTER PARTITION <name> <command>

Commands:
  -- Alter a single partition which exists on any of a table's indexes.
  ALTER PARTITION <partition> OF TABLE <tablename> CONFIGURE ZONE <zoneconfig>

  -- Alter a partition of a specific index.
  ALTER PARTITION <partition> OF INDEX <tablename>@<indexname> CONFIGURE ZONE <zoneconfig>

  -- Alter all partitions with the same name across a table's indexes.
  ALTER PARTITION <partition> OF INDEX <tablename>@* CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1232
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1237
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1238
		Category: hDDL,
		//line sql.y: 1239
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1241
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1248
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1249
		Category: hDDL,
		//line sql.y: 1250
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1273
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1274
		Category: hPriv,
		//line sql.y: 1275
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1277
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1282
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1283
		Category: hDDL,
		//line sql.y: 1284
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1286
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1294
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1295
		Category: hDDL,
		//line sql.y: 1296
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1308
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1313
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1314
		Category: hDDL,
		//line sql.y: 1315
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER INDEX ... UNSPLIT AT <selectclause>
  ALTER INDEX ... UNSPLIT ALL
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1331
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1833
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1834
		Category: hCCL,
		//line sql.y: 1835
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1852
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1864
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1865
		Category: hCCL,
		//line sql.y: 1866
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1882
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1920
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1921
		Category: hCCL,
		//line sql.y: 1922
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   DELIMITED
   MYSQLDUMP
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1950
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1994
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1995
		Category: hCCL,
		//line sql.y: 1996
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 2005
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2099
	`CANCEL`: {
		//line sql.y: 2100
		Category: hGroup,
		//line sql.y: 2101
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2108
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2109
		Category: hMisc,
		//line sql.y: 2110
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2113
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2131
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2132
		Category: hMisc,
		//line sql.y: 2133
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2136
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2167
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2168
		Category: hMisc,
		//line sql.y: 2169
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2172
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2242
	`CREATE`: {
		//line sql.y: 2243
		Category: hGroup,
		//line sql.y: 2244
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2325
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2326
		Category: hMisc,
		//line sql.y: 2327
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2470
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2471
		Category: hDML,
		//line sql.y: 2472
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2476
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2496
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2497
		Category: hCfg,
		//line sql.y: 2498
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2510
	`DROP`: {
		//line sql.y: 2511
		Category: hGroup,
		//line sql.y: 2512
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2529
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2530
		Category: hDDL,
		//line sql.y: 2531
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2532
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2544
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2545
		Category: hDDL,
		//line sql.y: 2546
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2547
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2559
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2560
		Category: hDDL,
		//line sql.y: 2561
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2562
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2574
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2575
		Category: hDDL,
		//line sql.y: 2576
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2577
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2597
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2598
		Category: hDDL,
		//line sql.y: 2599
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2600
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2620
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2621
		Category: hPriv,
		//line sql.y: 2622
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2623
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2635
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2636
		Category: hPriv,
		//line sql.y: 2637
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2638
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2662
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2663
		Category: hMisc,
		//line sql.y: 2664
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2677
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2760
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2761
		Category: hMisc,
		//line sql.y: 2762
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2763
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2794
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2795
		Category: hMisc,
		//line sql.y: 2796
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2797
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2827
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2828
		Category: hMisc,
		//line sql.y: 2829
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2830
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2850
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2851
		Category: hPriv,
		//line sql.y: 2852
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2865
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2881
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2882
		Category: hPriv,
		//line sql.y: 2883
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2896
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2950
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2951
		Category: hCfg,
		//line sql.y: 2952
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2953
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2965
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2966
		Category: hCfg,
		//line sql.y: 2967
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2968
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2977
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2978
		Category: hCfg,
		//line sql.y: 2979
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2982
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3003
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 3004
		Category: hExperimental,
		//line sql.y: 3005
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3013
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3019
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3020
		Category: hExperimental,
		//line sql.y: 3021
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3029
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3037
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3038
		Category: hExperimental,
		//line sql.y: 3039
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 3050
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3105
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3106
		Category: hCfg,
		//line sql.y: 3107
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3108
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3129
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3130
		Category: hCfg,
		//line sql.y: 3131
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3137
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3154
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3155
		Category: hTxn,
		//line sql.y: 3156
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3163
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3355
	`SHOW`: {
		//line sql.y: 3356
		Category: hGroup,
		//line sql.y: 3357
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3393
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3394
		Category: hCfg,
		//line sql.y: 3395
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3396
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3417
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3418
		Category: hExperimental,
		//line sql.y: 3419
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3426
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3439
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3440
		Category: hExperimental,
		//line sql.y: 3441
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3445
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3458
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3459
		Category: hCCL,
		//line sql.y: 3460
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3461
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3500
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3501
		Category: hCfg,
		//line sql.y: 3502
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3505
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3531
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3532
		Category: hDDL,
		//line sql.y: 3533
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3534
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3542
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3543
		Category: hDDL,
		//line sql.y: 3544
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3545
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3565
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3566
		Category: hDDL,
		//line sql.y: 3567
		Text: `SHOW DATABASES
`,
		//line sql.y: 3568
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3576
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3577
		Category: hPriv,
		//line sql.y: 3578
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3584
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3597
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3598
		Category: hDDL,
		//line sql.y: 3599
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3600
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3630
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3631
		Category: hDDL,
		//line sql.y: 3632
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3633
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3646
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3647
		Category: hMisc,
		//line sql.y: 3648
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3649
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3670
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3671
		Category: hMisc,
		//line sql.y: 3672
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3675
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3715
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3716
		Category: hMisc,
		//line sql.y: 3717
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3719
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3742
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3743
		Category: hMisc,
		//line sql.y: 3744
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3745
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3758
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3759
		Category: hDDL,
		//line sql.y: 3760
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3761
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3793
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3794
		Category: hDDL,
		//line sql.y: 3795
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3807
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3808
		Category: hDDL,
		//line sql.y: 3809
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3821
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3822
		Category: hMisc,
		//line sql.y: 3823
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3832
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3833
		Category: hCfg,
		//line sql.y: 3834
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3835
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3854
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3855
		Category: hDDL,
		//line sql.y: 3856
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3857
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3875
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3876
		Category: hPriv,
		//line sql.y: 3877
		Text: `SHOW USERS
`,
		//line sql.y: 3878
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3886
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3887
		Category: hPriv,
		//line sql.y: 3888
		Text: `SHOW ROLES
`,
		//line sql.y: 3889
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3945
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3946
		Category: hMisc,
		//line sql.y: 3947
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3968
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3969
		Category: hMisc,
		//line sql.y: 3970
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4207
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4208
		Category: hMisc,
		//line sql.y: 4209
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4212
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4229
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4230
		Category: hDDL,
		//line sql.y: 4231
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> ) [USING HASH WITH BUCKET_COUNT = <shard_buckets>]
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4258
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5002
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5003
		Category: hDDL,
		//line sql.y: 5004
		Text: `
CREATE [TEMPORARY | TEMP] SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 5014
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5079
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5080
		Category: hDML,
		//line sql.y: 5081
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5082
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5090
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5091
		Category: hPriv,
		//line sql.y: 5092
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 5093
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5122
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5123
		Category: hPriv,
		//line sql.y: 5124
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5125
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5143
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5144
		Category: hDDL,
		//line sql.y: 5145
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5146
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5193
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5194
		Category: hDDL,
		//line sql.y: 5195
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5203
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5434
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5435
		Category: hTxn,
		//line sql.y: 5436
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5437
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5445
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5446
		Category: hMisc,
		//line sql.y: 5447
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5450
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5467
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5468
		Category: hTxn,
		//line sql.y: 5469
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5470
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5485
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5486
		Category: hTxn,
		//line sql.y: 5487
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5495
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5508
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5509
		Category: hTxn,
		//line sql.y: 5510
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5513
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5537
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5538
		Category: hTxn,
		//line sql.y: 5539
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5540
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5658
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5659
		Category: hDDL,
		//line sql.y: 5660
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5661
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5730
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5731
		Category: hDML,
		//line sql.y: 5732
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5737
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5756
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5757
		Category: hDML,
		//line sql.y: 5758
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5762
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5873
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5874
		Category: hDML,
		//line sql.y: 5875
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5882
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6107
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6108
		Category: hDML,
		//line sql.y: 6109
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6120
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6121
		Category: hDML,
		//line sql.y: 6122
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 6134
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6209
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6210
		Category: hDML,
		//line sql.y: 6211
		Text: `TABLE <tablename>
`,
		//line sql.y: 6212
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6534
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6535
		Category: hDML,
		//line sql.y: 6536
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6537
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6646
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6647
		Category: hDML,
		//line sql.y: 6648
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'
  '{' IGNORE_FOREIGN_KEYS [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 6670
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
