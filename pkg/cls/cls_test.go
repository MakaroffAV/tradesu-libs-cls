// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package cls

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"errors"
	"testing"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type TestCaseLog struct {
	a map[string][]Log
}

// ------------------------------------------------------------------------ //

var (
	testCasesLog = TestCaseLog{
		a: map[string][]Log{
			"INFO": {
				Log{
					Msg: "InfoMsg1",
				},
				Log{
					Msg: "InfoMsg2",
					Add: map[string]string{"InfoMsg2Key1": "InfoMsg2Val1"},
				},
				Log{
					Msg: `
					i
					am here`,
					Add: map[string]string{"WarnMsg2Key1": "WarnMsg2Val1"},
				},
			},
			"WARN": {
				Log{
					Msg: "WarnMsg1",
					Err: errors.New("WarnMsg1Err"),
				},
				Log{
					Msg: "WarnMsg2",
					Err: errors.New("WarnMsg2Err"),
					Add: map[string]string{"WarnMsg2Key1": "WarnMsg2Val1"},
				},
				Log{
					Msg: `
					i
					am here`,
					Err: errors.New("WarnMsg2Err"),
					Add: map[string]string{"WarnMsg2Key1": "WarnMsg2Val1"},
				},
			},
			"FAIL": {
				Log{
					Msg: "FailMsg1",
					Err: errors.New("FailMsg1Err"),
				},
				Log{
					Msg: "FailMsg2",
					Err: errors.New("FailMsg2Err"),
					Add: map[string]string{"FailMsg2ErrKey1": "FailMsg2ErrVal1"},
				},
				Log{
					Msg: `
					i
					am here`,
					Err: errors.New("WarnMsg2Err"),
					Add: map[string]string{"WarnMsg2Key1": "WarnMsg2Val1"},
				},
			},
		},
	}
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func TestInfo(t *testing.T) {
	for _, testCase := range testCasesLog.a["INFO"] {
		Info(testCase)
	}
}

// ------------------------------------------------------------------------ //

func TestWarn(t *testing.T) {
	for _, testCase := range testCasesLog.a["WARN"] {
		Warn(testCase)
	}
}

// ------------------------------------------------------------------------ //

func TestFail(t *testing.T) {
	for _, testCase := range testCasesLog.a["FAIL"] {
		Fail(testCase)
	}
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func TestRemExtraSpace(t *testing.T) {

	testCases := []struct {
		want string
		arg  string
	}{
		{
			want: "here",
			arg: `
			here`,
		},
		{
			want: "here with trail",
			arg: `
			here
			with trail`,
		},
	}

	for _, c := range testCases {
		r := remExtraSpace(c.arg)
		if r != c.want {
			t.Errorf(
				`
				Test failed: 	
								should be "%s", not "%s"
				`,
				c.want,
				r,
			)
		}
	}
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
