package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "CUST001",
			Name:        "Alice Tan",
			Email:       "alice.tan@example.com",
			Phone:       "+60123456789",
			Address:     "123 Jalan Merdeka, Kuala Lumpur",
			DateOfBirth: "1990-01-15",
			Status:      "active",
		},
		{
			Id:          "CUST002",
			Name:        "Ben Lim",
			Email:       "ben.lim@example.com",
			Phone:       "+60129876543",
			Address:     "456 Jalan Tun Razak, Kuala Lumpur",
			DateOfBirth: "1985-03-22",
			Status:      "inactive",
		},
		{
			Id:          "CUST003",
			Name:        "Chong Mei Ling",
			Email:       "chong.ml@example.com",
			Phone:       "+60134567890",
			Address:     "789 Jalan Ampang, Selangor",
			DateOfBirth: "1992-07-30",
			Status:      "active",
		},
		{
			Id:          "CUST004",
			Name:        "David Wong",
			Email:       "david.wong@example.com",
			Phone:       "+60131234567",
			Address:     "12 Lorong Damai, Penang",
			DateOfBirth: "1988-11-05",
			Status:      "active",
		},
		{
			Id:          "CUST005",
			Name:        "Elaine Goh",
			Email:       "elaine.goh@example.com",
			Phone:       "+60137654321",
			Address:     "88 Jalan Bunga Raya, Melaka",
			DateOfBirth: "1995-09-10",
			Status:      "inactive",
		},
		{
			Id:          "CUST006",
			Name:        "Farid Ahmad",
			Email:       "farid.ahmad@example.com",
			Phone:       "+60135554433",
			Address:     "21 Jalan Bayu, Johor Bahru",
			DateOfBirth: "1983-04-28",
			Status:      "active",
		},
		{
			Id:          "CUST007",
			Name:        "Gwen Lee",
			Email:       "gwen.lee@example.com",
			Phone:       "+60131231231",
			Address:     "9 Taman Intan, Kedah",
			DateOfBirth: "1991-12-01",
			Status:      "active",
		},
		{
			Id:          "CUST008",
			Name:        "Hafiz Rahman",
			Email:       "hafiz.rahman@example.com",
			Phone:       "+60134443322",
			Address:     "33 Jalan Seremban, Negeri Sembilan",
			DateOfBirth: "1986-06-14",
			Status:      "inactive",
		},
		{
			Id:          "CUST009",
			Name:        "Ivy Chan",
			Email:       "ivy.chan@example.com",
			Phone:       "+60135559988",
			Address:     "14 Jalan Kuching, Sarawak",
			DateOfBirth: "1994-08-18",
			Status:      "active",
		},
		{
			Id:          "CUST010",
			Name:        "James Tan",
			Email:       "james.tan@example.com",
			Phone:       "+60138887766",
			Address:     "7 Taman Hijau, Sabah",
			DateOfBirth: "1989-10-03",
			Status:      "active",
		},
	}
	return CustomerRepositoryStub{customers}
}
