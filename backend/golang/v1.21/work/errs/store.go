package errs

type HttpResponseCode string

const (
	E0000 HttpResponseCode = "E0000" // JSON unmarsh error
	E0001 HttpResponseCode = "E0001" // user name error
	E0002 HttpResponseCode = "E0002" // user password error
	E0003 HttpResponseCode = "E0003" // user nickname error
	E0004 HttpResponseCode = "E0004" // gen token error
	E0005 HttpResponseCode = "E0005" // get token error
	E0006 HttpResponseCode = "E0006" // check admin lv error
	E0007 HttpResponseCode = "E0006" // File write error
)

const (
	E1000 HttpResponseCode = "E1000" // sql register error
	E1001 HttpResponseCode = "E1001" // sql login error
	E1002 HttpResponseCode = "E1002" // sql is_admin error
	E1003 HttpResponseCode = "E1003" // sql add shop error
	E1004 HttpResponseCode = "E1004" // sql get shop error
)

const (
	E2000 HttpResponseCode = "E2000" // redis set token error
	E2001 HttpResponseCode = "E2001" // redis del token error
	E2002 HttpResponseCode = "E2002" // redis get token error
)
