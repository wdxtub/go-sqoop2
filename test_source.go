package gosqoop2


var JSONVersion = `{
	"source-revision": "435d5e61b922a32d7bce567fe5fb1a9c0d9b1bbb",
	"source-url": "git:\/\/mbp.abrahamfine.com\/Users\/abefine\/cloudera_code\/sqoop-clean\/common",
	"build-date": "Tue Jul 19 16:08:27 PDT 2016",
	"user": "abefine",
	"api-versions": ["v1"],
	"build-version": "1.99.7"
}`

var JSONAllLink = `{
	"links": [{
		"link-config-values": {
			"configs": [{
				"validators": [],
				"inputs": [{
					"size": 128,
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.jdbcDriver",
					"id": 1,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "com.mysql.jdbc.Driver"
				}, {
					"size": 2000,
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.connectionString",
					"id": 2,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "jdbc:mysql://127.0.0.1:3306/yiplatform"
				}, {
					"size": 40,
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.username",
					"id": 3,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "root"
				}, {
					"size": 40,
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.password",
					"id": 4,
					"sensitive": true,
					"overrides": "",
					"type": "STRING"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.fetchSize",
					"id": 5,
					"sensitive": false,
					"overrides": "",
					"type": "INTEGER"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.jdbcProperties",
					"id": 6,
					"sensitive": false,
					"overrides": "",
					"type": "MAP",
					"value": {
						"protocol": "tcp"
					},
					"sensitive-pattern": ""
				}],
				"name": "linkConfig",
				"id": 1,
				"type": "LINK"
			}, {
				"validators": [],
				"inputs": [{
					"size": 5,
					"editable": "ANY",
					"validators": [],
					"name": "dialect.identifierEnclose",
					"id": 7,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "++"
				}],
				"name": "dialect",
				"id": 2,
				"type": "LINK"
			}],
			"validators": []
		},
		"creation-user": "yiplatform",
		"name": "mysql_link",
		"creation-date": 1525853205183,
		"connector-name": "generic-jdbc-connector",
		"id": 1,
		"update-date": 1526021615814,
		"enabled": true,
		"update-user": "yiplatform"
	}, {
		"link-config-values": {
			"configs": [{
				"validators": [],
				"inputs": [{
					"size": 255,
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.uri",
					"id": 65,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "hdfs://10.4.1.131:9000"
				}, {
					"size": 255,
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.confDir",
					"id": 66,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "/data/yiplatform/soft/bigdata/hadoop-2.7.2/etc/hadoop"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "linkConfig.configOverrides",
					"id": 67,
					"sensitive": false,
					"overrides": "",
					"type": "MAP",
					"sensitive-pattern": ""
				}],
				"name": "linkConfig",
				"id": 14,
				"type": "LINK"
			}],
			"validators": []
		},
		"creation-user": "yiplatform",
		"name": "hdfs_link",
		"creation-date": 1525853439596,
		"connector-name": "hdfs-connector",
		"id": 2,
		"update-date": 1526024384154,
		"enabled": true,
		"update-user": "yiplatform"
	}]
}`

