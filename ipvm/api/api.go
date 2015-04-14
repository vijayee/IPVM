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
	//only functions have dependencies
	if typedef(value) != "func"{
		a.Define(pkg, name, value)
		return
	}
	t := reflect.TypeOf(value)
	//abort if function does not have enough parameters to receive dependencies
	if t.NumIn() < len(deps){
		return
	}
	//does not support multiple return values
	if t.NumOut() > 1){
		return
	}
	init()
	if !a.contains(pkg){
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
	var d map[string]interface{}
	mdl:= registry[pkg].require(name)
	deps:= mdl.Dependencies
	ds:= deps.ToSlice()
	val := reflect.ValueOf(mdl.Exports)	
	typ := reflect.TypeOf(mdl.Exports)
	n:= typ.NumIn() 
	for _ , dep:= range ds{
		d[dep] := a.require(newDeps,pkg,dep, d)
	}
	vmmdl:= func(call otto.FunctionCall) vm.Value { 
		if(n != len(args.ArgumentList))
		if mdl.ModuleType == "func" {
			if(len(ds) > 0){
				
			else{				
				if n != len(call.ArgumentList)
				{
					return vm.toValue(nil)
				}
				else{
					var p []reflect.Value
					for i, arg := range.ArgumentList{
						p[i]:= reflect.ValueOf(arg.export())						
					}
					for i:= 0; i < n; i++{
						
					}
				}
			}
		}
		else{
			return vm.toValue(mdl.Exports)
		}
	}
	return vm.ToValue(registry[pkg].require(name))

}
func(a *API) require(deps *Set, pkg string, name string,  map[string]interface{}) interface{}{
	mdl:= registry[pkg].require(name, value)

	newDeps := deps.Union(mdl.Dependencies)
	for dep:= range deps.Difference(mdl.Dependencies).ToSlice(){
		d[dep] := a.require(newDeps,dep)
	}
	newDeps
}
func (a *API) set(k string, v interface {}){
	a.registry[k]= &module{}
}
func typedef(value interface{}) string {
	t := reflect.ValueOf(value)
	return t.Kind().String()
}

func parameterize(call otto.FunctionCall, method reflect.Value) []reflect.Value {
	m := method.Type()
	var a []reflect.Value
	a = make([]reflect.Value, m.NumIn(), m.NumIn())
	for i := 0; i < m.NumIn(); i++ {
		arg, _ := call.Argument(i).Export()
		a[i] = reflect.ValueOf(arg).Convert(m.In(i))
	}
	return a
}