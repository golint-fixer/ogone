package ogone

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

func shaInAliasCompose(params map[string]string, secret string) string {
	normalizedParams := normalizeParams(params, shaInAliasAllowedParams)
	return shaCompose(normalizedParams, secret)
}

func shaInCompose(params map[string]string, secret string) string {
	normalizedParams := normalizeParams(params, shaInAllowedParams)
	return shaCompose(normalizedParams, secret)
}

func shaOutCompose(params map[string]string, secret string) string {
	normalizedParams := normalizeParams(params, shaOutAllowedParams)
	return shaCompose(normalizedParams, secret)
}

func shaCompose(params map[string]string, secret string) string {
	keys := make([]string, len(params))

	keyCounter := 0

	for k := range params {
		keys[keyCounter] = k
		keyCounter++
	}

	sort.Strings(keys)

	sign := ""

	for _, key := range keys {
		sign += key + "=" + params[key] + secret
	}

	sum := sha1.Sum([]byte(sign))

	return strings.ToUpper(hex.EncodeToString(sum[:]))
}

// Make keys to uppercase and sort keys alphabetically, also intersect with allowed keys
func normalizeParams(params map[string]string, allowed []string) map[string]string {
	values := make(map[string]string)

	for k, v := range params {
		key := strings.ToUpper(k)

		isAllowed := allowed == nil

		if isAllowed == false {
			for _, allowedKey := range allowed {
				if allowedKey == key {
					isAllowed = true
					break
				}
			}
		}

		if isAllowed == true {
			values[key] = v
		}
	}

	return values
}

