package cmd

// var applicationName string
// var tslSecret string

// var createIngressCommand = &cobra.Command{
// 	Use:   "ingress",
// 	Short: "Create an ingress",
// 	Long: AddAppName(`Create an ingress
//     Usage:
//     $AppName create ingress <domain_name> --application=<application_name> --tls-secret=<tls_secret>`),
// 	SilenceUsage:  true,
// 	SilenceErrors: true,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		if len(args) < 1 {
// 			return fmt.Errorf("Not enough arguments - ingress domain name needed")
// 		}
// 		domainName := args[0]

// 		logging.Info("Creating ingress", logging.NewPair("domain-name", domainName))

// 		client, err := adapters.NewVampCloudHttpClient(ApiVersion, Config)
// 		if err != nil {
// 			return err
// 		}

// 		createIngress := usecase.NewCreateIngressUsecase(client)

// 		var secret *string

// 		if tslSecret != "" {
// 			secret = &tslSecret
// 		}

// 		id, err := createIngress(applicationName, domainName, secret)
// 		if err != nil {
// 			return err
// 		}

// 		logging.Info("Created ingress", logging.NewPair("domain-name", domainName), logging.NewPair("id", id))

// 		fmt.Printf("Ingress for domain name '%s' has been created\n", domainName)

// 		return nil
// 	},
// }

// func init() {
// 	createCmd.AddCommand(createIngressCommand)

// 	createIngressCommand.Flags().StringVar(&applicationName, "application", "", "Vamp cloud ingress application name")
// 	createIngressCommand.MarkFlagRequired("application")

// 	createIngressCommand.Flags().StringVar(&tslSecret, "tls-secret", "", "Vamp cloud ingress tls secret")
// }
