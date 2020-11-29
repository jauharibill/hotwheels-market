package domain

type UserInterface interface {
	Create()
	Index()
	Show()
	Delete()
	Update()
}

type UserUsecase struct {}

var user_usecase = UserUsecase{}

func (_u *UserUsecase) Create() {

}

func (_u *UserUsecase) Index() {
	
}

func (_u *UserUsecase) Show() {
	
}

func (_u *UserUsecase) Delete() {
	
}

func (_u *UserUsecase) Update()  {

}