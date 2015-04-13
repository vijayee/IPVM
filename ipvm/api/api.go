package api
import(
	vm "github.com/robertkrimen/otto"
	"error"
	"github.com/deckarep/golang-set"
	"reflect"
)

//Holds all api's in individual namespaces
type API struct{
	registry map[string]namespace
}
// the single instance of the API type
var instance *API = nil

//initilialize the singleton
func init() {
	if instance == nil{
		instance = &API{make(map[string]namespace)}
	}
}

//return true if a namespace's key is contained in the map
func (a *API) contains(k string) bool{
	_, ok := a.registry[k]
	return ok
}

//return true if a module key is contained in the map
func (a *API) Contains(k string) bool{
	for _, value := range a.registry{

		ok:= value.contains(k)
	}
	return ok
}
// add namespace key to teh registry
func (a *API) add(k string) {
		a.registry[k] = make(namespace)
}

//Define a new module  with dependencies at pkg namespace
func (a *API)Define(deps []string, pkg string, name string, value interface {}){
	if (typedef(value) != reflect.Func){
		a.Define(pkg, name, value)
	}
	init()
	if (!a.contains(pkg)){
		a.add(pkg)
	}
	a.registry[pkg].define(deps, name, value)
}
//Define a new module  without dependencies at pkg namespace
func (a *API)Define(pkg string, name string, value interface {}){
	init()
	if (!a.contains(pkg)){
		a.add(pkg)
	}
	a.registry[pkg].define(name, value)
}

func (a *API) Require(pkg string, name string) vm.value {
	var d []interface{}
	mdl:= registry[pkg].require(name)
	deps:= mdl.Dependencies
	for dep:= range deps.ToSlice(){
		d[len(d)++] := a.require(newDeps,pkg,dep)
	}

	return vm.ToValue(registry[pkg].require(name))

}
func(a *API) require(deps *Set, pkg string, name string, d []interface{}) interface{}{
	mdl:= registry[pkg].require(name, value)

	newDeps := deps.Union(mdl.Dependencies)
	for dep:= range deps.Difference(mdl.Dependencies).ToSlice(){
		d[len(d)++] := a.require(newDeps,dep)
	}
}
func (a *API) set(k string, v interface {}){
	a.registry[k]= &module{}
}
func typedef( value interface{}) string{
	value := reflect.ValueOf(value)
	return value
}
