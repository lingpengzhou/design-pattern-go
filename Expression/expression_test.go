package Expression

import "testing"

func Test_Expression(t1 *testing.T) {
	robert := TerminalExpression{"Robert"}
	john := TerminalExpression{"John"}
	isMale := SetOrExpression(&robert, &john)

	julie := TerminalExpression{"Julie"}
	married := TerminalExpression{"Married"}
	isMarriedWomen := SetAndExpression(&julie, &married)

	println("John is male?", isMale.interpret("John"))
	println("Julie is a married women?", isMarriedWomen.interpret("Married Julie"))
}
