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

func (dmx *Demuxer) FindBestAudioStream() (id int32, err error) {
	id, err = AvFindBestStream(dmx.fmtctx, AVMEDIA_TYPE_AUDIO)
	dmx.AudioStreamIndex = id
	return
}

func (dmx *Demuxer) FindBestVideoStream() (id int32, err error) {
	id, err = AvFindBestStream(dmx.fmtctx, AVMEDIA_TYPE_VIDEO)
	dmx.VideoStreamIndex = id
	return
}

func (dmx *Demuxer) DumpFormat() {
	dmx.fmtctx.DumpFormat(dmx.file)
}

func (dmx *Demuxer) ReadFrame(pkt *AVPacket) int32 {
	return AvReadFrame(dmx.fmtctx, pkt)
}

func (dmx *Demuxer) OpenAudioDecoder() (*AudioDecoder, error) {
	streamIndex, err := dmx.FindBestAudioStream()
	if streamIndex < 0 {
		return nil, err
	}

	avctx, err := AvcodecOpenContext(dmx.fmtctx, streamIndex)
	if err != nil {
		return nil, err
	}

	dec := &AudioDecoder{
		avctx,
	}

	return dec, nil

}
