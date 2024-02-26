package pupil

type Pupil struct{
	Name string
	Age int
	Class int
	Wallet  Cash
	CheatThing CheatThing
	Uniform  int
	Exam  TpExam  
}
type TpExam struct{
	TypeEx, WhichEx string
}

type CheatThing struct{
	CheatSheets bool
	Phone bool
}
type Cash struct{
	Money int
	Card  bool
}