package util
// 一个yaml解析工具，目前的作用是代替目前所有的服务器配置json文件

import (
	"encoding/json"
	"fmt"
)

const (
	Server_Gate   = "gatesvr"
	Server_Admin  = "adminsvr"
	Server_Login  = "loginsvr"
	Server_Hall   = "hallsvr"
	Server_Game   = "gamesvr"
	Server_Api    = "apisvr"
	Server_Robots = "robotsvr"
)

type Server_Yamler interface {
	Name() string
	Etc() interface{}
}

func LoadServeYaml(server Server_Yamler, need_service bool) error {
	return new(ServerYaml).yaml2JsonStruct(server.Name(), need_service, server.Etc())
}

/*
 * @Description //把yaml配置数据转换为二进制数据在json解析到服务器配置结构体
 * @Date 17:22 2019/2/14
 * @Param courseName 进程名
 * @Param b_service  是否需要第三方服务配置
 * @Param structdst  结构体指针
 * @return
 **/
func (yaml *ServerYaml) yaml2JsonStruct(courseName string, b_service bool, structdst interface{}) error {
	if err := yaml.loader(); err != nil {
		return fmt.Errorf("Yaml2JsonStruct:%s", err)
	}
	data, err := yaml.configByName(courseName)
	if err != nil {
		return fmt.Errorf("Yaml2JsonStruct:%s", err)
	}
	if b_service {
		if err := json.Unmarshal(yaml.serviceConfig(), structdst); err != nil {
			return fmt.Errorf("Yaml2JsonStruct:%s", err)
		}
	}
	return json.Unmarshal(data, structdst)
}


type ServerYaml struct {
	// 第三方服务配置
	Service struct {
		Wxappid     string            `yaml:"wxappid" json:"wxappid"`
		Wxappsecert string            `yaml:"wxappsecert" json:"wxappsecert"`
		Smshost     string            `yaml:"smshost" json:"smshost"`
		Adminhost   string            `yaml:"adminhost" json:"adminhost"`
		Redbaghost  string            `yaml:"redbaghost" json:"redbaghost"`
		DrobotToken map[string]string `yaml:"drobotToken" json:"drobotToken"`
	} `yaml:"service"`
	// 各服务器配置 进程名> 映射 >配置结构
	ServerConfig map[string]*struct {
		ID              int    `yaml:"id" json:"id"`
		Port            int    `yaml:"port" json:"port"`
		Safehost        string `yaml:"safehost" json:"safehost"`
		Usesafeip       int    `yaml:"usesafeip" json:"usesafeip"`
		Host            string `yaml:"host" json:"host"`
		Inhost          string `yaml:"inhost" json:"inhost"`
		Gate            string `yaml:"gate" json:"gate"`
		Login           string `yaml:"login" json:"login"`
		SafeHall        string `yaml:"safehall" json:"safehall"`
		Hall            string `yaml:"hall" json:"hall"`
		Api             string `yaml:"api" json:"api"`
		Inhall          string `yaml:"inhall" json:"inhall"`
		EnabledYk       bool   `yaml:"enabledyk" json:"enabledyk"`
		Encode          int    `yaml:"encode" json:"encode"`
		Encodeclientkey string `yaml:"encodeclientkey" json:"encodeclientkey"`
		Encodephpkey    string `yaml:"encodephpkey" json:"encodephpkey"`
		Redis           string `yaml:"redis" json:"redis"`
		Redisdb         int    `yaml:"redisdb" json:"redisdb"`
		Redisauth       string `yaml:"redisauth" json:"redisauth"`
		Rcredis         string `yaml:"rcredis" json:"rcredis"`
		Rcredisdb       int    `yaml:"rcredisdb" json:"rcredisdb"`
		DB              string `yaml:"db" json:"db"`
		DB_mysql        struct {
			User      string `yaml:"user"`
			PassWord  string `yaml:"password"`
			Port      int    `yaml:"port"`
			Addr      string `yaml:"addr"`
			DBname    string `yaml:"dbname"`
			TimeOut   string `yaml:"timeout"`
			ParseTime bool   `yaml:"parsetime"`
			Loc       string `yaml:"loc"`
			Charset   string `yaml:"charset"`
		} `yaml:"db_mysql"`
		DB_redis struct {
			Auth     string `yaml:"auth"`
			SelectDB int    `yaml:"db"`
			Addr     string `yaml:"addr"`
			Port     int    `yaml:"port"`
		} `yaml:"db_redis"`
		PubGame         bool   `yaml:"pubgame" json:"pubgame"`
		FollowHall      string `yaml:"followhall" json:"followhall"`
		FollowHallPort  string `yaml:"followhallport" json:"followhallport"`
		DBFollow        string `yaml:"dbfollow" json:"dbfollow"`
		DB_mysql_Follow struct {
			User      string `yaml:"user"`
			PassWord  string `yaml:"password"`
			Port      int    `yaml:"port"`
			Addr      string `yaml:"addr"`
			DBname    string `yaml:"dbname"`
			TimeOut   string `yaml:"timeout"`
			ParseTime bool   `yaml:"parsetime"`
			Loc       string `yaml:"loc"`
			Charset   string `yaml:"charset"`
		} `yaml:"db_mysql_follow"`
		DB_redis_Follow struct {
			Auth     string `yaml:"auth"`
			SelectDB int    `yaml:"db"`
			Addr     string `yaml:"addr"`
			Port     int    `yaml:"port"`
		} `yaml:"db_redis_follow"`
		RedisFollow     string `json:"redis_follow"`      // 同服游戏
		RedisFollowDB   int    `json:"redis_follow_db"`   // 同服游戏
		RedisFollowAuth string `json:"redis_follow_auth"` // 同服游戏
	} `yaml:"server_config"`
}

