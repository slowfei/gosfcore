//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	net 辅助工具
//
package SFNetUtil

import (
	"net"
)

//	获取本地内网IP
//	@return net.IP
func LocalIP() net.IP {
	var result net.IP = nil

	netInters, _ := net.Interfaces()
	for _, netInter := range netInters {

		addrs, _ := netInter.Addrs()
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			v4 := ipnet.IP.To4()
			if v4 == nil || v4[0] == 127 { // loopback address
				continue
			}
			return v4
		}
	}
	return result
}
