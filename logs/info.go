package logs

type (
	Info interface {
		Message() string
		Data() map[string]interface{}
	}

	info struct {
		data map[string]interface{}
		msg  string
	}

	ApplicableInfo interface {
		Apply(si *subinfo)
	}

	subinfo struct {
		data map[string]interface{}
	}

	subinfoMap map[string]interface{}

	subinfoKV struct {
		key   string
		value interface{}
	}
)

func NewInfo(msg string, subinfos ...ApplicableInfo) Info {
	s := &subinfo{
		data: make(map[string]interface{}),
	}

	length := len(subinfos)
	for i := 0; i < length; i++ {
		subinfos[i].Apply(s)
	}
	return &info{
		data: s.data,
		msg:  msg,
	}
}

func (i *info) Message() string {
	return i.msg
}
func (i *info) Data() map[string]interface{} {
	return i.data
}

func (s subinfoMap) Apply(si *subinfo) {
	for key, value := range s {
		si.data[key] = value
	}
}

func (s subinfoKV) Apply(si *subinfo) {
	si.data[s.key] = s.value
}

func KeyValue(key string, value interface{}) ApplicableInfo {
	return subinfoKV{key, value}
}

func Map(m map[string]interface{}) ApplicableInfo {
	return subinfoMap(m)
}
