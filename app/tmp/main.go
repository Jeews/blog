// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "blog/app"
	controllers "blog/app/controllers"
	tests "blog/tests"
	controllers1 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Main",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					14: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Blogger)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "BloggerPage",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					14: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Login)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "SignIn",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					21: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SignInHandler",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "passwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SignUp",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					49: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SignUpHandler",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "passwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Main)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Main",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Main",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					72: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