var JSONAllJob = `{
	"jobs": [{
		"to-config-values": {
			"configs": [{
				"validators": [],
				"inputs": [{
					"editable": "ANY",
					"validators": [],
					"name": "toJobConfig.overrideNullValue",
					"id": 73,
					"sensitive": false,
					"overrides": "",
					"type": "BOOLEAN"
				}, {
					"size": 255,
					"editable": "ANY",
					"validators": [],
					"name": "toJobConfig.nullValue",
					"id": 74,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}, {
					"editable": "ANY",
					"validators": [],
					"values": "TEXT_FILE,SEQUENCE_FILE,PARQUET_FILE",
					"name": "toJobConfig.outputFormat",
					"id": 75,
					"sensitive": false,
					"overrides": "",
					"type": "ENUM",
					"value": "TEXT_FILE"
				}, {
					"editable": "ANY",
					"validators": [],
					"values": "NONE,DEFAULT,DEFLATE,GZIP,BZIP2,LZO,LZ4,SNAPPY,CUSTOM",
					"name": "toJobConfig.compression",
					"id": 76,
					"sensitive": false,
					"overrides": "",
					"type": "ENUM",
					"value": "NONE"
				}, {
					"size": 255,
					"editable": "ANY",
					"validators": [],
					"name": "toJobConfig.customCompression",
					"id": 77,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}, {
					"size": 255,
					"editable": "ANY",
					"validators": [],
					"name": "toJobConfig.outputDirectory",
					"id": 78,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "hdfs://10.4.1.131:9000/sqoop/test"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "toJobConfig.appendMode",
					"id": 79,
					"sensitive": false,
					"overrides": "",
					"type": "BOOLEAN"
				}],
				"name": "toJobConfig",
				"id": 17,
				"type": "JOB"
			}],
			"validators": []
		},
		"from-config-values": {
			"configs": [{
				"validators": [],
				"inputs": [{
					"size": 50,
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.schemaName",
					"id": 8,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "yiplatform"
				}, {
					"size": 50,
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.tableName",
					"id": 9,
					"sensitive": false,
					"overrides": "",
					"type": "STRING",
					"value": "operator"
				}, {
					"size": 2000,
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.sql",
					"id": 10,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.columnList",
					"id": 11,
					"sensitive": false,
					"overrides": "",
					"type": "LIST"
				}, {
					"size": 50,
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.partitionColumn",
					"id": 12,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.allowNullValueInPartitionColumn",
					"id": 13,
					"sensitive": false,
					"overrides": "",
					"type": "BOOLEAN"
				}, {
					"size": 50,
					"editable": "ANY",
					"validators": [],
					"name": "fromJobConfig.boundaryQuery",
					"id": 14,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}],
				"name": "fromJobConfig",
				"id": 3,
				"type": "JOB"
			}, {
				"validators": [],
				"inputs": [{
					"size": 50,
					"editable": "ANY",
					"validators": [],
					"name": "incrementalRead.checkColumn",
					"id": 15,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}, {
					"size": -1,
					"editable": "ANY",
					"validators": [],
					"name": "incrementalRead.lastValue",
					"id": 16,
					"sensitive": false,
					"overrides": "",
					"type": "STRING"
				}],
				"name": "incrementalRead",
				"id": 4,
				"type": "JOB"
			}],
			"validators": []
		},
		"to-connector-name": "hdfs-connector",
		"from-link-name": "mysql_link",
		"enabled": true,
		"from-connector-name": "generic-jdbc-connector",
		"to-link-name": "hdfs_link",
		"driver-config-values": {
			"configs": [{
				"validators": [],
				"inputs": [{
					"editable": "ANY",
					"validators": [],
					"name": "throttlingConfig.numExtractors",
					"id": 88,
					"sensitive": false,
					"overrides": "",
					"type": "INTEGER",
					"value": "2"
				}, {
					"editable": "ANY",
					"validators": [],
					"name": "throttlingConfig.numLoaders",
					"id": 89,
					"sensitive": false,
					"overrides": "",
					"type": "INTEGER",
					"value": "2"
				}],
				"name": "throttlingConfig",
				"id": 22,
				"type": "JOB"
			}, {
				"validators": [],
				"inputs": [{
					"editable": "ANY",
					"validators": [],
					"name": "jarConfig.extraJars",
					"id": 90,
					"sensitive": false,
					"overrides": "",
					"type": "LIST"
				}],
				"name": "jarConfig",
				"id": 23,
				"type": "JOB"
			}],
			"validators": []
		},
		"creation-user": "yiplatform",
		"name": "job_test",
		"creation-date": 1526021804999,
		"id": 1,
		"update-date": 1526022294435,
		"update-user": "yiplatform"
	}]
}`

