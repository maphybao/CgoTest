package main

import "fmt"

/*
#cgo CFLAGS: -Ihikvision/include
#cgo LDFLAGS: -Lhikvision/lib -lHCNetSDK

#include <stdio.h>
#include <stdbool.h>
#include "HCNetSDK_c.h"

NET_DVR_DEVICEINFO_V40 deviceInfo;

LONG Login()
{
	NET_DVR_USER_LOGIN_INFO loginInfo =
	{
		.sDeviceAddress = "192.168.1.4",
		.sUserName = "admin",
		.sPassword = "HIKVISION520",
		.bUseAsynLogin = 0,
		.wPort = 8000
	};

	return NET_DVR_Login_V40(&loginInfo, &deviceInfo);
}

bool CaptureImage(LONG userID, char* path)
{
    NET_DVR_JPEGPARA strPicPara =
    {
    	.wPicQuality = 2,
    	.wPicSize = 0
    };
    int iRet;
    iRet = NET_DVR_CaptureJPEGPicture(userID, deviceInfo.struDeviceV30.byStartChan, &strPicPara, path);
    if (!iRet)
    {
        printf("pyd1---NET_DVR_CaptureJPEGPicture error, %d\n", NET_DVR_GetLastError());
        return false;
    }
}
*/
import "C"

func main() {
	initState := C.NET_DVR_Init()
	fmt.Println("Init State: ", initState)

	userId := C.Login()
	if userId == -1 {
		fmt.Println("登录失败")
		C.NET_DVR_Cleanup()
		return
	}

	//	C.NET_DVR_CaptureJPEGPicture_NEW(userId)

	C.NET_DVR_Logout(userId)
	C.NET_DVR_Cleanup()
}
