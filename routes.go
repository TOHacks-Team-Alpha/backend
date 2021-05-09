package main

import (
	"context"
	"log"
	"strconv"

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

	// user := &User{ID: c.GetString("uid")}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err)
		return
	}

	log.Printf("[get drive by id] | %v", id)

	drive := &Drive{}

	err = repo.conn.QueryRow(context.Background(), selectDriveByID, id).Scan(&drive.DriveID, &drive.DriverID, &drive.Time, &drive.SpaceAvailable, &drive.StartAddress, &drive.DestAddress, &drive.StartLat, &drive.StartLng, &drive.DestLat, &drive.DestLng)
	if err != nil {
		log.Printf("[get drive by id] | %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, &drive)
}

func getDrives(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())

	startLat, startLng, startRadius, destLat, destLng, destRadius := c.Query("start_lat"), c.Query("start_lng"), c.Query("start_radius"), c.Query("dest_lat"), c.Query("dest_lng"), c.Query("dest_radius")
	rows, err := repo.conn.Query(context.Background(), selectDrives, startLat, startLng, startRadius, destLat, destLng, destRadius, 20)
	if err != nil {
		log.Printf("[GET DRIVES] %v", err)
		c.JSON(500, err)
		return
	}

	type DriveResponse struct {
		Drive
		StartDistance float64 `json:"start_distance"`
		DestDistance  float64 `json:"dest_distance"`
		User
	}

	driveList := make([]*DriveResponse, 0)

	for rows.Next() {
		drive := &DriveResponse{}
		err = rows.Scan(
			&drive.DriveID,
			&drive.DriverID,
			&drive.Time,
			&drive.SpaceAvailable,
			&drive.StartAddress,
			&drive.DestAddress,
			&drive.StartLat,
			&drive.StartLng,
			&drive.DestLat,
			&drive.DestLng,
			&drive.StartDistance,
			&drive.DestDistance,
			&drive.ID,
			&drive.Name,
			&drive.Coins,
			&drive.NumTripsDriven,
			&drive.NumTripsRidden,
			&drive.TotalDistance,
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

	_, err = repo.conn.Exec(context.Background(), createDrive, &drive.DriverID, &drive.Time, &drive.SpaceAvailable, &drive.StartAddress, &drive.DestAddress, &drive.StartLat, &drive.StartLng, &drive.DestLat, &drive.DestLng) //TODO

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
	log.Printf("[GET DRIVE REQUESTS] %v", c.GetString("uid"))
	rows, err := repo.conn.Query(context.Background(), getDriveReqsForDriver, c.GetString("uid"))
	if err != nil {
		log.Printf("[GET DRIVE REQUESTS] %v", err)
		c.JSON(500, err)
		return
	}

	type DriveResponse struct {
		Req   DriveRequest
		Drive Drive
	}

	driveRequestList := make([]*DriveResponse, 0)

	for rows.Next() {
		drive := &DriveResponse{}
		err = rows.Scan(
			&drive.Req.DriveID,
			&drive.Drive.DriverID,
			&drive.Req.RiderID,
			&drive.Req.Status,
			&drive.Drive.Time,
			&drive.Drive.SpaceAvailable,
			&drive.Drive.StartAddress,
			&drive.Drive.DestAddress,
			&drive.Drive.StartLat,
			&drive.Drive.StartLng,
			&drive.Drive.DestLat,
			&drive.Drive.DestLng,
		)
		if err != nil {
			log.Printf("[GET DRIVE REQUESTS 2] %v", err)
			c.JSON(501, err)
			return
		}
		drive.Drive.DriveID = drive.Req.DriveID
		driveRequestList = append(driveRequestList, drive)
	}

	c.JSON(200, &driveRequestList)
}

func getRideRequests(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	rows, err := repo.conn.Query(context.Background(), getDriveReqsForRider, c.GetString("uid"))
	if err != nil {
		log.Printf("[GET DRIVE REQUESTS] %v", err)
		c.JSON(500, err)
		return
	}

	type DriveResponse struct {
		Req   DriveRequest
		Drive Drive
	}

	driveRequestList := make([]*DriveResponse, 0)

	for rows.Next() {
		drive := &DriveResponse{}
		err = rows.Scan(
			&drive.Req.DriveID,
			&drive.Drive.DriverID,
			&drive.Req.RiderID,
			&drive.Req.Status,
			&drive.Drive.Time,
			&drive.Drive.SpaceAvailable,
			&drive.Drive.StartAddress,
			&drive.Drive.DestAddress,
			&drive.Drive.StartLat,
			&drive.Drive.StartLng,
			&drive.Drive.DestLat,
			&drive.Drive.DestLng,
		)
		if err != nil {
			log.Printf("[GET DRIVE REQUESTS 2] %v", err)
			c.JSON(501, err)
			return
		}
		drive.Drive.DriveID = drive.Req.DriveID
		driveRequestList = append(driveRequestList, drive)
	}

	c.JSON(200, &driveRequestList)
}

func postDriveRequest(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	var driveReq *DriveRequest
	err := c.Bind(&driveReq)
	if err != nil {
		log.Printf("[POST DRIVE REQ] %v", err)
		c.JSON(501, err)
		return
	}

	driveReq.RiderID = c.GetString("uid")

	_, err = repo.conn.Exec(context.Background(), createDriveRequest, &driveReq.DriveID, &driveReq.RiderID, "sent")

	if err != nil {
		log.Printf("[POST DRIVE REQ 2] %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, driveReq)
}

func putDriveRequest(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	var driveReq *DriveRequest
	err := c.Bind(&driveReq)
	if err != nil {
		log.Printf("[PUT DRIVE REQ] %v", err)
		c.JSON(501, err)
		return
	}

	drive := &Drive{}
	err = repo.conn.QueryRow(context.Background(), selectDriveByID, driveReq.DriveID).Scan(&drive.DriveID, &drive.DriverID, &drive.Time, &drive.SpaceAvailable, &drive.StartAddress, &drive.DestAddress, &drive.StartLat, &drive.StartLng, &drive.DestLat, &drive.DestLng)
	if err != nil {
		log.Printf("[PUT DRIVE REQ] %v", err)
		c.JSON(500, err)
		return
	}

	userWhoSentRequest := c.GetString("uid")

	if drive.DriverID == userWhoSentRequest {
		log.Printf("[PUT DRIVE REQ - DRIVER] %v", userWhoSentRequest)
		_, err = repo.conn.Exec(context.Background(), updateDriveReq, driveReq.Status, drive.DriveID, driveReq.RiderID) //TODO:
		if err != nil {
			log.Printf("[PUT DRIVE REQ - DRIVER] %v", err)
			c.JSON(500, err)
			return
		}

		switch driveReq.Status {
		case "accepted":
			// driver accepts new rider
			// decrement space available in drive
			_, err = repo.conn.Exec(context.Background(), updateDrive, -1, drive.DriveID) //TODO:
			if err != nil {
				log.Printf("[POST DRIVE REQ 2] %v", err)
				c.JSON(500, err)
				return
			}
		case "rejected":
			//  driver says fuck the new rider
		case "cancelled":
			// driver cancels the drive
			// delete the drive
			_, err = repo.conn.Exec(context.Background(), deleteDrive, drive.DriveID) //TODO:
			if err != nil {
				log.Printf("[POST DRIVE REQ 2] %v", err)
				c.JSON(500, err)
				return
			}
		case "complete":
			//  driver is done zooming
			//  update users drive ridden
			_, err = repo.conn.Exec(context.Background(), updateUserTripsRidden, 1, driveReq.RiderID)
			if err != nil {
				log.Printf("[POST DRIVE REQ 2] %v", err)
				c.JSON(500, err)
				return
			}
			// add to drives driven
			_, err = repo.conn.Exec(context.Background(), updateUserTripsDriven, 1, drive.DriveID, 1)
			if err != nil {
				log.Printf("[POST DRIVE REQ 2] %v", err)
				c.JSON(500, err)
				return
			}
			// TODO: update Total Distance

		}

	} else if driveReq.RiderID == userWhoSentRequest {
		log.Printf("[PUT DRIVE REQ - RIDER] %v", userWhoSentRequest)
		_, err = repo.conn.Exec(context.Background(), updateDriveReq, driveReq.Status, drive.DriveID, driveReq.RiderID) //TODO:
		if err != nil {
			log.Printf("[PUT DRIVE REQ - RIDER] %v", err)
			c.JSON(500, err)
			return
		}
		switch driveReq.Status {
		case "cancelled":
			// rider cancels
			// increment space available
			_, err = repo.conn.Exec(context.Background(), updateDrive, 1, drive.DriveID) //TODO:
			if err != nil {
				log.Printf("[PUT] DRIVE REQ 2] %v", err)
				c.JSON(500, err)
				return
			}
		case "complete":
			// rider arrives at Dest
			// add to drives ridden

			//TODO: update the distance
			// update Total Distance
			// case "sent":
			// 	// rider is begging to ride
			// 	_, err = repo.conn.Exec(context.Background(), insertDriveReq, 1, driveReq.RiderID, 1)
		}
	} else {
		log.Printf("[PUT DRIVE REQ - WHO???] | %v", userWhoSentRequest)
	} //==> its the rider; else it is the driver probably?

	c.JSON(200, driveReq)
}

const (
	selectAllUsers        = "select * from users;"
	selectUserByID        = "SELECT * FROM users WHERE id = $1"
	updateUserbyID        = "INSERT INTO users VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET (name) = ($2);"
	updateUserTripsDriven = "UPDATE users SET num_trips_driven  = num_trips_driven  + $1 WHERE id = $2;"
	updateUserTripsRidden = "UPDATE users SET num_trips_ridden  = num_trips_driven  + $1 WHERE id = $2;"
	createDrive           = "INSERT INTO drives (driver_id, time, space_available, start_address, dest_address, start_lat, start_lng, dest_lat, dest_lng) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	createDriveRequest    = "INSERT INTO drive_reqs VALUES ($1, $2, $3);"
	updateUserCoins       = "UPDATE users AS u SET coins = u.coins + u2.coins FROM ( VALUES ($1, $2), ($3, $4)) AS u2(id, coins) WHERE u2.id=u.id;"
	getDriveReqs          = "SELECT * FROM drive_reqs where drive_id = $1;"
	// getUserDriveReqs      = "SELECT * FROM drive_reqs where rider_id = $1;"
	getDriveReqsForDriver = "SELECT drive_reqs.drive_id, driver_id, rider_id, status, time, space_available, start_address, dest_address, start_lat, start_lng, dest_lat, dest_lng FROM drive_reqs INNER JOIN drives ON drives.drive_id = drive_reqs.drive_id WHERE driver_id = $1;"
	getDriveReqsForRider  = "SELECT drive_reqs.drive_id, driver_id, rider_id, status, time, space_available, start_address, dest_address, start_lat, start_lng, dest_lat, dest_lng FROM drive_reqs INNER JOIN drives ON drives.drive_id = drive_reqs.drive_id WHERE rider_id = $1;"
	updateDriveReq        = "UPDATE drive_reqs SET status = $1 WHERE drive_id = $2 AND rider_id = $3;"
	insertDriveReq        = "INSERT INTO drive_reqs VALUES ($1, $2);"
	insertDrive           = "INSERT INTO drives (driver_id, time, space_available, start_address, dest_address, start_lat, start_lng, dest_lat, dest_lng) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	selectAllDrives       = "SELECT * FROM drives;"
	selectDriveByID       = "SELECT * FROM drives WHERE drive_id = $1"
	selectDrives          = "WITH cte AS ( " +
		"SELECT *, " +
		"( " +
		"6371 * " +
		"acos(cos(radians($1)) * " +
		"cos(radians(start_lat)) *  " +
		"cos(radians(start_lng) -  " +
		"radians($2)) +  " +
		"sin(radians($1)) * " +
		"sin(radians(start_lat))) " +
		") AS start_distance, " +
		"( " +
		"6371 * " +
		"acos(cos(radians($4)) * " +
		"cos(radians(dest_lat)) *  " +
		"cos(radians(dest_lng) -  " +
		"radians($5)) +  " +
		"sin(radians($4)) * " +
		"sin(radians(dest_lat))) " +
		") AS dest_distance  " +
		"FROM drives ), " +
		"dists AS (SELECT * FROM cte " +
		"WHERE cte.start_distance < $3 AND cte.dest_distance < $6 " +
		"ORDER BY cte.start_distance, cte.dest_distance LIMIT $7) " +
		"SELECT * " +
		"FROM dists " +
		"INNER JOIN users ON dists.driver_id = users.id;"
	updateDrive = "UPDATE drives SET space_available = space_available + $1 WHERE drive_id = $2;"
	deleteDrive = "DELETE drives WHERE drive_id = $1"
)
