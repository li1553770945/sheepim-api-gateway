// Code generated by thriftgo (0.3.1). DO NOT EDIT.

package user

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type GetUserInfoReq struct {
	// 添加 api 注解为方便进行参数绑定
	UserId int64 `thrift:"userId,1,required" json:"userId,required" query:"userId,required"`
}

func NewGetUserInfoReq() *GetUserInfoReq {
	return &GetUserInfoReq{}
}

func (p *GetUserInfoReq) GetUserId() (v int64) {
	return p.UserId
}

var fieldIDToName_GetUserInfoReq = map[int16]string{
	1: "userId",
}

func (p *GetUserInfoReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUserId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetUserId = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetUserId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetUserInfoReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetUserInfoReq[fieldId]))
}

func (p *GetUserInfoReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserId = v
	}
	return nil
}

func (p *GetUserInfoReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfoReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetUserInfoReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("userId", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetUserInfoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserInfoReq(%+v)", *p)
}

type GetUserInfoRespData struct {
	Username string `thrift:"username,1" form:"username" json:"username" query:"username"`
	Nickname string `thrift:"nickname,2" form:"nickname" json:"nickname" query:"nickname"`
	Role     string `thrift:"role,3" form:"role" json:"role" query:"role"`
}

func NewGetUserInfoRespData() *GetUserInfoRespData {
	return &GetUserInfoRespData{}
}

func (p *GetUserInfoRespData) GetUsername() (v string) {
	return p.Username
}

func (p *GetUserInfoRespData) GetNickname() (v string) {
	return p.Nickname
}

func (p *GetUserInfoRespData) GetRole() (v string) {
	return p.Role
}

var fieldIDToName_GetUserInfoRespData = map[int16]string{
	1: "username",
	2: "nickname",
	3: "role",
}

func (p *GetUserInfoRespData) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetUserInfoRespData[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GetUserInfoRespData) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Username = v
	}
	return nil
}

func (p *GetUserInfoRespData) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Nickname = v
	}
	return nil
}

func (p *GetUserInfoRespData) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Role = v
	}
	return nil
}

func (p *GetUserInfoRespData) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfoRespData"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetUserInfoRespData) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Username); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetUserInfoRespData) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("nickname", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Nickname); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *GetUserInfoRespData) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("role", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Role); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *GetUserInfoRespData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserInfoRespData(%+v)", *p)
}

type GetUserInfoResp struct {
	Code    int32                `thrift:"code,1,required" form:"code,required" json:"code,required" query:"code,required"`
	Message string               `thrift:"message,2,required" form:"message,required" json:"message,required" query:"message,required"`
	Data    *GetUserInfoRespData `thrift:"data,3,optional" form:"data" json:"data,omitempty" query:"data"`
}

func NewGetUserInfoResp() *GetUserInfoResp {
	return &GetUserInfoResp{}
}

func (p *GetUserInfoResp) GetCode() (v int32) {
	return p.Code
}

func (p *GetUserInfoResp) GetMessage() (v string) {
	return p.Message
}

var GetUserInfoResp_Data_DEFAULT *GetUserInfoRespData

func (p *GetUserInfoResp) GetData() (v *GetUserInfoRespData) {
	if !p.IsSetData() {
		return GetUserInfoResp_Data_DEFAULT
	}
	return p.Data
}

var fieldIDToName_GetUserInfoResp = map[int16]string{
	1: "code",
	2: "message",
	3: "data",
}

func (p *GetUserInfoResp) IsSetData() bool {
	return p.Data != nil
}

func (p *GetUserInfoResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetCode bool = false
	var issetMessage bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetCode = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetMessage = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetMessage {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetUserInfoResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetUserInfoResp[fieldId]))
}

func (p *GetUserInfoResp) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}

func (p *GetUserInfoResp) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Message = v
	}
	return nil
}

