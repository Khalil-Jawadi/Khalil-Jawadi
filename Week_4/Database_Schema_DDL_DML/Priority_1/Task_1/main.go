-- Create User Table
CREATE TABLE User (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Points INT DEFAULT 0
);

-- Create Bread Table
CREATE TABLE Bread (
    BreadID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Price DECIMAL(10, 2) NOT NULL
);

-- Create PaymentMethod Table
CREATE TABLE PaymentMethod (
    PaymentMethodID INT AUTO_INCREMENT PRIMARY KEY,
    MethodName VARCHAR(255) NOT NULL
);

-- Create Purchase Table
CREATE TABLE Purchase (
    PurchaseID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT,
    PurchaseDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    TotalAmount DECIMAL(10, 2),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

-- Create PurchaseDetail Table
CREATE TABLE PurchaseDetail (
    PurchaseDetailID INT AUTO_INCREMENT PRIMARY KEY,
    PurchaseID INT,
    BreadID INT,
    Quantity INT NOT NULL,
    Price DECIMAL(10, 2),
    FOREIGN KEY (PurchaseID) REFERENCES Purchase(PurchaseID),
    FOREIGN KEY (BreadID) REFERENCES Bread(BreadID)
);

-- Create Ingredient Table
CREATE TABLE Ingredient (
    IngredientID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    QuantityAvailable DECIMAL(10, 2) NOT NULL
);

-- Create BreadIngredient Table
CREATE TABLE BreadIngredient (
    BreadID INT,
    IngredientID INT,
    QuantityRequired DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (BreadID, IngredientID),
    FOREIGN KEY (BreadID) REFERENCES Bread(BreadID),
    FOREIGN KEY (IngredientID) REFERENCES Ingredient(IngredientID)
);