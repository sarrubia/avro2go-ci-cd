/*This file has been autogenerated, please do not edit it.*/

package com_sarrubia_eventbus_avro_example_embedded

// EventMetadataNewAVSC Avro schema to send to Schema Registry
const EventMetadataNewAVSC = "{\"type\":\"record\",\"name\":\"EventMetadataNew\",\"namespace\":\"com.sarrubia.eventbus.avro.example.embedded\",\"doc\":\"\",\"fields\":[{\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"},{\"name\":\"key\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}],\"default\":null},{\"name\":\"version\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"},{\"name\":\"creationLong\",\"type\":\"long\",\"default\":0},{\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"default\":\"NONE\"}]}"

// EventMetadataNewSubject schema registry subject
const EventMetadataNewSubject = "com.sarrubia.eventbus.avro.example.embedded.EventMetadataNew"

// EventMetadataNewBase
type EventMetadataNewBase struct {
	Id           string  `avro:"id"`
	Key          *string `avro:"key"`
	Version      string  `avro:"version"`
	CreationLong int64   `avro:"creationLong"`
	Type         string  `avro:"type"`
}

// Schema [REQUIRED] returns the schema const that belongs to EventMetadataNew
// Important: the receiver MUST NOT be a pointer in order to assert true against schema.EventSchemaInterface
func (e EventMetadataNewBase) Schema() string {
	return EventMetadataNewAVSC
}