CREATE TABLE Users
(
    id SERIAL,
    firstName TEXT NOT NULL,
    secondName TEXT NOT NULL,
    email TEXT NOT NULL,
    phoneNumber TEXT NOT NULL,
    balance TEXT NOT NULL,
    address TEXT NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)

CREATE TABLE Transactions
(
    senderUserId TEXT NOT NULL ,
    receiverUserId TEXT NOT NULL,
    amount TEXT NOT NULL
)