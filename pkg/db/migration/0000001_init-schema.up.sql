CREATE TABLE IF NOT EXISTS Users (
  UserId SERIAL PRIMARY KEY,
  Email varchar(255) UNIQUE NOT NULL,
  Password varchar(255) NOT NULL,
  Status varchar(30) NOT NULL,
  CreatedAt timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS Consumers (
  ConsumerId SERIAL PRIMARY KEY,
  NIK varchar(60) NULL UNIQUE,
  FullName varchar(255) NOT NULL,
  LegalName varchar(255) DEFAULT '',
  BirthPlace varchar(255) DEFAULT '',
  BirthDate date DEFAULT CURRENT_TIMESTAMP,
  Salary decimal DEFAULT 0,
  Address varchar(255) DEFAULT '',
  KTPPhoto varchar(255) DEFAULT '',
  SelfiePhoto varchar(255) DEFAULT '',
  UserID int NOT NULL,
  FOREIGN KEY (UserId) REFERENCES Users(UserId)
);

CREATE TABLE IF NOT EXISTS Products (
  ProductId SERIAL PRIMARY KEY,
  Name varchar(255) NOT NULL,
  Category varchar(75) NOT NULL,
  PriceOTR decimal NOT NULL
);

CREATE TABLE IF NOT EXISTS Transactions (
  TransactionId SERIAL PRIMARY KEY,
  ContractNumber varchar(100) NOT NULL,
  AdminFee decimal NOT NULL,
  InstallmentAmount decimal NOT NULL,
  InterestAmount decimal NOT NULL,
  ConsumerId int NOT NULL,
  ProductId int NOT NULL,
  FOREIGN KEY (ConsumerId) REFERENCES Consumers(ConsumerId),
  FOREIGN KEY (ProductId) REFERENCES Products(ProductId)
);

CREATE TABLE IF NOT EXISTS Loans (
  LoanId SERIAL PRIMARY KEY,
  OneMonthLimit decimal NOT NULL,
  TwoMonthLimit decimal NOT NULL,
  ThreeMonthLimit decimal NOT NULL,
  SixMonthLimit decimal NOT NULL,
  ConsumerId int NOT NULL,
  FOREIGN KEY (ConsumerId) REFERENCES Consumers(ConsumerId)
);