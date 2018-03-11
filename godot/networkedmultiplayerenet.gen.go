package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

// NetworkedMultiplayerENetCompressionMode is an enum for CompressionMode values.
type NetworkedMultiplayerENetCompressionMode int

const (
	NetworkedMultiplayerENetCompressFastlz     NetworkedMultiplayerENetCompressionMode = 2
	NetworkedMultiplayerENetCompressNone       NetworkedMultiplayerENetCompressionMode = 0
	NetworkedMultiplayerENetCompressRangeCoder NetworkedMultiplayerENetCompressionMode = 1
	NetworkedMultiplayerENetCompressZlib       NetworkedMultiplayerENetCompressionMode = 3
	NetworkedMultiplayerENetCompressZstd       NetworkedMultiplayerENetCompressionMode = 4
)

//func NewNetworkedMultiplayerENetFromPointer(ptr gdnative.Pointer) NetworkedMultiplayerENet {
func newNetworkedMultiplayerENetFromPointer(ptr gdnative.Pointer) NetworkedMultiplayerENet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NetworkedMultiplayerENet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type NetworkedMultiplayerENet struct {
	NetworkedMultiplayerPeer
	owner gdnative.Object
}

func (o *NetworkedMultiplayerENet) BaseClass() string {
	return "NetworkedMultiplayerENet"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *NetworkedMultiplayerENet) CloseConnection() {
	//log.Println("Calling NetworkedMultiplayerENet.CloseConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "close_connection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ip String} { false port int} {0 true in_bandwidth int} {0 true out_bandwidth int}], Returns: enum.Error
*/
func (o *NetworkedMultiplayerENet) CreateClient(ip gdnative.String, port gdnative.Int, inBandwidth gdnative.Int, outBandwidth gdnative.Int) gdnative.Error {
	//log.Println("Calling NetworkedMultiplayerENet.CreateClient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(ip)
	ptrArguments[1] = gdnative.NewPointerFromInt(port)
	ptrArguments[2] = gdnative.NewPointerFromInt(inBandwidth)
	ptrArguments[3] = gdnative.NewPointerFromInt(outBandwidth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "create_client")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false port int} {32 true max_clients int} {0 true in_bandwidth int} {0 true out_bandwidth int}], Returns: enum.Error
*/
func (o *NetworkedMultiplayerENet) CreateServer(port gdnative.Int, maxClients gdnative.Int, inBandwidth gdnative.Int, outBandwidth gdnative.Int) gdnative.Error {
	//log.Println("Calling NetworkedMultiplayerENet.CreateServer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromInt(maxClients)
	ptrArguments[2] = gdnative.NewPointerFromInt(inBandwidth)
	ptrArguments[3] = gdnative.NewPointerFromInt(outBandwidth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "create_server")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.NetworkedMultiplayerENet::CompressionMode
*/
func (o *NetworkedMultiplayerENet) GetCompressionMode() NetworkedMultiplayerENetCompressionMode {
	//log.Println("Calling NetworkedMultiplayerENet.GetCompressionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_compression_mode")

	// Call the parent method.
	// enum.NetworkedMultiplayerENet::CompressionMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return NetworkedMultiplayerENetCompressionMode(ret)
}

/*
        Undocumented
	Args: [{ false ip String}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetBindIp(ip gdnative.String) {
	//log.Println("Calling NetworkedMultiplayerENet.SetBindIp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(ip)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_bind_ip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetCompressionMode(mode gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerENet.SetCompressionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_compression_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// NetworkedMultiplayerENetImplementer is an interface that implements the methods
// of the NetworkedMultiplayerENet class.
type NetworkedMultiplayerENetImplementer interface {
	NetworkedMultiplayerPeerImplementer
	CloseConnection()
	SetBindIp(ip gdnative.String)
	SetCompressionMode(mode gdnative.Int)
}