var shaInAliasAllowedParams = []string{
	"ACCEPTANCE",
	"ACCEPTURL",
	"ADDMATCH",
	"ADDRMATCH",
	"AIACTIONNUMBER",
	"AIAGIATA",
	"AIAIRNAME",
	"AIAIRTAX",
	"AIBOOKIND*XX*",
	"AICARRIER*XX*",
	"AICHDET",
	"AICLASS*XX*",
	"AICONJTI",
	"AIDEPTCODE",
	"AIDESTCITY*XX*",
	"AIDESTCITYL*XX*",
	"AIEXTRAPASNAME*XX*",
	"AIEYCD",
	"AIFLDATE*XX*",
	"AIFLNUM*XX*",
	"AIGLNUM",
	"AIINVOICE",
	"AIIRST",
	"AIORCITY*XX*",
	"AIORCITYL*XX*",
	"AIPASNAME",
	"AIPROJNUM",
	"AISTOPOV*XX*",
	"AITIDATE",
	"AITINUM",
	"AITINUML*XX*",
	"AITYPCH",
	"AIVATAMNT",
	"AIVATAPPL",
	"ALIAS",
	"ALIASOPERATION",
	"ALIASPERSISTEDAFTERUSE",
	"ALIASUSAGE",
	"ALLOWCORRECTION",
	"AMOUNT",
	"AMOUNT*XX*",
	"AMOUNTHTVA",
	"AMOUNTTVA",
	"ARP_TRN",
	"BACKURL",
	"BATCHID",
	"BGCOLOR",
	"BLVERNUM",
	"BIC",
	"BIN",
	"BRAND",
	"BRANDVISUAL",
	"BUTTONBGCOLOR",
	"BUTTONTXTCOLOR",
	"CANCELURL",
	"CATALOGURL",
	"CAVV_3D",
	"CAVVALGORITHM_3D",
	"CERTID",
	"CHECK_AAV",
	"CIVILITY",
	"COM",
	"COMPLUS",
	"CONVCCY",
	"COSTCENTER",
	"COSTCODE",
	"CREDITCODE",
	"CREDITDEBIT",
	"CUID",
	"CURRENCY",
	"DATA",
	"DATATYPE",
	"DATEIN",
	"DATEOUT",
	"DBXML",
	"DCC_COMMPERC",
	"DCC_CONVAMOUNT",
	"DCC_CONVCCY",
	"DCC_EXCHRATE",
	"DCC_EXCHRATETS",
	"DCC_INDICATOR",
	"DCC_MARGINPERC",
	"DCC_REF",
	"DCC_SOURCE",
	"DCC_VALID",
	"DECLINEURL",
	"DELIVERYDATE",
	"DEVICE",
	"DISCOUNTRATE",
	"DISPLAYMODE",
	"ECI",
	"ECI_3D",
	"ECOM_BILLTO_COMPANY",
	"ECOM_BILLTO_POSTAL_CITY",
	"ECOM_BILLTO_POSTAL_COUNTRYCODE",
	"ECOM_BILLTO_POSTAL_COUNTY",
	"ECOM_BILLTO_POSTAL_NAME_FIRST",
	"ECOM_BILLTO_POSTAL_NAME_LAST",
	"ECOM_BILLTO_POSTAL_NAME_PREFIX",
	"ECOM_BILLTO_POSTAL_POSTALCODE",
	"ECOM_BILLTO_POSTAL_STREET_LINE1",
	"ECOM_BILLTO_POSTAL_STREET_LINE2",
	"ECOM_BILLTO_POSTAL_STREET_LINE3",
	"ECOM_BILLTO_POSTAL_STREET_NUMBER",
	"ECOM_BILLTO_TELECOM_MOBILE_NUMBER",
	"ECOM_BILLTO_TELECOM_PHONE_NUMBER",
	"ECOM_CONSUMERID",
	"ECOM_CONSUMER_GENDER",
	"ECOM_CONSUMEROGID",
	"ECOM_CONSUMERORDERID",
	"ECOM_CONSUMERUSERALIAS",
	"ECOM_CONSUMERUSERPWD",
	"ECOM_CONSUMERUSERID",
	"ECOM_ESTIMATEDDELIVERYDATE",
	"ECOM_ESTIMATEDELIVERYDATE",
	"ECOM_SHIPMETHOD",
	"ECOM_SHIPMETHODDETAILS",
	"ECOM_SHIPMETHODSPEED",
	"ECOM_SHIPMETHODTYPE",
	"ECOM_SHIPTO_COMPANY",
	"ECOM_SHIPTO_DOB",
	"ECOM_SHIPTO_ONLINE_EMAIL",
	"ECOM_SHIPTO_POSTAL_CITY",
	"ECOM_SHIPTO_POSTAL_COUNTRYCODE",
	"ECOM_SHIPTO_POSTAL_COUNTY",
	"ECOM_SHIPTO_POSTAL_NAME_FIRST",
	"ECOM_SHIPTO_POSTAL_NAME_LAST",
	"ECOM_SHIPTO_POSTAL_NAME_PREFIX",
	"ECOM_SHIPTO_POSTAL_POSTALCODE",
	"ECOM_SHIPTO_POSTAL_STATE",
	"ECOM_SHIPTO_POSTAL_STREET_LINE1",
	"ECOM_SHIPTO_POSTAL_STREET_LINE2",
	"ECOM_SHIPTO_POSTAL_STREET_NUMBER",
	"ECOM_SHIPTO_TELECOM_FAX_NUMBER",
	"ECOM_SHIPTO_TELECOM_MOBILE_NUMBER",
	"ECOM_SHIPTO_TELECOM_PHONE_NUMBER",
	"ECOM_SHIPTO_TVA",
	"ED",
	"EMAIL",
	"EXCEPTIONURL",
	"EXCLPMLIST",
	"EXECUTIONDATE*XX*",
	"FACEXCL*XX*",
	"FACTOTAL*XX*",
	"FIRSTCALL",
	"FLAG3D",
	"FONTTYPE",
	"FORCECODE1",
	"FORCECODE2",
	"FORCECODEHASH",
	"FORCEPROCESS",
	"FORCETP",
	"FP_ACTIV",
	"GENERIC_BL",
	"GIROPAY_ACCOUNT_NUMBER",
	"GIROPAY_BLZ",
	"GIROPAY_OWNER_NAME",
	"GLOBORDERID",
	"GUID",
	"HDFONTTYPE",
	"HDTBLBGCOLOR",
	"HDTBLTXTCOLOR",
	"HEIGHTFRAME",
	"HOMEURL",
	"HTTP_ACCEPT",
	"HTTP_USER_AGENT",
	"INCLUDE_BIN",
	"INCLUDE_COUNTRIES",
	"INITIAL_REC_TRN",
	"INVDATE",
	"INVDISCOUNT",
	"INVLEVEL",
	"INVORDERID",
	"ISSUERID",
	"IST_MOBILE",
	"ITEM_COUNT",
	"ITEMATTRIBUTES*XX*",
	"ITEMCATEGORY*XX*",
	"ITEMCOMMENTS*XX*",
	"ITEMDESC*XX*",
	"ITEMDISCOUNT*XX*",
	"ITEMFDMPRODUCTCATEG*XX*",
	"ITEMID*XX*",
	"ITEMNAME*XX*",
	"ITEMPRICE*XX*",
	"ITEMQUANT*XX*",
	"ITEMQUANTORIG*XX*",
	"ITEMUNITOFMEASURE*XX*",
	"ITEMVAT*XX*",
	"ITEMVATCODE*XX*",
	"ITEMWEIGHT*XX*",
	"LANGUAGE",
	"LEVEL1AUTHCPC",
	"LIDEXCL*XX*",
	"LIMITCLIENTSCRIPTUSAGE",
	"LINE_REF",
	"LINE_REF1",
	"LINE_REF2",
	"LINE_REF3",
	"LINE_REF4",
	"LINE_REF5",
	"LINE_REF6",
	"LIST_BIN",
	"LIST_COUNTRIES",
	"LOGO",
	"MANDATEID",
	"MAXITEMQUANT*XX*",
	"MERCHANTID",
	"MODE",
	"MTIME",
	"MVER",
	"NETAMOUNT",
	"OPERATION",
	"ORDERID",
	"ORDERSHIPCOST",
	"ORDERSHIPMETH",
	"ORDERSHIPTAX",
	"ORDERSHIPTAXCODE",
	"ORIG",
	"OR_INVORDERID",
	"OR_ORDERID",
	"OWNERADDRESS",
	"OWNERADDRESS2",
	"OWNERCTY",
	"OWNERTELNO",
	"OWNERTELNO2",
	"OWNERTOWN",
	"OWNERZIP",
	"PAIDAMOUNT",
	"PARAMPLUS",
	"PARAMVAR",
	"PAYID",
	"PAYMETHOD",
	"PM",
	"PMLIST",
	"PMLISTPMLISTTYPE",
	"PMLISTTYPE",
	"PMLISTTYPEPMLIST",
	"PMTYPE",
	"POPUP",
	"POST",
	"PSPID",
	"PSWD",
	"RECIPIENTACCOUNTNUMBER",
	"RECIPIENTDOB",
	"RECIPIENTLASTNAME",
	"RECIPIENTZIP",
	"REF",
	"REFER",
	"REFID",
	"REFKIND",
	"REF_CUSTOMERID",
	"REF_CUSTOMERREF",
	"REGISTRED",
	"REMOTE_ADDR",
	"REQGENFIELDS",
	"RNPOFFERT",
	"RTIMEOUT",
	"RTIMEOUTREQUESTEDTIMEOUT",
	"SCORINGCLIENT",
	"SEQUENCETYPE",
	"SETT_BATCH",
	"SID",
	"SIGNDATE",
	"STATUS_3D",
	"SUBSCRIPTION_ID",
	"SUB_AM",
	"SUB_AMOUNT",
	"SUB_COM",
	"SUB_COMMENT",
	"SUB_CUR",
	"SUB_ENDDATE",
	"SUB_ORDERID",
	"SUB_PERIOD_MOMENT",
	"SUB_PERIOD_MOMENT_M",
	"SUB_PERIOD_MOMENT_WW",
	"SUB_PERIOD_NUMBER",
	"SUB_PERIOD_NUMBER_D",
	"SUB_PERIOD_NUMBER_M",
	"SUB_PERIOD_NUMBER_WW",
	"SUB_PERIOD_UNIT",
	"SUB_STARTDATE",
	"SUB_STATUS",
	"TAAL",
	"TAXINCLUDED*XX*",
	"TBLBGCOLOR",
	"TBLTXTCOLOR",
	"TID",
	"TITLE",
	"TOTALAMOUNT",
	"TP",
	"TRACK2",
	"TXTBADDR2",
	"TXTCOLOR",
	"TXTOKEN",
	"TXTOKENTXTOKENPAYPAL",
	"TXSHIPPING",
	"TXSHIPPINGLOCATIONPROFILE",
	"TXURL",
	"TXVERIFIER",
	"TYPE_COUNTRY",
	"UCAF_AUTHENTICATION_DATA",
	"USERID",
	"USERTYPE",
	"VERSION",
	"WBTU_MSISDN",
	"WBTU_ORDERID",
	"WEIGHTUNIT",
	"WIN3DS",
	"WITHROOT",
}