func (p *GetUserInfoResp) ReadField3(iprot thrift.TProtocol) error {
	p.Data = NewGetUserInfoRespData()
	if err := p.Data.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GetUserInfoResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfoResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetUserInfoResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetUserInfoResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Message); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *GetUserInfoResp) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetData() {
		if err = oprot.WriteFieldBegin("data", thrift.STRUCT, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Data.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *GetUserInfoResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserInfoResp(%+v)", *p)
}

type UserController interface {
	GetUserInfo(ctx context.Context, request *GetUserInfoReq) (r *GetUserInfoResp, err error)

	GetSelfInfo(ctx context.Context, request *GetUserInfoReq) (r *GetUserInfoResp, err error)
}

type UserControllerClient struct {
	c thrift.TClient
}

func NewUserControllerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UserControllerClient {
	return &UserControllerClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewUserControllerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UserControllerClient {
	return &UserControllerClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewUserControllerClient(c thrift.TClient) *UserControllerClient {
	return &UserControllerClient{
		c: c,
	}
}

func (p *UserControllerClient) Client_() thrift.TClient {
	return p.c
}

func (p *UserControllerClient) GetUserInfo(ctx context.Context, request *GetUserInfoReq) (r *GetUserInfoResp, err error) {
	var _args UserControllerGetUserInfoArgs
	_args.Request = request
	var _result UserControllerGetUserInfoResult
	if err = p.Client_().Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *UserControllerClient) GetSelfInfo(ctx context.Context, request *GetUserInfoReq) (r *GetUserInfoResp, err error) {
	var _args UserControllerGetSelfInfoArgs
	_args.Request = request
	var _result UserControllerGetSelfInfoResult
	if err = p.Client_().Call(ctx, "GetSelfInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type UserControllerProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      UserController
}

func (p *UserControllerProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *UserControllerProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *UserControllerProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewUserControllerProcessor(handler UserController) *UserControllerProcessor {
	self := &UserControllerProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GetUserInfo", &userControllerProcessorGetUserInfo{handler: handler})
	self.AddToProcessorMap("GetSelfInfo", &userControllerProcessorGetSelfInfo{handler: handler})
	return self
}
func (p *UserControllerProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type userControllerProcessorGetUserInfo struct {
	handler UserController
}

func (p *userControllerProcessorGetUserInfo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := UserControllerGetUserInfoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetUserInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := UserControllerGetUserInfoResult{}
	var retval *GetUserInfoResp
	if retval, err2 = p.handler.GetUserInfo(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetUserInfo: "+err2.Error())
		oprot.WriteMessageBegin("GetUserInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetUserInfo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type userControllerProcessorGetSelfInfo struct {
	handler UserController
}

func (p *userControllerProcessorGetSelfInfo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := UserControllerGetSelfInfoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetSelfInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := UserControllerGetSelfInfoResult{}
	var retval *GetUserInfoResp
	if retval, err2 = p.handler.GetSelfInfo(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetSelfInfo: "+err2.Error())
		oprot.WriteMessageBegin("GetSelfInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetSelfInfo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type UserControllerGetUserInfoArgs struct {
	Request *GetUserInfoReq `thrift:"request,1"`
}

func NewUserControllerGetUserInfoArgs() *UserControllerGetUserInfoArgs {
	return &UserControllerGetUserInfoArgs{}
}

var UserControllerGetUserInfoArgs_Request_DEFAULT *GetUserInfoReq

func (p *UserControllerGetUserInfoArgs) GetRequest() (v *GetUserInfoReq) {
	if !p.IsSetRequest() {
		return UserControllerGetUserInfoArgs_Request_DEFAULT
	}
	return p.Request
}

var fieldIDToName_UserControllerGetUserInfoArgs = map[int16]string{
	1: "request",
}

func (p *UserControllerGetUserInfoArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *UserControllerGetUserInfoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserControllerGetUserInfoArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserControllerGetUserInfoArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewGetUserInfoReq()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UserControllerGetUserInfoArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserControllerGetUserInfoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserControllerGetUserInfoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserControllerGetUserInfoArgs(%+v)", *p)
}

type UserControllerGetUserInfoResult struct {
	Success *GetUserInfoResp `thrift:"success,0,optional"`
}

func NewUserControllerGetUserInfoResult() *UserControllerGetUserInfoResult {
	return &UserControllerGetUserInfoResult{}
}

var UserControllerGetUserInfoResult_Success_DEFAULT *GetUserInfoResp

func (p *UserControllerGetUserInfoResult) GetSuccess() (v *GetUserInfoResp) {
	if !p.IsSetSuccess() {
		return UserControllerGetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_UserControllerGetUserInfoResult = map[int16]string{
	0: "success",
}

func (p *UserControllerGetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserControllerGetUserInfoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserControllerGetUserInfoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserControllerGetUserInfoResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewGetUserInfoResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UserControllerGetUserInfoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserControllerGetUserInfoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *UserControllerGetUserInfoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserControllerGetUserInfoResult(%+v)", *p)
}

type UserControllerGetSelfInfoArgs struct {
	Request *GetUserInfoReq `thrift:"request,1"`
}

func NewUserControllerGetSelfInfoArgs() *UserControllerGetSelfInfoArgs {
	return &UserControllerGetSelfInfoArgs{}
}

var UserControllerGetSelfInfoArgs_Request_DEFAULT *GetUserInfoReq

func (p *UserControllerGetSelfInfoArgs) GetRequest() (v *GetUserInfoReq) {
	if !p.IsSetRequest() {
		return UserControllerGetSelfInfoArgs_Request_DEFAULT
	}
	return p.Request
}

var fieldIDToName_UserControllerGetSelfInfoArgs = map[int16]string{
	1: "request",
}

func (p *UserControllerGetSelfInfoArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *UserControllerGetSelfInfoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserControllerGetSelfInfoArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserControllerGetSelfInfoArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewGetUserInfoReq()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UserControllerGetSelfInfoArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetSelfInfo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserControllerGetSelfInfoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserControllerGetSelfInfoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserControllerGetSelfInfoArgs(%+v)", *p)
}

type UserControllerGetSelfInfoResult struct {
	Success *GetUserInfoResp `thrift:"success,0,optional"`
}

func NewUserControllerGetSelfInfoResult() *UserControllerGetSelfInfoResult {
	return &UserControllerGetSelfInfoResult{}
}

var UserControllerGetSelfInfoResult_Success_DEFAULT *GetUserInfoResp

func (p *UserControllerGetSelfInfoResult) GetSuccess() (v *GetUserInfoResp) {
	if !p.IsSetSuccess() {
		return UserControllerGetSelfInfoResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_UserControllerGetSelfInfoResult = map[int16]string{
	0: "success",
}

func (p *UserControllerGetSelfInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserControllerGetSelfInfoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserControllerGetSelfInfoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserControllerGetSelfInfoResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewGetUserInfoResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UserControllerGetSelfInfoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetSelfInfo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserControllerGetSelfInfoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *UserControllerGetSelfInfoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserControllerGetSelfInfoResult(%+v)", *p)
}
