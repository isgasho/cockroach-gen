// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1053
	`ALTER`: {
		//line sql.y: 1054
		Category: hGroup,
		//line sql.y: 1055
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1069
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1070
		Category: hDDL,
		//line sql.y: 1071
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1093
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1105
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1106
		Category: hDDL,
		//line sql.y: 1107
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1109
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1116
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1117
		Category: hDDL,
		//line sql.y: 1118
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
	//line sql.y: 1141
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1142
		Category: hPriv,
		//line sql.y: 1143
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1145
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1150
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1151
		Category: hDDL,
		//line sql.y: 1152
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1154
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1165
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1166
		Category: hDDL,
		//line sql.y: 1167
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1175
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1476
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1477
		Category: hCCL,
		//line sql.y: 1478
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
		//line sql.y: 1495
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1503
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1504
		Category: hCCL,
		//line sql.y: 1505
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
		//line sql.y: 1521
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1539
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1540
		Category: hCCL,
		//line sql.y: 1541
		Text: `
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   comma = '...'          [CSV-specific]
   comment = '...'        [CSV-specific]
   nullif = '...'         [CSV-specific]

`,
		//line sql.y: 1559
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1650
	`CANCEL`: {
		//line sql.y: 1651
		Category: hGroup,
		//line sql.y: 1652
		Text: `CANCEL JOB, CANCEL QUERY
`,
	},
	//line sql.y: 1658
	`CANCEL JOB`: {
		ShortDescription: `cancel a background job`,
		//line sql.y: 1659
		Category: hMisc,
		//line sql.y: 1660
		Text: `CANCEL JOB <jobid>
`,
		//line sql.y: 1661
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOB
`,
	},
	//line sql.y: 1669
	`CANCEL QUERY`: {
		ShortDescription: `cancel a running query`,
		//line sql.y: 1670
		Category: hMisc,
		//line sql.y: 1671
		Text: `CANCEL QUERY <queryid>
`,
		//line sql.y: 1672
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1696
	`CREATE`: {
		//line sql.y: 1697
		Category: hGroup,
		//line sql.y: 1698
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1719
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 1720
		Category: hMisc,
		//line sql.y: 1721
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 1736
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1737
		Category: hDML,
		//line sql.y: 1738
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1742
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1757
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1758
		Category: hCfg,
		//line sql.y: 1759
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1771
	`DROP`: {
		//line sql.y: 1772
		Category: hGroup,
		//line sql.y: 1773
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 1789
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1790
		Category: hDDL,
		//line sql.y: 1791
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1792
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1804
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 1805
		Category: hDDL,
		//line sql.y: 1806
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1807
		SeeAlso: `DROP
`,
	},
	//line sql.y: 1819
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 1820
		Category: hDDL,
		//line sql.y: 1821
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1822
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 1834
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 1835
		Category: hDDL,
		//line sql.y: 1836
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1837
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1857
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 1858
		Category: hDDL,
		//line sql.y: 1859
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 1860
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 1880
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 1881
		Category: hPriv,
		//line sql.y: 1882
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 1883
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 1895
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 1896
		Category: hPriv,
		//line sql.y: 1897
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 1898
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 1920
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 1921
		Category: hMisc,
		//line sql.y: 1922
		Text: `
EXPLAIN <statement>
EXPLAIN [( [PLAN ,] <planoptions...> )] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, EXPRS, METADATA, QUALIFY, INDENT, VERBOSE, DIST_SQL

`,
		//line sql.y: 1933
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 1994
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 1995
		Category: hMisc,
		//line sql.y: 1996
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 1997
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2019
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2020
		Category: hMisc,
		//line sql.y: 2021
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2022
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2045
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2046
		Category: hMisc,
		//line sql.y: 2047
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2048
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2068
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2069
		Category: hPriv,
		//line sql.y: 2070
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
		//line sql.y: 2083
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2099
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2100
		Category: hPriv,
		//line sql.y: 2101
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
		//line sql.y: 2114
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2183
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2184
		Category: hCfg,
		//line sql.y: 2185
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2186
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2198
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2199
		Category: hCfg,
		//line sql.y: 2200
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2201
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2227
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2228
		Category: hExperimental,
		//line sql.y: 2229
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2237
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2243
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2244
		Category: hExperimental,
		//line sql.y: 2245
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2253
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2261
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2262
		Category: hExperimental,
		//line sql.y: 2263
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
		//line sql.y: 2274
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2329
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2330
		Category: hCfg,
		//line sql.y: 2331
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2332
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2353
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2354
		Category: hCfg,
		//line sql.y: 2355
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }

`,
		//line sql.y: 2360
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2377
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2378
		Category: hTxn,
		//line sql.y: 2379
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2386
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2538
	`SHOW`: {
		//line sql.y: 2539
		Category: hGroup,
		//line sql.y: 2540
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW CREATE SEQUENCE, SHOW USERS,
SHOW TRANSACTION, SHOW BACKUP, SHOW JOBS, SHOW QUERIES, SHOW ROLES, SHOW SESSIONS, SHOW SYNTAX,
SHOW TRACE
`,
	},
	//line sql.y: 2572
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2573
		Category: hCfg,
		//line sql.y: 2574
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2575
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2596
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics`,
		//line sql.y: 2597
		Category: hMisc,
		//line sql.y: 2598
		Text: `SHOW STATISTICS FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
`,
		//line sql.y: 2603
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2611
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram`,
		//line sql.y: 2612
		Category: hMisc,
		//line sql.y: 2613
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2617
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2630
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2631
		Category: hCCL,
		//line sql.y: 2632
		Text: `SHOW BACKUP <location>
`,
		//line sql.y: 2633
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2641
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2642
		Category: hCfg,
		//line sql.y: 2643
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2646
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2663
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2664
		Category: hDDL,
		//line sql.y: 2665
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2666
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2674
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2675
		Category: hDDL,
		//line sql.y: 2676
		Text: `SHOW DATABASES
`,
		//line sql.y: 2677
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2685
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2686
		Category: hPriv,
		//line sql.y: 2687
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 2693
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2706
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2707
		Category: hDDL,
		//line sql.y: 2708
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2709
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2727
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2728
		Category: hDDL,
		//line sql.y: 2729
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2730
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2743
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2744
		Category: hMisc,
		//line sql.y: 2745
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2746
		SeeAlso: `CANCEL QUERY
`,
	},
	//line sql.y: 2762
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2763
		Category: hMisc,
		//line sql.y: 2764
		Text: `SHOW JOBS
`,
		//line sql.y: 2765
		SeeAlso: `CANCEL JOB, PAUSE JOB, RESUME JOB
`,
	},
	//line sql.y: 2773
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 2774
		Category: hMisc,
		//line sql.y: 2775
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
SHOW [COMPACT] [KV] TRACE FOR <statement>
`,
		//line sql.y: 2778
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 2807
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 2808
		Category: hMisc,
		//line sql.y: 2809
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
	},
	//line sql.y: 2825
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 2826
		Category: hDDL,
		//line sql.y: 2827
		Text: `SHOW TABLES [FROM <databasename>]
`,
		//line sql.y: 2828
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 2840
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 2841
		Category: hMisc,
		//line sql.y: 2842
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 2851
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 2852
		Category: hCfg,
		//line sql.y: 2853
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 2854
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 2873
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 2874
		Category: hDDL,
		//line sql.y: 2875
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 2876
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 2884
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 2885
		Category: hDDL,
		//line sql.y: 2886
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 2887
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 2895
	`SHOW CREATE SEQUENCE`: {
		ShortDescription: `display the CREATE SEQUENCE statement for a sequence`,
		//line sql.y: 2896
		Category: hDDL,
		//line sql.y: 2897
		Text: `SHOW CREATE SEQUENCE <seqname>
`,
	},
	//line sql.y: 2905
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 2906
		Category: hPriv,
		//line sql.y: 2907
		Text: `SHOW USERS
`,
		//line sql.y: 2908
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 2916
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 2917
		Category: hPriv,
		//line sql.y: 2918
		Text: `SHOW ROLES
`,
		//line sql.y: 2919
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3009
	`PAUSE JOB`: {
		ShortDescription: `pause a background job`,
		//line sql.y: 3010
		Category: hMisc,
		//line sql.y: 3011
		Text: `PAUSE JOB <jobid>
`,
		//line sql.y: 3012
		SeeAlso: `SHOW JOBS, CANCEL JOB, RESUME JOB
`,
	},
	//line sql.y: 3020
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3021
		Category: hDDL,
		//line sql.y: 3022
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS <expr> STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3049
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 3551
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 3552
		Category: hDDL,
		//line sql.y: 3553
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]

`,
		//line sql.y: 3561
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 3611
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 3612
		Category: hDML,
		//line sql.y: 3613
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 3614
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 3622
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 3623
		Category: hPriv,
		//line sql.y: 3624
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 3625
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 3647
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 3648
		Category: hPriv,
		//line sql.y: 3649
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 3650
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 3662
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 3663
		Category: hDDL,
		//line sql.y: 3664
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 3665
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 3679
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 3680
		Category: hDDL,
		//line sql.y: 3681
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3689
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 3883
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 3884
		Category: hTxn,
		//line sql.y: 3885
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 3886
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3894
	`RESUME JOB`: {
		ShortDescription: `resume a background job`,
		//line sql.y: 3895
		Category: hMisc,
		//line sql.y: 3896
		Text: `RESUME JOB <jobid>
`,
		//line sql.y: 3897
		SeeAlso: `SHOW JOBS, CANCEL JOB, PAUSE JOB
`,
	},
	//line sql.y: 3905
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 3906
		Category: hTxn,
		//line sql.y: 3907
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 3908
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3923
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 3924
		Category: hTxn,
		//line sql.y: 3925
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3933
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 3946
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 3947
		Category: hTxn,
		//line sql.y: 3948
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 3951
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 3975
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 3976
		Category: hTxn,
		//line sql.y: 3977
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 3978
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4091
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4092
		Category: hDDL,
		//line sql.y: 4093
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4094
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4163
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4164
		Category: hDML,
		//line sql.y: 4165
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4170
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4189
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4190
		Category: hDML,
		//line sql.y: 4191
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4195
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4300
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4301
		Category: hDML,
		//line sql.y: 4302
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4309
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 4479
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 4480
		Category: hDML,
		//line sql.y: 4481
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 4492
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 4493
		Category: hDML,
		//line sql.y: 4494
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
		//line sql.y: 4506
		SeeAlso: `WEBDOCS/select.html
`,
	},
	//line sql.y: 4581
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 4582
		Category: hDML,
		//line sql.y: 4583
		Text: `TABLE <tablename>
`,
		//line sql.y: 4584
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4851
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 4852
		Category: hDML,
		//line sql.y: 4853
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 4854
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4955
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 4956
		Category: hDML,
		//line sql.y: 4957
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 4975
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
