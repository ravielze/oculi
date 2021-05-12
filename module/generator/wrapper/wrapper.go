package generator_wrapper

import (
	_ "embed"
	"fmt"
	"strings"

	u "github.com/ravielze/oculi/module/generator/utils"
)

type InterfaceWrapper struct {
	Implemented []MethodWrapper
	Declared    []MethodWrapper
}

type MethodWrapper struct {
	Name      string
	Return    string
	Parameter string
	Body      string
}

type MethodWrapperSorter []MethodWrapper

func (a MethodWrapperSorter) Len() int {
	return len(a)
}

func (a MethodWrapperSorter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a MethodWrapperSorter) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

//go:embed template/func.txt
var MethodRawContent string

func MakePlaceholders(structName, methodName, parameter, result, body string) []u.Placeholder {
	return []u.Placeholder{
		u.NewPlaceholder("$$struct$$", structName+" "),
		u.NewPlaceholder("$$method_name$$", methodName),
		u.NewPlaceholder("$$parameter$$", parameter),
		u.NewPlaceholder("$$result$$", result),
		u.NewPlaceholder("$$body$$", body),
	}
}

func (m MethodWrapper) String(structName string) string {
	if len(m.Body) == 0 {
		m.Body = "panic(\"not implemented\")"
	}
	if strings.Contains(m.Return, ",") {
		m.Return = fmt.Sprintf("(%s)", m.Return)
	}
	if len((m.Return)) == 0 {
		return "\n\n" + u.Replacer(MethodRawContent, MakePlaceholders(structName, m.Name, m.Parameter, "", m.Body))
	}
	return "\n\n" + u.Replacer(MethodRawContent, MakePlaceholders(structName, m.Name, m.Parameter, " "+m.Return, m.Body))
}

func (m MethodWrapper) Equal(other MethodWrapper) bool {
	return m.Name == other.Name && m.Return == other.Return && m.Parameter == other.Parameter
}

func (implemented *MethodWrapper) ReplaceIfNeeded(declared MethodWrapper) {
	if implemented.Name != declared.Name || implemented.Equal(declared) {
		return
	}
	implemented.Name = declared.Name
	implemented.Return = declared.Return
	implemented.Parameter = declared.Parameter
}

func (m MethodWrapper) Find(mw *[]MethodWrapper, l, r int) int {
	if r >= l {
		mid := l + (r-l)/2
		if ((*mw)[mid]).Name == m.Name {
			return mid
		} else if ((*mw)[mid]).Name > m.Name {
			return m.Find(mw, l, mid-1)
		} else {
			return m.Find(mw, mid+1, r)
		}
	}
	return -1
}

func (iw InterfaceWrapper) String(structName string) string {
	result := ""
	for _, declared := range iw.Declared {
		if idx := declared.Find(&iw.Implemented, 0, len(iw.Implemented)-1); idx == -1 {
			result += declared.String(structName)
		} else {
			impl := iw.Implemented[idx]
			if !impl.Equal(declared) {
				impl.ReplaceIfNeeded(declared)
				result += impl.String(structName)
			}
		}
	}
	return result
}