/*
 * @Description //根据进程名得到其配置的二进制数据
 * @Date 17:06 2019/2/14
 * @Param courseName 进程名
 * @return
 **/
func (self *ServerYaml) configByName(courseName string) ([]byte, error) {
	svr, ok := self.ServerConfig[courseName]
	if !ok {
		return nil, fmt.Errorf("configByName:No config named: %s", courseName)
	}
	return HF_JtoB(svr), nil
}

/*
 * @Description //得到配置的第三方服务数据
 * @Date 17:06 2019/2/14
 * @return
 **/
func (self *ServerYaml) serviceConfig() []byte {
	return HF_JtoB(&self.Service)
}

/*
 * @Description //加载配置
 * @Date 17:06 2019/2/14
 * @Param
 * @return
 **/
func (self *ServerYaml) loader() error {
	if err := HF_ReadYaml("../etc", "cty_server.yaml", self); err != nil {
		return fmt.Errorf("loader:%s", err)
	}
	return self.convertor()
}

/*
 * @Description //转换配置 使兼容当前需要配置
 * @Date 17:06 2019/2/14
 * @Param
 * @return
 **/
func (self *ServerYaml) convertor() error {
	var hallhost, hallInhost, loginhost, gatehost, apihost string
	for k, v := range self.ServerConfig {
		// 服务器主机地址
		v.Host = fmt.Sprintf("%s:%d", v.Host, v.Port)
		// 服务器内网地址
		v.Inhost = fmt.Sprintf("%s:%d", v.Inhost, v.Port)
		// 服务器安全地址
		v.Safehost = fmt.Sprintf("%s:%d", v.Safehost, v.Port)
		// redis 地址
		v.Redis = fmt.Sprintf("%s:%d", v.DB_redis.Addr, v.DB_redis.Port)
		// redis 库
		v.Redisdb = v.DB_redis.SelectDB
		// redis auth
		v.Redisauth = v.DB_redis.Auth

		v.RedisFollow = fmt.Sprintf("%s:%d", v.DB_redis_Follow.Addr, v.DB_redis_Follow.Port)
		v.RedisFollowAuth = v.DB_redis_Follow.Auth
		v.RedisFollowDB = v.DB_redis_Follow.SelectDB
		v.FollowHall = fmt.Sprintf("%s:%s", v.FollowHall, v.FollowHallPort)
		// 数据库
		v.DB = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=%t&loc=%s",
			v.DB_mysql.User,
			v.DB_mysql.PassWord,
			v.DB_mysql.Addr,
			v.DB_mysql.Port,
			v.DB_mysql.DBname,
			v.DB_mysql.Charset,
			v.DB_mysql.TimeOut,
			v.DB_mysql.ParseTime,
			v.DB_mysql.Loc,
		)
		v.DBFollow = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=%t&loc=%s",
			v.DB_mysql_Follow.User,
			v.DB_mysql_Follow.PassWord,
			v.DB_mysql_Follow.Addr,
			v.DB_mysql_Follow.Port,
			v.DB_mysql_Follow.DBname,
			v.DB_mysql_Follow.Charset,
			v.DB_mysql_Follow.TimeOut,
			v.DB_mysql_Follow.ParseTime,
			v.DB_mysql_Follow.Loc,
		)
		// v.DB = "root:root@tcp(192.168.1.169:3306)/kygame?charset=utf8mb4&timeout=10s&parseTime=true&loc=Local"
		if k == Server_Hall {
			hallhost = v.Host
			hallInhost = v.Inhost
		}

		if k == Server_Login {
			loginhost = v.Host
		}

		if k == Server_Gate {
			gatehost = v.Host
		}

		if k == Server_Api {
			apihost = v.Host
		}
	}

	for _, v := range self.ServerConfig {
		// 大厅地址
		v.Hall = hallhost
		// 内网大厅地址
		v.Inhall = hallInhost
		// 登录服地址
		v.Login = loginhost
		// 大厅服安全地址
		v.SafeHall = hallhost
		// 网关地址
		v.Gate = gatehost
		// api 服务器地址
		v.Api = apihost
		//fmt.Println(fmt.Sprintf("%+v", v))
	}

	return nil
}
