CREATE TABLE Data_test (
    `ID`        INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `pID`       INT NOT NULL,
    `bmpTemp`   FLOAT NOT NULL,
    `mpuTemp`   FLOAT NOT NULL,
    `pressure`  FLOAT NOT NULL,
    `ax`        FLOAT NOT NULL,
    `ay`        FLOAT NOT NULL,
    `az`        FLOAT NOT NULL,
    `gx`        FLOAT NOT NULL,
    `gz`        FLOAT NOT NULL,
    `gy`        FLOAT NOT NULL,
    `latitude`  FLOAT NOT NULL,
    `longitude` FLOAT NOT NULL,
    `gpsAltitude`  FLOAT NOT NULL,
    `mpuAltitude`  FLOAT NOT NULL,
    `time` VARCHAR(4) NOT NULL,
);