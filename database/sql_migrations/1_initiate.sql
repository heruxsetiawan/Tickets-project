-- *migrate UP
-- *migrate StetementBegin
-- +migrate Up
CREATE TABLE Users (
    UserID SERIAL PRIMARY KEY,
    Username VARCHAR(50) NOT NULL UNIQUE,
    Email VARCHAR(100) NOT NULL UNIQUE,
    PasswordHash VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE Tickets (
    TicketID SERIAL PRIMARY KEY,
    CreatorID INT NOT NULL,
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    Status VARCHAR(50) NOT NULL DEFAULT 'Pending',
    StartTime TIMESTAMP,
    EndTime TIMESTAMP,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (CreatorID) REFERENCES Users(UserID)
);
CREATE TABLE Tasks (
    TaskID SERIAL PRIMARY KEY,
    TicketID INT NOT NULL,
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    Status VARCHAR(50) NOT NULL DEFAULT 'Pending',
    StartTime TIMESTAMP,
    EndTime TIMESTAMP,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (TicketID) REFERENCES Tickets(TicketID)
);
CREATE TABLE TicketAssignments (
    AssignmentID SERIAL PRIMARY KEY,
    TicketID INT NOT NULL UNIQUE,
    AssigneeID INT NOT NULL,
    AssignedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (TicketID) REFERENCES Tickets(TicketID),
    FOREIGN KEY (AssigneeID) REFERENCES Users(UserID)
);


-- *migrate StetementEnd