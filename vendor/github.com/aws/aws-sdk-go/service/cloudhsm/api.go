// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudhsm provides a client for Amazon CloudHSM.
package cloudhsm

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opCreateHapg = "CreateHapg"

// CreateHapgRequest generates a request for the CreateHapg operation.
func (c *CloudHSM) CreateHapgRequest(input *CreateHapgInput) (req *request.Request, output *CreateHapgOutput) {
	op := &request.Operation{
		Name:       opCreateHapg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHapgInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateHapgOutput{}
	req.Data = output
	return
}

// Creates a high-availability partition group. A high-availability partition
// group is a group of partitions that spans multiple physical HSMs.
func (c *CloudHSM) CreateHapg(input *CreateHapgInput) (*CreateHapgOutput, error) {
	req, out := c.CreateHapgRequest(input)
	err := req.Send()
	return out, err
}

const opCreateHsm = "CreateHsm"

// CreateHsmRequest generates a request for the CreateHsm operation.
func (c *CloudHSM) CreateHsmRequest(input *CreateHsmInput) (req *request.Request, output *CreateHsmOutput) {
	op := &request.Operation{
		Name:       opCreateHsm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHsmInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateHsmOutput{}
	req.Data = output
	return
}

// Creates an uninitialized HSM instance.
//
// There is an upfront fee charged for each HSM instance that you create with
// the CreateHsm operation. If you accidentally provision an HSM and want to
// request a refund, delete the instance using the DeleteHsm operation, go to
// the AWS Support Center (https://console.aws.amazon.com/support/home#/), create
// a new case, and select Account and Billing Support.
//
//   It can take up to 20 minutes to create and provision an HSM. You can monitor
// the status of the HSM with the DescribeHsm operation. The HSM is ready to
// be initialized when the status changes to RUNNING.
func (c *CloudHSM) CreateHsm(input *CreateHsmInput) (*CreateHsmOutput, error) {
	req, out := c.CreateHsmRequest(input)
	err := req.Send()
	return out, err
}

const opCreateLunaClient = "CreateLunaClient"

// CreateLunaClientRequest generates a request for the CreateLunaClient operation.
func (c *CloudHSM) CreateLunaClientRequest(input *CreateLunaClientInput) (req *request.Request, output *CreateLunaClientOutput) {
	op := &request.Operation{
		Name:       opCreateLunaClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLunaClientInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateLunaClientOutput{}
	req.Data = output
	return
}

// Creates an HSM client.
func (c *CloudHSM) CreateLunaClient(input *CreateLunaClientInput) (*CreateLunaClientOutput, error) {
	req, out := c.CreateLunaClientRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteHapg = "DeleteHapg"

// DeleteHapgRequest generates a request for the DeleteHapg operation.
func (c *CloudHSM) DeleteHapgRequest(input *DeleteHapgInput) (req *request.Request, output *DeleteHapgOutput) {
	op := &request.Operation{
		Name:       opDeleteHapg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHapgInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteHapgOutput{}
	req.Data = output
	return
}

// Deletes a high-availability partition group.
func (c *CloudHSM) DeleteHapg(input *DeleteHapgInput) (*DeleteHapgOutput, error) {
	req, out := c.DeleteHapgRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteHsm = "DeleteHsm"

// DeleteHsmRequest generates a request for the DeleteHsm operation.
func (c *CloudHSM) DeleteHsmRequest(input *DeleteHsmInput) (req *request.Request, output *DeleteHsmOutput) {
	op := &request.Operation{
		Name:       opDeleteHsm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHsmInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteHsmOutput{}
	req.Data = output
	return
}

// Deletes an HSM. After completion, this operation cannot be undone and your
// key material cannot be recovered.
func (c *CloudHSM) DeleteHsm(input *DeleteHsmInput) (*DeleteHsmOutput, error) {
	req, out := c.DeleteHsmRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteLunaClient = "DeleteLunaClient"

// DeleteLunaClientRequest generates a request for the DeleteLunaClient operation.
func (c *CloudHSM) DeleteLunaClientRequest(input *DeleteLunaClientInput) (req *request.Request, output *DeleteLunaClientOutput) {
	op := &request.Operation{
		Name:       opDeleteLunaClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLunaClientInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteLunaClientOutput{}
	req.Data = output
	return
}

// Deletes a client.
func (c *CloudHSM) DeleteLunaClient(input *DeleteLunaClientInput) (*DeleteLunaClientOutput, error) {
	req, out := c.DeleteLunaClientRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeHapg = "DescribeHapg"

// DescribeHapgRequest generates a request for the DescribeHapg operation.
func (c *CloudHSM) DescribeHapgRequest(input *DescribeHapgInput) (req *request.Request, output *DescribeHapgOutput) {
	op := &request.Operation{
		Name:       opDescribeHapg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHapgInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeHapgOutput{}
	req.Data = output
	return
}

// Retrieves information about a high-availability partition group.
func (c *CloudHSM) DescribeHapg(input *DescribeHapgInput) (*DescribeHapgOutput, error) {
	req, out := c.DescribeHapgRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeHsm = "DescribeHsm"

// DescribeHsmRequest generates a request for the DescribeHsm operation.
func (c *CloudHSM) DescribeHsmRequest(input *DescribeHsmInput) (req *request.Request, output *DescribeHsmOutput) {
	op := &request.Operation{
		Name:       opDescribeHsm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHsmInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeHsmOutput{}
	req.Data = output
	return
}

// Retrieves information about an HSM. You can identify the HSM by its ARN or
// its serial number.
func (c *CloudHSM) DescribeHsm(input *DescribeHsmInput) (*DescribeHsmOutput, error) {
	req, out := c.DescribeHsmRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeLunaClient = "DescribeLunaClient"

// DescribeLunaClientRequest generates a request for the DescribeLunaClient operation.
func (c *CloudHSM) DescribeLunaClientRequest(input *DescribeLunaClientInput) (req *request.Request, output *DescribeLunaClientOutput) {
	op := &request.Operation{
		Name:       opDescribeLunaClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLunaClientInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeLunaClientOutput{}
	req.Data = output
	return
}

// Retrieves information about an HSM client.
func (c *CloudHSM) DescribeLunaClient(input *DescribeLunaClientInput) (*DescribeLunaClientOutput, error) {
	req, out := c.DescribeLunaClientRequest(input)
	err := req.Send()
	return out, err
}

const opGetConfig = "GetConfig"

// GetConfigRequest generates a request for the GetConfig operation.
func (c *CloudHSM) GetConfigRequest(input *GetConfigInput) (req *request.Request, output *GetConfigOutput) {
	op := &request.Operation{
		Name:       opGetConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConfigInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetConfigOutput{}
	req.Data = output
	return
}

// Gets the configuration files necessary to connect to all high availability
// partition groups the client is associated with.
func (c *CloudHSM) GetConfig(input *GetConfigInput) (*GetConfigOutput, error) {
	req, out := c.GetConfigRequest(input)
	err := req.Send()
	return out, err
}

const opListAvailableZones = "ListAvailableZones"

// ListAvailableZonesRequest generates a request for the ListAvailableZones operation.
func (c *CloudHSM) ListAvailableZonesRequest(input *ListAvailableZonesInput) (req *request.Request, output *ListAvailableZonesOutput) {
	op := &request.Operation{
		Name:       opListAvailableZones,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAvailableZonesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListAvailableZonesOutput{}
	req.Data = output
	return
}

// Lists the Availability Zones that have available AWS CloudHSM capacity.
func (c *CloudHSM) ListAvailableZones(input *ListAvailableZonesInput) (*ListAvailableZonesOutput, error) {
	req, out := c.ListAvailableZonesRequest(input)
	err := req.Send()
	return out, err
}

const opListHapgs = "ListHapgs"

// ListHapgsRequest generates a request for the ListHapgs operation.
func (c *CloudHSM) ListHapgsRequest(input *ListHapgsInput) (req *request.Request, output *ListHapgsOutput) {
	op := &request.Operation{
		Name:       opListHapgs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListHapgsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListHapgsOutput{}
	req.Data = output
	return
}

// Lists the high-availability partition groups for the account.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListHapgs to retrieve the next
// set of items.
func (c *CloudHSM) ListHapgs(input *ListHapgsInput) (*ListHapgsOutput, error) {
	req, out := c.ListHapgsRequest(input)
	err := req.Send()
	return out, err
}

const opListHsms = "ListHsms"

// ListHsmsRequest generates a request for the ListHsms operation.
func (c *CloudHSM) ListHsmsRequest(input *ListHsmsInput) (req *request.Request, output *ListHsmsOutput) {
	op := &request.Operation{
		Name:       opListHsms,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListHsmsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListHsmsOutput{}
	req.Data = output
	return
}

// Retrieves the identifiers of all of the HSMs provisioned for the current
// customer.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListHsms to retrieve the next set
// of items.
func (c *CloudHSM) ListHsms(input *ListHsmsInput) (*ListHsmsOutput, error) {
	req, out := c.ListHsmsRequest(input)
	err := req.Send()
	return out, err
}

const opListLunaClients = "ListLunaClients"

// ListLunaClientsRequest generates a request for the ListLunaClients operation.
func (c *CloudHSM) ListLunaClientsRequest(input *ListLunaClientsInput) (req *request.Request, output *ListLunaClientsOutput) {
	op := &request.Operation{
		Name:       opListLunaClients,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListLunaClientsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListLunaClientsOutput{}
	req.Data = output
	return
}

// Lists all of the clients.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListLunaClients to retrieve the
// next set of items.
func (c *CloudHSM) ListLunaClients(input *ListLunaClientsInput) (*ListLunaClientsOutput, error) {
	req, out := c.ListLunaClientsRequest(input)
	err := req.Send()
	return out, err
}

const opModifyHapg = "ModifyHapg"

// ModifyHapgRequest generates a request for the ModifyHapg operation.
func (c *CloudHSM) ModifyHapgRequest(input *ModifyHapgInput) (req *request.Request, output *ModifyHapgOutput) {
	op := &request.Operation{
		Name:       opModifyHapg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyHapgInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ModifyHapgOutput{}
	req.Data = output
	return
}

// Modifies an existing high-availability partition group.
func (c *CloudHSM) ModifyHapg(input *ModifyHapgInput) (*ModifyHapgOutput, error) {
	req, out := c.ModifyHapgRequest(input)
	err := req.Send()
	return out, err
}

const opModifyHsm = "ModifyHsm"

// ModifyHsmRequest generates a request for the ModifyHsm operation.
func (c *CloudHSM) ModifyHsmRequest(input *ModifyHsmInput) (req *request.Request, output *ModifyHsmOutput) {
	op := &request.Operation{
		Name:       opModifyHsm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyHsmInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ModifyHsmOutput{}
	req.Data = output
	return
}

// Modifies an HSM.
//
//   This operation can result in the HSM being offline for up to 15 minutes
// while the AWS CloudHSM service is reconfigured. If you are modifying a production
// HSM, you should ensure that your AWS CloudHSM service is configured for high
// availability, and consider executing this operation during a maintenance
// window.
func (c *CloudHSM) ModifyHsm(input *ModifyHsmInput) (*ModifyHsmOutput, error) {
	req, out := c.ModifyHsmRequest(input)
	err := req.Send()
	return out, err
}

const opModifyLunaClient = "ModifyLunaClient"

// ModifyLunaClientRequest generates a request for the ModifyLunaClient operation.
func (c *CloudHSM) ModifyLunaClientRequest(input *ModifyLunaClientInput) (req *request.Request, output *ModifyLunaClientOutput) {
	op := &request.Operation{
		Name:       opModifyLunaClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyLunaClientInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ModifyLunaClientOutput{}
	req.Data = output
	return
}

// Modifies the certificate used by the client.
//
// This action can potentially start a workflow to install the new certificate
// on the client's HSMs.
func (c *CloudHSM) ModifyLunaClient(input *ModifyLunaClientInput) (*ModifyLunaClientOutput, error) {
	req, out := c.ModifyLunaClientRequest(input)
	err := req.Send()
	return out, err
}

// Contains the inputs for the CreateHapgRequest action.
type CreateHapgInput struct {
	// The label of the new high-availability partition group.
	Label *string `type:"string" required:"true"`

	metadataCreateHapgInput `json:"-" xml:"-"`
}

type metadataCreateHapgInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateHapgInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHapgInput) GoString() string {
	return s.String()
}

// Contains the output of the CreateHAPartitionGroup action.
type CreateHapgOutput struct {
	// The ARN of the high-availability partition group.
	HapgArn *string `type:"string"`

	metadataCreateHapgOutput `json:"-" xml:"-"`
}

type metadataCreateHapgOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateHapgOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHapgOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the CreateHsm operation.
type CreateHsmInput struct {
	// A user-defined token to ensure idempotence. Subsequent calls to this operation
	// with the same token will be ignored.
	ClientToken *string `locationName:"ClientToken" type:"string"`

	// The IP address to assign to the HSM's ENI.
	//
	// If an IP address is not specified, an IP address will be randomly chosen
	// from the CIDR range of the subnet.
	EniIp *string `locationName:"EniIp" type:"string"`

	// The external ID from IamRoleArn, if present.
	ExternalId *string `locationName:"ExternalId" type:"string"`

	// The ARN of an IAM role to enable the AWS CloudHSM service to allocate an
	// ENI on your behalf.
	IamRoleArn *string `locationName:"IamRoleArn" type:"string" required:"true"`

	// The SSH public key to install on the HSM.
	SshKey *string `locationName:"SshKey" type:"string" required:"true"`

	// The identifier of the subnet in your VPC in which to place the HSM.
	SubnetId *string `locationName:"SubnetId" type:"string" required:"true"`

	// Specifies the type of subscription for the HSM.
	//
	//   PRODUCTION - The HSM is being used in a production environment.  TRIAL
	// - The HSM is being used in a product trial.
	SubscriptionType *string `locationName:"SubscriptionType" type:"string" required:"true" enum:"SubscriptionType"`

	// The IP address for the syslog monitoring server. The AWS CloudHSM service
	// only supports one syslog monitoring server.
	SyslogIp *string `locationName:"SyslogIp" type:"string"`

	metadataCreateHsmInput `json:"-" xml:"-"`
}

type metadataCreateHsmInput struct {
	SDKShapeTraits bool `locationName:"CreateHsmRequest" type:"structure"`
}

// String returns the string representation
func (s CreateHsmInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHsmInput) GoString() string {
	return s.String()
}

// Contains the output of the CreateHsm operation.
type CreateHsmOutput struct {
	// The ARN of the HSM.
	HsmArn *string `type:"string"`

	metadataCreateHsmOutput `json:"-" xml:"-"`
}

type metadataCreateHsmOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateHsmOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateHsmOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the CreateLunaClient action.
type CreateLunaClientInput struct {
	// The contents of a Base64-Encoded X.509 v3 certificate to be installed on
	// the HSMs used by this client.
	Certificate *string `type:"string" required:"true"`

	// The label for the client.
	Label *string `type:"string"`

	metadataCreateLunaClientInput `json:"-" xml:"-"`
}

type metadataCreateLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateLunaClientInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateLunaClientInput) GoString() string {
	return s.String()
}

// Contains the output of the CreateLunaClient action.
type CreateLunaClientOutput struct {
	// The ARN of the client.
	ClientArn *string `type:"string"`

	metadataCreateLunaClientOutput `json:"-" xml:"-"`
}

type metadataCreateLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateLunaClientOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateLunaClientOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the DeleteHapg action.
type DeleteHapgInput struct {
	// The ARN of the high-availability partition group to delete.
	HapgArn *string `type:"string" required:"true"`

	metadataDeleteHapgInput `json:"-" xml:"-"`
}

type metadataDeleteHapgInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteHapgInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHapgInput) GoString() string {
	return s.String()
}

// Contains the output of the DeleteHapg action.
type DeleteHapgOutput struct {
	// The status of the action.
	Status *string `type:"string" required:"true"`

	metadataDeleteHapgOutput `json:"-" xml:"-"`
}

type metadataDeleteHapgOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteHapgOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHapgOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the DeleteHsm operation.
type DeleteHsmInput struct {
	// The ARN of the HSM to delete.
	HsmArn *string `locationName:"HsmArn" type:"string" required:"true"`

	metadataDeleteHsmInput `json:"-" xml:"-"`
}

type metadataDeleteHsmInput struct {
	SDKShapeTraits bool `locationName:"DeleteHsmRequest" type:"structure"`
}

// String returns the string representation
func (s DeleteHsmInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHsmInput) GoString() string {
	return s.String()
}

// Contains the output of the DeleteHsm operation.
type DeleteHsmOutput struct {
	// The status of the operation.
	Status *string `type:"string" required:"true"`

	metadataDeleteHsmOutput `json:"-" xml:"-"`
}

type metadataDeleteHsmOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteHsmOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHsmOutput) GoString() string {
	return s.String()
}

type DeleteLunaClientInput struct {
	// The ARN of the client to delete.
	ClientArn *string `type:"string" required:"true"`

	metadataDeleteLunaClientInput `json:"-" xml:"-"`
}

type metadataDeleteLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteLunaClientInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLunaClientInput) GoString() string {
	return s.String()
}

type DeleteLunaClientOutput struct {
	// The status of the action.
	Status *string `type:"string" required:"true"`

	metadataDeleteLunaClientOutput `json:"-" xml:"-"`
}

type metadataDeleteLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteLunaClientOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLunaClientOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the DescribeHapg action.
type DescribeHapgInput struct {
	// The ARN of the high-availability partition group to describe.
	HapgArn *string `type:"string" required:"true"`

	metadataDescribeHapgInput `json:"-" xml:"-"`
}

type metadataDescribeHapgInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeHapgInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHapgInput) GoString() string {
	return s.String()
}

// Contains the output of the DescribeHapg action.
type DescribeHapgOutput struct {
	// The ARN of the high-availability partition group.
	HapgArn *string `type:"string"`

	// The serial number of the high-availability partition group.
	HapgSerial *string `type:"string"`

	// Contains a list of ARNs that identify the HSMs.
	HsmsLastActionFailed []*string `type:"list"`

	// Contains a list of ARNs that identify the HSMs.
	HsmsPendingDeletion []*string `type:"list"`

	// Contains a list of ARNs that identify the HSMs.
	HsmsPendingRegistration []*string `type:"list"`

	// The label for the high-availability partition group.
	Label *string `type:"string"`

	// The date and time the high-availability partition group was last modified.
	LastModifiedTimestamp *string `type:"string"`

	// The list of partition serial numbers that belong to the high-availability
	// partition group.
	PartitionSerialList []*string `type:"list"`

	// The state of the high-availability partition group.
	State *string `type:"string" enum:"CloudHsmObjectState"`

	metadataDescribeHapgOutput `json:"-" xml:"-"`
}

type metadataDescribeHapgOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeHapgOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHapgOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the DescribeHsm operation.
type DescribeHsmInput struct {
	// The ARN of the HSM. Either the HsmArn or the SerialNumber parameter must
	// be specified.
	HsmArn *string `type:"string"`

	// The serial number of the HSM. Either the HsmArn or the HsmSerialNumber parameter
	// must be specified.
	HsmSerialNumber *string `type:"string"`

	metadataDescribeHsmInput `json:"-" xml:"-"`
}

type metadataDescribeHsmInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeHsmInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHsmInput) GoString() string {
	return s.String()
}

// Contains the output of the DescribeHsm operation.
type DescribeHsmOutput struct {
	// The Availability Zone that the HSM is in.
	AvailabilityZone *string `type:"string"`

	// The identifier of the elastic network interface (ENI) attached to the HSM.
	EniId *string `type:"string"`

	// The IP address assigned to the HSM's ENI.
	EniIp *string `type:"string"`

	// The ARN of the HSM.
	HsmArn *string `type:"string"`

	// The HSM model type.
	HsmType *string `type:"string"`

	// The ARN of the IAM role assigned to the HSM.
	IamRoleArn *string `type:"string"`

	// The list of partitions on the HSM.
	Partitions []*string `type:"list"`

	// The serial number of the HSM.
	SerialNumber *string `type:"string"`

	// The date and time that the server certificate was last updated.
	ServerCertLastUpdated *string `type:"string"`

	// The URI of the certificate server.
	ServerCertUri *string `type:"string"`

	// The HSM software version.
	SoftwareVersion *string `type:"string"`

	// The date and time that the SSH key was last updated.
	SshKeyLastUpdated *string `type:"string"`

	// The public SSH key.
	SshPublicKey *string `type:"string"`

	// The status of the HSM.
	Status *string `type:"string" enum:"HsmStatus"`

	// Contains additional information about the status of the HSM.
	StatusDetails *string `type:"string"`

	// The identifier of the subnet that the HSM is in.
	SubnetId *string `type:"string"`

	// The subscription end date.
	SubscriptionEndDate *string `type:"string"`

	// The subscription start date.
	SubscriptionStartDate *string `type:"string"`

	// Specifies the type of subscription for the HSM.
	//
	//   PRODUCTION - The HSM is being used in a production environment.  TRIAL
	// - The HSM is being used in a product trial.
	SubscriptionType *string `type:"string" enum:"SubscriptionType"`

	// The name of the HSM vendor.
	VendorName *string `type:"string"`

	// The identifier of the VPC that the HSM is in.
	VpcId *string `type:"string"`

	metadataDescribeHsmOutput `json:"-" xml:"-"`
}

type metadataDescribeHsmOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeHsmOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHsmOutput) GoString() string {
	return s.String()
}

type DescribeLunaClientInput struct {
	// The certificate fingerprint.
	CertificateFingerprint *string `type:"string"`

	// The ARN of the client.
	ClientArn *string `type:"string"`

	metadataDescribeLunaClientInput `json:"-" xml:"-"`
}

type metadataDescribeLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeLunaClientInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLunaClientInput) GoString() string {
	return s.String()
}

type DescribeLunaClientOutput struct {
	// The certificate installed on the HSMs used by this client.
	Certificate *string `type:"string"`

	// The certificate fingerprint.
	CertificateFingerprint *string `type:"string"`

	// The ARN of the client.
	ClientArn *string `type:"string"`

	// The label of the client.
	Label *string `type:"string"`

	// The date and time the client was last modified.
	LastModifiedTimestamp *string `type:"string"`

	metadataDescribeLunaClientOutput `json:"-" xml:"-"`
}

type metadataDescribeLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeLunaClientOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLunaClientOutput) GoString() string {
	return s.String()
}

type GetConfigInput struct {
	// The ARN of the client.
	ClientArn *string `type:"string" required:"true"`

	// The client version.
	ClientVersion *string `type:"string" required:"true" enum:"ClientVersion"`

	// A list of ARNs that identify the high-availability partition groups that
	// are associated with the client.
	HapgList []*string `type:"list" required:"true"`

	metadataGetConfigInput `json:"-" xml:"-"`
}

type metadataGetConfigInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s GetConfigInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConfigInput) GoString() string {
	return s.String()
}

type GetConfigOutput struct {
	// The certificate file containing the server.pem files of the HSMs.
	ConfigCred *string `type:"string"`

	// The chrystoki.conf configuration file.
	ConfigFile *string `type:"string"`

	// The type of credentials.
	ConfigType *string `type:"string"`

	metadataGetConfigOutput `json:"-" xml:"-"`
}

type metadataGetConfigOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s GetConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConfigOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the ListAvailableZones action.
type ListAvailableZonesInput struct {
	metadataListAvailableZonesInput `json:"-" xml:"-"`
}

type metadataListAvailableZonesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListAvailableZonesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAvailableZonesInput) GoString() string {
	return s.String()
}

type ListAvailableZonesOutput struct {
	// The list of Availability Zones that have available AWS CloudHSM capacity.
	AZList []*string `type:"list"`

	metadataListAvailableZonesOutput `json:"-" xml:"-"`
}

type metadataListAvailableZonesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListAvailableZonesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAvailableZonesOutput) GoString() string {
	return s.String()
}

type ListHapgsInput struct {
	// The NextToken value from a previous call to ListHapgs. Pass null if this
	// is the first call.
	NextToken *string `type:"string"`

	metadataListHapgsInput `json:"-" xml:"-"`
}

type metadataListHapgsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListHapgsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListHapgsInput) GoString() string {
	return s.String()
}

type ListHapgsOutput struct {
	// The list of high-availability partition groups.
	HapgList []*string `type:"list" required:"true"`

	// If not null, more results are available. Pass this value to ListHapgs to
	// retrieve the next set of items.
	NextToken *string `type:"string"`

	metadataListHapgsOutput `json:"-" xml:"-"`
}

type metadataListHapgsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListHapgsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListHapgsOutput) GoString() string {
	return s.String()
}

type ListHsmsInput struct {
	// The NextToken value from a previous call to ListHsms. Pass null if this is
	// the first call.
	NextToken *string `type:"string"`

	metadataListHsmsInput `json:"-" xml:"-"`
}

type metadataListHsmsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListHsmsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListHsmsInput) GoString() string {
	return s.String()
}

// Contains the output of the ListHsms operation.
type ListHsmsOutput struct {
	// The list of ARNs that identify the HSMs.
	HsmList []*string `type:"list"`

	// If not null, more results are available. Pass this value to ListHsms to retrieve
	// the next set of items.
	NextToken *string `type:"string"`

	metadataListHsmsOutput `json:"-" xml:"-"`
}

type metadataListHsmsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListHsmsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListHsmsOutput) GoString() string {
	return s.String()
}

type ListLunaClientsInput struct {
	// The NextToken value from a previous call to ListLunaClients. Pass null if
	// this is the first call.
	NextToken *string `type:"string"`

	metadataListLunaClientsInput `json:"-" xml:"-"`
}

type metadataListLunaClientsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListLunaClientsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListLunaClientsInput) GoString() string {
	return s.String()
}

type ListLunaClientsOutput struct {
	// The list of clients.
	ClientList []*string `type:"list" required:"true"`

	// If not null, more results are available. Pass this to ListLunaClients to
	// retrieve the next set of items.
	NextToken *string `type:"string"`

	metadataListLunaClientsOutput `json:"-" xml:"-"`
}

type metadataListLunaClientsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListLunaClientsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListLunaClientsOutput) GoString() string {
	return s.String()
}

type ModifyHapgInput struct {
	// The ARN of the high-availability partition group to modify.
	HapgArn *string `type:"string" required:"true"`

	// The new label for the high-availability partition group.
	Label *string `type:"string"`

	// The list of partition serial numbers to make members of the high-availability
	// partition group.
	PartitionSerialList []*string `type:"list"`

	metadataModifyHapgInput `json:"-" xml:"-"`
}

type metadataModifyHapgInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ModifyHapgInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHapgInput) GoString() string {
	return s.String()
}

type ModifyHapgOutput struct {
	// The ARN of the high-availability partition group.
	HapgArn *string `type:"string"`

	metadataModifyHapgOutput `json:"-" xml:"-"`
}

type metadataModifyHapgOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ModifyHapgOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHapgOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the ModifyHsm operation.
type ModifyHsmInput struct {
	// The new IP address for the elastic network interface (ENI) attached to the
	// HSM.
	//
	// If the HSM is moved to a different subnet, and an IP address is not specified,
	// an IP address will be randomly chosen from the CIDR range of the new subnet.
	EniIp *string `locationName:"EniIp" type:"string"`

	// The new external ID.
	ExternalId *string `locationName:"ExternalId" type:"string"`

	// The ARN of the HSM to modify.
	HsmArn *string `locationName:"HsmArn" type:"string" required:"true"`

	// The new IAM role ARN.
	IamRoleArn *string `locationName:"IamRoleArn" type:"string"`

	// The new identifier of the subnet that the HSM is in. The new subnet must
	// be in the same Availability Zone as the current subnet.
	SubnetId *string `locationName:"SubnetId" type:"string"`

	// The new IP address for the syslog monitoring server. The AWS CloudHSM service
	// only supports one syslog monitoring server.
	SyslogIp *string `locationName:"SyslogIp" type:"string"`

	metadataModifyHsmInput `json:"-" xml:"-"`
}

type metadataModifyHsmInput struct {
	SDKShapeTraits bool `locationName:"ModifyHsmRequest" type:"structure"`
}

// String returns the string representation
func (s ModifyHsmInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHsmInput) GoString() string {
	return s.String()
}

// Contains the output of the ModifyHsm operation.
type ModifyHsmOutput struct {
	// The ARN of the HSM.
	HsmArn *string `type:"string"`

	metadataModifyHsmOutput `json:"-" xml:"-"`
}

type metadataModifyHsmOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ModifyHsmOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHsmOutput) GoString() string {
	return s.String()
}

type ModifyLunaClientInput struct {
	// The new certificate for the client.
	Certificate *string `type:"string" required:"true"`

	// The ARN of the client.
	ClientArn *string `type:"string" required:"true"`

	metadataModifyLunaClientInput `json:"-" xml:"-"`
}

type metadataModifyLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ModifyLunaClientInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyLunaClientInput) GoString() string {
	return s.String()
}

type ModifyLunaClientOutput struct {
	// The ARN of the client.
	ClientArn *string `type:"string"`

	metadataModifyLunaClientOutput `json:"-" xml:"-"`
}

type metadataModifyLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ModifyLunaClientOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyLunaClientOutput) GoString() string {
	return s.String()
}

