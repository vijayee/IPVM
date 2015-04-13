package ipvm
import(
	vm "github.com/robertkrimen/otto"
	"github.com/vijayee/IPVM/ipvm/api"
)

type IPVM struct {
	vm *vm
	api API
}

func New() *IPVM {
	return &IPVM{vm: vm.New() vm: api: api.New()}

}
// Set up a Dependency Injector
func (v *IPVM) injectAPI(m map[string] interface {}){
	v.vm.Set("require",func(call vm.FunctionCall) vm.Value {
	k:= call.Argument(0).String()
	return v.api.Require(k)
})
}
