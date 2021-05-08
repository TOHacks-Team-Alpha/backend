package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	// defer repo.conn.Close(c.Request.Context())

	user := &User{ID: c.GetString("uid")}

	log.Printf("[getUser] | %v", user.ID)

	err := repo.conn.QueryRow(context.Background(), selectUserByID, user.ID).Scan(&user.ID, &user.Name, &user.Coins, &user.NumTripsDriven, &user.NumTripsRidden, &user.TotalDistance)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, &user)
}

func putUser(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	// defer repo.conn.Close(c.Request.Context())

	var user *User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(501, err)
		return
	}

	user.ID = c.GetString("uid")
	log.Printf("[updateUser] | %v", user)

	_, err = repo.conn.Exec(context.Background(), updateUserbyID, &user.ID, &user.Name)
	if err != nil {
		log.Printf("[Error] [updateUser] | %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, &user)
}

func getDriveByID(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	// defer repo.conn.Close(c.Request.Context())

	user := &User{ID: c.GetString("uid")}

	log.Printf("[getUser] | %v", user.ID)

	err := repo.conn.QueryRow(context.Background(), selectUserByID, user.ID).Scan(&user.ID, &user.Name, &user.Coins, &user.NumTripsDriven, &user.NumTripsRidden, &user.TotalDistance)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, &user)
}

func getDrives(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	rows, err := repo.conn.Query(context.Background(), selectDrives, c.GetString("uid"))
	if err != nil {
		log.Printf("[GET DRIVES] %v", err)
		c.JSON(500, err)
		return
	}

	driveList := make([]*Drive, 0)

	for rows.Next() {
		drive := &Drive{}
		err = rows.Scan(
		//TODO
		)
		if err != nil {
			log.Printf("[GET DRIVES 2] %v", err)
			c.JSON(501, err)
			return
		}
		driveList = append(driveList, drive)
	}

	c.JSON(200, &driveList)
}

func postDrive(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	var drive *Drive
	err := c.Bind(&drive)
	if err != nil {
		log.Printf("[POST DRIVE] %v", err)
		c.JSON(501, err)
		return
	}

	drive.DriverID = c.GetString("uid")

	_, err = repo.conn.Exec(context.Background(), createDrive) //TODO

	if err != nil {
		log.Printf("[POST DRIVE] %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, true)
}

func getDriveRequests(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	rows, err := repo.conn.Query(context.Background(), selectDriveRequestsByID, c.GetString("uid"))
	if err != nil {
		log.Printf("[GET DRIVE REQUESTS] %v", err)
		c.JSON(500, err)
		return
	}

	driveRequestList := make([]*DriveRequest, 0)

	for rows.Next() {
		driveRequest := &DriveRequest{}
		err = rows.Scan(
			&driveRequest.DriveID,
			//TODO
		)
		if err != nil {
			log.Printf("[GET DRIVE REQUESTS 2] %v", err)
			c.JSON(501, err)
			return
		}
		driveRequestList = append(driveRequestList, driveRequest)
	}

	c.JSON(200, &driveRequestList)
}

func getRideRequests(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	rows, err := repo.conn.Query(context.Background(), selectRideRequestsByID, c.GetString("uid"))
	if err != nil {
		log.Printf("[GET DRIVE REQUESTS] %v", err)
		c.JSON(500, err)
		return
	}

	driveRequestList := make([]*DriveRequest, 0)

	for rows.Next() {
		driveRequest := &DriveRequest{}
		err = rows.Scan(
			&driveRequest.DriveID,
			//TODO
		)
		if err != nil {
			log.Printf("[GET DRIVE REQUESTS 2] %v", err)
			c.JSON(501, err)
			return
		}
		driveRequestList = append(driveRequestList, driveRequest)
	}

	c.JSON(200, &driveRequestList)
}

func postDriveRequest(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	var drive *Drive
	err := c.Bind(&drive)
	if err != nil {
		log.Printf("[POST DRIVE REQ] %v", err)
		c.JSON(501, err)
		return
	}

	drive.DriverID = c.GetString("uid")

	_, err = repo.conn.Exec(context.Background(), createDriveRequest) //TODO

	if err != nil {
		log.Printf("[POST DRIVE REQ 2] %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, drive)
}

func putDriveRequest(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	var drive *Drive
	err := c.Bind(&drive)
	if err != nil {
		log.Printf("[POST DRIVE REQ] %v", err)
		c.JSON(501, err)
		return
	}

	drive.DriverID = c.GetString("uid")

	_, err = repo.conn.Exec(context.Background(), updateDriveRequest) //TODO

	if err != nil {
		log.Printf("[POST DRIVE REQ 2] %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, drive)
}

const (
	selectUserByID = "SELECT * FROM useriles WHERE uid = $1"
	updateUserbyID = "INSERT INTO useriles (uid, name, age, weight, low, high) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (uid) DO UPDATE SET (name, age, weight, low, high) = ($2, $3, $4, $5, $6);"
	selectData     = "SELECT * FROM data WHERE uid = $1 ORDER BY time"
	postDataStmt   = "INSERT INTO data (uid, time, bg) VALUES ($1, $2, $3)"
)
