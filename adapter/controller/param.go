package controller

type ParamControllerPreset struct {
	Param string
}

func (v *ParamControllerPreset) ReadParam() (string, error) {
	return v.Param, nil
}
