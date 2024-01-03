package data

type UserModel struct {
	ID         int
	Email      string
	Cpf        string
	Password   string
	CreditCard CreditCard
}
