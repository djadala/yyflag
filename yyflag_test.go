package yyflag

import (
	"fmt"
	"testing"
	"time"
)

var now = time.Date(2013, time.June, 10, 23, 1, 2, 3, time.UTC)

const tf = `2006-01-02 15:04:05.999999999 -0700 MST`

type testData struct {
	in       string
	now, out string
}

var ta = []testData{
	{
		in:  "2017-11-03",
		out: `2017-11-03 00:00:00 +0000 UTC`,
		//now: `2013-06-10 00:00:00 +0000 UTC`,
	},
	{
		in:  "917-11-03",
		out: `1917-11-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "01-11-03",
		out: `2001-11-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "1-11-03",
		out: `2011-11-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "10-03",
		out: `2013-10-03 00:00:00 +0000 UTC`,
	},

	{
		in:  "03",
		out: `2013-06-03 00:00:00 +0000 UTC`,
	},

	{
		in:  "-20",
		out: `2013-05-21 00:00:00 +0000 UTC`,
	},
	{
		in:  "+40",
		out: `2013-07-20 00:00:00 +0000 UTC`,
	},

	{
		in:  "123",
		out: `2013-05-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "9-123",
		out: `2009-05-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "99-123",
		out: `1999-05-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "199-123",
		out: `2199-05-03 00:00:00 +0000 UTC`,
	},
	{
		in:  "1964-123",
		out: `1964-05-02 00:00:00 +0000 UTC`,
	},

	{
		in:  "y1234",
		out: `1234-01-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "y123",
		out: `2123-01-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "y98",
		out: `1998-01-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "y9",
		out: `2009-01-01 00:00:00 +0000 UTC`,
	},

	{
		in:  "y5678-02",
		out: `5678-02-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "y804-04",
		out: `1804-04-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "y44-03",
		out: `2044-03-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "y7-11",
		out: `2017-11-01 00:00:00 +0000 UTC`,
	},

	{
		in:  "m10",
		out: `2013-10-01 00:00:00 +0000 UTC`,
	},
	{
		in:  "02-29",
		out: `2012-02-29 00:00:00 +0000 UTC`,
	},
	{
		in:  "31",
		out: `2013-01-31 00:00:00 +0000 UTC`,
		now: `2013-02-01 00:00:00 +0000 UTC`,
	},

	///////////////////////////////////////////////////////////////////////////
	// time

	{
		in:  "#11:22:33.1234",
		out: `2013-06-10 11:22:33.1234 +0000 UTC`,
	},
	{
		in:  "T11:22:33",
		out: `2013-06-10 11:22:33 +0000 UTC`,
	},
	{
		in:  "T11:22",
		out: `2013-06-10 11:22:00 +0000 UTC`,
	},
	{
		in:  "T11",
		out: `2013-06-10 11:00:00 +0000 UTC`,
	},

	{
		in:  "#11:22:33.5678@EET",
		out: `2013-06-10 11:22:33.5678 +0300 EST`,
	},
	{
		in:  "@EET",
		out: `2013-06-10 00:00:00 +0300 EST`,
	},
	{
		in:  "@z",
		out: `2013-06-10 00:00:00 +0000 UTC`,
	},
	{
		in:  "@+0300",
		out: `2013-06-10 00:00:00 +0300 EET`,
	},
	{
		in:  "@+03:00",
		out: `2013-06-10 00:00:00 +0300 EET`,
	},
	{
		in:  "@l",
		now: time.Date(2013, time.June, 10, 23, 1, 2, 3, time.UTC).String(),
		out: time.Date(2013, time.June, 10, 0, 0, 0, 0, time.Local).String(),
	},
}

func Test_01(t *testing.T) {

	var (
		fld DT
		n   time.Time
	)

	for i := range ta {

		if ta[i].now == "" {
			n = now
		} else {
			var err error
			n, err = time.Parse(tf, ta[i].now)
			if err != nil {
				t.Fatal(err)
			}
		}
		fld.tt = n

		err := fld.Set(ta[i].in)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		o, err := time.Parse(tf, ta[i].out)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(fld.dt)
		if !fld.tt.Equal(o) {
			fmt.Println(fld.tt, o)
			t.Error("times dont match")
			t.Fail()

		}

	}

}
