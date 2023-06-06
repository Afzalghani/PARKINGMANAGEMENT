CREATE TABLE car (
    carId INT PRIMARY KEY,
    name VARCHAR(30),
    model VARCHAR(12),
    owner VARCHAR(30)
);

CREATE TABLE parkingslot (
    slotId INT PRIMARY KEY,
    size ENUM('1', '2', '3')
);

CREATE TABLE parker (
    parkerId INT PRIMARY KEY,
    name VARCHAR(30)
);

CREATE TABLE payment (
    paymentId INT PRIMARY KEY,
    paymentAmount DECIMAL(10, 2),
    paymentDate varchar(30),
    paymentLocation VARCHAR(34)
);

CREATE TABLE parkedcar (
    parkedcarId INT PRIMARY KEY,
    carId INT,
    slotId INT,
    parkerId INT,
    status ENUM('parked', 'left'),
    paymentId INT,
    FOREIGN KEY (carId) REFERENCES car (carId),
    FOREIGN KEY (slotId) REFERENCES parkingslot (slotId),
    FOREIGN KEY (parkerId) REFERENCES parker (parkerId),
    FOREIGN KEY (paymentId) REFERENCES payment (paymentId)
);
