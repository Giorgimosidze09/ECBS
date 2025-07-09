package config_manager

// import (
// 	"context"
// 	"fmt"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
// )

// type AWSAppConfig struct {
// 	service string
// 	app     string
// 	env     string
// }

// func NewAWSAppConfigManager(appName, service, env string) AWSAppConfig {
// 	return AWSAppConfig{
// 		app:     appName,
// 		service: service,
// 		env:     env,
// 	}
// }

// func (c AWSAppConfig) GetConfig(ctx context.Context) ([]byte, error) {
// 	configProfile := fmt.Sprintf("%s-%s", c.service, c.env)

// 	return fetchAppConfig(ctx, c.app, c.env, configProfile)
// }

// func fetchAppConfig(
// 	ctx context.Context,
// 	application,
// 	env,
// 	profile string,
// ) ([]byte, error) {
// 	cfg, err := config.LoadDefaultConfig(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load AWS config: %w", err)
// 	}

// 	appconfigClient := appconfigdata.NewFromConfig(cfg)

// 	sessionOutput, err := appconfigClient.StartConfigurationSession(ctx, &appconfigdata.StartConfigurationSessionInput{
// 		ApplicationIdentifier:          aws.String(application),
// 		EnvironmentIdentifier:          aws.String(env),
// 		ConfigurationProfileIdentifier: aws.String(profile),
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to start configuration session: %w", err)
// 	}

// 	latestServiceConfig, err := appconfigClient.GetLatestConfiguration(ctx, &appconfigdata.GetLatestConfigurationInput{
// 		ConfigurationToken: sessionOutput.InitialConfigurationToken,
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get latest configuration: %w", err)
// 	}

// 	if latestServiceConfig.Configuration == nil {
// 		return nil, fmt.Errorf("configuration data is nil")
// 	}

// 	return latestServiceConfig.Configuration, nil
// }
