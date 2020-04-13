package av

func (fmtctx *AVFormatContext) DumpFormat(file string) {
	avDumpFormat(fmtctx, file)
}
