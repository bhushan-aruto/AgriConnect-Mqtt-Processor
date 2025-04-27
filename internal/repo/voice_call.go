package repo

type VoiceCallRepo interface {
	MakeAlertCall(callAnswerApiUrl, callFrom, callTo string) error
}
