package main


type user struct { // #1

id int  // #2

username int // #3

password int // #4

}


type userservice struct { // #5

t []user // #6

}


func (u userservice) getallusers() []user { // #7

return u.t // 

}


func (u userservice) getuserbyid(id int) user { // #8

for _, r := range u.t { 

if id == r.id { 

return r 

}

}


return user{} 

}

/*
a. Kekurangan Code tersebut sebanyak 8 baris
b. Bagian yang banyak kekurangannya yaitu pada penamaan kedua struct yang tidak jelas, 
   penamaan variable dan tipe data struct yang tidak sesuai, serta penamaan kedua methodnya yang tidak konsisten

c. - Pada penanda #1 dan #5 terjadi kesalahan dalam penamaan structnya, dimana sebaiknya penamaan tersebut diawali dengan huruf kapital
     dan disambung dengan salah satu gaya penulisan case apabila struct tersebut terdiri dari lebih dari satu kata
	 contohnya : type UserService struct {}
   
   - Pada penanda #2 - #4 dan #6 terjadi kesalahan dalam penamaan variable dan tipe data struct yang tidak sesuai, dimana seharusnya
     penamaan variable harus jelas mewakilkan sesuatu dan ditulis dengan huruf kapital, serta type datanya harus sesuang dengan variable
	 yang mewakilkan, contohnya :
	 username int ==> seharusnya menjadi seperti ini : Username string
	 password int ==> seharusnya menjadi seperti ini : Password string
	 t []user ==> seharusnya menjadi seperti ini : Users []User (dikarenakan nama 't' memberikan informasi tidak jelas)
   
   - Pada penanda #7 dan #8 ini, method tidak menggunakan konvensi penamaan umum pada GO, dimana seharusnya penamaan tersebut menggunakan
     salah satu case contohnya :
		getallusers ==>> GetAllUsers
		getuserbyid ==>> GetUserById
*/

