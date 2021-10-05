package nc

type NoConnectionType struct{}

func (e *NoConnectionType) Error() string {
	return "connection type is not specified, use -t for TCP and -u for UDP"
}

type ArgsError struct{}

func (e *ArgsError) Error() string {
	return "argument quantity isn't correct, request should look like: nc -t(TCP)/-u(UDP) URL PORT"
}
