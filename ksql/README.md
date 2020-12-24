# KSQL 

## Start standalone kafka
- [zookeeper]
  - bin/zookeeper-server-start.sh config/zookeeper.properties >> /dev/null &
- [kafka server]
  - bin/kafka-server-start.sh config/server.properties >> /dev/null &

## Start KsqlDB and Ksql CLI
- [ksqlDb server]
-v $PWD/plugin/confluentinc-kafka-connect-aws-dynamodb:/usr/share/kafka/plugins/dynamodb
  - docker run -it -p 8088:8088 --env-file ./ksql_server.list -v $PWD/plugin/confluentinc-kafka-connect-aws-dynamodb:/usr/share/kafka/plugins/dynamodb -v $PWD/plugin/confluentinc-kafka-connect-jdbc:/usr/share/kafka/plugins/jdbc  confluentinc/ksqldb-server:0.13.0
  - docker run -it -p 8088:8088 -e KSQL_LISTENERS=http://0.0.0.0:8088 -e KSQL_BOOTSTRAP_SERVERS=3.0.182.67:9092 -e KSQL_KSQL_LOGGING_PROCESSING_STREAM_AUTO_CREATE=true -e KSQL_KSQL_LOGGING_PROCESSING_TOPIC_AUTO_CREATE=true confluentinc/ksqldb-server:0.13.0
- [ksql cli]
  - docker run -it confluentinc/ksqldb-cli:0.13.0 ksql http://192.168.0.18:8088
- [ksql connect]
  - docker run -it docker pull confluentinc/cp-kafka-connect   
  
  
## Create USER topic
- bin/kafka-topics.sh --create --topic user --partitions 10 --replication-factor 1 --bootstrap-server 3.0.182.67:9092
## Produce data in topic
- bin/kafka-console-producer.sh --topic USERS --bootstrap-server 3.0.182.67:9092  
## DELETE topic
- bin/kafka-topics.sh --bootstrap-server 3.0.182.67:9092 --delete --topic USERPROFILE

## Ksql CLI commands
### List topics
- show topics;
- list topics;

```
KSQL is case sensitive.
```
### Streams
```
A stream in kafka is the full history from start of time.
```

### Messages
- Constantly arriving & added to topic.
- They are independent and never ending.
- Arrive in time ordered manner.
- Do not have relation with each other.
- Can be processed independently.

### Show events of topics
- print 'USERS';
- print 'USERS' from beginning;
- [Print top 2 events]
    - print 'USERS' from beginning limit 2;
- [Print events with an interval from top]
    - print 'USERS' from beginning interval 2 limit 2;    

### Push Queries
- Constantly query and output result.
- Keep getting result until you terminate query or hit result limit.
- 'emit changes' indicates the query is a push query.

### Create stream with CSV
- create stream users_stream (name VARCHAR, countryCode VARCHAR) WITH (KAFKA_TOPIC='USERS', VALUE_FORMAT='DELIMITED');
- Valid values for VALUE_FORMAT are 
    - DELIMITED (CSV)
    - JSON (JSON)
    - AVRO (AVRO)
- list streams; [List all created streams]
- drop stream ACCOUNTBALANCES; [drop created streams]   
- select * from users_stream emit changes; [select events from a stream] 
- SET 'auto.offset.reset' = 'earliest'; [tell the ksql current session to read streams from the beginning]
- select * from users_stream emit changes limit 2; [select from stream with limit]
- select countrycode, count(*) as count from users_stream group by countrycode emit changes; [aggregates in stream]
- drop stream if exists users_stream delete topic; [drop stream followed by delete topic]

### Create stream with JSON
- create stream user_stream (userid VARCHAR, firstname VARCHAR, lastname VARCHAR, countrycode VARCHAR, rating DOUBLE) with (KAFKA_TOPIC='USERPROFILE',VALUE_FORMAT='JSON');
- select * from user_stream emit changes;
- select rowtime,* from user_stream emit changes; [select rowtime along with data]  

