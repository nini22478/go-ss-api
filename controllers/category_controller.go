package controllers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ErrorPage(c *gin.Context) {
	var title = c.Query("title")
	var href = c.Query("href")
	var err = c.Query("err")
	c.HTML(http.StatusOK, "error.html", gin.H{"title": title, "href": href, "err": err})
}

func AddUser(c *gin.Context) {
	port := c.PostForm("port")
	passwd := c.PostForm("passwd")
	ports, _ := strconv.Atoi(port)
	datas := map[string]interface{}{
		"server_port": ports,
		"password":    passwd,
	}

	jsonData, err := json.Marshal(datas)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  200,
			"massg": "发送数据失败，err:",
			"err":   err,
		})
	}
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8088,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code":  200,
			"massg": "链接数据失败，err:",
			"err":   err,
		})
		return
	}
	defer socket.Close()
	sendData := []byte("add: " + string(jsonData))
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		c.JSON(200, gin.H{
			"code":  200,
			"massg": "发送数据失败，err:",
			"err":   err,
		})
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		c.JSON(200, gin.H{
			"code":  200,
			"massg": "接收数据失败，err:",
			"err":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"addr":  remoteAddr,
		"count": n,
		"recv":  string(data[:n]),
	})
}

func DelUser(c *gin.Context) {
	port := c.PostForm("port")

	datas := map[string]interface{}{
		"server_port": port,
	}

	jsonData, err := json.Marshal(datas)
	if err != nil {
		panic(err)
	}
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8088,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("remove: " + string(jsonData))
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"addr":  remoteAddr,
		"count": n,
		"recv":  string(data[:n]),
	})
}

func AllUser(c *gin.Context) {

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8088,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("all")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"addr":  remoteAddr,
		"count": n,
		"recv":  string(data[:n]),
	})
}

func Ping(c *gin.Context) {

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8088,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("ping")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"addr":  remoteAddr,
		"count": n,
		"recv":  string(data[:n]),
	})
}
