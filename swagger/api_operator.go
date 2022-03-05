/*
 * TradeBits
 *
 * In this API description you can find description of ways markets communicate with each other and users. Each sign of message user send to market is additionally reinforced by market name as indtent to ensure markets will not try to proceed same operations with users on other markets. That is not mention in API, but implemented on server and client sides. Each procedure on specific market need to be reinforced by market name in senders message, to prevent mimic behavior by markets collecting data. That will guarantee, that only owner of the private key is able to process transactions on specific market.  ## API notions: - **ukey** - public key of the user, which is used by user to send messages and process operations on a related market  - **mkey** - public key of the market, which can be used to encrypt and verify messages - **hkey** - key of host added for additional security, for message verification. This is the public key of the node accepting the request 
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
)

func OperatorDepositPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func OperatorMessagePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func OperatorWithdrawalPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
