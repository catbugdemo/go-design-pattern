package abstractfactory

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderDetailDAO().SaveOrderDetail()
	factory.CreateOrderMainDAO().SaveOrderMain()
}

func TestE(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}
