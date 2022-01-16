package polypsv

import (
    "github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
    "reflect"
    "testing"
    "time"
)

func TestNewFileFromPath(t *testing.T) {
    type args struct {
        path string
    }
    tests := []struct {
        name    string
        args    args
        want    *FileInfo
        wantErr bool
    }{
        {
            name: "ok compressed",
            args: args{path: "polygon/us/stocks/trades/2020/2020-03/us-stocks-trades-2020-03-26.psv.gz"},
            want: &FileInfo{
                locale:     polyconst.LOC_USA,
                market:     polyconst.MKT_Stocks,
                dataType:   polyconst.DT_Trades,
                date:       time.Date(2020, 03, 26, 0, 0, 0, 0, polyconst.NYCTime),
                compressed: true,
            },
        },
        {
            name: "ok uncompressed",
            args: args{path: "polygon/us/stocks/trades/2020/2020-03/us-stocks-trades-2020-03-26.psv"},
            want: &FileInfo{
                locale:     polyconst.LOC_USA,
                market:     polyconst.MKT_Stocks,
                dataType:   polyconst.DT_Trades,
                date:       time.Date(2020, 03, 26, 0, 0, 0, 0, polyconst.NYCTime),
                compressed: false,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := NewFileFromPath(tt.args.path)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewFileFromPath() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewFileFromPath() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestMakeDirPath(t *testing.T) {
    type args struct {
        file *FileInfo
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "ok",
            args: args{file: &FileInfo{
                locale:     polyconst.LOC_USA,
                market:     polyconst.MKT_Stocks,
                dataType:   polyconst.DT_Trades,
                date:       time.Date(2000, 01, 01, 0, 0, 0, 0, polyconst.NYCTime),
                compressed: true,
            }},
            want: "polygon/us/stocks/trades/2000/2000-01",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MakeDirPath(tt.args.file); got != tt.want {
                t.Errorf("MakeDirPath() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_makeFileName(t *testing.T) {
    type args struct {
        file       *FileInfo
        compressed bool
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "ok compressed",
            args: args{
                file: &FileInfo{
                    locale:   polyconst.LOC_USA,
                    market:   polyconst.MKT_Stocks,
                    dataType: polyconst.DT_Trades,
                    date:     time.Date(2000, 01, 01, 0, 0, 0, 0, polyconst.NYCTime),
                },
                compressed: true,
            },
            want: "us-stocks-trades-2000-01-01.psv.gz",
        },
        {
            name: "ok uncompressed",
            args: args{
                file: &FileInfo{
                    locale:   polyconst.LOC_USA,
                    market:   polyconst.MKT_Stocks,
                    dataType: polyconst.DT_Trades,
                    date:     time.Date(2000, 01, 01, 0, 0, 0, 0, polyconst.NYCTime),
                },
                compressed: false,
            },
            want: "us-stocks-trades-2000-01-01.psv",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MakeFileName(tt.args.file, tt.args.compressed); got != tt.want {
                t.Errorf("makeFileName() = %v, want %v", got, tt.want)
            }
        })
    }
}
