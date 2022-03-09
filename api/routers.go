package api

import (
	"net/http"
	"strings"

	"tradebits/api/info"
	"tradebits/api/market"
	"tradebits/api/operator"
	"tradebits/api/user"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"InfoMarketGet",
		strings.ToUpper("Get"),
		"/info/market",
		info.InfoMarketGet,
	},

	Route{
		"InfoNetGet",
		strings.ToUpper("Get"),
		"/info/net",
		info.InfoNetGet,
	},

	Route{
		"MarketCloseDelete",
		strings.ToUpper("Delete"),
		"/market/close",
		market.MarketCloseDelete,
	},

	Route{
		"MarketCreatePut",
		strings.ToUpper("Put"),
		"/market/create",
		market.MarketCreatePut,
	},

	Route{
		"MarketDecreasePost",
		strings.ToUpper("Post"),
		"/market/decrease",
		market.MarketDecreasePost,
	},

	Route{
		"MarketRelatedGet",
		strings.ToUpper("Get"),
		"/market/related",
		market.MarketRelatedGet,
	},

	Route{
		"OperatorDepositPost",
		strings.ToUpper("Post"),
		"/operator/deposit",
		operator.OperatorDepositPost,
	},

	Route{
		"OperatorMessagePut",
		strings.ToUpper("Put"),
		"/operator/message",
		operator.OperatorMessagePut,
	},

	Route{
		"OperatorWithdrawalPost",
		strings.ToUpper("Post"),
		"/operator/withdrawal",
		operator.OperatorWithdrawalPost,
	},

	Route{
		"OperatorNewmessagesPost",
		strings.ToUpper("Post"),
		"/operator/newmessages",
		operator.OperatorNewmessagesPost,
	},

	Route{
		"UserBalanceGet",
		strings.ToUpper("Get"),
		"/user/balance",
		user.UserBalanceGet,
	},

	Route{
		"UserCancelordersPost",
		strings.ToUpper("Post"),
		"/user/cancelorders",
		user.UserCancelordersPost,
	},

	Route{
		"UserCreatePut",
		strings.ToUpper("Put"),
		"/user/create",
		user.UserCreatePut,
	},

	Route{
		"UserMessagePut",
		strings.ToUpper("Put"),
		"/user/message",
		user.UserMessagePut,
	},

	Route{
		"UserMessagesGet",
		strings.ToUpper("Get"),
		"/user/messages",
		user.UserMessagesGet,
	},

	Route{
		"UserOrderPost",
		strings.ToUpper("Post"),
		"/user/order",
		user.UserOrderPost,
	},

	Route{
		"UserTradesGet",
		strings.ToUpper("Get"),
		"/user/trades",
		user.UserTradesGet,
	},
}
