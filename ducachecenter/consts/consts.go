package consts

const (
	Separator = byte(1)
	//安卓查询业务
	AndQueryCustomerInfoByCid    = "AndQueryCustomerInfoByCid"
	AndQueryCustomerInfoByCidUrl = "/v1/and/cid_customer_info"

	AndQueryProductInfoByPkg    = "AndQueryProductInfoByPkg"
	AndQueryProductInfoByPkgUrl = "/v1/and/pkg_product_info"

	//IOS查询业务
	IosQueryCustomerInfoByCid    = "IosQueryCustomerInfoByCid"
	IosQueryCustomerInfoByCidUrl = "/v1/ios/cid_customer_info"

	IosQueryProductInfoByPkg    = "IosQueryProductInfoByPkg"
	IosQueryProductInfoByPkgUrl = "/v1/ios/pkg_product_info"

	IosQueryCidByAndCid    = "IosQueryCidByAndCid"
	AndQueryCidByIosCid    = "AndQueryCidByIosCid"
	IosQueryCidByAndCidUrl = "/v1/ios/and_cid_to_ios_cid"

	IosQueryCustomerInfoByPkg    = "IosQueryCustomerInfoByPkg"
	IosQueryCustomerInfoByPkgUrl = "/v1/ios/pkg_customer_info"

	//Applet查询业务
	AppletQueryProductInfoByAppid    = "AppletQueryProductInfoByAppid"
	AppletQueryProductInfoByAppidUrl = "/v1/applet/appid_product_info"

	//H5查询业务
	H5QueryProductInfoByAppId    = "H5QueryProductInfoByAppId"
	H5QueryProductInfoByAppIdUrl = "/v1/h5/appid_product_info"

	//前端Front查询业务
	FrontQueryCustomerInfoByCid    = "FrontQueryCustomerInfoByCid"
	FrontQueryCustomerInfoByCidUrl = "/v1/front/get_activate_info"
)
