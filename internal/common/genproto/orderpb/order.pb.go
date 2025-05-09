// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: orderpb/order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OrderStatus represents the current status of an order
type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_UNSPECIFIED OrderStatus = 0
	OrderStatus_ORDER_STATUS_PENDING     OrderStatus = 1
	OrderStatus_ORDER_STATUS_PAID        OrderStatus = 2
	OrderStatus_ORDER_STATUS_SHIPPED     OrderStatus = 3
	OrderStatus_ORDER_STATUS_DELIVERED   OrderStatus = 4
	OrderStatus_ORDER_STATUS_CANCELLED   OrderStatus = 5
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_STATUS_UNSPECIFIED",
		1: "ORDER_STATUS_PENDING",
		2: "ORDER_STATUS_PAID",
		3: "ORDER_STATUS_SHIPPED",
		4: "ORDER_STATUS_DELIVERED",
		5: "ORDER_STATUS_CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_UNSPECIFIED": 0,
		"ORDER_STATUS_PENDING":     1,
		"ORDER_STATUS_PAID":        2,
		"ORDER_STATUS_SHIPPED":     3,
		"ORDER_STATUS_DELIVERED":   4,
		"ORDER_STATUS_CANCELLED":   5,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_orderpb_order_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_orderpb_order_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{0}
}

// Order represents a customer order
type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	TotalAmount   float64                `protobuf:"fixed64,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Status        OrderStatus            `protobuf:"varint,5,opt,name=status,proto3,enum=orderpb.OrderStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_orderpb_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// OrderItem represents an item in an order
type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitPrice     float64                `protobuf:"fixed64,4,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Subtotal      float64                `protobuf:"fixed64,5,opt,name=subtotal,proto3" json:"subtotal,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_orderpb_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *OrderItem) GetSubtotal() float64 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

// CreateOrderRequest represents a request to create a new order
type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_orderpb_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// CreateOrderResponse represents a response to creating a new order
type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_orderpb_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

// GetOrderRequest represents a request to get an order by ID
type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_orderpb_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

// GetOrderResponse represents a response to getting an order
type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_orderpb_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

