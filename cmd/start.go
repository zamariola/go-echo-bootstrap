package cmd

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zamariola/go-echo-bootstrap/internal"
)

var (
	port                                   string
	dbhost, dbport, dbuser, dbpass, dbname string
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

// Execute starts the application
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
	rootCmd.PersistentFlags().StringVar(&port, "port", "", "Server port, default: 9011")
	rootCmd.PersistentFlags().StringVar(&dbhost, "db_host", "", "Database host")
	rootCmd.PersistentFlags().StringVar(&dbport, "db_port", "", "Database port")
	rootCmd.PersistentFlags().StringVar(&dbuser, "db_user", "", "Database user")
	rootCmd.PersistentFlags().StringVar(&dbpass, "db_pass", "", "Database password")
	rootCmd.PersistentFlags().StringVar(&dbname, "db_name", "", "Database name")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("db_host", rootCmd.PersistentFlags().Lookup("db_host"))
	viper.BindPFlag("db_port", rootCmd.PersistentFlags().Lookup("db_port"))
	viper.BindPFlag("db_user", rootCmd.PersistentFlags().Lookup("db_user"))
	viper.BindPFlag("db_pass", rootCmd.PersistentFlags().Lookup("db_pass"))
	viper.BindPFlag("db_name", rootCmd.PersistentFlags().Lookup("db_name"))

	viper.SetDefault("port", "9011")
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
