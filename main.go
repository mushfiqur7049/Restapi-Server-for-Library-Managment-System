///*
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Author struct including Firstname and Lastname of Author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book struct including the information about ID, Isbn, Title, and Author information of book
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

/// User profile struct which including User id and information about loaned book
type User_profile struct {
	User_ID string `json:"user_id"`
	Book    *Book  `json:"book"`
}

// Request loan struct which including Requested user id and requested book information
type Request_loan struct {
	User_ID string `json:"user_id"`
	Book    *Book  `json:"book"`
}

///Init books var as a slice of book struct which is used for preserving all books information.
/// When we browes book we find all books from here.
///Delete book, Update book and create book information are done here.
var books []Book

///Init authors var as a slice of author struct wich user for preserving the information of all author.
var authors []Author

///Init request_book var as a slice of Request loan struct
///When any user request of loan for book, this request are being store here.
/// Admin view request of loan for book from here and when accept & reject request then it is delete from here
var request_book []Request_loan

///Init user_profile var as a slice of user profile struct
///When admin accept request of loan for book then updated user profile by requested book are accepted by user here
///Admin_Update_Book_When_Return function update user_profile storage or array or table
var user_profile []User_profile

///----------------------------------------------------------------------------------------------------------------------///
///*********** User Functionalities********************///

// Browse and Get All books
// Browse_Books function user for finding all books
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Browse_Books(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(books)

}

///Browse and Get All Authors
// Browse_Authors function user for finding all books
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Browse_Authors(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(authors)

}

// Search  book
//Search_book function is used to find expected book from library
// After searching  data id find from books array or table
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Search_book(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	json.NewEncoder(response).Encode(&Book{})

}

/// Request for book
///Request_for_book_loan function is used for requesting book_loan to admin
/// After requesting the data are store into request_book array or table
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Request_For_Book_Loan(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.

	for _, item := range books {
		if item.ID == params["id"] {
			//added request_for_book_loan to request_book array or table
			request_book = append(request_book, Request_loan{User_ID: params["user_id"], Book: &Book{ID: item.ID, Isbn: item.Isbn, Title: item.Title, Author: &Author{Firstname: item.Author.Firstname, Lastname: item.Author.Lastname}}})
			return
		}
	}
	json.NewEncoder(response).Encode(&Request_loan{})

}

///User Profile view by user
///User_profile_view function is used for viewing user profile.
///Its need user id to correctly identify user profile from user_profile array or table
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func User_Profile_View(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.

	for _, item := range user_profile {
		if item.User_ID == params["user_id"] {
			json.NewEncoder(response).Encode(item)

		}
	}
	json.NewEncoder(response).Encode(&User_profile{})

}

///***************************End of User Functionalities****************************///
///---------------------------------------------------------------------------------------------------------------------///

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

///***************************Admin Functionalities**********************************///

// Create_Books
//Create_Books function is used by admin for create new book and added new book to books arry or table
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Create_Books(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var book Book // declare book variable
	_ = json.NewDecoder(request.Body).Decode(&book)
	//book.ID = strconv.Itoa(rand.Intn(10000000)) //MOCK ID set
	books = append(books, book) // added new book to books array or table
	json.NewEncoder(response).Encode(book)

}

// Update_Books
//Update_Books function is a admin control function for updating book information into books array or table.
//Its need books id to find book's information from books array or table then update information there.
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Update_Books(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // delete old information for books array or table
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book) // add updated information to books array or table
			json.NewEncoder(response).Encode(book)
			return
		}
	}

}

// Delete_Books
// Delete_Books function is a admin control function which is used for deleting book's information from books array or table
//Its need book id to find book from books table or array then delete it from there.
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Delete_Books(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // delete book information for books array or table
			break
		}
	}
	json.NewEncoder(response).Encode(books)

}

///Admin view request book for loan
/// when user request for new book then the requested information are store into request_book array or table .
///And the Admin_view_Request_for_book_loan function is user to viewing all requested information from request_book array or table
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Admin_View_Request_For_Book_Loan(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(request_book)

}

///Admin_View_All_Loaned_Book
///The Admin_View_All_Loaned_Book function is user to viewing all loaned book information from user_profile array or table
// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Admin_View_All_Loaned_Book(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(user_profile)

}

