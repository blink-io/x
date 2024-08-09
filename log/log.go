package log

import "github.com/phuslu/log"

type (
	Level = log.Level

	Context = log.Context

	Entry = log.Entry

	Fields = log.Fields

	Logger = log.Logger

	TSVEntry = log.TSVEntry

	FormatterArgs = log.FormatterArgs

	IOWriteCloser = log.IOWriteCloser

	LogfmtFormatter = log.LogfmtFormatter

	TSVLogger = log.TSVLogger

	ObjectMarshaler = log.ObjectMarshaler

	Writer = log.Writer

	AsyncWriter = log.AsyncWriter

	ConsoleWriter = log.ConsoleWriter

	FileWriter = log.FileWriter

	JournalWriter = log.JournalWriter

	IOWriter = log.IOWriter

	MultiEntryWriter = log.MultiEntryWriter

	MultiIOWriter = log.MultiIOWriter

	MultiLevelWriter = log.MultiLevelWriter

	MultiWriter = log.MultiWriter

	SyslogWriter = log.SyslogWriter

	WriterFunc = log.WriterFunc

	XID = log.XID
)

const (
	TraceLevel       = log.TraceLevel
	DebugLevel       = log.DebugLevel
	InfoLevel        = log.InfoLevel
	WarnLevel        = log.WarnLevel
	ErrorLevel       = log.ErrorLevel
	FatalLevel       = log.FatalLevel
	PanicLevel       = log.PanicLevel
	NoLevel    Level = 8

	ErrInvalidXID = log.ErrInvalidXID

	TimeFormatUnix = log.TimeFormatUnix

	TimeFormatUnixMs = log.TimeFormatUnixMs

	TimeFormatUnixWithMs = log.TimeFormatUnixWithMs
)

var (
	DefaultLogger = log.DefaultLogger

	SlogNewJSONHandler = log.SlogNewJSONHandler

	Printf = log.Printf

	Fastrandn = log.Fastrandn

	IsTerminal = log.IsTerminal

	ErrAsyncWriterFull = log.ErrAsyncWriterFull

	ParseLevel = log.ParseLevel
)
