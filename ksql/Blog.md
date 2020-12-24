# Building Data Wrangling zone using ksqlDB, kTable, kStream and Kafka Connect

## What is Data Wrangling ?
- Data Wrangling is the task of taking and standardising disorganised or incomplete raw data so that it can be obtained, consolidated, and analysed easily. It also requires mapping from source to destination data fields. For example, data wrangling could target a sector, row, or column in a dataset and execute an action to generate the necessary performance, such as joining, parsing, cleaning, consolidating, or filtering.
- It helps improves data accessibility by converting it to make it consistent with the end system as complex and intricate databases can obstruct data analysis and business processes. It has to be transformed and structured according to the specifications of the target system to make data available for the end-processes.

## Data Wrangling Use-Cases
- Real-time monitoring and real-time analytics
- Online Data Integration
- Materialized Cache
- Streaming ETL pipeline
- Event Driven Microservices

## Introducing ksqlDB as a platform component for handling Data Wrangling
- KsqlDB is an event streaming database. Constructed for applications that deals with stream processing.
- Events can be transformed in the form of tables (kTable) and streams (kStream).
- On these tables/streams, SQL operations are applied to transform or aggregate information and push it into another Kafka topic. 
## How it works ?
- A database for event streaming is a specific form of database that lets you create applications for stream processing. In virtually every event streaming architecture, it consolidates the many components contained.
- To do what you need, ksqlDB aims to provide one mental model. A full streaming app can be designed against ksqlDB, which in turn has only one dependency, i.e. Kafka.
- KSQL operates on continuous transformations of queries that run continuously as new data passes through them in Kafka topics with data streams. 
- For each topic partition processed by a given ksqlDB server, Kafka Streams generates one RocksDB state store instance for aggregates and joins. Each instance of the RocksDB state store has a 50 MB memory overhead for its cache plus the data actually stored.
- To prevent I/O operations, Kafka Streams/RocksDB attempts to hold the working set of a state store in memory for aggregates and joins. This takes more memory if there are several keys. 
## Setup
- Kafka: the event streaming platform.
  - Zookeeper
  ```  
    - bin/zookeeper-server-start.sh config/zookeeper.properties >> /dev/null &
  ```
  - Kafka
  ```
    - bin/kafka-server-start.sh config/server.properties >> /dev/null &
  ```
- ksqlDB Server:
  - ksql_server.list
  ```
    KSQL_LISTENERS=http://0.0.0.0:8088
    KSQL_BOOTSTRAP_SERVERS=KAFKA_BROKER_IP:9092
    KSQL_KSQL_LOGGING_PROCESSING_STREAM_AUTO_CREATE=true
    KSQL_KSQL_LOGGING_PROCESSING_TOPIC_AUTO_CREATE=true
    KSQL_KSQL_CONNECT_WORKER_CONFIG=/connect/connect.properties
    KSQL_CONNECT_REST_ADVERTISED_HOST_NAME=PUBLIC_IP_KAFKA_INSTANCE
    KSQL_CONNECT_GROUP_ID=ksql-connect-cluster
    KSQL_CONNECT_BOOTSTRAP_SERVERS=KAFKA_BROKER_IP:9092
    KSQL_CONNECT_KEY_CONVERTER=org.apache.kafka.connect.storage.StringConverter
    KSQL_CONNECT_VALUE_CONVERTER=org.apache.kafka.connect.json.JsonConverter
    KSQL_CONNECT_VALUE_CONVERTER_SCHEMAS_ENABLE=false
    KSQL_CONNECT_CONFIG_STORAGE_TOPIC=ksql-connect-configs
    KSQL_CONNECT_OFFSET_STORAGE_TOPIC=ksql-connect-offsets
    KSQL_CONNECT_STATUS_STORAGE_TOPIC=ksql-connect-statuses
    KSQL_CONNECT_CONFIG_STORAGE_REPLICATION_FACTOR=1
    KSQL_CONNECT_OFFSET_STORAGE_REPLICATION_FACTOR=1
    KSQL_CONNECT_STATUS_STORAGE_REPLICATION_FACTOR=1
    KSQL_CONNECT_PLUGIN_PATH=/usr/share/kafka/plugins
  ```
  - Server
  ```
   - docker run -it -p 8088:8088 --env-file ./ksql_server.list  confluentinc/ksqldb-server:0.13.0 
  ```  
- ksql CLI
 - This is command line utility that act as an interface for ksqlDB server and allow SQL operations to be executed interactively. 
 ```
   - docker run -it confluentinc/ksqldb-cli:0.13.0 ksql http://KSQLDB_SERVER_IP.18:8088
 ```
## What are kStream and kTable ?
### kStream
  - It is a structured but infinite series of events emitting out of a topic. 
  - It is immutable and can be created by specifying the format of incoming events like DELIMITED (CSV), JSON, AVRO etc. 
#### Create Stream
  ```
  create stream users_stream (name VARCHAR, countryCode VARCHAR) WITH (KAFKA_TOPIC='USERS', VALUE_FORMAT='JSON');
  ```  
#### Select from Stream
  ```
  select rowtime,* from user_stream emit changes;
  ```  
### kTable
  - kTables are defined by a primary key.
  - The events in kTable are updatable and can be deleted. 
#### Create Table
  ```
  create table countrytable (countrycode VARCHAR PRIMARY KEY, countryname VARCHAR) WITH (KAFKA_TOPIC='COUNTRY-CSV',VALUE_FORMAT='DELIMITED');
  ```  