var shaInAllowedParams = append(
	shaInAliasAllowedParams,
	"CARDNO",
	"CN",
	"CVC",
	"CVCFLAG",
	"ECOM_PAYMENT_CARD_EXPDATE_MONTH",
	"ECOM_PAYMENT_CARD_EXPDATE_YEAR",
	"ECOM_PAYMENT_CARD_NAME",
	"ECOM_PAYMENT_CARD_VERIFICATION",
	"UCAF_PAYMENT_CARD_CVC2",
	"UCAF_PAYMENT_CARD_EXPDATE_MONTH",
	"UCAF_PAYMENT_CARD_EXPDATE_YEAR",
	"UCAF_PAYMENT_CARD_NUMBER",
)

var shaOutAllowedParams = []string{
	"AAVADDRESS",
	"AAVCHECK",
	"AAVMAIL",
	"AAVNAME",
	"AAVPHONE",
	"AAVZIP",
	"ACCEPTANCE",
	"ALIAS",
	"AMOUNT",
	"BIC",
	"BIN",
	"BRAND",
	"CARDNO",
	"CCCTY",
	"CN",
	"COLLECTOR_BIC",
	"COLLECTOR_IBAN",
	"COMPLUS",
	"CREATION_STATUS",
	"CREDITDEBIT",
	"CURRENCY",
	"CVCCHECK",
	"DCC_COMMPERCENTAGE",
	"DCC_CONVAMOUNT",
	"DCC_CONVCCY",
	"DCC_EXCHRATE",
	"DCC_EXCHRATESOURCE",
	"DCC_EXCHRATETS",
	"DCC_INDICATOR",
	"DCC_MARGINPERCENTAGE",
	"DCC_VALIDHOURS",
	"DEVICEID",
	"DIGESTCARDNO",
	"ECI",
	"ED",
	"EMAIL",
	"ENCCARDNO",
	"FXAMOUNT",
	"FXCURRENCY",
	"IP",
	"IPCTY",
	"MANDATEID",
	"MOBILEMODE",
	"NBREMAILUSAGE",
	"NBRIPUSAGE",
	"NBRIPUSAGE_ALLTX",
	"NBRUSAGE",
	"NCERROR",
	"ORDERID",
	"PAYID",
	"PAYMENT_REFERENCE",
	"PM",
	"SCO_CATEGORY",
	"SCORING",
	"SEQUENCETYPE",
	"SIGNDATE",
	"STATUS",
	"SUBBRAND",
	"SUBSCRIPTION_ID",
	"TRXDATE",
	"VC",
}
