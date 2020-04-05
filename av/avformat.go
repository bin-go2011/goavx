package av

func (fmtctx *AVFormatContext) DumpFormat(file string) {
	AvDumpFormat(fmtctx, file)
}