#### Select from Table
  ```
  select * from countrytable where countrycode='GB'  emit changes limit 1;
  ```
## Type of Joins
   - Stream to Stream
   ```
   CREATE STREAM s3 AS SELECT s1.c1, s2.c2 FROM s1 JOIN s2 WITHIN 5 MINUTES ON s1.c1 = s2.c1 EMIT CHANGES;
   ```
   - Stream to Table
   ```
   CREATE STREAM s3 AS SELECT my_stream.c1, my_table.c2 FROM my_stream JOIN my_table ON s1.c1 = s2.c1 EMIT CHANGES;
   ```
   - Table to Table
   ```
   SELECT M.ID, M.TITLE, M.RELEASE_YEAR, L.ACTOR_NAME FROM MOVIES M INNER JOIN LEAD_ACTOR L ON M.TITLE = L.TITLE EMIT CHANGES LIMIT 3;
   ```
### Supported Join Combinations
   ```
   |Name         |Type        |INNER    |LEFT OUTER|FULL OUTER   |
   |-------------|------------|---------|----------|-------------|
   |Stream-Stream|Windowed    |Supported|Supported |Supported    |
   |Table-Table  |Non-windowed|Supported|Supported |Supported    |
   |Stream-Table |Non-windowed|Supported|Supported |Not Supported|
   ```   
## Add-on features
### Embedded kafka connect
 - ksqlDB connect management takes the responsibility to read from and write to between topic and external data source.
 - This functionality is very helpful if you don't want to write glue code to do it.
 - Download preferred sink/source connector jar from <a href='https://www.confluent.io/hub/'>Connectors</a>.
 - Copy/Mount the connector jar in docker volume '/usr/share/kafka/plugins', like '-v ./confluent-hub-components/debezium-debezium-connector-postgres:/usr/share/kafka/plugins/debezium-postgres'
 - Create connector configuration primarily takes 2 important input: 
   - Connection string of datasource.
   - Topic name.
   
         ```
         CREATE SOURCE/SINK CONNECTOR `jdbc-connector` WITH(
             "connector.class"='io.confluent.connect.jdbc.JdbcSourceConnector',
             "connection.url"='jdbc:postgresql://localhost:5432/my.db',
             "mode"='bulk',
             "topic.prefix"='jdbc-',
             "table.whitelist"='users',
             "key"='username');
         ```
### UDF (User Defined Functions) and UDAF (User Defined Aggregated Functions)  
#### UDF 
 - Extending ksql using its programming interface and create scalar functions.
 - For an input parameter return one output.
#### UDAF 
 - For many input rows return one output.
 - State of input rows is preserved and aggregated output is returned.
 - UDF & UDAF are implemented as custom jars.
 - Jars copied to 'ext' directory of the KSQL server.
 
 ```
 @UdfDescription(name = "SIMPLE_INTEREST", description = "Return simple interest calculated")
 public class SimpleInterest {
 
    @Udf(description = "Given principal, rate of interest and time return simple interest")
    public double simple_interest(final double principal, final double rate, final int time){
        ...
        ...
    }
 
 }
 ``` 
 
## KSQL Vs SparkSQL

```
|KSQL                                                                                 |SparkSQL                                                              |
|-------------------------------------------------------------------------------------|----------------------------------------------------------------------|
|KSQL is a streaming SQL engine that is fully interactive.                            |Spark SQL is not an interactive platform for Streaming SQL. You have  |
|One can interactively do sophisticated stream processing operations using SQL        |to switch between writing code using Java/Scala/Python and SQL        |
|statements alone.                                                                    |statements to do stream processing.                                   |
|                                                                                     |                                                                      |
|KSQL is a real Streaming SQL Engine event-at-a-time.                                 |Spark SQL is micro-batch.                                             |
|                                                                                     |                                                                      |
|KSQL belongs to "Stream Processing" category of the tech stack.                      |Apache Spark can be primarily classified under "Big Data Tools".      |
```

## Final Words
- Spark has been a battle tested tool since many years in the field of Data Streaming specially working with Kafka, but there is a cost associated with it which have been observed to be bear by every project.
  - As Spark is natively written in scala, all its latest releases and bug fixes first appear in scala because of Spark being its first class citizen.
  - Learning curve of scala is considered high as compared to other supported languages.
  - Teams may opt for other supported languages like Python, Java, R for Spark development. That will not only implicitly bring trans-compilation delays in building and deploying Spark jobs 
    but also may experience delays in the release of new features and bug fixes.
  - Going cloud native for setting up spark cluster over managed services like Gluu, EMR (Elastic Map Reduce) can be super expensive as they only allow sequential execution of spark jobs.
    Parallel execution expects a complete new cluster to be provisioned.
  - Dependency resolution is one of the another challenge, where teams spend good amount of time. It gets even more difficult where the dependency management repos are outside the firewall 
    periphery of the build environment.
  - Developers need to focus not only on the core data streaming logic but also on clean coding practices, unit tests, code coverage, CI/CD etc.
  
- I find KSQL (once deployed), only expects developers to understand SQL. 
  - It expose REST endpoint to accept updates in streaming queries in the form of SQL format.   
  - Can be deployed on any container orchestration platform. With horizontally scaled instances.
  - Natively support kafka, and provide embedded Kafka connect to ship data from/to multiple data sources.
  - Extensions to UDF (User Defined Functions), to implement complicated transformation logic.