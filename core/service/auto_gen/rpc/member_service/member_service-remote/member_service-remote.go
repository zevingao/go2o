// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "github.com/apache/thrift/lib/go/thrift"
	"go2o/core/service/auto_gen/rpc/ttype"
	"go2o/core/service/auto_gen/rpc/message_service"
        "go2o/core/service/auto_gen/rpc/member_service"
)

var _ = ttype.GoUnusedProtection__
var _ = message_service.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  Result RegisterMemberV2(string user, string pwd, i32 flag, string name, string phone, string email, string avatar,  extend)")
  fmt.Fprintln(os.Stderr, "  Result CheckLogin(string user, string pwd, bool update)")
  fmt.Fprintln(os.Stderr, "  Result CheckTradePwd(i64 memberId, string tradePwd)")
  fmt.Fprintln(os.Stderr, "  i64 SwapMemberId(ECredentials cred, string value)")
  fmt.Fprintln(os.Stderr, "   LevelList()")
  fmt.Fprintln(os.Stderr, "  STrustedInfo GetTrustInfo(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SubmitTrustInfo(i64 memberId, STrustedInfo info)")
  fmt.Fprintln(os.Stderr, "  Result ReviewTrustedInfo(i64 memberId, bool reviewPass, string remark)")
  fmt.Fprintln(os.Stderr, "  SLevel GetLevel(i32 id)")
  fmt.Fprintln(os.Stderr, "  SLevel GetLevelBySign(string sign)")
  fmt.Fprintln(os.Stderr, "  SMember GetMember(i64 id)")
  fmt.Fprintln(os.Stderr, "  SMember GetMemberByUser(string user)")
  fmt.Fprintln(os.Stderr, "  SProfile GetProfile(i64 id)")
  fmt.Fprintln(os.Stderr, "  Result Active(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result Lock(i64 memberId, i32 minutes, string remark)")
  fmt.Fprintln(os.Stderr, "  Result Unlock(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result GrantFlag(i64 memberId, i32 flag)")
  fmt.Fprintln(os.Stderr, "  SComplexMember Complex(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SendCode(i64 memberId, string operation, EMessageChannel msgType)")
  fmt.Fprintln(os.Stderr, "  Result CompareCode(i64 memberId, string code)")
  fmt.Fprintln(os.Stderr, "   ReceiptsCodes(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SaveReceiptsCode(i64 memberId, SReceiptsCode code)")
  fmt.Fprintln(os.Stderr, "   Bankcards(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SaveBankcard(i64 memberId, SBankcard card)")
  fmt.Fprintln(os.Stderr, "  Result CheckProfileComplete(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  SMemberLevelInfo MemberLevelInfo(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result UpdateLevel(i64 memberId, i32 level, bool review, i64 paymentOrderId)")
  fmt.Fprintln(os.Stderr, "  Result ChangePhone(i64 memberId, string phone)")
  fmt.Fprintln(os.Stderr, "  Result ChangeUser(i64 memberId, string usr)")
  fmt.Fprintln(os.Stderr, "  Result ModifyPwd(i64 memberId, string old, string pwd)")
  fmt.Fprintln(os.Stderr, "  Result ModifyTradePwd(i64 memberId, string old, string pwd)")
  fmt.Fprintln(os.Stderr, "  Result Premium(i64 memberId, i32 v, i64 expires)")
  fmt.Fprintln(os.Stderr, "  string GetToken(i64 memberId, bool reset)")
  fmt.Fprintln(os.Stderr, "  bool CheckToken(i64 memberId, string token)")
  fmt.Fprintln(os.Stderr, "  void RemoveToken(i64 memberId)")
  fmt.Fprintln(os.Stderr, "   GetAddressList(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  SAddress GetAddress(i64 memberId, i64 addrId)")
  fmt.Fprintln(os.Stderr, "  SAccount GetAccount(i64 memberId)")
  fmt.Fprintln(os.Stderr, "   InviterArray(i64 memberId, i32 depth)")
  fmt.Fprintln(os.Stderr, "  i32 InviteMembersQuantity(i64 memberId, i32 depth)")
  fmt.Fprintln(os.Stderr, "  i32 QueryInviteQuantity(i64 memberId,  data)")
  fmt.Fprintln(os.Stderr, "   QueryInviteArray(i64 memberId,  data)")
  fmt.Fprintln(os.Stderr, "  Result AccountCharge(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountConsume(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountDiscount(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountRefund(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountAdjust(i64 memberId, i32 account, i32 value, i64 relateUser, string remark)")
  fmt.Fprintln(os.Stderr, "  Result B4EAuth(i64 memberId, string action,  data)")
  fmt.Fprintln(os.Stderr, "  SPagingResult PagingAccountLog(i64 memberId, i32 accountType, SPagingParams params)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := member_service.NewMemberServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "RegisterMemberV2":
    if flag.NArg() - 1 != 8 {
      fmt.Fprintln(os.Stderr, "RegisterMemberV2 requires 8 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err116 := (strconv.Atoi(flag.Arg(3)))
    if err116 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    argvalue6 := flag.Arg(7)
    value6 := argvalue6
    arg121 := flag.Arg(8)
    mbTrans122 := thrift.NewTMemoryBufferLen(len(arg121))
    defer mbTrans122.Close()
    _, err123 := mbTrans122.WriteString(arg121)
    if err123 != nil { 
      Usage()
      return
    }
    factory124 := thrift.NewTJSONProtocolFactory()
    jsProt125 := factory124.GetProtocol(mbTrans122)
    containerStruct7 := member_service.NewMemberServiceRegisterMemberV2Args()
    err126 := containerStruct7.ReadField8(jsProt125)
    if err126 != nil {
      Usage()
      return
    }
    argvalue7 := containerStruct7.Extend
    value7 := argvalue7
    fmt.Print(client.RegisterMemberV2(context.Background(), value0, value1, value2, value3, value4, value5, value6, value7))
    fmt.Print("\n")
    break
  case "CheckLogin":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CheckLogin requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    fmt.Print(client.CheckLogin(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "CheckTradePwd":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CheckTradePwd requires 2 args")
      flag.Usage()
    }
    argvalue0, err130 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err130 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CheckTradePwd(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "SwapMemberId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SwapMemberId requires 2 args")
      flag.Usage()
    }
    tmp0, err := (strconv.Atoi(flag.Arg(1)))
    if err != nil {
      Usage()
     return
    }
    argvalue0 := member_service.ECredentials(tmp0)
    value0 := member_service.ECredentials(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SwapMemberId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "LevelList":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "LevelList requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.LevelList(context.Background()))
    fmt.Print("\n")
    break
  case "GetTrustInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTrustInfo requires 1 args")
      flag.Usage()
    }
    argvalue0, err133 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err133 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTrustInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SubmitTrustInfo":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SubmitTrustInfo requires 2 args")
      flag.Usage()
    }
    argvalue0, err134 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err134 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg135 := flag.Arg(2)
    mbTrans136 := thrift.NewTMemoryBufferLen(len(arg135))
    defer mbTrans136.Close()
    _, err137 := mbTrans136.WriteString(arg135)
    if err137 != nil {
      Usage()
      return
    }
    factory138 := thrift.NewTJSONProtocolFactory()
    jsProt139 := factory138.GetProtocol(mbTrans136)
    argvalue1 := member_service.NewSTrustedInfo()
    err140 := argvalue1.Read(jsProt139)
    if err140 != nil {
      Usage()
      return
    }
    value1 := member_service.STrustedInfo(argvalue1)
    fmt.Print(client.SubmitTrustInfo(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ReviewTrustedInfo":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ReviewTrustedInfo requires 3 args")
      flag.Usage()
    }
    argvalue0, err141 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err141 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.ReviewTrustedInfo(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "GetLevel":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLevel requires 1 args")
      flag.Usage()
    }
    tmp0, err144 := (strconv.Atoi(flag.Arg(1)))
    if err144 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.GetLevel(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetLevelBySign":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLevelBySign requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetLevelBySign(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetMember":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetMember requires 1 args")
      flag.Usage()
    }
    argvalue0, err146 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err146 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetMember(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetMemberByUser":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetMemberByUser requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetMemberByUser(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetProfile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetProfile requires 1 args")
      flag.Usage()
    }
    argvalue0, err148 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err148 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetProfile(context.Background(), value0))
    fmt.Print("\n")
    break
  case "Active":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Active requires 1 args")
      flag.Usage()
    }
    argvalue0, err149 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err149 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Active(context.Background(), value0))
    fmt.Print("\n")
    break
  case "Lock":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Lock requires 3 args")
      flag.Usage()
    }
    argvalue0, err150 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err150 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err151 := (strconv.Atoi(flag.Arg(2)))
    if err151 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.Lock(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "Unlock":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Unlock requires 1 args")
      flag.Usage()
    }
    argvalue0, err153 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err153 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Unlock(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GrantFlag":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GrantFlag requires 2 args")
      flag.Usage()
    }
    argvalue0, err154 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err154 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err155 := (strconv.Atoi(flag.Arg(2)))
    if err155 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.GrantFlag(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "Complex":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Complex requires 1 args")
      flag.Usage()
    }
    argvalue0, err156 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err156 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Complex(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SendCode":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SendCode requires 3 args")
      flag.Usage()
    }
    argvalue0, err157 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err157 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err := (strconv.Atoi(flag.Arg(3)))
    if err != nil {
      Usage()
     return
    }
    argvalue2 := member_service.EMessageChannel(tmp2)
    value2 := argvalue2
    fmt.Print(client.SendCode(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "CompareCode":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CompareCode requires 2 args")
      flag.Usage()
    }
    argvalue0, err159 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err159 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CompareCode(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ReceiptsCodes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ReceiptsCodes requires 1 args")
      flag.Usage()
    }
    argvalue0, err161 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err161 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ReceiptsCodes(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SaveReceiptsCode":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SaveReceiptsCode requires 2 args")
      flag.Usage()
    }
    argvalue0, err162 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err162 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg163 := flag.Arg(2)
    mbTrans164 := thrift.NewTMemoryBufferLen(len(arg163))
    defer mbTrans164.Close()
    _, err165 := mbTrans164.WriteString(arg163)
    if err165 != nil {
      Usage()
      return
    }
    factory166 := thrift.NewTJSONProtocolFactory()
    jsProt167 := factory166.GetProtocol(mbTrans164)
    argvalue1 := member_service.NewSReceiptsCode()
    err168 := argvalue1.Read(jsProt167)
    if err168 != nil {
      Usage()
      return
    }
    value1 := member_service.SReceiptsCode(argvalue1)
    fmt.Print(client.SaveReceiptsCode(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "Bankcards":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Bankcards requires 1 args")
      flag.Usage()
    }
    argvalue0, err169 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err169 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Bankcards(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SaveBankcard":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SaveBankcard requires 2 args")
      flag.Usage()
    }
    argvalue0, err170 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err170 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg171 := flag.Arg(2)
    mbTrans172 := thrift.NewTMemoryBufferLen(len(arg171))
    defer mbTrans172.Close()
    _, err173 := mbTrans172.WriteString(arg171)
    if err173 != nil {
      Usage()
      return
    }
    factory174 := thrift.NewTJSONProtocolFactory()
    jsProt175 := factory174.GetProtocol(mbTrans172)
    argvalue1 := member_service.NewSBankcard()
    err176 := argvalue1.Read(jsProt175)
    if err176 != nil {
      Usage()
      return
    }
    value1 := member_service.SBankcard(argvalue1)
    fmt.Print(client.SaveBankcard(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "CheckProfileComplete":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CheckProfileComplete requires 1 args")
      flag.Usage()
    }
    argvalue0, err177 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err177 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CheckProfileComplete(context.Background(), value0))
    fmt.Print("\n")
    break
  case "MemberLevelInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "MemberLevelInfo requires 1 args")
      flag.Usage()
    }
    argvalue0, err178 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err178 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.MemberLevelInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "UpdateLevel":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "UpdateLevel requires 4 args")
      flag.Usage()
    }
    argvalue0, err179 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err179 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err180 := (strconv.Atoi(flag.Arg(2)))
    if err180 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3, err182 := (strconv.ParseInt(flag.Arg(4), 10, 64))
    if err182 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    fmt.Print(client.UpdateLevel(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "ChangePhone":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ChangePhone requires 2 args")
      flag.Usage()
    }
    argvalue0, err183 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err183 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ChangePhone(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ChangeUser":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ChangeUser requires 2 args")
      flag.Usage()
    }
    argvalue0, err185 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err185 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ChangeUser(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ModifyPwd":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ModifyPwd requires 3 args")
      flag.Usage()
    }
    argvalue0, err187 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err187 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.ModifyPwd(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "ModifyTradePwd":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ModifyTradePwd requires 3 args")
      flag.Usage()
    }
    argvalue0, err190 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err190 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.ModifyTradePwd(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "Premium":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Premium requires 3 args")
      flag.Usage()
    }
    argvalue0, err193 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err193 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err194 := (strconv.Atoi(flag.Arg(2)))
    if err194 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2, err195 := (strconv.ParseInt(flag.Arg(3), 10, 64))
    if err195 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.Premium(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "GetToken":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetToken requires 2 args")
      flag.Usage()
    }
    argvalue0, err196 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err196 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    fmt.Print(client.GetToken(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "CheckToken":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CheckToken requires 2 args")
      flag.Usage()
    }
    argvalue0, err198 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err198 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CheckToken(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "RemoveToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveToken requires 1 args")
      flag.Usage()
    }
    argvalue0, err200 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err200 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RemoveToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetAddressList":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetAddressList requires 1 args")
      flag.Usage()
    }
    argvalue0, err201 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err201 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetAddressList(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetAddress":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetAddress requires 2 args")
      flag.Usage()
    }
    argvalue0, err202 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err202 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1, err203 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err203 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.GetAddress(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GetAccount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetAccount requires 1 args")
      flag.Usage()
    }
    argvalue0, err204 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err204 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetAccount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "InviterArray":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InviterArray requires 2 args")
      flag.Usage()
    }
    argvalue0, err205 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err205 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err206 := (strconv.Atoi(flag.Arg(2)))
    if err206 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.InviterArray(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InviteMembersQuantity":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InviteMembersQuantity requires 2 args")
      flag.Usage()
    }
    argvalue0, err207 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err207 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err208 := (strconv.Atoi(flag.Arg(2)))
    if err208 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.InviteMembersQuantity(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "QueryInviteQuantity":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "QueryInviteQuantity requires 2 args")
      flag.Usage()
    }
    argvalue0, err209 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err209 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg210 := flag.Arg(2)
    mbTrans211 := thrift.NewTMemoryBufferLen(len(arg210))
    defer mbTrans211.Close()
    _, err212 := mbTrans211.WriteString(arg210)
    if err212 != nil { 
      Usage()
      return
    }
    factory213 := thrift.NewTJSONProtocolFactory()
    jsProt214 := factory213.GetProtocol(mbTrans211)
    containerStruct1 := member_service.NewMemberServiceQueryInviteQuantityArgs()
    err215 := containerStruct1.ReadField2(jsProt214)
    if err215 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.Data
    value1 := argvalue1
    fmt.Print(client.QueryInviteQuantity(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "QueryInviteArray":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "QueryInviteArray requires 2 args")
      flag.Usage()
    }
    argvalue0, err216 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err216 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg217 := flag.Arg(2)
    mbTrans218 := thrift.NewTMemoryBufferLen(len(arg217))
    defer mbTrans218.Close()
    _, err219 := mbTrans218.WriteString(arg217)
    if err219 != nil { 
      Usage()
      return
    }
    factory220 := thrift.NewTJSONProtocolFactory()
    jsProt221 := factory220.GetProtocol(mbTrans218)
    containerStruct1 := member_service.NewMemberServiceQueryInviteArrayArgs()
    err222 := containerStruct1.ReadField2(jsProt221)
    if err222 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.Data
    value1 := argvalue1
    fmt.Print(client.QueryInviteArray(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "AccountCharge":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountCharge requires 6 args")
      flag.Usage()
    }
    argvalue0, err223 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err223 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err224 := (strconv.Atoi(flag.Arg(2)))
    if err224 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err226 := (strconv.Atoi(flag.Arg(4)))
    if err226 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountCharge(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountConsume":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountConsume requires 6 args")
      flag.Usage()
    }
    argvalue0, err229 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err229 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err230 := (strconv.Atoi(flag.Arg(2)))
    if err230 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err232 := (strconv.Atoi(flag.Arg(4)))
    if err232 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountConsume(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountDiscount":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountDiscount requires 6 args")
      flag.Usage()
    }
    argvalue0, err235 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err235 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err236 := (strconv.Atoi(flag.Arg(2)))
    if err236 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err238 := (strconv.Atoi(flag.Arg(4)))
    if err238 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountDiscount(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountRefund":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountRefund requires 6 args")
      flag.Usage()
    }
    argvalue0, err241 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err241 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err242 := (strconv.Atoi(flag.Arg(2)))
    if err242 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err244 := (strconv.Atoi(flag.Arg(4)))
    if err244 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountRefund(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountAdjust":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "AccountAdjust requires 5 args")
      flag.Usage()
    }
    argvalue0, err247 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err247 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err248 := (strconv.Atoi(flag.Arg(2)))
    if err248 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err249 := (strconv.Atoi(flag.Arg(3)))
    if err249 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    argvalue3, err250 := (strconv.ParseInt(flag.Arg(4), 10, 64))
    if err250 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.AccountAdjust(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "B4EAuth":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "B4EAuth requires 3 args")
      flag.Usage()
    }
    argvalue0, err252 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err252 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg254 := flag.Arg(3)
    mbTrans255 := thrift.NewTMemoryBufferLen(len(arg254))
    defer mbTrans255.Close()
    _, err256 := mbTrans255.WriteString(arg254)
    if err256 != nil { 
      Usage()
      return
    }
    factory257 := thrift.NewTJSONProtocolFactory()
    jsProt258 := factory257.GetProtocol(mbTrans255)
    containerStruct2 := member_service.NewMemberServiceB4EAuthArgs()
    err259 := containerStruct2.ReadField3(jsProt258)
    if err259 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Data
    value2 := argvalue2
    fmt.Print(client.B4EAuth(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "PagingAccountLog":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "PagingAccountLog requires 3 args")
      flag.Usage()
    }
    argvalue0, err260 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err260 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err261 := (strconv.Atoi(flag.Arg(2)))
    if err261 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    arg262 := flag.Arg(3)
    mbTrans263 := thrift.NewTMemoryBufferLen(len(arg262))
    defer mbTrans263.Close()
    _, err264 := mbTrans263.WriteString(arg262)
    if err264 != nil {
      Usage()
      return
    }
    factory265 := thrift.NewTJSONProtocolFactory()
    jsProt266 := factory265.GetProtocol(mbTrans263)
    argvalue2 := ttype.NewSPagingParams()
    err267 := argvalue2.Read(jsProt266)
    if err267 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.PagingAccountLog(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
