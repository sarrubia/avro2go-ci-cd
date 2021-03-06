/*This file has been autogenerated, please do not edit it.*/

package com_sarrubia_eventbus_avro_example_level2

// EventMetadataLevel2Bis1AVSC Avro schema to send to Schema Registry
const EventMetadataLevel2Bis1AVSC = "{\"type\":\"record\",\"name\":\"EventMetadataLevel2Bis1\",\"namespace\":\"com.sarrubia.eventbus.avro.example.level2\",\"doc\":\"\",\"fields\":[{\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"},{\"name\":\"key\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}],\"default\":null},{\"name\":\"versionX\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"},{\"name\":\"creationLong\",\"type\":\"long\",\"default\":0},{\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"}]}"

// EventMetadataLevel2Bis1Subject schema registry subject
const EventMetadataLevel2Bis1Subject = "com.sarrubia.eventbus.avro.example.level2.EventMetadataLevel2Bis1"

// EventMetadataLevel2Bis1Base
type EventMetadataLevel2Bis1Base struct {
	Id           string  `avro:"id"`
	Key          *string `avro:"key"`
	VersionX     string  `avro:"versionX"`
	CreationLong int64   `avro:"creationLong"`
	Type         string  `avro:"type"`
}

// Schema [REQUIRED] returns the schema const that belongs to EventMetadataLevel2Bis1
// Important: the receiver MUST NOT be a pointer in order to assert true against schema.EventSchemaInterface
func (e EventMetadataLevel2Bis1Base) Schema() string {
	return EventMetadataLevel2Bis1AVSC
}
