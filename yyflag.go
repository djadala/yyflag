// Package yyflag implement flag.Getter interface for time.Time values,
// allowing date/time to be entered as flags at command line.
// yyflag allow all incomplete formats of package djadala/yy to be used.
//
// yyflag recognize following formats:
//  ['date'][#'time'][@'zone'] or ['date'][T'time'][@'zone']
//
//  'date' is one of:
//   +ddd               relative days
//   -ddd               relative days
//   jjj                julian year day
//   y-jjj              year 1 digit + julian year day
//   yy-jjj             year 2 digit + julian year day
//   yyy-jjj            year 3 digit + julian year day
//   yyyy-jjj           year + julian year day
//   dd                 day
//   mm-dd              month + day
//   y-mm-dd            year 1 digit + month + day
//   yy-mm-dd           year 2 digit + month + day
//   yyy-mm-dd          year 3 digit + month + day
//   yyyy-mm-dd         year + month + day
//   mm                 month
//   y-mm               year 1 digit + month
//   yy-mm              year 2 digit + month
//   yyy-mm             year 3 digit + month
//   yyyy-mm            year + month
//   y                  year 1 digit
//   yy                 year 2 digit
//   yyy                year 3 digit
//   yyyy               year
//
//  'time' is one of:
//   hh
//   hh:mm
//   hh:mm:ss
//   hh:mm:ss.d+
//
//  'timezone' is one of:
//   +/-hh:mm
//   +/-hhmm
//   timezoneName
//   z                  UTC
//   l                  Local
//
package yyflag

import (
	"fmt"
	"regexp"
	"time"

	"github.com/djadala/yy"
)

//go:generate sh -c "go run makeregexp.go > regexp.go"

var reAll = regexp.MustCompile(reStr)

type dt struct {
	tt time.Time
}

// return new flag.Getter for time.Time, with default time t
func New(t time.Time) dt {
	return dt{tt: t}
}

// return time.Time value
func (d *dt) Time() time.Time {
	return d.tt
}

// implement flag.Value interface
func (d *dt) String() string {
	return d.tt.String()
}

// implement flag.Value interface
func (d *dt) Set(v string) error {
	result := reAll.FindSubmatch([]byte(v))
	if len(result) == 0 {
		return fmt.Errorf("invalid date format %s", v)
	}
	var (
		par yy.IDate
	)
	for i := range result {
		if len(result[i]) != 0 && len(reAll.SubexpNames()[i]) != 0 {
			var err error
			switch reAll.SubexpNames()[i] {
			case `rel`:
				err = par.R.Set(result[i+1])
			case `jjj`:
				err = par.J.Set(result[i+1])
			case `yjjj`:
				err = checkError(par.Y.Set(result[i+1]), par.J.Set(result[i+2]))
			case `yyjjj`:
				err = checkError(par.Y.Set(result[i+1]), par.J.Set(result[i+2]))
			case `yyyjjj`:
				err = checkError(par.Y.Set(result[i+1]), par.J.Set(result[i+2]))
			case `yyyyjjj`:
				err = checkError(par.Y.Set(result[i+1]), par.J.Set(result[i+2]))
			case `dd`:
				err = par.D.Set(result[i+1])
			case `mmdd`:
				err = checkError(par.Mo.Set(result[i+1]), par.D.Set(result[i+2]))
			case `ymmdd`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]), par.D.Set(result[i+3]))
			case `yymmdd`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]), par.D.Set(result[i+3]))
			case `yyymmdd`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]), par.D.Set(result[i+3]))
			case `yyyymmdd`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]), par.D.Set(result[i+3]))
			case `yyyy`:
				err = par.Y.Set(result[i+1])

			case `yyy`:
				err = par.Y.Set(result[i+1])
			case `yy`:
				err = par.Y.Set(result[i+1])
			case `y`:
				err = par.Y.Set(result[i+1])

			case `yyyymm`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]))

			case `yyymm`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]))
			case `yymm`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]))
			case `ymm`:
				err = checkError(par.Y.Set(result[i+1]), par.Mo.Set(result[i+2]))

			case `mm`:
				err = par.Mo.Set(result[i+1])

			case `hh`:
				err = par.H.Set(result[i+1])
			case `hhmm`:
				err = checkError(par.H.Set(result[i+1]), par.M.Set(result[i+2]))
			case `hhmmss`:
				err = checkError(par.H.Set(result[i+1]), par.M.Set(result[i+2]), par.S.Set(result[i+3]))
			case `hhmmssf`:
				err = checkError(
					par.H.Set(result[i+1]),
					par.M.Set(result[i+2]),
					par.S.Set(result[i+3]),
					par.F.Set(result[i+4]),
				)

			case `tzn`:
				err = par.L.SetHHMM(result[i+1], result[i+2])
			case `tzs`:
				err = par.L.SetName(result[i+1])
			default:
				panic("SubexpNames ??:" + reAll.SubexpNames()[i])
			}
			if err != nil {
				return err
			}
		}
	}
	r, e := yy.Convert(d.tt, &par)
	if e == nil {
		d.tt = r
	}
	return e
}

// implement flag.Getter interface
func (d *dt) Get() interface{} {
	return d.tt
}

func checkError(err ...error) error {
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return nil
}
