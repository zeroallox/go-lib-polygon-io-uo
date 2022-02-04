package polyrest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parseResponseBody(t *testing.T) {

	const expectedErrorString = "Unknown database error ( Error Code: 001 )"

	type args struct {
		src  []byte
		dest *Response
	}
	tests := []struct {
		name              string
		args              args
		wantErr           bool
		wantResult        []byte
		wantARErrorCode   string
		wantARErrorString string
		wantARErrorError  error
	}{
		{
			name: "ok one string",
			args: args{
				src: []byte(`
                    {
                        "error": "Unknown database error ( Error Code: 001 )",
                        "errorcode": "1",
                        "status": "ERROR",
                        "success": false
                    }
                `),
				dest: new(Response),
			},
			wantResult:        nil,
			wantErr:           false,
			wantARErrorCode:   "1",
			wantARErrorString: expectedErrorString,
			wantARErrorError:  errors.New(expectedErrorString),
		},
		{
			name: "ok zerozeroone string",
			args: args{
				src: []byte(`
                    {
                        "error": "Unknown database error ( Error Code: 001 )",
                        "errorcode": "001",
                        "status": "ERROR",
                        "success": false
                    }
                `),
				dest: new(Response),
			},
			wantResult:        nil,
			wantErr:           false,
			wantARErrorCode:   "001",
			wantARErrorString: expectedErrorString,
			wantARErrorError:  errors.New(expectedErrorString),
		},
		{
			name: "ok one int",
			args: args{
				src: []byte(`
                    {
                        "error": "Unknown database error ( Error Code: 001 )",
                        "errorcode": 1,
                        "status": "ERROR",
                        "success": false
                    }
                `),
				dest: new(Response),
			},
			wantResult:        nil,
			wantErr:           true,
			wantARErrorCode:   "1",
			wantARErrorString: expectedErrorString,
			wantARErrorError:  errors.New(expectedErrorString),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := parseResponseBody(tt.args.src, tt.args.dest)

			if tt.wantErr == false {
				require.Nil(t, err)
			}

			require.Equal(t, results, tt.wantResult)
			require.Equal(t, tt.wantARErrorString, tt.args.dest.ErrorString())
			require.Equal(t, tt.wantARErrorCode, tt.args.dest.ErrorCode())
			require.Equal(t, tt.wantARErrorError, tt.args.dest.Error())
		})
	}
}
