package Database

import (
	"database/sql"
	"fmt"
	"github.com/Autonomeasure/GroundStation/pkg/Radio"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) Open() error {
	d, err := sql.Open("mysql", "root:Test123@unix(/var/run/mysqld/mysqld.sock)/CanSat")
	if err != nil {
		return err
	}
	db.DB = d

	return nil
}

func (db *Database) Query(query string) (*sql.Rows, error) {
	rows, err := db.DB.Query(query)

	return rows, err
}

func (db *Database) Exec(query string, params ...interface{}) (sql.Result, error) {
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := statement.Exec(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *Database) SaveRadioPacket(packet Radio.Packet) error {
	statement, err := db.DB.Prepare("INSERT INTO Data_test (pID, bmpTemp, mpuTemp, pressure, ax, ay, az, gx, gy, gz, latitude, longitude, altitude, gpsSpeed) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")


	if err != nil {
		return err
	}

	_, err = statement.Exec(packet.ID, packet.Temperature.BMP, packet.Temperature.MPU, packet.Pressure, packet.Acceleration.X, packet.Acceleration.Y, packet.Acceleration.Z, packet.Gyroscope.X, packet.Gyroscope.Y, packet.Gyroscope.Z, packet.GPS.Latitude, packet.GPS.Longitude, packet.GPS.Altitude, packet.GPS.Speed)

	if err != nil {
		return err
	}

	// Everything went correctly, no error found so return nil
	return nil
}

func (db *Database) GetRadioPacket(packetID uint32) (Radio.Packet, error) {
	fmt.Println("Checking packet ID: ", packetID)
	rows, err := db.DB.Query("SElECT * FROM Data_test WHERE pID = ?", packetID)
	var p Radio.Packet
	if err != nil {
		return p, err
	}
	var id int
	rows.Next()
	rows.Scan(&id, &p.ID, &p.Temperature.BMP, &p.Temperature.MPU, &p.Pressure, &p.Acceleration.X, &p.Acceleration.Y, &p.Acceleration.Z, &p.Gyroscope.X, &p.Gyroscope.Y, &p.Gyroscope.Z, &p.GPS.Latitude, &p.GPS.Longitude, &p.GPS.Altitude, &p.GPS.Speed)
	return p, nil
}