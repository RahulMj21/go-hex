package ports

type GRPCPorts interface {
	Run()
	GetAddition()
	GetSubstraction()
	GetMultiplication()
	GetDivision()
}
