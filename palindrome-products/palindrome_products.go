// Copied from test file.

/* API to implement:

type Product struct {
	Product int // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
 }

 func Products(fmin, fmax int) (pmin, pmax Product, error)
*/

package palindrome

import (
	"fmt"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin Product, pmax Product, err error) {

	if fmin > fmax {
		return pmin, pmax, fmt.Errorf("fmin > fmax")
	}

	products := computeProducts(fmin, fmax)

	var palindromes []Product

	for i := 0; i < len(products); i++ {
		var current_product Product = products[i]

		if checkIfPalindrome(current_product.Product) {

			palindromes = append(palindromes, current_product)
		}
	}

	if len(palindromes) < 1 {
		return pmin, pmax, fmt.Errorf("no palindromes...")
	}

	var maximumPalindromeProduct Product = Product{0, [][2]int{{0, 0}}}
	var minimumPalindromeProduct Product = Product{fmax * fmin, [][2]int{{fmin, fmax}}}

	for i := 0; i < len(palindromes); i++ {
		currentProduct := palindromes[i]

		if currentProduct.Product > maximumPalindromeProduct.Product {
			maximumPalindromeProduct = currentProduct

		}

		if currentProduct.Product < minimumPalindromeProduct.Product {
			minimumPalindromeProduct = currentProduct

		}

	}

	return minimumPalindromeProduct, maximumPalindromeProduct, nil

}

// Go is fucking retarded and seems to not have a minimum/max function for integers?
// https://golang.org/pkg/math/

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeProducts(min int, max int) []Product {

	productsArray := make([]Product, 0)

	productsMap := make(map[int][][2]int)

	for i := min; i <= max; i++ {
		for j := max; j >= min; j-- {
			prod := i * j
			if value, ok := productsMap[prod]; ok {
				var doesFactorExist bool = false
				// If the product already exists in the map, then simply append the
				// factors to the factorization field -- but only if the factors
				// are new.
				for i := 0; i < len(value); i++ {
					// Iterate over the list of factors found thus far.
					a := value[i][0]

					// We only need to check if one of the two factors appears in the
					// list, since if a * b = X, then it's not possible for a * c = X
					// as well if c != b.
					if a == i || a == j {
						//fmt.Println("Oops, duplicate factors found -- ", a, b)
						doesFactorExist = true
						break
					}

				}
				if doesFactorExist == false {
					productsMap[i*j] = append(value, [2]int{minimum(i, j), maximum(i, j)})
				}
			} else {
				// If the product doesn't exist in the map, insert it
				productsMap[i*j] = [][2]int{{minimum(i, j), maximum(i, j)}}
			}

		}
	}

	// iterate through the map, and use it to populate the array of products

	for key, value := range productsMap {
		productsArray = append(productsArray, Product{key, value})
	}

	return productsArray
}

func checkIfPalindrome(number int) bool {
	// Create two iterators, one from the beginning and one from the end;
	// Increment the forward iterator while decrementing the reverse iterator.
	// If they're ever different then you can't have a palindrome.
	input := strconv.Itoa(number)

	mid := len(input) / 2
	last := len(input) - 1
	for i := 0; i < mid; i++ {
		if input[i] != input[last-i] {
			return false
		}
	}
	return true
}
