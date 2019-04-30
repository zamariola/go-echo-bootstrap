package cmd

import (
	"fmt"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zamariola/go-echo-bootstrap/internal"
	"os"
)

var (
	port string
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Server starts the webserver application ",
	Run: func(cmd *cobra.Command, args []string) {

		//configure echo
		log.Infof("Starting web application with flags %s", os.Args[1:])
		echo := internal.WireRoutes()
		echo.Use(middleware.Logger())
		echo.Use(middleware.Recover())

		//define port using flags, env and config file
		sPort, ok := viper.Get("port").(string)
		if !ok {
			panic(fmt.Errorf("Error while parsing port number %s", viper.Get("port")))
		}

		//Start echo server
		log.Fatal(echo.Start(":" + sPort))
	},
}

//Execute starts the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	setupConfig()
	setupLogrus()
}

func setupConfig() {
	cobra.OnInitialize(setupViper)
	rootCmd.PersistentFlags().StringVar(&port, "port", "", "Server port, default: 9010")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", "9010")
}

func setupViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func setupLogrus() {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)
}