var JSONAllConnector = `{
	"connectors": [{
		"job-config": {
			"FROM": {
				"configs": [{
					"validators": [],
					"inputs": [{
						"size": 50,
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.schemaName",
						"id": 8,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"size": 50,
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.tableName",
						"id": 9,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"size": 2000,
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.sql",
						"id": 10,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.columnList",
						"id": 11,
						"sensitive": false,
						"overrides": "",
						"type": "LIST"
					}, {
						"size": 50,
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.partitionColumn",
						"id": 12,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.allowNullValueInPartitionColumn",
						"id": 13,
						"sensitive": false,
						"overrides": "",
						"type": "BOOLEAN"
					}, {
						"size": 50,
						"editable": "ANY",
						"validators": [],
						"name": "fromJobConfig.boundaryQuery",
						"id": 14,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}],
					"name": "fromJobConfig",
					"id": 3,
					"type": "JOB"
				}, {
					"validators": [],
					"inputs": [{
						"size": 50,
						"editable": "ANY",
						"validators": [],
						"name": "incrementalRead.checkColumn",
						"id": 15,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"size": -1,
						"editable": "ANY",
						"validators": [],
						"name": "incrementalRead.lastValue",
						"id": 16,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}],
					"name": "incrementalRead",
					"id": 4,
					"type": "JOB"
				}],
				"validators": []
			},
			"TO": {
				"configs": [{
					"validators": [],
					"inputs": [{
						"size": 50,
						"editable": "ANY",
						"validators": [],
						"name": "toJobConfig.schemaName",
						"id": 17,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"size": 2000,
						"editable": "ANY",
						"validators": [],
						"name": "toJobConfig.tableName",
						"id": 18,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"editable": "ANY",
						"validators": [],
						"name": "toJobConfig.columnList",
						"id": 19,
						"sensitive": false,
						"overrides": "",
						"type": "LIST"
					}, {
						"size": 2000,
						"editable": "ANY",
						"validators": [],
						"name": "toJobConfig.stageTableName",
						"id": 20,
						"sensitive": false,
						"overrides": "",
						"type": "STRING"
					}, {
						"editable": "ANY",
						"validators": [],
						"name": "toJobConfig.shouldClearStageTable",
						"id": 21,
						"sensitive": false,
						"overrides": "",
						"type": "BOOLEAN"
					}],
					"name": "toJobConfig",
					"id": 5,
					"type": "JOB"
				}],
				"validators": []
			}
		},
		"name": "generic-jdbc-connector",
		"all-config-resources": {
			"toJobConfig.schemaName.help": "Schema name if the table is not stored in default schema. Note: Not all database systems understands the concept of schema.",
			"linkConfig.password.help": "Password to be used for connection to the database server.",
			"fromJobConfig.boundaryQuery.example": "select min(id), max(id) from input_table",
			"incrementalRead.checkColumn.example": "last_update_date",
			"fromJobConfig.schemaName.label": "Schema name",
			"linkConfig.connectionString.example": "jdbc:mysql:\/\/mysql.server\/sqoop",
			"fromJobConfig.allowNullValueInPartitionColumn.help": "Set true if partition column can contain NULL value.",
			"toJobConfig.stageTableName.label": "Staging table",
			"fromJobConfig.tableName.example": "input_table",
			"fromJobConfig.columnList.label": "Column names",
			"fromJobConfig.partitionColumn.example": "id",
			"dialect.label": "SQL Dialect",
			"fromJobConfig.sql.example": "select * from input_table where ${CONDITIONS}",
			"linkConfig.fetchSize.label": "Fetch Size",
			"fromJobConfig.allowNullValueInPartitionColumn.example": "true",
			"toJobConfig.schemaName.example": "my_schema",
			"fromJobConfig.help": "Specifies source and way how the data should be fetched from source database.",
			"fromJobConfig.sql.label": "SQL statement",
			"linkConfig.username.help": "Username to be used for connection to the database server.",
			"toJobConfig.label": "Database target",
			"toJobConfig.shouldClearStageTable.label": "Clear stage table",
			"linkConfig.username.label": "Username",
			"incrementalRead.checkColumn.help": "Column that is checked during incremental read for new values.",
			"toJobConfig.help": "Describes target destination and way how data should be persisted on the RDBMS system.",
			"toJobConfig.tableName.example": "target_table",
			"fromJobConfig.schemaName.example": "my_schema",
			"incrementalRead.lastValue.label": "Last value",
			"incrementalRead.lastValue.example": "19870202",
			"connector.name": "Generic JDBC Connector",
			"incrementalRead.help": "Configures optional incremental read from the database where source data are changing over time and only new changes need to be re-imported.",
			"incrementalRead.checkColumn.label": "Check column",
			"fromJobConfig.partitionColumn.label": "Partition column",
			"incrementalRead.label": "Incremental read",
			"fromJobConfig.sql.help": "Import data from given query's results set rather then static table.",
			"linkConfig.help": "Contains configuration that is required to establish connection with your database server.",
			"toJobConfig.shouldClearStageTable.help": "If set to true, staging table will be wiped out upon job start.",
			"linkConfig.jdbcDriver.label": "Driver class",
			"toJobConfig.schemaName.label": "Schema name",
			"fromJobConfig.boundaryQuery.help": "Customize query to retrieve minimal and maximal value of partition column.",
			"fromJobConfig.tableName.help": "Input table name from from which data will be retrieved.",
			"fromJobConfig.columnList.help": "Subset of columns that should be retrieved from source table.",
			"linkConfig.label": "Database connection",
			"linkConfig.connectionString.help": "JDBC connection string associated with your database server.",
			"linkConfig.connectionString.label": "Connection String",
			"linkConfig.username.example": "sqoop-user",
			"toJobConfig.columnList.example": "id,text,city",
			"toJobConfig.tableName.label": "Table name",
			"linkConfig.password.label": "Password",
			"linkConfig.password.example": "Sup3rS3cr3t!",
			"toJobConfig.stageTableName.help": "Name of table with same structure as final table that should be used as a staging destination. Data will be directly written to final table if no staging table is specified.",
			"linkConfig.jdbcDriver.example": "com.mysql.jdbc.Driver",
			"toJobConfig.stageTableName.example": "staging_target_table",
			"linkConfig.fetchSize.example": "1000",
			"dialect.identifierEnclose.label": "Identifier enclose",
			"dialect.identifierEnclose.example": "''" ,
"linkConfig.jdbcProperties.label": "Connection Properties",
"fromJobConfig.allowNullValueInPartitionColumn.label": "Partition column nullable",
"linkConfig.jdbcProperties.example": "useCompression=true",
"toJobConfig.tableName.help": "Destination table name to store transfer results.",
"toJobConfig.shouldClearStageTable.example": "true",
"dialect.help": "Database dialect that should be used for generated queries.",
"fromJobConfig.columnList.example": "id,text,city",
"incrementalRead.lastValue.help": "Last imported value, job will read only newer values.",
"fromJobConfig.label": "Database source",
"dialect.identifierEnclose.help": "Character(s) that should be used to enclose table name, schema or column names.",
"fromJobConfig.partitionColumn.help": "Input column that should be use to split the import into independent parallel processes. This column will be used in condition of generated queries.",
"fromJobConfig.tableName.label": "Table name",
"fromJobConfig.boundaryQuery.label": "Boundary query",
"fromJobConfig.schemaName.help": "Schema name if the table is not stored in default schema. Note: Not all database systems understands the concept of schema.",
"linkConfig.jdbcProperties.help": "Key-value pairs that should be passed down to JDBC driver when establishing connection.",
"toJobConfig.columnList.help": "Subset of columns that will will be written to. Omitted columns have to either allow NULL values or have defined default value.",
"toJobConfig.columnList.label": "Column names",
"linkConfig.jdbcDriver.help": "Fully qualified class name of the JDBC driver that will be used for establishing this connection. Check documentation for instructions how to make the driver's jar files available to Sqoop 2 server.",
"linkConfig.fetchSize.help": "Optional hint specifying requested JDBC fetch size."
},
"id": 1,
"class": "org.apache.sqoop.connector.jdbc.GenericJdbcConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 128,
"editable": "ANY",
"validators": [],
"name": "linkConfig.jdbcDriver",
"id": 1,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "linkConfig.connectionString",
"id": 2,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 40,
"editable": "ANY",
"validators": [],
"name": "linkConfig.username",
"id": 3,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 40,
"editable": "ANY",
"validators": [],
"name": "linkConfig.password",
"id": 4,
"sensitive": true,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "linkConfig.fetchSize",
"id": 5,
"sensitive": false,
"overrides": "",
"type": "INTEGER"
}, {
"editable": "ANY",
"validators": [],
"name": "linkConfig.jdbcProperties",
"id": 6,
"sensitive": false,
"overrides": "",
"type": "MAP",
"sensitive-pattern": ""
}],
"name": "linkConfig",
"id": 1,
"type": "LINK"
}, {
"validators": [],
"inputs": [{
"size": 5,
"editable": "ANY",
"validators": [],
"name": "dialect.identifierEnclose",
"id": 7,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "dialect",
"id": 2,
"type": "LINK"
}],
"validators": []
}
}, {
"job-config": {
"FROM": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.uri",
"id": 24,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "fromJobConfig",
"id": 7,
"type": "JOB"
}],
"validators": []
},
"TO": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.uri",
"id": 25,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"values": "CSV,AVRO,PARQUET",
"name": "toJobConfig.fileFormat",
"id": 26,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}],
"name": "toJobConfig",
"id": 8,
"type": "JOB"
}],
"validators": []
}
},
"name": "kite-connector",
"all-config-resources": {
"fromJobConfig.help": "Configuration options relevant to source dataset.",
"linkConfig.help": "Global configuration options that will be used for both from and to sides.",
"linkConfig.confDir.example": "\/etc\/hadoop\/conf\/",
"linkConfig.confDir.label": "Hadoop conf directory",
"toJobConfig.label": "Target configuration",
"fromJobConfig.uri.label": "dataset:hdfs:\/\/namespace\/table",
"toJobConfig.uri.label": "Dataset URI",
"fromJobConfig.label": "Source configuration",
"toJobConfig.uri.help": "Kite Dataset URI where should be data written to.",
"linkConfig.label": "Global configuration",
"linkConfig.authority.label": "HDFS host and port",
"linkConfig.authority.help": "Optional to override HDFS file system location.",
"toJobConfig.fileFormat.help": "Storage format that should be used when creating new dataset.",
"toJobConfig.help": "Configuration options relevant to target dataset.",
"fromJobConfig.uri.help": "Kite Dataset URI from which data will be read.",
"connector.name": "Kite connector",
"linkConfig.confDir.help": "Directory with Hadoop configuration files. This directory will be added to the classpath.",
"linkConfig.authority.example": "namenode.sqoop.org:8020",
"toJobConfig.fileFormat.label": "File format",
"toJobConfig.fileFormat.example": "PARQUET",
"toJobConfig.uri.example": "dataset:hdfs:\/\/namespace\/table"
},
"id": 2,
"class": "org.apache.sqoop.connector.kite.KiteConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "linkConfig.authority",
"id": 22,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "linkConfig.confDir",
"id": 23,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "linkConfig",
"id": 6,
"type": "LINK"
}],
"validators": []
}
}, {
"job-config": {
"FROM": {
"configs": [{
"validators": [],
"inputs": [{
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.tableName",
"id": 37,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.columns",
"id": 38,
"sensitive": false,
"overrides": "",
"type": "LIST"
}, {
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.consistentRead",
"id": 39,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.consistentReadScn",
"id": 40,
"sensitive": false,
"overrides": "",
"type": "LONG"
}, {
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.partitionList",
"id": 41,
"sensitive": false,
"overrides": "",
"type": "LIST"
}, {
"editable": "ANY",
"validators": [],
"values": "ROWID,PARTITION",
"name": "fromJobConfig.dataChunkMethod",
"id": 42,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"editable": "ANY",
"validators": [],
"values": "ROUNDROBIN,SEQUENTIAL,RANDOM",
"name": "fromJobConfig.dataChunkAllocationMethod",
"id": 43,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"editable": "ANY",
"validators": [],
"values": "SUBSPLIT,SPLIT",
"name": "fromJobConfig.whereClauseLocation",
"id": 44,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.omitLobColumns",
"id": 45,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.queryHint",
"id": 46,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.conditions",
"id": 47,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "fromJobConfig",
"id": 10,
"type": "JOB"
}],
"validators": []
},
"TO": {
"configs": [{
"validators": [],
"inputs": [{
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.tableName",
"id": 48,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.columns",
"id": 49,
"sensitive": false,
"overrides": "",
"type": "LIST"
}, {
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.templateTable",
"id": 50,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.partitioned",
"id": 51,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.nologging",
"id": 52,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.updateKey",
"id": 53,
"sensitive": false,
"overrides": "",
"type": "LIST"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.updateMerge",
"id": 54,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.dropTableIfExists",
"id": 55,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.storageClause",
"id": 56,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 2000,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.temporaryStorageClause",
"id": 57,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"values": "AUTO,ON,OFF",
"name": "toJobConfig.appendValuesHint",
"id": 58,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.parallel",
"id": 59,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}],
"name": "toJobConfig",
"id": 11,
"type": "JOB"
}],
"validators": []
}
},
"name": "oracle-jdbc-connector",
"all-config-resources": {
"connectionConfig.jdbcProperties.help": "Enter any JDBC properties that should be supplied during the creation of connection.",
"connectionConfig.password.help": "Enter the password to be used for connecting to the database.",
"fromJobConfig.consistentRead.help": "consistentRead",
"fromJobConfig.columns.label": "Columns",
"toJobConfig.updateKey.help": "updateKey",
"toJobConfig.updateMerge.label": "Merge updates",
"connectionConfig.fetchSize.label": "JDBC fetch size",
"toJobConfig.parallel.label": "Parallel",
"connectionConfig.username.label": "Username",
"connectionConfig.username.help": "Enter the username to be used for connecting to the database.",
"connectionConfig.initializationStatements.label": "Session initialization statements",
"toJobConfig.templateTable.help": "templateTable",
"connectionConfig.connectionString.label": "JDBC connection string",
"fromJobConfig.consistentReadScn.label": "Consistent read SCN",
"fromJobConfig.conditions.help": "conditions",
"fromJobConfig.help": "You must supply the information requested in order to create the FROM part of the job object.",
"connectionConfig.jdbcProperties.label": "JDBC connection properties",
"toJobConfig.nologging.label": "Nologging",
"toJobConfig.label": "To database configuration",
"fromJobConfig.dataChunkMethod.label": "Data chunk method",
"fromJobConfig.partitionList.help": "partitionList",
"toJobConfig.help": "You must supply the information requested in order to create the TO part of the job object.",
"fromJobConfig.dataChunkAllocationMethod.label": "Data chunk allocation method",
"fromJobConfig.omitLobColumns.label": "Omit LOB columns",
"connectionConfig.jdbcUrlVerbatim.help": "jdbcUrlVerbatim",
"toJobConfig.partitioned.help": "partitioned",
"connectionConfig.help": "You must supply the information requested in order to create an Oracle connection object.",
"toJobConfig.dropTableIfExists.help": "dropTableIfExists",
"fromJobConfig.tableName.help": "tableName",
"fromJobConfig.whereClauseLocation.label": "Where clause location",
"fromJobConfig.whereClauseLocation.help": "whereClauseLocation",
"toJobConfig.nologging.help": "nologging",
"connectionConfig.jdbcUrlVerbatim.label": "Use JDBC connection string verbatim",
"toJobConfig.dropTableIfExists.label": "Drop table if exists",
"toJobConfig.updateMerge.help": "updateMerge",
"fromJobConfig.consistentReadScn.help": "consistentReadScn",
"toJobConfig.storageClause.label": "Template table storage clause",
"toJobConfig.tableName.label": "Table name",
"toJobConfig.temporaryStorageClause.help": "temporaryStorageClause",
"fromJobConfig.columns.help": "Columns",
"connectionConfig.timeZone.help": "timeZone",
"connectionConfig.racServiceName.label": "RAC service name",
"fromJobConfig.queryHint.help": "queryHint",
"connectionConfig.actionName.label": "Session action name",
"connectionConfig.timeZone.label": "Session time zone",
"toJobConfig.updateKey.label": "Update key columns",
"connectionConfig.racServiceName.help": "racServiceName",
"toJobConfig.appendValuesHint.label": "Append values hint usage",
"toJobConfig.temporaryStorageClause.label": "Temporary table storage clause",
"toJobConfig.tableName.help": "Table name to write data into",
"fromJobConfig.omitLobColumns.help": "omitLobColumns",
"fromJobConfig.consistentRead.label": "Consistent read",
"fromJobConfig.label": "From Oracle configuration",
"toJobConfig.partitioned.label": "Partitioned",
"connectionConfig.connectionString.help": "Enter the value of JDBC connection string to be used by this connector for creating Oracle connections.",
"connectionConfig.label": "Oracle connection configuration",
"toJobConfig.storageClause.help": "storageClause",
"fromJobConfig.tableName.label": "Table name",
"connectionConfig.fetchSize.help": "fetchSize",
"fromJobConfig.conditions.label": "Conditions",
"connectionConfig.password.label": "Password",
"toJobConfig.columns.label": "Columns",
"fromJobConfig.dataChunkMethod.help": "dataChunkMethod",
"fromJobConfig.queryHint.label": "Query hint",
"connectionConfig.initializationStatements.help": "initializationStatements",
"toJobConfig.appendValuesHint.help": "appendValuesHint",
"connectionConfig.actionName.help": "actionName",
"toJobConfig.parallel.help": "parallel",
"fromJobConfig.partitionList.label": "Partitions",
"fromJobConfig.dataChunkAllocationMethod.help": "dataChunkAllocationMethod",
"toJobConfig.columns.help": "Columns",
"toJobConfig.templateTable.label": "Template table name"
},
"id": 3,
"class": "org.apache.sqoop.connector.jdbc.oracle.OracleJdbcConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 128,
"editable": "ANY",
"validators": [],
"name": "connectionConfig.connectionString",
"id": 27,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 40,
"editable": "ANY",
"validators": [],
"name": "connectionConfig.username",
"id": 28,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 40,
"editable": "ANY",
"validators": [],
"name": "connectionConfig.password",
"id": 29,
"sensitive": true,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "connectionConfig.jdbcProperties",
"id": 30,
"sensitive": false,
"overrides": "",
"type": "MAP",
"sensitive-pattern": ""
}, {
"size": -1,
"editable": "ANY",
"validators": [],
"name": "connectionConfig.timeZone",
"id": 31,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": -1,
"editable": "ANY",
"validators": [],
"name": "connectionConfig.actionName",
"id": 32,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "connectionConfig.fetchSize",
"id": 33,
"sensitive": false,
"overrides": "",
"type": "INTEGER"
}, {
"editable": "ANY",
"validators": [],
"name": "connectionConfig.initializationStatements",
"id": 34,
"sensitive": false,
"overrides": "",
"type": "LIST"
}, {
"editable": "ANY",
"validators": [],
"name": "connectionConfig.jdbcUrlVerbatim",
"id": 35,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"size": -1,
"editable": "ANY",
"validators": [],
"name": "connectionConfig.racServiceName",
"id": 36,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "connectionConfig",
"id": 9,
"type": "LINK"
}],
"validators": []
}
}, {
"job-config": {
"TO": {
"configs": [{
"validators": [],
"inputs": [{
"size": 260,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.outputDirectory",
"id": 64,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "toJobConfig",
"id": 13,
"type": "JOB"
}],
"validators": []
}
},
"name": "ftp-connector",
"all-config-resources": {
"linkConfig.help": "Parameters required to connect to an FTP server.",
"linkConfig.username.help": "Username that will be used to authenticate connection to the FTP Server.",
"linkConfig.password.help": "Password that will be used to authenticate connection to the FTP Server.",
"linkConfig.server.example": "ftp.apache.org",
"toJobConfig.label": "Output configuration",
"linkConfig.username.label": "Username",
"linkConfig.label": "FTP Server configuration",
"toJobConfig.help": "Parameters required to store data on the FTP server.",
"toJobConfig.outputDirectory.help": "Directory on the FTP server to write data to.",
"linkConfig.username.example": "sqoop",
"toJobConfig.outputDirectory.label": "Output directory",
"connector.name": "FTP Connector",
"linkConfig.password.example": "Sup3rS3cr3t!",
"linkConfig.password.label": "Password",
"toJobConfig.outputDirectory.example": "\/user\/sqoop\/data",
"linkConfig.server.help": "Hostname for the FTP server.",
"linkConfig.server.label": "Hostname",
"linkConfig.port.help": "Port for the FTP server. Connector will use 21 if omitted.",
"linkConfig.port.example": "21",
"linkConfig.port.label": "Port"
},
"id": 4,
"class": "org.apache.sqoop.connector.ftp.FtpConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 256,
"editable": "ANY",
"validators": [],
"name": "linkConfig.server",
"id": 60,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "linkConfig.port",
"id": 61,
"sensitive": false,
"overrides": "",
"type": "INTEGER"
}, {
"size": 256,
"editable": "ANY",
"validators": [],
"name": "linkConfig.username",
"id": 62,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 256,
"editable": "ANY",
"validators": [],
"name": "linkConfig.password",
"id": 63,
"sensitive": true,
"overrides": "",
"type": "STRING"
}],
"name": "linkConfig",
"id": 12,
"type": "LINK"
}],
"validators": []
}
}, {
"job-config": {
"FROM": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.inputDirectory",
"id": 68,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.overrideNullValue",
"id": 69,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "fromJobConfig.nullValue",
"id": 70,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "fromJobConfig",
"id": 15,
"type": "JOB"
}, {
"validators": [],
"inputs": [{
"editable": "ANY",
"validators": [],
"values": "NONE,NEW_FILES",
"name": "incremental.incrementalType",
"id": 71,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"editable": "ANY",
"validators": [],
"name": "incremental.lastImportedDate",
"id": 72,
"sensitive": false,
"overrides": "",
"type": "DATETIME"
}],
"name": "incremental",
"id": 16,
"type": "JOB"
}],
"validators": []
},
"TO": {
"configs": [{
"validators": [],
"inputs": [{
"editable": "ANY",
"validators": [],
"name": "toJobConfig.overrideNullValue",
"id": 73,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.nullValue",
"id": 74,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"values": "TEXT_FILE,SEQUENCE_FILE,PARQUET_FILE",
"name": "toJobConfig.outputFormat",
"id": 75,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"editable": "ANY",
"validators": [],
"values": "NONE,DEFAULT,DEFLATE,GZIP,BZIP2,LZO,LZ4,SNAPPY,CUSTOM",
"name": "toJobConfig.compression",
"id": 76,
"sensitive": false,
"overrides": "",
"type": "ENUM"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.customCompression",
"id": 77,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.outputDirectory",
"id": 78,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "toJobConfig.appendMode",
"id": 79,
"sensitive": false,
"overrides": "",
"type": "BOOLEAN"
}],
"name": "toJobConfig",
"id": 17,
"type": "JOB"
}],
"validators": []
}
},
"name": "hdfs-connector",
"all-config-resources": {
"incremental.incrementalType.label": "Incremental type",
"incremental.lastImportedDate.example": "1987-02-02 13:15:27",
"linkConfig.uri.help": "Namenode URI for your cluster.",
"incremental.incrementalType.help": "Type of incremental import.",
"fromJobConfig.nullValue.label": "Null value",
"linkConfig.uri.example": "hdfs:\/\/nn1.example.com\/",
"incremental.help": "Information relevant for incremental reading from HDFS.",
"toJobConfig.overrideNullValue.example": "true",
"fromJobConfig.overrideNullValue.label": "Override null value",
"toJobConfig.appendMode.help": "If set to false, job will fail if output directory already exists. If set to true then imported data will be stored to already existing and possibly non empty directory.",
"linkConfig.confDir.help": "Directory on Sqoop server machine with hdfs configuration files (hdfs-site.xml, ...). This connector will load all files ending with -site.xml.",
"toJobConfig.outputDirectory.example": "\/user\/jarcec\/output-dir",
"fromJobConfig.nullValue.help": "For file formats that doesn't have native representation of NULL (as for example text file) use this particular string to decode NULL value.",
"toJobConfig.overrideNullValue.help": "If set to true, then the null value will be overridden with the value set in Null value.",
"toJobConfig.nullValue.label": "Null value",
"toJobConfig.customCompression.example": "org.apache.hadoop.io.compress.SnappyCodec",
"incremental.lastImportedDate.help": "Datetime stamp of last read file. Next job execution will read only files that have been created after this point in time.",
"toJobConfig.outputFormat.example": "PARQUET_FILE",
"fromJobConfig.help": "Specifies information required to get data from HDFS.",
"toJobConfig.customCompression.label": "Custom codec",
"linkConfig.confDir.example": "\/etc\/hadoop\/conf\/",
"fromJobConfig.overrideNullValue.example": "true",
"linkConfig.confDir.label": "Conf directory",
"toJobConfig.label": "Target configuration",
"toJobConfig.help": "Configuration describing where and how the resulting data should be stored.",
"toJobConfig.compression.help": "Compression codec that should be use to compress transferred data.",
"toJobConfig.outputDirectory.help": "HDFS directory where transferred data will be written to.",
"connector.name": "HDFS Connector",
"linkConfig.configOverrides.help": "Additional configuration that will be set on HDFS Configuration object, possibly overriding any keys loaded from configuration files.",
"linkConfig.uri.label": "URI",
"fromJobConfig.inputDirectory.label": "Input directory",
"toJobConfig.appendMode.example": "true",
"toJobConfig.compression.example": "SNAPPY",
"linkConfig.help": "Contains configuration required to connect to your HDFS cluster.",
"incremental.label": "Incremental import",
"linkConfig.label": "HDFS cluster",
"incremental.incrementalType.example": "NEW_FILES",
"toJobConfig.overrideNullValue.label": "Override null value",
"fromJobConfig.inputDirectory.example": "\/user\/jarcec\/input-dir",
"incremental.lastImportedDate.label": "Last imported date",
"toJobConfig.compression.label": "Compression codec",
"fromJobConfig.overrideNullValue.help": "If set to true, then the null value will be overridden with the value set in Null value.",
"toJobConfig.nullValue.example": "N",
"fromJobConfig.nullValue.example": "N",
"linkConfig.configOverrides.label": "Additional configs:",
"fromJobConfig.label": "Input configuration",
"toJobConfig.nullValue.help": "For file formats that doesn't have native representation of NULL (as for example text file) use this particular string to encode NULL value.",
"linkConfig.configOverrides.example": "custom.key=custom.value",
"fromJobConfig.inputDirectory.help": "Input directory containing files that should be transferred.",
"toJobConfig.outputFormat.help": "File format that should be used for transferred data.",
"toJobConfig.customCompression.help": "Fully qualified class name with Hadoop codec implementation that should be used if none of the build-in options are suitable.",
"toJobConfig.outputDirectory.label": "Output directory",
"toJobConfig.appendMode.label": "Append mode",
"toJobConfig.outputFormat.label": "File format"
},
"id": 5,
"class": "org.apache.sqoop.connector.hdfs.HdfsConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "linkConfig.uri",
"id": 65,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "linkConfig.confDir",
"id": 66,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "linkConfig.configOverrides",
"id": 67,
"sensitive": false,
"overrides": "",
"type": "MAP",
"sensitive-pattern": ""
}],
"name": "linkConfig",
"id": 14,
"type": "LINK"
}],
"validators": []
}
}, {
"job-config": {
"TO": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.topic",
"id": 82,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "toJobConfig",
"id": 19,
"type": "JOB"
}],
"validators": []
}
},
"name": "kafka-connector",
"all-config-resources": {
"linkConfig.help": "Configuration options describing Kafka cluster.",
"toJobConfig.topic.label": "Topic",
"toJobConfig.topic.example": "sqoop_topic",
"toJobConfig.label": "Output configuration",
"linkConfig.brokerList.example": "broker-1.kafka.org:9092,broker-2.kafka.org:9092",
"toJobConfig.topic.help": "Name of Kafka topic where data will be written into.",
"linkConfig.brokerList.help": "Comma-separated list of Kafka brokers in the form of host:port. It doesn't need to contain all brokers, but at least two are recommended for high availability",
"linkConfig.label": "Kafka cluster",
"linkConfig.brokerList.label": "Kafka brokers",
"toJobConfig.help": "Configuration necessary when writing data to Kafka.",
"linkConfig.zookeeperConnect.example": "zk-1.kafka.org:2181,zk-2.kafka.org:2182",
"connector.name": "Kafka connector",
"linkConfig.zookeeperConnect.label": "Zookeeper quorum",
"linkConfig.zookeeperConnect.help": "Address of Zookeeper used by the Kafka cluster. Usually host:port. Multiple zookeeper nodes are supported. If Kafka is stored in its own znode use host:portkafka"
},
"id": 6,
"class": "org.apache.sqoop.connector.kafka.KafkaConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 1024,
"editable": "ANY",
"validators": [],
"name": "linkConfig.brokerList",
"id": 80,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 255,
"editable": "ANY",
"validators": [],
"name": "linkConfig.zookeeperConnect",
"id": 81,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "linkConfig",
"id": 18,
"type": "LINK"
}],
"validators": []
}
}, {
"job-config": {
"TO": {
"configs": [{
"validators": [],
"inputs": [{
"size": 260,
"editable": "ANY",
"validators": [],
"name": "toJobConfig.outputDirectory",
"id": 87,
"sensitive": false,
"overrides": "",
"type": "STRING"
}],
"name": "toJobConfig",
"id": 21,
"type": "JOB"
}],
"validators": []
}
},
"name": "sftp-connector",
"all-config-resources": {
"linkConfig.help": "Parameters required to connect to an SFTP server.",
"linkConfig.username.help": "Username that will be used to authenticate connection to SFTP serer.",
"linkConfig.password.help": "Password that will be used to authenticate connection to the FTP Server.",
"linkConfig.server.example": "sftp.apache.org",
"toJobConfig.label": "Output configuration",
"linkConfig.username.label": "Username",
"linkConfig.label": "FTP Server configuration",
"toJobConfig.help": "Parameters required to store data on the SFTP server.",
"toJobConfig.outputDirectory.help": "Directory on the SFTP server to write data to.",
"linkConfig.username.example": "sqoop",
"toJobConfig.outputDirectory.label": "Output directory",
"connector.name": "SFTP Connector",
"linkConfig.password.example": "Sup3rS3cr3t!",
"linkConfig.password.label": "Password",
"toJobConfig.outputDirectory.example": "\/user\/sqoop\/data",
"linkConfig.server.help": "Hostname of the SFTP server.",
"linkConfig.server.label": "Hostname",
"linkConfig.port.help": "Port for the SFTP server. Connector will use 22 if omitted.",
"linkConfig.port.example": "22",
"linkConfig.port.label": "Port"
},
"id": 7,
"class": "org.apache.sqoop.connector.sftp.SftpConnector",
"version": "1.99.7",
"link-config": {
"configs": [{
"validators": [],
"inputs": [{
"size": 255,
"editable": "ANY",
"validators": [],
"name": "linkConfig.server",
"id": 83,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"editable": "ANY",
"validators": [],
"name": "linkConfig.port",
"id": 84,
"sensitive": false,
"overrides": "",
"type": "INTEGER"
}, {
"size": 256,
"editable": "ANY",
"validators": [],
"name": "linkConfig.username",
"id": 85,
"sensitive": false,
"overrides": "",
"type": "STRING"
}, {
"size": 256,
"editable": "ANY",
"validators": [],
"name": "linkConfig.password",
"id": 86,
"sensitive": true,
"overrides": "",
"type": "STRING"
}],
"name": "linkConfig",
"id": 20,
"type": "LINK"
}],
"validators": []
}
}]
}`