const (
	// @enum ClientVersion
	ClientVersion51 = "5.1"
	// @enum ClientVersion
	ClientVersion53 = "5.3"
)

const (
	// @enum CloudHsmObjectState
	CloudHsmObjectStateReady = "READY"
	// @enum CloudHsmObjectState
	CloudHsmObjectStateUpdating = "UPDATING"
	// @enum CloudHsmObjectState
	CloudHsmObjectStateDegraded = "DEGRADED"
)

const (
	// @enum HsmStatus
	HsmStatusPending = "PENDING"
	// @enum HsmStatus
	HsmStatusRunning = "RUNNING"
	// @enum HsmStatus
	HsmStatusUpdating = "UPDATING"
	// @enum HsmStatus
	HsmStatusSuspended = "SUSPENDED"
	// @enum HsmStatus
	HsmStatusTerminating = "TERMINATING"
	// @enum HsmStatus
	HsmStatusTerminated = "TERMINATED"
	// @enum HsmStatus
	HsmStatusDegraded = "DEGRADED"
)

// Specifies the type of subscription for the HSM.
//
//   PRODUCTION - The HSM is being used in a production environment.  TRIAL
// - The HSM is being used in a product trial.
const (
	// @enum SubscriptionType
	SubscriptionTypeProduction = "PRODUCTION"
)
