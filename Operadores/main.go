package main

import "fmt"

func main() {
	/*
	 *	Opereadores Aritmeticos.
	 *	+	Suma
	 *	-	Resta
	 *	*	Multiplicacion
	 *	/	Division
	 *	%	Resto
	 */

	a := 5
	b := 3

	fmt.Println("-------------------------")
	fmt.Println("5 + 3 = ", a+b)
	fmt.Println("5 - 3 = ", a-b)
	fmt.Println("5 * 3 = ", a*b)
	fmt.Println("5 / 3 = ", a/b)
	fmt.Println("5 % 3 = ", a%b)

	/*
	 *	Operadores de comparación.
	 *	==	Igual que
	 *	!=	Diferente que
	 *	<	Menor que
	 *	<=	Menor o Igual que
	 *	>	Mayor que
	 *	>=	Mayor o Igual que
	 */

	fmt.Println("-------------------------")
	fmt.Println("5 == 3 = ", a == b)
	fmt.Println("5 != 3 = ", a != b)
	fmt.Println("5 < 3 = ", a < b)
	fmt.Println("5 <= 3 = ", a <= b)
	fmt.Println("5 > 3 = ", a > b)
	fmt.Println("5 >= 3 = ", a >= b)

	/*
	 *	Operadores de asignación.
	 *	|	+=	|	a += b	|	a = a + b	|
	 *	|	-=	|	a -= b	|	a = a + b	|
	 *	|	*=	|	a *= b	|	a = a + b	|
	 *	|	/=	|	a /= b	|	a = a + b	|
	 *	|	%=	|	a %= b	|	a = a + b	|
	 */

	c := 5
	c += b

	fmt.Println("-------------------------")
	fmt.Println("5 += 3 = ", c)
	c -= b
	fmt.Println("5 -= 3 = ", c)
	c *= b
	fmt.Println("5 *= 3 = ", c)
	c /= b
	fmt.Println("5 /= 3 = ", c)
	c %= b
	fmt.Println("5 %= 3 = ", c)

	/*
	 *	Operadores logicos
	 *	&&	AND
	 *	||	OR
	 *	!	NOT
	 */

	fmt.Println("-------------------------")
	fmt.Println("[AND]")
	// fmt.Println("true && true =", true && true)
	fmt.Println("false && true =", false && true)
	fmt.Println("true && false =", true && false)
	// fmt.Println("false && false =", false && false)
	fmt.Println("[OR]")
	// fmt.Println("true || true =", true || true)
	fmt.Println("false || true =", false || true)
	fmt.Println("true || false =", true || false)
	// fmt.Println("false || false =", false || false)

	/*
	 *	Jerarquia de operadores
	 *	()
	 *	*		/
	 *	+		-
	 *	==		!=		<		<=		>		>=
	 *	&&
	 *	||
	 */

	fmt.Println("-------------------------")
	fmt.Println("5 - 3 * 5 = ", 5-3*5)
	fmt.Println("(5 - 3) * 5 = ", (5-3)*5)

	/*
	 *	Operadores de incremento y decremento
	 *	++	Sumar uno
	 *	--	Restar uno
	 */

	fmt.Println("-------------------------")
	a++
	fmt.Println("a++ = ", a)
}
