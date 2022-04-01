package main

import (
	"category/common"
	"category/controller"
	"category/domain/repository"
	"category/domain/service"
	"category/proto/category"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "ferry")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		//注册中心地址
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul注册中心
		micro.Registry(consulRegistry),
	)
	//获取mysql配置,路径中不用带前缀
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig,"mysql")

	//初始化数据库 "root:123456@/micro?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open("mysql",mysqlConfig.User+":"+mysqlConfig.Pwd+"@/"+
		mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
    //禁止复表
	db.SingularTable(true)
	//初始化服务
	srv.Init()
	//初始化操作数据库层面的repository
	categoryRepository := repository.NewCategoryRepository(db)
	//执行一次，初始化创建数据库表
	/*categoryRepository.InitTable()*/

	//初始化service
	categoryService := service.NewCategoryDataService(categoryRepository)
	//服务绑定指定Service
	err = category.RegisterCategoryHandler(srv.Server(), &controller.CategoryController{CategoryService: categoryService})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
