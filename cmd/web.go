package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rwirdemann/databasedragon/web/app"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

func init() {
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Starts local web application",
	Run: func(cmd *cobra.Command, args []string) {
		router := mux.NewRouter()
		app.RegisterHandler(router)
		log.Printf("Listening on :%d...", app.Conf.Web.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", app.Conf.Web.Port), router); err != nil {
			log.Fatal(err)
		}
	},
}