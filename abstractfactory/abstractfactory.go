package abstractfactory

import "fmt"

// 抽象方法模式 -- 工厂方法模式的多种方法结合在一起的模式
// https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/abstract_factory.html

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//RDBMainDAP 为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct {
}

var _ OrderMainDAO = (*RDBMainDAO)(nil)

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

type RDBDetailDAO struct {
}

var _ OrderDetailDAO = (*RDBDetailDAO)(nil)

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct {
}

var _ OrderDetailDAO = (*RDBDAOFactory)(nil)

func (*RDBDAOFactory) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}
