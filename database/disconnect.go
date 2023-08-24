package database

func Disconnect() {
	println("Disconnecting database...")

	dbInstance, _ := Database.DB.DB()

	_ = dbInstance.Close()
}
