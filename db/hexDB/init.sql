CREATE TABLE OperationsHistory (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Result VARCHAR(255),
    Operation VARCHAR(255),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);