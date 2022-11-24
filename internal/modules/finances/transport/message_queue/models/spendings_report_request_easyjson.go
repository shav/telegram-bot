// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package finance_transport_mq_models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	date "github.com/shav/telegram-bot/internal/common/date"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF9ddb652DecodeGithubComShavTelegramBotInternalModulesFinancesTransportMessageQueueModels(in *jlexer.Lexer, out *SpendingReportRequestMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "UserId":
			out.UserId = int64(in.Int64())
		case "PeriodName":
			out.PeriodName = string(in.String())
		case "StartDate":
			easyjsonF9ddb652DecodeGithubComShavTelegramBotInternalCommonDate(in, &out.StartDate)
		case "EndDate":
			easyjsonF9ddb652DecodeGithubComShavTelegramBotInternalCommonDate(in, &out.EndDate)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF9ddb652EncodeGithubComShavTelegramBotInternalModulesFinancesTransportMessageQueueModels(out *jwriter.Writer, in SpendingReportRequestMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.UserId))
	}
	{
		const prefix string = ",\"PeriodName\":"
		out.RawString(prefix)
		out.String(string(in.PeriodName))
	}
	{
		const prefix string = ",\"StartDate\":"
		out.RawString(prefix)
		easyjsonF9ddb652EncodeGithubComShavTelegramBotInternalCommonDate(out, in.StartDate)
	}
	{
		const prefix string = ",\"EndDate\":"
		out.RawString(prefix)
		easyjsonF9ddb652EncodeGithubComShavTelegramBotInternalCommonDate(out, in.EndDate)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SpendingReportRequestMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF9ddb652EncodeGithubComShavTelegramBotInternalModulesFinancesTransportMessageQueueModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SpendingReportRequestMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF9ddb652EncodeGithubComShavTelegramBotInternalModulesFinancesTransportMessageQueueModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SpendingReportRequestMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF9ddb652DecodeGithubComShavTelegramBotInternalModulesFinancesTransportMessageQueueModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SpendingReportRequestMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF9ddb652DecodeGithubComShavTelegramBotInternalModulesFinancesTransportMessageQueueModels(l, v)
}
func easyjsonF9ddb652DecodeGithubComShavTelegramBotInternalCommonDate(in *jlexer.Lexer, out *date.Date) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Year":
			out.Year = int(in.Int())
		case "Month":
			out.Month = int(in.Int())
		case "Day":
			out.Day = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF9ddb652EncodeGithubComShavTelegramBotInternalCommonDate(out *jwriter.Writer, in date.Date) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Year\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Year))
	}
	{
		const prefix string = ",\"Month\":"
		out.RawString(prefix)
		out.Int(int(in.Month))
	}
	{
		const prefix string = ",\"Day\":"
		out.RawString(prefix)
		out.Int(int(in.Day))
	}
	out.RawByte('}')
}
