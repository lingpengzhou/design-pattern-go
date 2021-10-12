package FilterChain

import "testing"

func Test_Filter_Chain(t *testing.T) {
	target := Target{}
	filterManager := FilterManager{}
	filterManager.SetFilterManager(target)
	authenticationFilter := AuthenticationFilter{}
	debugFilter := DebugFilter{}
	filterManager.SetFilter(&authenticationFilter)
	filterManager.SetFilter(&debugFilter)
	client := Client{filterManager: filterManager}
	client.sendRequest("HOME")
}
