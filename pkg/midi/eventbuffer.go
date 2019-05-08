package midi

//EncodeEventBuffer encodes a slice of midi.Event into a byte buffer
func EncodeEventBuffer(eventBuffer []Event) []byte {
	byteBuffer := make([]byte, 0, len(eventBuffer)<<3+1) // approx required buffer size

	for _, event := range eventBuffer {
		eventData := event.Data()
		eventBytes := make([]byte, 5, len(eventData.Data)+5)
		eventBytes[0] = byte(len(eventData.Data))
		eventBytes = append(eventBytes, eventData.Data...)
		byteBuffer = append(byteBuffer, eventBytes...)
	}
	byteBuffer = append(byteBuffer, 0) // termination byte

	return byteBuffer
}

//DecodeByteBuffer decodes a byte buffer to a slice of midi.Event
func DecodeByteBuffer(byteBuffer []byte) []Event {
	eventBuffer := make([]Event, 0, (len(byteBuffer)-1)>>3) // approx required buffer size
	byteIndex := 0

	for byteCount := int(byteBuffer[byteIndex]); byteCount > 0; byteCount = int(byteBuffer[byteIndex]) {
		byteSlice := byteBuffer[byteIndex+5 : byteIndex+5+byteCount]

		event := MakeMidiEvent(0, byteSlice)
		eventBuffer = append(eventBuffer, event)
		byteIndex = byteIndex + byteCount + 5
	}

	return eventBuffer
}