## Manipulate stream
- select TIMESTAMPTOSTRING(ROWTIME, 'yyyy-MM-dd HH:mm:ss.SSS') as usertime,* from user_stream emit changes; [convert rowtime (epoch) in human readable format]
- select TIMESTAMPTOSTRING(ROWTIME, 'yyyy-MM-dd HH:mm:ss.SSS') as usertime, firstname+' '+ucase(lastname) as full_name from user_stream  emit changes; [concat fields in ksql and use of scalar function ucase :Uppercase]

## Streams from streams
- create stream user_stream_pretty_try as select firstname+' '+lastname+' from '+countrycode+' has a rating of '+cast(rating as varchar)+' stars. '+  case when rating < 2.5 then 'Poor' when rating between 2.5 and 4.2 then 'Good' else 'Excellent' end as description from user_stream; [case when then example]
- describe extended user_stream; [get the complete definition of stream]
- drop stream user_profile_pretty; [to drop the stream]
```
        Cannot drop USER_PROFILE_PRETTY.
        The following queries write into this source: [CSAS_USER_PROFILE_PRETTY_15].
```

This message means that there is already a query running over this stream that we are trying to drop. So, we need to terminate the query first.

- terminate CSAS_USER_PROFILE_PRETTY_15;
 
## Introducing Tables
- The table in kafka represent the current state of the stream.
- The message updates the previous message with the same key.
- Adds a new message when there is no new key in the table.
- bin/kafka-topics.sh --create --topic COUNTRY-CSV --partitions 1 --replication-factor 1 --bootstrap-server 3.0.182.67:9092
- bin/kafka-console-producer.sh --topic COUNTRY-CSV --property "parse.key=true" --property "key.separator=:" --bootstrap-server 3.0.182.67:9092
```
>AU:AU,Australia
>IN:IN,India
>GB:GB,United Kingdom 
>US:US,United States
```
- Instead above, use bin/kafka-console-producer.sh --topic COUNTRY-CSV --property "parse.key=true" --property "key.separator=," --bootstrap-server 3.0.182.67:9092
```
>AU,Australia
>IN,India
>GB,United Kingdom 
>US,United States
```
- create table countrytable (countrycode VARCHAR PRIMARY KEY, countryname VARCHAR) WITH (KAFKA_TOPIC='COUNTRY-CSV',VALUE_FORMAT='DELIMITED'); [create table with primary key]
- CREATE TABLE COUNTRYTABLE (countrycode VARCHAR PRIMARY KEY, countryname VARCHAR) WITH (KAFKA_TOPIC='COUNTRY-CSV', VALUE_FORMAT='KAFKA'); [there is a problem with csv values in creating table]
- show tables;
- describe countrytable;
- select * from countrytable where countrycode='GB'  emit changes limit 1; [select row from table] 

## Joins
- The result of ksql join is either a stream or a new table.
- stream + stream => new stream.
- table + table => new table.
    - for tables, key must be varchar or string.
    - message key must be same as the content of the column set in key.
- stream + table => new stream.
    - eg.  select up.firstname, up.lastname, up.countrycode,up.rating,c.countryname from user_stream up left join countrytable c on up.countrycode=c.countrycode emit changes;

## Connectors
- [DynamoDB]

