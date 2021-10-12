package FilterChain

import "log"

type Filter interface {
	execute(request string)
}

type Target struct {
}

func (target *Target) execute(request string) {
	println("target is requesting", request)
}

type AuthenticationFilter struct {
}

func (authenticationFilter *AuthenticationFilter) execute(request string) {
	log.Println("Authenticating request:" + request)
}

type DebugFilter struct {
}

func (debugFilter *DebugFilter) execute(request string) {
	log.Println("DebugFiltering request:" + request)
}

type FilterChain struct {
	filterList []*Filter
	target     Target
}

func (filterChain *FilterChain) addFilter(filter *Filter) {
	filterChain.filterList = append(filterChain.filterList, filter)
}

func (filterChain *FilterChain) execute(request string) {
	for _, filter := range filterChain.filterList {
		(*filter).execute(request)
	}
	filterChain.target.execute(request)
}

type FilterManager struct {
	filterChain FilterChain
}

func (filterManager *FilterManager) SetFilterManager(target Target) {
	filterManager.filterChain.target = target

}

func (filterManager *FilterManager) SetFilter(filter Filter) {
	filterManager.filterChain.addFilter(&filter)
}

func (filterManager *FilterManager) sendRequest(request string) {
	filterManager.filterChain.execute(request)
}

type Client struct {
	filterManager FilterManager
}

func (client *Client) sendRequest(request string) {
	client.filterManager.sendRequest(request)
}
