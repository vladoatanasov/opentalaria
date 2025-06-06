package protocol

import "github.com/google/uuid"

// PacketEncoder is the interface providing helpers for writing with Kafka's encoding rules.
// Types implementing Encoder only need to worry about calling methods like PutString,
// not about how a string is represented in Kafka.
type packetEncoder interface {
	// Primitives
	putInt8(in int8)
	putInt16(in int16)
	putUint16(in uint16)
	putInt32(in int32)
	putUint32(in uint32)
	putInt64(in int64)
	putFloat64(in float64)
	putVarint(in int64)
	putUVarint(in uint64)
	putCompactArrayLength(in int) error
	putArrayLength(in int) error
	putBool(in bool)

	// Collections
	putBytes(in []byte) error
	putVarintBytes(in []byte) error
	putCompactBytes(in []byte) error
	putRawBytes(in []byte) error
	putCompactString(in string) error
	putNullableCompactString(in *string) error
	putString(in string) error
	putNullableString(in *string) error
	putStringArray(in []string) error
	putCompactStringArray(in []string) error
	putCompactInt8Array(in []int8) error
	putCompactInt16Array(in []int16) error
	putCompactInt32Array(in []int32) error
	putNullableCompactInt32Array(in []int32) error
	putInt8Array(in []int8) error
	putInt16Array(in []int16) error
	putInt32Array(in []int32) error
	putInt64Array(in []int64) error
	putEmptyTaggedFieldArray()
	putUUID(in uuid.UUID) error
	putUUIDArray(in []uuid.UUID) error

	// Provide the current offset to record the batch size metric
	offset() int

	// Stacks, see PushEncoder
	push(in pushEncoder)
	pop() error
}

// PushEncoder is the interface for encoding fields like CRCs and lengths where the value
// of the field depends on what is encoded after it in the packet. Start them with PacketEncoder.Push() where
// the actual value is located in the packet, then PacketEncoder.Pop() them when all the bytes they
// depend upon have been written.
type pushEncoder interface {
	// Saves the offset into the input buffer as the location to actually write the calculated value when able.
	saveOffset(in int)

	// Returns the length of data to reserve for the output of this encoder (eg 4 bytes for a CRC32).
	reserveLength() int

	// Indicates that all required data is now available to calculate and write the field.
	// SaveOffset is guaranteed to have been called first. The implementation should write ReserveLength() bytes
	// of data to the saved offset, based on the data between the saved offset and curOffset.
	run(curOffset int, buf []byte) error
}

// dynamicPushEncoder extends the interface of pushEncoder for uses cases where the length of the
// fields itself is unknown until its value was computed (for instance varint encoded length
// fields).
type dynamicPushEncoder interface {
	pushEncoder

	// Called during pop() to adjust the length of the field.
	// It should return the difference in bytes between the last computed length and current length.
	adjustLength(currOffset int) int
}

func FlexibleEncoderFrom(pe packetEncoder) packetEncoder {
	fe := &flexibleEncoder{
		parent: pe,
	}

	return fe
}
