package main

//2311102309
import (
    "fmt"
    "math/big"
)

func factorial(n int64) *big.Int {
    result := big.NewInt(1)
    for i := int64(2); i <= n; i++ {
        result.Mul(result, big.NewInt(i))
    }
    return result
}
func perm(n, r int64) *big.Int {
    if r > n {
        return big.NewInt(0)
    }
    return new(big.Int).Div(factorial(n), factorial(n-r))
}

func comb(n, r int64) *big.Int {
    if r > n {
        return big.NewInt(0)
    }

    return new(big.Int).Div(factorial(n), new(big.Int).Mul(factorial(r),
        factorial(n-r)))
}

func main() {
    var a, b, c, d int64
    fmt.Print("Masukkan empat angka :")
    fmt.Scan(&a, &b, &c, &d)
    P_ac := perm(a, c)
    C_ac := comb(a, c)
    P_bd := perm(b, d)
    C_bd := comb(b, d)
    fmt.Printf("Permutasi (%d, %d) = %d\n", a, c, P_ac)
    fmt.Printf("Kombinasi (%d, %d) = %d\n", a, c, C_ac)
    fmt.Printf("Permutasi (%d, %d) = %d\n", b, d, P_bd)

    fmt.Printf("Kombinasi (%d, %d) = %d\n", b, d, C_bd)
}