/// Admin Accept book request
/// When user request for book_loan then admin use Admin_Accept_Request_For_Book_Loan function for accept request from request_book array or table.
///After accepting request, the request are delete from requset_book table or arraay and again store into user_profile array or table.
/// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Admin_Accept_Request_For_Book_Loan(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.
	for index, item := range request_book {
		if item.Book.ID == params["id"] {
			/// updated user_profile array or table. (requested book are provided by admin and include it to user_profile)
			user_profile = append(user_profile, User_profile{User_ID: params["user_id"], Book: &Book{ID: item.Book.ID, Isbn: item.Book.Isbn, Title: item.Book.Title, Author: &Author{Firstname: item.Book.Author.Firstname, Lastname: item.Book.Author.Lastname}}})
			// delete from request_book array or table.(requested book are accepted by admin and delete the request)
			request_book = append(request_book[:index], request_book[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(request_book)

}

/// Admin Reject Request for book
/// When user request for book_loan then admin use Admin_Reject_Request_For_Book_Loan function for reject request from request_book array or table.
///After rejecting request, the request are delete from requset_book table or array.
/// response and request are the function parameter for taking data from http.Responsewritter and *http.Request
func Admin_Reject_Request_For_Book_Loan(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.
	for index, item := range request_book {
		if item.Book.ID == params["id"] {
			//delete request from request_book array or table.(requested book are rejected by admin and delete the request)
			request_book = append(request_book[:index], request_book[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(request_book)

}

/// Admin Update Book loan when return book
///Admin_Update_Book_When_Return function used when return book then delete the return_book information from user_profile array or table.
func Admin_Update_Book_When_Return(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // get param from http Request and mux is used as a third-party router for routing data.
	for index, item := range user_profile {
		if item.Book.ID == params["id"] {
			// delete from user_profile array or table. (Return book are delete from user_profile by admin)
			user_profile = append(user_profile[:index], user_profile[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(user_profile)

}

///***********************End of Admin Functionalities************************///
func main() {
	// Init Router
	router := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "1111", Title: "Book one", Author: &Author{Firstname: "Rifat", Lastname: "Fakir"}})
	authors = append(authors, Author{Firstname: "Rifat", Lastname: "Fakir"})
	books = append(books, Book{ID: "2", Isbn: "1112", Title: "Book two", Author: &Author{Firstname: "Rifat2", Lastname: "Fakir2"}})
	authors = append(authors, Author{Firstname: "Rifat2", Lastname: "Fakir2"})
	books = append(books, Book{ID: "3", Isbn: "1113", Title: "Book three", Author: &Author{Firstname: "Rifat3", Lastname: "Fakir3"}})
	authors = append(authors, Author{Firstname: "Rifat3", Lastname: "Fakir3"})
	books = append(books, Book{ID: "4", Isbn: "1114", Title: "Book four", Author: &Author{Firstname: "Rifat4", Lastname: "Fakir4"}})
	authors = append(authors, Author{Firstname: "Rifat4", Lastname: "Fakir4"})
	books = append(books, Book{ID: "5", Isbn: "1115", Title: "Book five", Author: &Author{Firstname: "Rifat5", Lastname: "Fakir5"}})
	authors = append(authors, Author{Firstname: "Rifat5", Lastname: "Fakir5"})
	books = append(books, Book{ID: "6", Isbn: "1116", Title: "Book six", Author: &Author{Firstname: "Rifat6", Lastname: "Fakir6"}})
	authors = append(authors, Author{Firstname: "Rifat6", Lastname: "Fakir6"})
	books = append(books, Book{ID: "7", Isbn: "1117", Title: "Book seven", Author: &Author{Firstname: "Rifat7", Lastname: "Fakir7"}})
	authors = append(authors, Author{Firstname: "Rifat7", Lastname: "Fakir7"})
	books = append(books, Book{ID: "8", Isbn: "1118", Title: "Book eight", Author: &Author{Firstname: "Rifat8", Lastname: "Fakir8"}})
	authors = append(authors, Author{Firstname: "Rifat8", Lastname: "Fakir8"})
	//Mock data end

	// Route handles & endpoints
	// *******Route handles & endpoints for User (user functionalities)*******
	router.HandleFunc("/User/Browse_Books", Browse_Books).Methods("GET")
	router.HandleFunc("/User/Browse_Authors", Browse_Authors).Methods("GET")
	router.HandleFunc("/User/Search_book/{id}", Search_book).Methods("GET")                               /// need books_id (id) for searching book correctly
	router.HandleFunc("/User/Request_For_Book_Loan/{user_id}/{id}", Request_For_Book_Loan).Methods("PUT") /// need user_id , book_id (id) to insert request_book array or table. book_id(id) are needed to spacify expected book that are requested by user (user_id).
	router.HandleFunc("/User/User_Profile_View/{user_id}", User_Profile_View).Methods("GET")              ///need user_id to correctly spacify user than show his/her profile.

	// *******Route handles & endpoints for Admin (user functionalities)*******
	router.HandleFunc("/Admin/Create_Books", Create_Books).Methods("POST")
	router.HandleFunc("/Admin/Update_Books/{id}", Update_Books).Methods("PUT")    /// book_id (id) are needed for spacify expected book than update its other information.
	router.HandleFunc("/Admin/Delete_Books/{id}", Delete_Books).Methods("DELETE") /// book_id (id) are needed for spacify expected book than delete its information.
	router.HandleFunc("/Admin/Admin_View_Request_For_Book_Loan", Admin_View_Request_For_Book_Loan).Methods("GET")
	router.HandleFunc("/Admin/Admin_View_All_Loaned_Book", Admin_View_All_Loaned_Book).Methods("GET")
	router.HandleFunc("/Admin/Admin_Accept_Request_For_Book_Loan/{user_id}/{id}", Admin_Accept_Request_For_Book_Loan).Methods("DELETE") /// need user_id , book_id (id). book_id (id) is provided to this user (user_id).
	router.HandleFunc("/Admin/Admin_Reject_Request_For_Book_Loan/{user_id}/{id}", Admin_Reject_Request_For_Book_Loan).Methods("DELETE") /// need user_id , book_id (id). book_id (id) is not provide to this user (user_id).
	router.HandleFunc("/Admin/Admin_Update_Book_When_Return/{user_id}/{id}", Admin_Update_Book_When_Return).Methods("DELETE")           /// need user_id , book_id (id). book_id (id) is return by this user (user_id).

	log.Fatal(http.ListenAndServe(":8000", router))

}
