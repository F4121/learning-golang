package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Empty ID"} // Karena error adalah interface maka bentuk balikan haruslah pointer
	}
	if id != "faizi" {
		return &notFoundError{"Data Not Found"} // Karena error adalah interface maka bentuk balikan haruslah pointer
	}

	// ok
	return nil
}

func main() {
	/**
	MEMBUAT CUSTOM ERROR
	- Karena error adalah sebuah interface, jadi jika kita ingin membuat error sendiri, kita bisa membuat struct yang mengikuti kontrak dari interface error
	*/

	err := SaveData("budi", nil)

	// Contoh IF
	if err != nil {
		// **** SAMPLE USING IF
		/**
		validationErr adalah variable (variable baru)
		ok := err.(*validationError) <== ini adalah type assertion / digunakan untuk mengkonversi ok menjadi validationError
		ok adalah bool yang mana digunakan untuk menentukan kondisi pada if statement
		*/
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error :", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error :", notFoundErr.Error())
		} else {
			fmt.Println("Unknown error", err.Error())
		}
		// **** END SAMPLE USING IF

		// **** SAMPLE USING SWITCH
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error :", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found error :", finalError.Error())
		default:
			fmt.Println("Unknown error", err.Error())
		}
		// **** END SAMPLE USING SWITCH

	} else {
		fmt.Println("Success")
	}

}