CREATE SINK CONNECTOR user_dynamodb_sink WITH (
'connector.class' = 'io.confluent.connect.aws.dynamodb.DynamoDbSinkConnector',
'topics': '',
'aws.dynamodb.pk.hash' = 'value.USERID',
'aws.dynamodb.pk.sort' = 'value.FIRSTNAME',
'aws.dynamodb.region' = 'ap-southeast-1',
'aws.dynamodb.endpoint' = 'https://dynamodb.ap-southeast-1.amazonaws.com',
'table.name.format' = 'kafka_${topic}'        

[account]
- create stream account_stream (account_created struct<account struct<id varchar, name varchar, status varchar, opening_timestamp varchar,stakeholder_ids ARRAY<STRING>,
  instance_param_vals struct<auto_save_percentage varchar,autosave_opt_in varchar,credit_limit varchar,rloc_interest_rate varchar,withdrawal_hierarchy varchar>,
  details struct<KEY varchar>,accounting struct<tside varchar>>>) with (kafka_topic='account',value_format='JSON');

-   create stream account_stream_final as select account_created->account->id as id,
    account_created->account->name as name,
    account_created->account->status as status,
    account_created->account->opening_timestamp as opening_timestamp,
    explode(account_created->account->stakeholder_ids) as stakeholder_ids,
    account_created->account->instance_param_vals->auto_save_percentage as auto_save_percentage,
    account_created->account->instance_param_vals->autosave_opt_in as autosave_opt_in,
    account_created->account->instance_param_vals->credit_limit as credit_limit,
    account_created->account->instance_param_vals->rloc_interest_rate as rloc_interest_rate,
    account_created->account->instance_param_vals->withdrawal_hierarchy as withdrawal_hierarchy,
    account_created->account->details->KEY as KEY,
    account_created->account->accounting->tside as tside from account_stream; 
    
[balances]
- create stream balance_stream (account_id varchar,balances array<struct<account_address varchar,phase varchar,asset varchar,denomination varchar,value_time varchar,amount varchar,total_debit varchar,total_credit varchar>>) with (kafka_topic='balance',value_format='JSON');
- create stream balance_detail_stream as select account_id as accountId, explode(balances) as balances from balance_stream;    
- select * from balance_detail_stream emit changes;
- create stream balance_detail_final as select ACCOUNTID, 
  balances->account_address as accountAddress, 
  balances->phase as phase, 
  balances->asset as asset, 
  balances->denomination as denomination, 
  balances->value_time as valueTime, 
  balances->amount as amount, 
  balances->total_debit as totalDebit, 
  balances->total_credit as totalCredit from balance_detail_stream;
- select * from balance_detail_final emit changes;



- create table balance_detail_final_tbl (ACCOUNTID VARCHAR, ACCOUNTADDRESS VARCHAR, PHASE VARCHAR, ASSET VARCHAR, DENOMINATION VARCHAR, VALUETIME VARCHAR, AMOUNT VARCHAR,
 TOTALDEBIT VARCHAR, TOTALCREDIT VARCHAR) with (kafka_topic='BALANCE_DETAIL_FINAL',value_format='JSON');
 
- create table balance_detail_final_tbl as select ACCOUNTID , ACCOUNTADDRESS, PHASE, ASSET, DENOMINATION, VALUETIME, CAST(AMOUNT as double),
   cast(TOTALDEBIT as double), cast(TOTALCREDIT as double) from BALANCE_DETAIL_FINAL; 
   
- create table ACCOUNT_STREAM_FINAL_TBL (ID VARCHAR primary key, NAME VARCHAR, STATUS VARCHAR, OPENING_TIMESTAMP VARCHAR, STAKEHOLDER_IDS VARCHAR,  AUTO_SAVE_PERCENTAGE VARCHAR, AUTOSAVE_OPT_IN VARCHAR, CREDIT_LIMIT VARCHAR,
 RLOC_INTEREST_RATE VARCHAR, WITHDRAWAL_HIERARCHY VARCHAR, KEY VARCHAR, TSIDE VARCHAR) with (kafka_topic='ACCOUNT_STREAM_FINAL',value_format='JSON');  
 
- select *,sum(cast(b.amount as double)) from  balance_detail_final b left join account_stream_final a within 24 hour on b.accountid = a.id where a.id='ab43a8c6-9438-5f9e-49dd-4b404c359b68' group by b.accountaddress  emit changes; 

- select b.accountaddress,sum(cast(b.amount as double)) from  balance_detail_final b left join account_stream_final a within 24 hour on b.accountid = a.id where a.id='ab43a8c6-9438-5f9e-49dd-4b404c359b68' group by b.accountaddress  emit changes;


- select b.ACCOUNTID,trim(b.accountAddress) as accountaddress, b.PHASE, b.ASSET, b.DENOMINATION, a.ID, a.NAME, a.STATUS, a.STAKEHOLDER_IDS, 
a.AUTO_SAVE_PERCENTAGE, a.AUTOSAVE_OPT_IN, a.CREDIT_LIMIT, a.RLOC_INTEREST_RATE, a.WITHDRAWAL_HIERARCHY, a.KEY, a.TSIDE,sum(cast(b.amount as double)) as AMOUNT,
sum(cast(b.TOTALDEBIT as double)) as TOTALDEBIT, sum(cast(b.TOTALCREDIT as double)) as TOTALCREDIT from  
balance_detail_final b left join account_stream_final a within 24 hour on b.accountid = a.id 
where a.id='26ba6068-4eec-496d-2e10-53c027f51f6e' and b.phase='POSTING_PHASE_COMMITTED'  
group by b.ACCOUNTID,trim(b.accountAddress), b.PHASE, 
b.ASSET, b.DENOMINATION, a.ID, a.NAME, a.STATUS, 
a.STAKEHOLDER_IDS, a.AUTO_SAVE_PERCENTAGE, a.AUTOSAVE_OPT_IN, a.CREDIT_LIMIT, a.RLOC_INTEREST_RATE, a.WITHDRAWAL_HIERARCHY, a.KEY, a.TSIDE  emit changes;

- select b.ACCOUNTID,b.accountAddress, b.PHASE, b.ASSET, b.DENOMINATION, a.ID, a.NAME, a.STATUS, a.STAKEHOLDER_IDS, 
  a.AUTO_SAVE_PERCENTAGE, a.AUTOSAVE_OPT_IN, a.CREDIT_LIMIT, a.RLOC_INTEREST_RATE, a.WITHDRAWAL_HIERARCHY, a.KEY, a.TSIDE,
  sum(cast(b.amount as double)) as AMOUNT,sum(cast(b.TOTALDEBIT as double)) as TOTALDEBIT, sum(cast(b.TOTALCREDIT as double)) as TOTALCREDIT 
  from balance_detail_final b left join account_stream_final a within 24 hour on b.accountid = a.id 
  where a.id='26ba6068-4eec-496d-2e10-53c027f51f6e' and b.phase='POSTING_PHASE_COMMITTED'  
  group by b.ACCOUNTID,b.accountAddress, b.PHASE, 
  b.ASSET, b.DENOMINATION, a.ID, a.NAME, a.STATUS, 
  a.STAKEHOLDER_IDS, a.AUTO_SAVE_PERCENTAGE, a.AUTOSAVE_OPT_IN, a.CREDIT_LIMIT, a.RLOC_INTEREST_RATE, a.WITHDRAWAL_HIERARCHY, a.KEY, a.TSIDE  emit changes;
  
- select b.ACCOUNTID,b.accountAddress,b.amount,b.TOTALDEBIT,TOTALCREDIT, b.PHASE, b.ASSET, b.DENOMINATION, a.ID, a.NAME, a.STATUS, a.STAKEHOLDER_IDS, 
  a.AUTO_SAVE_PERCENTAGE, a.AUTOSAVE_OPT_IN, a.CREDIT_LIMIT, a.RLOC_INTEREST_RATE, a.WITHDRAWAL_HIERARCHY, a.KEY, a.TSIDE
  from balance_detail_final b left join account_stream_final a within 24 hour on b.accountid = a.id 
  where a.id='26ba6068-4eec-496d-2e10-53c027f51f6e' and b.phase='POSTING_PHASE_COMMITTED' emit changes;