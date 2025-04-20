package controller


func (r *Controller) GetSession(token string) bool {
	return r.authLib.Validate(token) 
}
