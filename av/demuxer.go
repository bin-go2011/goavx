package av

type Demuxer struct {
	fmtctx           *AVFormatContext
	file             string
	AudioStreamIndex int32
	VideoStreamIndex int32
}

func NewDemuxer(file string) (*Demuxer, error) {
	fmtctx, err := AvformatOpenInput(file)
	if fmtctx == nil {
		return nil, err
	}
	dmx := &Demuxer{
		fmtctx,
		file,
		-1,
		-1,
	}
	return dmx, nil
}

func (dmx *Demuxer) Release() {
	AvformatCloseInput(dmx.fmtctx)
}

func (dmx *Demuxer) FindBestAudioStream() int32 {
	return AvFindBestStream(dmx.fmtctx, 1)
}

func (dmx *Demuxer) FindBestVideoStream() int32 {
	return AvFindBestStream(dmx.fmtctx, 0)
}

func (dmx *Demuxer) DumpFormat() {
	dmx.fmtctx.DumpFormat(dmx.file)
}

func (dmx *Demuxer) ReadFrame(pkt *AVPacket) int32 {
	return AvReadFrame(dmx.fmtctx, pkt)
}

func (dmx *Demuxer) OpenAudioDecoder() (*AudioDecoder, error) {
	audioStreamIndex := dmx.FindBestAudioStream()

	avctx, err := AvcodecOpenContext(dmx.fmtctx, audioStreamIndex)
	if err != nil {
		return nil, err
	}

	dmx.AudioStreamIndex = audioStreamIndex

	dec := &AudioDecoder{
		avctx,
	}

	return dec, nil

}
