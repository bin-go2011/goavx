package av

func NewAVPacket() (*AVPacket, error) {
	pkt, err := AvAllocPacket()
	if err != nil {
		return nil, err
	}
	return pkt, nil
}

func (pkt *AVPacket) Release() {
	AvFreePacket(pkt)
}

func (pkt *AVPacket) StreamIndex() int32 {
	return pkt.stream_index
}
