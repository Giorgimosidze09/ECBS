package main

import (
	"config_manager"
	"context"
	"fmt"
	"log"
	"os"

	sm "shared/nats_client/subscribe-manager"

	db_config "database/config"
	db "database/db"
	handlers "database/handlers"
)

func main() {
	// cfgManager := config_manager.NewConfigManager[db_config.DBConfig](
	// 	config_manager.GetDefaultSettings(),
	// )

	dir, _ := os.Getwd()
	log.Println("Running from directory:", dir)

	cfgManager := config_manager.NewConfigManager[db_config.DBConfig]("./config.local.json")

	cfg, err := cfgManager.GetConfig(context.Background())
	if err != nil {
		panic(err)
	}

	db_config.Set(cfg)

	db.InitDB()

	subscriber := sm.InitManager(sm.SubscribeManagerConfig{
		Name:    "service.database",
		NatsURL: cfg.NatsURL,
	})

	subscriber.RegisterSyncListener("create", handlers.CreateUser)
	subscriber.RegisterSyncListener("create_card", handlers.CreateCard)
	subscriber.RegisterSyncListener("top_up", handlers.TopUpBalance)
	subscriber.RegisterSyncListener("devices", handlers.CreateDevices)

	subscriber.RegisterSyncListener("validate", handlers.ValidateCard)
	subscriber.RegisterSyncListener("validate", handlers.Authorization)
	subscriber.RegisterSyncListener("access_logs", handlers.SyncAccessLogs)

	subscriber.RegisterSyncListener("count_users", handlers.CountUsers)
	subscriber.RegisterSyncListener("count_cards", handlers.CountCards)
	subscriber.RegisterSyncListener("total_balance", handlers.TotalBalance)
	subscriber.RegisterSyncListener("users_list", handlers.UsersList)
	subscriber.RegisterSyncListener("cards_list", handlers.CardsList)
	subscriber.RegisterSyncListener("balande_list", handlers.BalanceList)
	subscriber.RegisterSyncListener("devices_list", handlers.DevicesList)
	subscriber.RegisterSyncListener("add_card_activation", handlers.AddCardActivation)

	subscriber.RegisterSyncListener("charges_list", handlers.ChargesList)
	subscriber.RegisterSyncListener("ride_cost", handlers.RideCost)

	subscriber.RegisterSyncListener("update_user", handlers.UpdateUser)
	subscriber.RegisterSyncListener("delete_user", handlers.SoftDeleteUser)
	subscriber.RegisterSyncListener("get_user_by_id", handlers.GetUserByID)

	subscriber.RegisterSyncListener("update_card", handlers.UpdateCard)
	subscriber.RegisterSyncListener("delete_card", handlers.SoftDeleteCard)
	subscriber.RegisterSyncListener("get_card_by_id", handlers.GetCardByID)

	subscriber.RegisterSyncListener("update_device", handlers.UpdateDevice)
	subscriber.RegisterSyncListener("delete_device", handlers.SoftDeleteDevice)
	subscriber.RegisterSyncListener("get_device_by_id", handlers.GetDeviceByID)

	fmt.Println("Database service ready")

	select {}
}
