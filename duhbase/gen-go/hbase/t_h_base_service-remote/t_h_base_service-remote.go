// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"git.du.com/cloud/du_component/duhbase/gen-go/hbase"
	"github.com/apache/thrift/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var _ = hbase.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  bool exists(string table, TGet tget)")
	fmt.Fprintln(os.Stderr, "   existsAll(string table,  tgets)")
	fmt.Fprintln(os.Stderr, "  TResult get(string table, TGet tget)")
	fmt.Fprintln(os.Stderr, "   getMultiple(string table,  tgets)")
	fmt.Fprintln(os.Stderr, "  void put(string table, TPut tput)")
	fmt.Fprintln(os.Stderr, "  bool checkAndPut(string table, string row, string family, string qualifier, string value, TPut tput)")
	fmt.Fprintln(os.Stderr, "  void putMultiple(string table,  tputs)")
	fmt.Fprintln(os.Stderr, "  void deleteSingle(string table, TDelete tdelete)")
	fmt.Fprintln(os.Stderr, "   deleteMultiple(string table,  tdeletes)")
	fmt.Fprintln(os.Stderr, "  bool checkAndDelete(string table, string row, string family, string qualifier, string value, TDelete tdelete)")
	fmt.Fprintln(os.Stderr, "  TResult increment(string table, TIncrement tincrement)")
	fmt.Fprintln(os.Stderr, "  TResult append(string table, TAppend tappend)")
	fmt.Fprintln(os.Stderr, "  i32 openScanner(string table, TScan tscan)")
	fmt.Fprintln(os.Stderr, "   getScannerRows(i32 scannerId, i32 numRows)")
	fmt.Fprintln(os.Stderr, "  void closeScanner(i32 scannerId)")
	fmt.Fprintln(os.Stderr, "  void mutateRow(string table, TRowMutations trowMutations)")
	fmt.Fprintln(os.Stderr, "   getScannerResults(string table, TScan tscan, i32 numRows)")
	fmt.Fprintln(os.Stderr, "  THRegionLocation getRegionLocation(string table, string row, bool reload)")
	fmt.Fprintln(os.Stderr, "   getAllRegionLocations(string table)")
	fmt.Fprintln(os.Stderr, "  bool checkAndMutate(string table, string row, string family, string qualifier, TCompareOp compareOp, string value, TRowMutations rowMutations)")
	fmt.Fprintln(os.Stderr, "  TTableDescriptor getTableDescriptor(TTableName table)")
	fmt.Fprintln(os.Stderr, "   getTableDescriptors( tables)")
	fmt.Fprintln(os.Stderr, "  bool tableExists(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "   getTableDescriptorsByPattern(string regex, bool includeSysTables)")
	fmt.Fprintln(os.Stderr, "   getTableDescriptorsByNamespace(string name)")
	fmt.Fprintln(os.Stderr, "   getTableNamesByPattern(string regex, bool includeSysTables)")
	fmt.Fprintln(os.Stderr, "   getTableNamesByNamespace(string name)")
	fmt.Fprintln(os.Stderr, "  void createTable(TTableDescriptor desc,  splitKeys)")
	fmt.Fprintln(os.Stderr, "  void deleteTable(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "  void truncateTable(TTableName tableName, bool preserveSplits)")
	fmt.Fprintln(os.Stderr, "  void enableTable(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "  void disableTable(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "  bool isTableEnabled(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "  bool isTableDisabled(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "  bool isTableAvailable(TTableName tableName)")
	fmt.Fprintln(os.Stderr, "  bool isTableAvailableWithSplit(TTableName tableName,  splitKeys)")
	fmt.Fprintln(os.Stderr, "  void addColumnFamily(TTableName tableName, TColumnFamilyDescriptor column)")
	fmt.Fprintln(os.Stderr, "  void deleteColumnFamily(TTableName tableName, string column)")
	fmt.Fprintln(os.Stderr, "  void modifyColumnFamily(TTableName tableName, TColumnFamilyDescriptor column)")
	fmt.Fprintln(os.Stderr, "  void modifyTable(TTableDescriptor desc)")
	fmt.Fprintln(os.Stderr, "  void createNamespace(TNamespaceDescriptor namespaceDesc)")
	fmt.Fprintln(os.Stderr, "  void modifyNamespace(TNamespaceDescriptor namespaceDesc)")
	fmt.Fprintln(os.Stderr, "  void deleteNamespace(string name)")
	fmt.Fprintln(os.Stderr, "  TNamespaceDescriptor getNamespaceDescriptor(string name)")
	fmt.Fprintln(os.Stderr, "   listNamespaceDescriptors()")
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
	client := hbase.NewTHBaseServiceClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "exists":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Exists requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
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
		argvalue1 := hbase.NewTGet()
		err215 := argvalue1.Read(context.Background(), jsProt214)
		if err215 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Exists(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "existsAll":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ExistsAll requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
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
		containerStruct1 := hbase.NewTHBaseServiceExistsAllArgs()
		err222 := containerStruct1.ReadField2(context.Background(), jsProt221)
		if err222 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tgets
		value1 := argvalue1
		fmt.Print(client.ExistsAll(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "get":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Get requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg224 := flag.Arg(2)
		mbTrans225 := thrift.NewTMemoryBufferLen(len(arg224))
		defer mbTrans225.Close()
		_, err226 := mbTrans225.WriteString(arg224)
		if err226 != nil {
			Usage()
			return
		}
		factory227 := thrift.NewTJSONProtocolFactory()
		jsProt228 := factory227.GetProtocol(mbTrans225)
		argvalue1 := hbase.NewTGet()
		err229 := argvalue1.Read(context.Background(), jsProt228)
		if err229 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Get(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "getMultiple":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetMultiple requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg231 := flag.Arg(2)
		mbTrans232 := thrift.NewTMemoryBufferLen(len(arg231))
		defer mbTrans232.Close()
		_, err233 := mbTrans232.WriteString(arg231)
		if err233 != nil {
			Usage()
			return
		}
		factory234 := thrift.NewTJSONProtocolFactory()
		jsProt235 := factory234.GetProtocol(mbTrans232)
		containerStruct1 := hbase.NewTHBaseServiceGetMultipleArgs()
		err236 := containerStruct1.ReadField2(context.Background(), jsProt235)
		if err236 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tgets
		value1 := argvalue1
		fmt.Print(client.GetMultiple(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "put":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Put requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg238 := flag.Arg(2)
		mbTrans239 := thrift.NewTMemoryBufferLen(len(arg238))
		defer mbTrans239.Close()
		_, err240 := mbTrans239.WriteString(arg238)
		if err240 != nil {
			Usage()
			return
		}
		factory241 := thrift.NewTJSONProtocolFactory()
		jsProt242 := factory241.GetProtocol(mbTrans239)
		argvalue1 := hbase.NewTPut()
		err243 := argvalue1.Read(context.Background(), jsProt242)
		if err243 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Put(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "checkAndPut":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "CheckAndPut requires 6 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := []byte(flag.Arg(3))
		value2 := argvalue2
		argvalue3 := []byte(flag.Arg(4))
		value3 := argvalue3
		argvalue4 := []byte(flag.Arg(5))
		value4 := argvalue4
		arg249 := flag.Arg(6)
		mbTrans250 := thrift.NewTMemoryBufferLen(len(arg249))
		defer mbTrans250.Close()
		_, err251 := mbTrans250.WriteString(arg249)
		if err251 != nil {
			Usage()
			return
		}
		factory252 := thrift.NewTJSONProtocolFactory()
		jsProt253 := factory252.GetProtocol(mbTrans250)
		argvalue5 := hbase.NewTPut()
		err254 := argvalue5.Read(context.Background(), jsProt253)
		if err254 != nil {
			Usage()
			return
		}
		value5 := argvalue5
		fmt.Print(client.CheckAndPut(context.Background(), value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "putMultiple":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PutMultiple requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg256 := flag.Arg(2)
		mbTrans257 := thrift.NewTMemoryBufferLen(len(arg256))
		defer mbTrans257.Close()
		_, err258 := mbTrans257.WriteString(arg256)
		if err258 != nil {
			Usage()
			return
		}
		factory259 := thrift.NewTJSONProtocolFactory()
		jsProt260 := factory259.GetProtocol(mbTrans257)
		containerStruct1 := hbase.NewTHBaseServicePutMultipleArgs()
		err261 := containerStruct1.ReadField2(context.Background(), jsProt260)
		if err261 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tputs
		value1 := argvalue1
		fmt.Print(client.PutMultiple(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "deleteSingle":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeleteSingle requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg263 := flag.Arg(2)
		mbTrans264 := thrift.NewTMemoryBufferLen(len(arg263))
		defer mbTrans264.Close()
		_, err265 := mbTrans264.WriteString(arg263)
		if err265 != nil {
			Usage()
			return
		}
		factory266 := thrift.NewTJSONProtocolFactory()
		jsProt267 := factory266.GetProtocol(mbTrans264)
		argvalue1 := hbase.NewTDelete()
		err268 := argvalue1.Read(context.Background(), jsProt267)
		if err268 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.DeleteSingle(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "deleteMultiple":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeleteMultiple requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg270 := flag.Arg(2)
		mbTrans271 := thrift.NewTMemoryBufferLen(len(arg270))
		defer mbTrans271.Close()
		_, err272 := mbTrans271.WriteString(arg270)
		if err272 != nil {
			Usage()
			return
		}
		factory273 := thrift.NewTJSONProtocolFactory()
		jsProt274 := factory273.GetProtocol(mbTrans271)
		containerStruct1 := hbase.NewTHBaseServiceDeleteMultipleArgs()
		err275 := containerStruct1.ReadField2(context.Background(), jsProt274)
		if err275 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tdeletes
		value1 := argvalue1
		fmt.Print(client.DeleteMultiple(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "checkAndDelete":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "CheckAndDelete requires 6 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := []byte(flag.Arg(3))
		value2 := argvalue2
		argvalue3 := []byte(flag.Arg(4))
		value3 := argvalue3
		argvalue4 := []byte(flag.Arg(5))
		value4 := argvalue4
		arg281 := flag.Arg(6)
		mbTrans282 := thrift.NewTMemoryBufferLen(len(arg281))
		defer mbTrans282.Close()
		_, err283 := mbTrans282.WriteString(arg281)
		if err283 != nil {
			Usage()
			return
		}
		factory284 := thrift.NewTJSONProtocolFactory()
		jsProt285 := factory284.GetProtocol(mbTrans282)
		argvalue5 := hbase.NewTDelete()
		err286 := argvalue5.Read(context.Background(), jsProt285)
		if err286 != nil {
			Usage()
			return
		}
		value5 := argvalue5
		fmt.Print(client.CheckAndDelete(context.Background(), value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "increment":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Increment requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg288 := flag.Arg(2)
		mbTrans289 := thrift.NewTMemoryBufferLen(len(arg288))
		defer mbTrans289.Close()
		_, err290 := mbTrans289.WriteString(arg288)
		if err290 != nil {
			Usage()
			return
		}
		factory291 := thrift.NewTJSONProtocolFactory()
		jsProt292 := factory291.GetProtocol(mbTrans289)
		argvalue1 := hbase.NewTIncrement()
		err293 := argvalue1.Read(context.Background(), jsProt292)
		if err293 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Increment(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "append":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Append requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg295 := flag.Arg(2)
		mbTrans296 := thrift.NewTMemoryBufferLen(len(arg295))
		defer mbTrans296.Close()
		_, err297 := mbTrans296.WriteString(arg295)
		if err297 != nil {
			Usage()
			return
		}
		factory298 := thrift.NewTJSONProtocolFactory()
		jsProt299 := factory298.GetProtocol(mbTrans296)
		argvalue1 := hbase.NewTAppend()
		err300 := argvalue1.Read(context.Background(), jsProt299)
		if err300 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Append(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "openScanner":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "OpenScanner requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg302 := flag.Arg(2)
		mbTrans303 := thrift.NewTMemoryBufferLen(len(arg302))
		defer mbTrans303.Close()
		_, err304 := mbTrans303.WriteString(arg302)
		if err304 != nil {
			Usage()
			return
		}
		factory305 := thrift.NewTJSONProtocolFactory()
		jsProt306 := factory305.GetProtocol(mbTrans303)
		argvalue1 := hbase.NewTScan()
		err307 := argvalue1.Read(context.Background(), jsProt306)
		if err307 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.OpenScanner(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "getScannerRows":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetScannerRows requires 2 args")
			flag.Usage()
		}
		tmp0, err308 := (strconv.Atoi(flag.Arg(1)))
		if err308 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err309 := (strconv.Atoi(flag.Arg(2)))
		if err309 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.GetScannerRows(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "closeScanner":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CloseScanner requires 1 args")
			flag.Usage()
		}
		tmp0, err310 := (strconv.Atoi(flag.Arg(1)))
		if err310 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.CloseScanner(context.Background(), value0))
		fmt.Print("\n")
		break
	case "mutateRow":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "MutateRow requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg312 := flag.Arg(2)
		mbTrans313 := thrift.NewTMemoryBufferLen(len(arg312))
		defer mbTrans313.Close()
		_, err314 := mbTrans313.WriteString(arg312)
		if err314 != nil {
			Usage()
			return
		}
		factory315 := thrift.NewTJSONProtocolFactory()
		jsProt316 := factory315.GetProtocol(mbTrans313)
		argvalue1 := hbase.NewTRowMutations()
		err317 := argvalue1.Read(context.Background(), jsProt316)
		if err317 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.MutateRow(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "getScannerResults":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetScannerResults requires 3 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg319 := flag.Arg(2)
		mbTrans320 := thrift.NewTMemoryBufferLen(len(arg319))
		defer mbTrans320.Close()
		_, err321 := mbTrans320.WriteString(arg319)
		if err321 != nil {
			Usage()
			return
		}
		factory322 := thrift.NewTJSONProtocolFactory()
		jsProt323 := factory322.GetProtocol(mbTrans320)
		argvalue1 := hbase.NewTScan()
		err324 := argvalue1.Read(context.Background(), jsProt323)
		if err324 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		tmp2, err325 := (strconv.Atoi(flag.Arg(3)))
		if err325 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.GetScannerResults(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "getRegionLocation":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetRegionLocation requires 3 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.GetRegionLocation(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "getAllRegionLocations":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAllRegionLocations requires 1 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		fmt.Print(client.GetAllRegionLocations(context.Background(), value0))
		fmt.Print("\n")
		break
	case "checkAndMutate":
		if flag.NArg()-1 != 7 {
			fmt.Fprintln(os.Stderr, "CheckAndMutate requires 7 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := []byte(flag.Arg(3))
		value2 := argvalue2
		argvalue3 := []byte(flag.Arg(4))
		value3 := argvalue3
		tmp4, err := (strconv.Atoi(flag.Arg(5)))
		if err != nil {
			Usage()
			return
		}
		argvalue4 := hbase.TCompareOp(tmp4)
		value4 := argvalue4
		argvalue5 := []byte(flag.Arg(6))
		value5 := argvalue5
		arg335 := flag.Arg(7)
		mbTrans336 := thrift.NewTMemoryBufferLen(len(arg335))
		defer mbTrans336.Close()
		_, err337 := mbTrans336.WriteString(arg335)
		if err337 != nil {
			Usage()
			return
		}
		factory338 := thrift.NewTJSONProtocolFactory()
		jsProt339 := factory338.GetProtocol(mbTrans336)
		argvalue6 := hbase.NewTRowMutations()
		err340 := argvalue6.Read(context.Background(), jsProt339)
		if err340 != nil {
			Usage()
			return
		}
		value6 := argvalue6
		fmt.Print(client.CheckAndMutate(context.Background(), value0, value1, value2, value3, value4, value5, value6))
		fmt.Print("\n")
		break
	case "getTableDescriptor":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTableDescriptor requires 1 args")
			flag.Usage()
		}
		arg341 := flag.Arg(1)
		mbTrans342 := thrift.NewTMemoryBufferLen(len(arg341))
		defer mbTrans342.Close()
		_, err343 := mbTrans342.WriteString(arg341)
		if err343 != nil {
			Usage()
			return
		}
		factory344 := thrift.NewTJSONProtocolFactory()
		jsProt345 := factory344.GetProtocol(mbTrans342)
		argvalue0 := hbase.NewTTableName()
		err346 := argvalue0.Read(context.Background(), jsProt345)
		if err346 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTableDescriptor(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getTableDescriptors":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTableDescriptors requires 1 args")
			flag.Usage()
		}
		arg347 := flag.Arg(1)
		mbTrans348 := thrift.NewTMemoryBufferLen(len(arg347))
		defer mbTrans348.Close()
		_, err349 := mbTrans348.WriteString(arg347)
		if err349 != nil {
			Usage()
			return
		}
		factory350 := thrift.NewTJSONProtocolFactory()
		jsProt351 := factory350.GetProtocol(mbTrans348)
		containerStruct0 := hbase.NewTHBaseServiceGetTableDescriptorsArgs()
		err352 := containerStruct0.ReadField1(context.Background(), jsProt351)
		if err352 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Tables
		value0 := argvalue0
		fmt.Print(client.GetTableDescriptors(context.Background(), value0))
		fmt.Print("\n")
		break
	case "tableExists":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TableExists requires 1 args")
			flag.Usage()
		}
		arg353 := flag.Arg(1)
		mbTrans354 := thrift.NewTMemoryBufferLen(len(arg353))
		defer mbTrans354.Close()
		_, err355 := mbTrans354.WriteString(arg353)
		if err355 != nil {
			Usage()
			return
		}
		factory356 := thrift.NewTJSONProtocolFactory()
		jsProt357 := factory356.GetProtocol(mbTrans354)
		argvalue0 := hbase.NewTTableName()
		err358 := argvalue0.Read(context.Background(), jsProt357)
		if err358 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TableExists(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getTableDescriptorsByPattern":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetTableDescriptorsByPattern requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetTableDescriptorsByPattern(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "getTableDescriptorsByNamespace":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTableDescriptorsByNamespace requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetTableDescriptorsByNamespace(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getTableNamesByPattern":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetTableNamesByPattern requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetTableNamesByPattern(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "getTableNamesByNamespace":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTableNamesByNamespace requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetTableNamesByNamespace(context.Background(), value0))
		fmt.Print("\n")
		break
	case "createTable":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CreateTable requires 2 args")
			flag.Usage()
		}
		arg365 := flag.Arg(1)
		mbTrans366 := thrift.NewTMemoryBufferLen(len(arg365))
		defer mbTrans366.Close()
		_, err367 := mbTrans366.WriteString(arg365)
		if err367 != nil {
			Usage()
			return
		}
		factory368 := thrift.NewTJSONProtocolFactory()
		jsProt369 := factory368.GetProtocol(mbTrans366)
		argvalue0 := hbase.NewTTableDescriptor()
		err370 := argvalue0.Read(context.Background(), jsProt369)
		if err370 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg371 := flag.Arg(2)
		mbTrans372 := thrift.NewTMemoryBufferLen(len(arg371))
		defer mbTrans372.Close()
		_, err373 := mbTrans372.WriteString(arg371)
		if err373 != nil {
			Usage()
			return
		}
		factory374 := thrift.NewTJSONProtocolFactory()
		jsProt375 := factory374.GetProtocol(mbTrans372)
		containerStruct1 := hbase.NewTHBaseServiceCreateTableArgs()
		err376 := containerStruct1.ReadField2(context.Background(), jsProt375)
		if err376 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.SplitKeys
		value1 := argvalue1
		fmt.Print(client.CreateTable(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "deleteTable":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteTable requires 1 args")
			flag.Usage()
		}
		arg377 := flag.Arg(1)
		mbTrans378 := thrift.NewTMemoryBufferLen(len(arg377))
		defer mbTrans378.Close()
		_, err379 := mbTrans378.WriteString(arg377)
		if err379 != nil {
			Usage()
			return
		}
		factory380 := thrift.NewTJSONProtocolFactory()
		jsProt381 := factory380.GetProtocol(mbTrans378)
		argvalue0 := hbase.NewTTableName()
		err382 := argvalue0.Read(context.Background(), jsProt381)
		if err382 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteTable(context.Background(), value0))
		fmt.Print("\n")
		break
	case "truncateTable":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TruncateTable requires 2 args")
			flag.Usage()
		}
		arg383 := flag.Arg(1)
		mbTrans384 := thrift.NewTMemoryBufferLen(len(arg383))
		defer mbTrans384.Close()
		_, err385 := mbTrans384.WriteString(arg383)
		if err385 != nil {
			Usage()
			return
		}
		factory386 := thrift.NewTJSONProtocolFactory()
		jsProt387 := factory386.GetProtocol(mbTrans384)
		argvalue0 := hbase.NewTTableName()
		err388 := argvalue0.Read(context.Background(), jsProt387)
		if err388 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.TruncateTable(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "enableTable":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EnableTable requires 1 args")
			flag.Usage()
		}
		arg390 := flag.Arg(1)
		mbTrans391 := thrift.NewTMemoryBufferLen(len(arg390))
		defer mbTrans391.Close()
		_, err392 := mbTrans391.WriteString(arg390)
		if err392 != nil {
			Usage()
			return
		}
		factory393 := thrift.NewTJSONProtocolFactory()
		jsProt394 := factory393.GetProtocol(mbTrans391)
		argvalue0 := hbase.NewTTableName()
		err395 := argvalue0.Read(context.Background(), jsProt394)
		if err395 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.EnableTable(context.Background(), value0))
		fmt.Print("\n")
		break
	case "disableTable":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DisableTable requires 1 args")
			flag.Usage()
		}
		arg396 := flag.Arg(1)
		mbTrans397 := thrift.NewTMemoryBufferLen(len(arg396))
		defer mbTrans397.Close()
		_, err398 := mbTrans397.WriteString(arg396)
		if err398 != nil {
			Usage()
			return
		}
		factory399 := thrift.NewTJSONProtocolFactory()
		jsProt400 := factory399.GetProtocol(mbTrans397)
		argvalue0 := hbase.NewTTableName()
		err401 := argvalue0.Read(context.Background(), jsProt400)
		if err401 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DisableTable(context.Background(), value0))
		fmt.Print("\n")
		break
	case "isTableEnabled":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "IsTableEnabled requires 1 args")
			flag.Usage()
		}
		arg402 := flag.Arg(1)
		mbTrans403 := thrift.NewTMemoryBufferLen(len(arg402))
		defer mbTrans403.Close()
		_, err404 := mbTrans403.WriteString(arg402)
		if err404 != nil {
			Usage()
			return
		}
		factory405 := thrift.NewTJSONProtocolFactory()
		jsProt406 := factory405.GetProtocol(mbTrans403)
		argvalue0 := hbase.NewTTableName()
		err407 := argvalue0.Read(context.Background(), jsProt406)
		if err407 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.IsTableEnabled(context.Background(), value0))
		fmt.Print("\n")
		break
	case "isTableDisabled":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "IsTableDisabled requires 1 args")
			flag.Usage()
		}
		arg408 := flag.Arg(1)
		mbTrans409 := thrift.NewTMemoryBufferLen(len(arg408))
		defer mbTrans409.Close()
		_, err410 := mbTrans409.WriteString(arg408)
		if err410 != nil {
			Usage()
			return
		}
		factory411 := thrift.NewTJSONProtocolFactory()
		jsProt412 := factory411.GetProtocol(mbTrans409)
		argvalue0 := hbase.NewTTableName()
		err413 := argvalue0.Read(context.Background(), jsProt412)
		if err413 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.IsTableDisabled(context.Background(), value0))
		fmt.Print("\n")
		break
	case "isTableAvailable":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "IsTableAvailable requires 1 args")
			flag.Usage()
		}
		arg414 := flag.Arg(1)
		mbTrans415 := thrift.NewTMemoryBufferLen(len(arg414))
		defer mbTrans415.Close()
		_, err416 := mbTrans415.WriteString(arg414)
		if err416 != nil {
			Usage()
			return
		}
		factory417 := thrift.NewTJSONProtocolFactory()
		jsProt418 := factory417.GetProtocol(mbTrans415)
		argvalue0 := hbase.NewTTableName()
		err419 := argvalue0.Read(context.Background(), jsProt418)
		if err419 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.IsTableAvailable(context.Background(), value0))
		fmt.Print("\n")
		break
	case "isTableAvailableWithSplit":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "IsTableAvailableWithSplit requires 2 args")
			flag.Usage()
		}
		arg420 := flag.Arg(1)
		mbTrans421 := thrift.NewTMemoryBufferLen(len(arg420))
		defer mbTrans421.Close()
		_, err422 := mbTrans421.WriteString(arg420)
		if err422 != nil {
			Usage()
			return
		}
		factory423 := thrift.NewTJSONProtocolFactory()
		jsProt424 := factory423.GetProtocol(mbTrans421)
		argvalue0 := hbase.NewTTableName()
		err425 := argvalue0.Read(context.Background(), jsProt424)
		if err425 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg426 := flag.Arg(2)
		mbTrans427 := thrift.NewTMemoryBufferLen(len(arg426))
		defer mbTrans427.Close()
		_, err428 := mbTrans427.WriteString(arg426)
		if err428 != nil {
			Usage()
			return
		}
		factory429 := thrift.NewTJSONProtocolFactory()
		jsProt430 := factory429.GetProtocol(mbTrans427)
		containerStruct1 := hbase.NewTHBaseServiceIsTableAvailableWithSplitArgs()
		err431 := containerStruct1.ReadField2(context.Background(), jsProt430)
		if err431 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.SplitKeys
		value1 := argvalue1
		fmt.Print(client.IsTableAvailableWithSplit(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "addColumnFamily":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AddColumnFamily requires 2 args")
			flag.Usage()
		}
		arg432 := flag.Arg(1)
		mbTrans433 := thrift.NewTMemoryBufferLen(len(arg432))
		defer mbTrans433.Close()
		_, err434 := mbTrans433.WriteString(arg432)
		if err434 != nil {
			Usage()
			return
		}
		factory435 := thrift.NewTJSONProtocolFactory()
		jsProt436 := factory435.GetProtocol(mbTrans433)
		argvalue0 := hbase.NewTTableName()
		err437 := argvalue0.Read(context.Background(), jsProt436)
		if err437 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg438 := flag.Arg(2)
		mbTrans439 := thrift.NewTMemoryBufferLen(len(arg438))
		defer mbTrans439.Close()
		_, err440 := mbTrans439.WriteString(arg438)
		if err440 != nil {
			Usage()
			return
		}
		factory441 := thrift.NewTJSONProtocolFactory()
		jsProt442 := factory441.GetProtocol(mbTrans439)
		argvalue1 := hbase.NewTColumnFamilyDescriptor()
		err443 := argvalue1.Read(context.Background(), jsProt442)
		if err443 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.AddColumnFamily(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "deleteColumnFamily":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeleteColumnFamily requires 2 args")
			flag.Usage()
		}
		arg444 := flag.Arg(1)
		mbTrans445 := thrift.NewTMemoryBufferLen(len(arg444))
		defer mbTrans445.Close()
		_, err446 := mbTrans445.WriteString(arg444)
		if err446 != nil {
			Usage()
			return
		}
		factory447 := thrift.NewTJSONProtocolFactory()
		jsProt448 := factory447.GetProtocol(mbTrans445)
		argvalue0 := hbase.NewTTableName()
		err449 := argvalue0.Read(context.Background(), jsProt448)
		if err449 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		fmt.Print(client.DeleteColumnFamily(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "modifyColumnFamily":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ModifyColumnFamily requires 2 args")
			flag.Usage()
		}
		arg451 := flag.Arg(1)
		mbTrans452 := thrift.NewTMemoryBufferLen(len(arg451))
		defer mbTrans452.Close()
		_, err453 := mbTrans452.WriteString(arg451)
		if err453 != nil {
			Usage()
			return
		}
		factory454 := thrift.NewTJSONProtocolFactory()
		jsProt455 := factory454.GetProtocol(mbTrans452)
		argvalue0 := hbase.NewTTableName()
		err456 := argvalue0.Read(context.Background(), jsProt455)
		if err456 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg457 := flag.Arg(2)
		mbTrans458 := thrift.NewTMemoryBufferLen(len(arg457))
		defer mbTrans458.Close()
		_, err459 := mbTrans458.WriteString(arg457)
		if err459 != nil {
			Usage()
			return
		}
		factory460 := thrift.NewTJSONProtocolFactory()
		jsProt461 := factory460.GetProtocol(mbTrans458)
		argvalue1 := hbase.NewTColumnFamilyDescriptor()
		err462 := argvalue1.Read(context.Background(), jsProt461)
		if err462 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.ModifyColumnFamily(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "modifyTable":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ModifyTable requires 1 args")
			flag.Usage()
		}
		arg463 := flag.Arg(1)
		mbTrans464 := thrift.NewTMemoryBufferLen(len(arg463))
		defer mbTrans464.Close()
		_, err465 := mbTrans464.WriteString(arg463)
		if err465 != nil {
			Usage()
			return
		}
		factory466 := thrift.NewTJSONProtocolFactory()
		jsProt467 := factory466.GetProtocol(mbTrans464)
		argvalue0 := hbase.NewTTableDescriptor()
		err468 := argvalue0.Read(context.Background(), jsProt467)
		if err468 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ModifyTable(context.Background(), value0))
		fmt.Print("\n")
		break
	case "createNamespace":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateNamespace requires 1 args")
			flag.Usage()
		}
		arg469 := flag.Arg(1)
		mbTrans470 := thrift.NewTMemoryBufferLen(len(arg469))
		defer mbTrans470.Close()
		_, err471 := mbTrans470.WriteString(arg469)
		if err471 != nil {
			Usage()
			return
		}
		factory472 := thrift.NewTJSONProtocolFactory()
		jsProt473 := factory472.GetProtocol(mbTrans470)
		argvalue0 := hbase.NewTNamespaceDescriptor()
		err474 := argvalue0.Read(context.Background(), jsProt473)
		if err474 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CreateNamespace(context.Background(), value0))
		fmt.Print("\n")
		break
	case "modifyNamespace":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ModifyNamespace requires 1 args")
			flag.Usage()
		}
		arg475 := flag.Arg(1)
		mbTrans476 := thrift.NewTMemoryBufferLen(len(arg475))
		defer mbTrans476.Close()
		_, err477 := mbTrans476.WriteString(arg475)
		if err477 != nil {
			Usage()
			return
		}
		factory478 := thrift.NewTJSONProtocolFactory()
		jsProt479 := factory478.GetProtocol(mbTrans476)
		argvalue0 := hbase.NewTNamespaceDescriptor()
		err480 := argvalue0.Read(context.Background(), jsProt479)
		if err480 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ModifyNamespace(context.Background(), value0))
		fmt.Print("\n")
		break
	case "deleteNamespace":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteNamespace requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.DeleteNamespace(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getNamespaceDescriptor":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetNamespaceDescriptor requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetNamespaceDescriptor(context.Background(), value0))
		fmt.Print("\n")
		break
	case "listNamespaceDescriptors":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "ListNamespaceDescriptors requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.ListNamespaceDescriptors(context.Background()))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}