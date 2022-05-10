/*This file has been autogenerated, please do not edit it.*/

package com_sarrubia_eventbus_avro_example_level2

// EventMetadataLevel2BisAVSC Avro schema to send to Schema Registry
const EventMetadataLevel2BisAVSC = "{\"type\":\"record\",\"name\":\"EventMetadataLevel2Bis\",\"namespace\":\"com.sarrubia.eventbus.avro.example.level2\",\"doc\":\"\",\"fields\":[{\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"},{\"name\":\"key\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}],\"default\":null},{\"name\":\"version\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"},{\"name\":\"creationLong\",\"type\":\"long\",\"default\":0},{\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"}]}"

// EventMetadataLevel2BisSubject schema registry subject
const EventMetadataLevel2BisSubject = "com.sarrubia.eventbus.avro.example.level2.EventMetadataLevel2Bis"

// EventMetadataLevel2BisBase
type EventMetadataLevel2BisBase struct {
	Id           string  `avro:"id"`
	Key          *string `avro:"key"`
	Version      string  `avro:"version"`
	CreationLong int64   `avro:"creationLong"`
	Type         string  `avro:"type"`
}

// Schema [REQUIRED] returns the schema const that belongs to EventMetadataLevel2Bis
// Important: the receiver MUST NOT be a pointer in order to assert true against schema.EventSchemaInterface
func (e EventMetadataLevel2BisBase) Schema() string {
	return EventMetadataLevel2BisAVSC
}