// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// PartitionResult_WriteShareGroupStateResponse contains the results for the partitions.
type PartitionResult_WriteShareGroupStateResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Partition contains the partition index.
	Partition int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
}

func (p *PartitionResult_WriteShareGroupStateResponse) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Partition)

	pe.putInt16(p.ErrorCode)

	if err := pe.putNullableString(p.ErrorMessage); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (p *PartitionResult_WriteShareGroupStateResponse) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Partition, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// WriteStateResult contains the write results.
type WriteStateResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicID contains the topic identifier.
	TopicID uuid.UUID
	// Partitions contains the results for the partitions.
	Partitions []PartitionResult_WriteShareGroupStateResponse
}

func (r *WriteStateResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putUUID(r.TopicID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Partitions)); err != nil {
		return err
	}
	for _, block := range r.Partitions {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *WriteStateResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicID, err = pd.getUUID(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		r.Partitions = make([]PartitionResult_WriteShareGroupStateResponse, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block PartitionResult_WriteShareGroupStateResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Partitions[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type WriteShareGroupStateResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Results contains the write results.
	Results []WriteStateResult
}

func (r *WriteShareGroupStateResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *WriteShareGroupStateResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResults > 0 {
		r.Results = make([]WriteStateResult, numResults)
		for i := 0; i < numResults; i++ {
			var block WriteStateResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Results[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *WriteShareGroupStateResponse) GetKey() int16 {
	return 85
}

func (r *WriteShareGroupStateResponse) GetVersion() int16 {
	return r.Version
}

func (r *WriteShareGroupStateResponse) SetVersion(version int16) {
	r.Version = version
}

func (r *WriteShareGroupStateResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *WriteShareGroupStateResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *WriteShareGroupStateResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
