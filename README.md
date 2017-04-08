# yyflag

Package yyflag implement flag.Getter interface for time.Time values,
allowing date/time to be entered as flags at command line.
yyflag allow all incomplete formats of package djadala/yy to be used.

yyflag recognize following formats:
 ['date'][#'time'][@'zone'] or ['date'][T'time'][@'zone']

 'date' is one of:
-  +ddd               relative days
-  -ddd               relative days
-  jjj                julian year day
-  y-jjj              year 1 digit + julian year day
-  yy-jjj             year 2 digit + julian year day
-  yyy-jjj            year 3 digit + julian year day
-  yyyy-jjj           year + julian year day
-  dd                 day
-  mm-dd              month + day
-  y-mm-dd            year 1 digit + month + day
-  yy-mm-dd           year 2 digit + month + day
-  yyy-mm-dd          year 3 digit + month + day
-  yyyy-mm-dd         year + month + day
-  mm                 month
-  y-mm               year 1 digit + month
-  yy-mm              year 2 digit + month
-  yyy-mm             year 3 digit + month
-  yyyy-mm            year + month
-  y                  year 1 digit
-  yy                 year 2 digit
-  yyy                year 3 digit
-  yyyy               year

 'time' is one of:
-  hh
-  hh:mm
-  hh:mm:ss
-  hh:mm:ss.d+

- 'timezone' is one of:
-  +/-hh:mm
-  +/-hhmm
-  timezoneName
-  z                  UTC
-  l                  Local