// ListOrdersRequest represents a request to list orders
type ListOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_orderpb_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrdersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListOrdersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListOrdersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListOrdersResponse represents a response to listing orders
type ListOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_orderpb_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *ListOrdersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// UpdateOrderStatusRequest represents a request to update an order's status
type UpdateOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        OrderStatus            `protobuf:"varint,2,opt,name=status,proto3,enum=orderpb.OrderStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusRequest) Reset() {
	*x = UpdateOrderStatusRequest{}
	mi := &file_orderpb_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusRequest) ProtoMessage() {}

func (x *UpdateOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrderStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderStatusRequest) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

// UpdateOrderStatusResponse represents a response to updating an order's status
type UpdateOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusResponse) Reset() {
	*x = UpdateOrderStatusResponse{}
	mi := &file_orderpb_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResponse) ProtoMessage() {}

func (x *UpdateOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateOrderStatusResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_orderpb_order_proto protoreflect.FileDescriptor

const file_orderpb_order_proto_rawDesc = "" +
	"\n" +
	"\x13orderpb/order.proto\x12\aorderpb\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa1\x02\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12(\n" +
	"\x05items\x18\x03 \x03(\v2\x12.orderpb.OrderItemR\x05items\x12!\n" +
	"\ftotal_amount\x18\x04 \x01(\x01R\vtotalAmount\x12,\n" +
	"\x06status\x18\x05 \x01(\x0e2\x14.orderpb.OrderStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\x95\x01\n" +
	"\tOrderItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\x12\x1d\n" +
	"\n" +
	"unit_price\x18\x04 \x01(\x01R\tunitPrice\x12\x1a\n" +
	"\bsubtotal\x18\x05 \x01(\x01R\bsubtotal\"W\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12(\n" +
	"\x05items\x18\x02 \x03(\v2\x12.orderpb.OrderItemR\x05items\";\n" +
	"\x13CreateOrderResponse\x12$\n" +
	"\x05order\x18\x01 \x01(\v2\x0e.orderpb.OrderR\x05order\",\n" +
	"\x0fGetOrderRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"8\n" +
	"\x10GetOrderResponse\x12$\n" +
	"\x05order\x18\x01 \x01(\v2\x0e.orderpb.OrderR\x05order\"h\n" +
	"\x11ListOrdersRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"d\n" +
	"\x12ListOrdersResponse\x12&\n" +
	"\x06orders\x18\x01 \x03(\v2\x0e.orderpb.OrderR\x06orders\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"c\n" +
	"\x18UpdateOrderStatusRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12,\n" +
	"\x06status\x18\x02 \x01(\x0e2\x14.orderpb.OrderStatusR\x06status\"A\n" +
	"\x19UpdateOrderStatusResponse\x12$\n" +
	"\x05order\x18\x01 \x01(\v2\x0e.orderpb.OrderR\x05order*\xae\x01\n" +
	"\vOrderStatus\x12\x1c\n" +
	"\x18ORDER_STATUS_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14ORDER_STATUS_PENDING\x10\x01\x12\x15\n" +
	"\x11ORDER_STATUS_PAID\x10\x02\x12\x18\n" +
	"\x14ORDER_STATUS_SHIPPED\x10\x03\x12\x1a\n" +
	"\x16ORDER_STATUS_DELIVERED\x10\x04\x12\x1a\n" +
	"\x16ORDER_STATUS_CANCELLED\x10\x052\xbc\x02\n" +
	"\fOrderService\x12H\n" +
	"\vCreateOrder\x12\x1b.orderpb.CreateOrderRequest\x1a\x1c.orderpb.CreateOrderResponse\x12?\n" +
	"\bGetOrder\x12\x18.orderpb.GetOrderRequest\x1a\x19.orderpb.GetOrderResponse\x12E\n" +
	"\n" +
	"ListOrders\x12\x1a.orderpb.ListOrdersRequest\x1a\x1b.orderpb.ListOrdersResponse\x12Z\n" +
	"\x11UpdateOrderStatus\x12!.orderpb.UpdateOrderStatusRequest\x1a\".orderpb.UpdateOrderStatusResponseB9Z7github.com/hw1513/ecomm/common/genproto/orderpb;orderpbb\x06proto3"

var (
	file_orderpb_order_proto_rawDescOnce sync.Once
	file_orderpb_order_proto_rawDescData []byte
)

func file_orderpb_order_proto_rawDescGZIP() []byte {
	file_orderpb_order_proto_rawDescOnce.Do(func() {
		file_orderpb_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orderpb_order_proto_rawDesc), len(file_orderpb_order_proto_rawDesc)))
	})
	return file_orderpb_order_proto_rawDescData
}

var file_orderpb_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orderpb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_orderpb_order_proto_goTypes = []any{
	(OrderStatus)(0),                  // 0: orderpb.OrderStatus
	(*Order)(nil),                     // 1: orderpb.Order
	(*OrderItem)(nil),                 // 2: orderpb.OrderItem
	(*CreateOrderRequest)(nil),        // 3: orderpb.CreateOrderRequest
	(*CreateOrderResponse)(nil),       // 4: orderpb.CreateOrderResponse
	(*GetOrderRequest)(nil),           // 5: orderpb.GetOrderRequest
	(*GetOrderResponse)(nil),          // 6: orderpb.GetOrderResponse
	(*ListOrdersRequest)(nil),         // 7: orderpb.ListOrdersRequest
	(*ListOrdersResponse)(nil),        // 8: orderpb.ListOrdersResponse
	(*UpdateOrderStatusRequest)(nil),  // 9: orderpb.UpdateOrderStatusRequest
	(*UpdateOrderStatusResponse)(nil), // 10: orderpb.UpdateOrderStatusResponse
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
}
var file_orderpb_order_proto_depIdxs = []int32{
	2,  // 0: orderpb.Order.items:type_name -> orderpb.OrderItem
	0,  // 1: orderpb.Order.status:type_name -> orderpb.OrderStatus
	11, // 2: orderpb.Order.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: orderpb.Order.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 4: orderpb.CreateOrderRequest.items:type_name -> orderpb.OrderItem
	1,  // 5: orderpb.CreateOrderResponse.order:type_name -> orderpb.Order
	1,  // 6: orderpb.GetOrderResponse.order:type_name -> orderpb.Order
	1,  // 7: orderpb.ListOrdersResponse.orders:type_name -> orderpb.Order
	0,  // 8: orderpb.UpdateOrderStatusRequest.status:type_name -> orderpb.OrderStatus
	1,  // 9: orderpb.UpdateOrderStatusResponse.order:type_name -> orderpb.Order
	3,  // 10: orderpb.OrderService.CreateOrder:input_type -> orderpb.CreateOrderRequest
	5,  // 11: orderpb.OrderService.GetOrder:input_type -> orderpb.GetOrderRequest
	7,  // 12: orderpb.OrderService.ListOrders:input_type -> orderpb.ListOrdersRequest
	9,  // 13: orderpb.OrderService.UpdateOrderStatus:input_type -> orderpb.UpdateOrderStatusRequest
	4,  // 14: orderpb.OrderService.CreateOrder:output_type -> orderpb.CreateOrderResponse
	6,  // 15: orderpb.OrderService.GetOrder:output_type -> orderpb.GetOrderResponse
	8,  // 16: orderpb.OrderService.ListOrders:output_type -> orderpb.ListOrdersResponse
	10, // 17: orderpb.OrderService.UpdateOrderStatus:output_type -> orderpb.UpdateOrderStatusResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_orderpb_order_proto_init() }
func file_orderpb_order_proto_init() {
	if File_orderpb_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orderpb_order_proto_rawDesc), len(file_orderpb_order_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderpb_order_proto_goTypes,
		DependencyIndexes: file_orderpb_order_proto_depIdxs,
		EnumInfos:         file_orderpb_order_proto_enumTypes,
		MessageInfos:      file_orderpb_order_proto_msgTypes,
	}.Build()
	File_orderpb_order_proto = out.File
	file_orderpb_order_proto_goTypes = nil
	file_orderpb_order_proto_depIdxs = nil
}
