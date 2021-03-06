/**
*  @file
*  @copyright defined in go-toolkit/LICENSE
 */

package rpc

import (
	"net/rpc"
)

// MetadataAPI is a default service for RegisterName.
const MetadataAPI = "rpc"

// Server represents a RPC server
type Server struct {
	rpc.Server
}

// API is a collection of methods for the RPC interface.
type API struct {
	// namespace of service
	Namespace string
	// api version
	Version string
	// the service methods holder
	Service interface{}
	// indication if the methods must be considered safe for public use
	Public bool
}

// RPCService offers meta information of the server.
//type RPCService struct {
//	server *Server
//}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{
		rpc.Server{},
	}

	// Not implemented service
	// register a default service which will provide meta information about the RPC service.
	//rpcService := &RPCService{server}
	//server.RegisterName(MetadataAPI, rpcService)

	return server
}
