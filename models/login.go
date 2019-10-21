package models

type LoginResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		User  Admin  `json:"user"`
		Token string `json:"token"`
		Name  string `json:"name"`
		UUid  string `json:"uuid"`
	} `json:"data"`
}
