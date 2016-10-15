package ckutils

// 서비스 전용으로 사용 할 에러값. 여기에 맞추어 에러 코드, 메시지가 출력 된다.
type CkError struct {
	Code	string
	Msg	string
}

// 서비스에서 사용할 에러를 만든다.
func NewError(code string, msg string) (*CkError) {
	return &CkError{
		Code: code,
		Msg: msg,
	}
}