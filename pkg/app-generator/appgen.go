package appgenerator

//Appgen holding application generator data
type Appgen struct {
	appName string
	repo    string
}

//MainFuncData used on main function template
type MainFuncData struct {
	AppName string
	Repo    string
}

//New create appgen instance
func New(appName string, repo string) (appgen *Appgen) {
	appgen = &Appgen{
		appName: appName,
		repo:    repo,
	}
	return
}
