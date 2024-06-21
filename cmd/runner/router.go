package runner

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"reflect"
	"runtime"
	"service/internal/app/api"
	"strconv"
	"strings"

	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/router"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "route-list",
		Long: "Running Route List",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()

			newRoute := mux.NewRouter()
			router.RegisterRouter(newRoute, api.Register)

			methodLen := 0
			pathLen := 0

			routeLists := make([][]string, 0)
			newRoute.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
				path, err := route.GetPathTemplate()
				if err != nil {
					return err
				}

				methods, _ := route.GetMethods()
				handlerName := getFunctionName(route.GetHandler())

				if handlerName != nil {
					method := fmt.Sprintf("[%s]", strings.Join(methods, ", "))
					if methodLenStr := len(method); methodLenStr > methodLen {
						methodLen = methodLenStr
					}

					if pathLenStr := len(path); pathLenStr > pathLen {
						pathLen = pathLenStr
					}

					routeLists = append(routeLists, []string{method, path, *handlerName})
				}

				return nil
			})

			methodFormat := "%-" + strconv.Itoa(methodLen+3) + "s"
			pathFormat := "%-" + strconv.Itoa(pathLen+5) + "s"

			for _, routeList := range routeLists {
				printMethod := fmt.Sprintf(methodFormat, routeList[0])
				printPath := fmt.Sprintf(pathFormat, routeList[1])

				fmt.Printf("%s %s %s\n", printMethod, printPath, routeList[2])
				xtremelog.Debug(routeList)
			}
		},
	})
}

func getFunctionName(f interface{}) *string {
	if f == nil {
		return nil
	}

	ptr := reflect.ValueOf(f).Pointer()
	funcName := runtime.FuncForPC(ptr).Name()

	if len(funcName) >= 4 && funcName[len(funcName)-3:] == "-fm" {
		funcName = funcName[:len(funcName)-3]
	}

	return &funcName
}
